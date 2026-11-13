#!/bin/bash
# Usage: carry-forward-snapshot-results.sh <results-dir> [<results-dir> ...]
#
# Repos skipped by analyze-repo.sh (unchanged HEAD + unchanged analyzer
# version) produce no results in this run. Copies their raw JSONs from the
# most recent snapshot on the snapshots branch into the results directories
# so aggregation and dated snapshots stay complete.
#
# Results directories are named "<platform>-results" (e.g. odh-results);
# the platform is derived from the directory name. Destinations and
# snapshot lookups are keyed by platform and repository short name: the
# same repository (e.g. odh-dashboard) can be configured under multiple
# platforms and must stay separate.
set -euo pipefail

if [ $# -lt 1 ]; then
    echo "Usage: $0 <results-dir> [<results-dir> ...]" >&2
    exit 1
fi

RESULTS_DIRS=("$@")
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ANALYZER_DIR="$(dirname "${SCRIPT_DIR}")"
SCAN_CONFIG="${ANALYZER_DIR}/scan-config.yaml"

valid_security_findings() {
    python3 -c 'import json,sys; value=json.load(open(sys.argv[1])); sys.exit(0 if isinstance(value, list) else 1)' "$1" >/dev/null 2>&1
}

valid_architecture() {
    python3 -c 'import json,sys; value=json.load(open(sys.argv[1])); valid=isinstance(value,dict) and isinstance(value.get("repo"),str) and bool(value.get("repo")) and isinstance(value.get("commit_sha"),str) and bool(value.get("commit_sha")); sys.exit(0 if valid else 1)' "$1" >/dev/null 2>&1
}

valid_snapshot_provenance() {
    python3 -c '
import json,sys
try:
    with open(sys.argv[1]) as f: metadata=json.load(f)
    with open(sys.argv[2]) as f: component=json.load(f)
except (OSError, ValueError, UnicodeError):
    sys.exit(1)
if not isinstance(metadata, dict) or not isinstance(component, dict):
    sys.exit(1)
repos=metadata.get("repos_analyzed")
repo_versions=metadata.get("repo_analyzer_versions")
repo_version=repo_versions.get(sys.argv[3]) if isinstance(repo_versions, dict) else None
analyzer_version=repo_version or metadata.get("analyzer_version")
valid=(isinstance(repos, dict)
    and repos.get(sys.argv[3]) == component.get("commit_sha")
    and analyzer_version == component.get("analyzer_version"))
sys.exit(0 if valid else 1)
' "$1" "$2" "$3" >/dev/null 2>&1
}

scan_state_matches() {
    python3 -c 'import json,sys; state=json.load(open(sys.argv[1])).get(sys.argv[2],{}); data=json.load(open(sys.argv[3])); ok=state.get("commit_sha") and state.get("commit_sha")==data.get("commit_sha") and state.get("analyzer_version")==data.get("analyzer_version"); sys.exit(0 if ok else 1)' \
        "${STATE_FILE}" "$1" "$2" >/dev/null 2>&1
}

if ! command -v yq &>/dev/null; then
    echo "::error::yq not available; cannot determine configured repos for carry-forward"
    exit 1
fi

# Snapshot raw layout: output/<date>/<platform>/<date>/raw/<repo-name>/
# (no org component). Map each configured org/repo to its platform and
# short name; short names are assumed unique within a platform, matching
# the snapshot publication logic. Only repos marked as intentionally
# skipped (an _skipped marker from analyze-repo.sh) are restored: a repo
# with no results and no marker failed its analysis and must not be
# silently masked by older data.
declare -A DEST_FOR_KEY
declare -A REPO_FOR_KEY
MISSING_KEYS=()
UNRESTORED_KEYS=()
UNRESTORED_UNAVAILABLE_KEYS=()
RESTORED_FALLBACKS=()
declare -A FALLBACK_FOR_KEY
mark_failed() {
    local key="$1"
    local dest="${DEST_FOR_KEY[${key}]:-}"
    if [ -n "$dest" ]; then
        mkdir -p "$dest"
        find "$dest" -type f \
            \( -name "component-architecture.json" -o -name "security-findings.json" \) \
            -delete 2>/dev/null || true
        python3 -c 'import json,os; from datetime import datetime,timezone; print(json.dumps({"timestamp":datetime.now(timezone.utc).isoformat(timespec="microseconds").replace("+00:00","Z"),"workflow_run_number":int(os.getenv("GITHUB_RUN_NUMBER","0")),"workflow_run_attempt":int(os.getenv("GITHUB_RUN_ATTEMPT","0"))}))' \
            > "${dest}/_failed" || touch "${dest}/_failed"
    fi
}

mark_missing_unrestorable() {
    local key
    for key in "${MISSING_KEYS[@]}"; do
        mark_failed "$key"
        if [ "${FALLBACK_FOR_KEY[${key}]:-}" = "true" ]; then
            UNRESTORED_UNAVAILABLE_KEYS+=("${key}")
        else
            UNRESTORED_KEYS+=("${key}")
        fi
    done
}

report_unavailable_exclusions() {
    if [ ${#UNRESTORED_UNAVAILABLE_KEYS[@]} -gt 0 ]; then
        echo "::warning::${#UNRESTORED_UNAVAILABLE_KEYS[@]} failed/unavailable repo(s) excluded because no verified previous snapshot exists: ${UNRESTORED_UNAVAILABLE_KEYS[*]}"
        if [ -n "${GITHUB_STEP_SUMMARY:-}" ]; then
            {
                echo "### Failed/unavailable repositories excluded from this snapshot"
                printf -- '- %s\n' "${UNRESTORED_UNAVAILABLE_KEYS[@]}"
            } >> "${GITHUB_STEP_SUMMARY}"
        fi
    fi
}

for results_dir in "${RESULTS_DIRS[@]}"; do
    platform="$(basename "${results_dir}" | sed 's/-results$//')"
    if ! repos=$(yq -r ".platforms.${platform}.orgs | to_entries[] | .key as \$org | .value[] | \"\(\$org)/\(.)\"" "${SCAN_CONFIG}" 2>/dev/null); then
        echo "::error::failed to read configured ${platform} repositories from ${SCAN_CONFIG}"
        exit 1
    fi
    while IFS= read -r repo; do
        [ -n "$repo" ] || continue
        short="${repo##*/}"
        key="${platform}/${short}"
        repo_dir="${results_dir}/${repo}"
        DEST_FOR_KEY["${key}"]="${repo_dir}"
        REPO_FOR_KEY["${key}"]="${repo}"
        component_json="$(find "${repo_dir}" -name component-architecture.json -type f -print -quit 2>/dev/null || true)"
        if [ -n "${component_json}" ]; then
            analysis_dir="$(dirname "${component_json}")"
            if ! valid_architecture "${component_json}" || \
                [ ! -f "${analysis_dir}/security-findings.json" ] || \
                ! valid_security_findings "${analysis_dir}/security-findings.json" || \
                [ -f "${repo_dir}/_failed" ] || [ -f "${analysis_dir}/_failed" ]; then
                mark_failed "${key}"
                MISSING_KEYS+=("${key}")
                FALLBACK_FOR_KEY["${key}"]=true
            fi
            continue
        fi
        if [ -f "${repo_dir}/security-findings.json" ] || [ -f "${repo_dir}/_failed" ]; then
            mark_failed "${key}"
            MISSING_KEYS+=("${key}")
            FALLBACK_FOR_KEY["${key}"]=true
            continue
        fi
        if [ -f "${repo_dir}/_unavailable" ]; then
            MISSING_KEYS+=("${key}")
            FALLBACK_FOR_KEY["${key}"]=true
            continue
        fi
        if [ ! -f "${repo_dir}/_skipped" ]; then
            mark_failed "${key}"
            MISSING_KEYS+=("${key}")
            FALLBACK_FOR_KEY["${key}"]=true
            continue
        fi
        # DEST_FOR_KEY was already populated above; every verified skip
        # must still be queued for restoration.
        MISSING_KEYS+=("${key}")
    done <<< "${repos}"
done

if [ ${#MISSING_KEYS[@]} -eq 0 ]; then
    echo "[*] No missing repos to carry forward"
    exit 0
fi

echo "[*] Carrying forward ${#MISSING_KEYS[@]} repo(s) from the latest snapshot"

# Fetch through the checked-out repository so its configured credentials are
# used, then read the snapshot branch from a sparse linked worktree.
SNAP_PARENT="$(mktemp -d)"
SNAP_DIR="${SNAP_PARENT}/snapshots-worktree"
cleanup_snapshots() {
    git -C "${ANALYZER_DIR}" worktree remove --force "${SNAP_DIR}" >/dev/null 2>&1 || true
    rm -rf "${SNAP_PARENT}"
}
trap cleanup_snapshots EXIT

PATTERNS=()
for key in "${MISSING_KEYS[@]}"; do
    PATTERNS+=("output/*/${key%%/*}/*/raw/${key##*/}/")
    PATTERNS+=("output/*/${key%%/*}/*/snapshot-metadata.json")
done

if ! git -C "${ANALYZER_DIR}" fetch origin \
        +refs/heads/snapshots:refs/remotes/origin/snapshots 2>/dev/null; then
    echo "::warning::snapshots branch unavailable; cannot carry forward ${#MISSING_KEYS[@]} marked repo(s)"
    mark_missing_unrestorable
    report_unavailable_exclusions
    if [ ${#UNRESTORED_KEYS[@]} -gt 0 ]; then
        echo "::error::${#UNRESTORED_KEYS[@]} verified skipped repo(s) cannot be restored"
        exit 1
    fi
    exit 0
fi

if ! git -C "${ANALYZER_DIR}" worktree add --detach --no-checkout \
        "${SNAP_DIR}" refs/remotes/origin/snapshots >/dev/null 2>&1; then
    echo "::warning::snapshots branch unavailable; cannot carry forward ${#MISSING_KEYS[@]} marked repo(s)"
    mark_missing_unrestorable
    report_unavailable_exclusions
    if [ ${#UNRESTORED_KEYS[@]} -gt 0 ]; then
        echo "::error::${#UNRESTORED_KEYS[@]} verified skipped repo(s) cannot be restored"
        exit 1
    fi
    exit 0
fi

(
    git -C "${SNAP_DIR}" sparse-checkout init --no-cone
    git -C "${SNAP_DIR}" sparse-checkout set --no-cone "${PATTERNS[@]}"
    git -C "${SNAP_DIR}" checkout --detach refs/remotes/origin/snapshots
)

# No snapshot data at all is fatal when repos must be restored: publishing
# without them while advancing the scan state would strand them.
if [ ! -d "${SNAP_DIR}/output" ]; then
    echo "::warning::no snapshot output directory; cannot carry forward ${#MISSING_KEYS[@]} marked repo(s)"
    mark_missing_unrestorable
    report_unavailable_exclusions
    if [ ${#UNRESTORED_KEYS[@]} -gt 0 ]; then
        echo "::error::${#UNRESTORED_KEYS[@]} verified skipped repo(s) cannot be restored"
        exit 1
    fi
    exit 0
fi

carried=0
STATE_FILE="${ANALYZER_DIR}/.last-scan-shas.json"
for key in "${MISSING_KEYS[@]}"; do
    platform="${key%%/*}"
    short="${key##*/}"
    # Try dated snapshots in publication order. Custom labels and future
    # dates stay excluded so they cannot override current dated data.
    today="$(date -u +%Y-%m-%d)"
    CANDIDATES=()
    while IFS= read -r candidate; do
        [ -n "${candidate}" ] && CANDIDATES+=("${candidate}")
    done < <(find "${SNAP_DIR}/output" -type d -path "*/${platform}/*/raw/${short}" 2>/dev/null \
        | grep -E "/output/[0-9]{4}-[0-9]{2}-[0-9]{2}/" \
        | awk -v today="$today" -v prefix="${SNAP_DIR}/output/" \
            'index($0, prefix) == 1 {
                 date = substr($0, length(prefix) + 1, 10)
                 if (date <= today) print
             }' \
        | python3 -c '
import json, os, sys
from datetime import datetime, timezone

def rank(candidate):
    metadata_path = candidate.rsplit("/raw/", 1)[0] + "/snapshot-metadata.json"
    try:
        with open(metadata_path) as f:
            metadata = json.load(f)
        run_number = int(metadata.get("workflow_run_number", 0))
        run_attempt = int(metadata.get("workflow_run_attempt", 0))
    except (OSError, ValueError, TypeError, AttributeError):
        run_number = run_attempt = 0
        metadata = {}
    timestamp = metadata.get("timestamp", "")
    timestamp_key = 0
    if isinstance(timestamp, str) and timestamp:
        try:
            parsed = datetime.fromisoformat(timestamp.replace("Z", "+00:00"))
            if parsed.tzinfo is not None and parsed.utcoffset() is not None:
                timestamp_key = parsed.astimezone(timezone.utc).timestamp()
        except ValueError:
            pass
    return run_number, run_attempt, timestamp_key, candidate

candidates = [line.rstrip("\n") for line in sys.stdin if line.rstrip("\n")]
for candidate in sorted(candidates, key=rank, reverse=True):
    print(candidate)
')

    src=""
    state_key="${platform}/${REPO_FOR_KEY[${key}]}"
    for candidate in "${CANDIDATES[@]}"; do
        if [ ! -f "${candidate}/component-architecture.json" ] || \
            ! valid_architecture "${candidate}/component-architecture.json" || \
            [ ! -f "${candidate}/security-findings.json" ] || \
            ! valid_security_findings "${candidate}/security-findings.json"; then
            continue
        fi
        if [ "${FALLBACK_FOR_KEY[${key}]:-}" = "true" ]; then
            snapshot_metadata="${candidate%/raw/*}/snapshot-metadata.json"
            if ! valid_snapshot_provenance \
                "${snapshot_metadata}" "${candidate}/component-architecture.json" "${REPO_FOR_KEY[${key}]}"; then
                continue
            fi
        else
            if [ ! -f "${STATE_FILE}" ] || \
                ! scan_state_matches "${state_key}" "${candidate}/component-architecture.json"; then
                continue
            fi
        fi
        src="${candidate}"
        break
    done

    if [ -z "${src}" ]; then
        if [ "${FALLBACK_FOR_KEY[${key}]:-}" = "true" ]; then
            echo "::warning::${key}: failed/unavailable repo has no provenance-matching prior snapshot; excluding it from this snapshot"
            mark_failed "${key}"
            UNRESTORED_UNAVAILABLE_KEYS+=("${key}")
        else
            echo "::warning::${key}: no snapshot matches the scan state entry"
            UNRESTORED_KEYS+=("${key}")
            mark_failed "${key}"
        fi
        continue
    fi
    dest="${DEST_FOR_KEY[${key}]}"
    mkdir -p "${dest}"
    if [ "${FALLBACK_FOR_KEY[${key}]:-}" = "true" ]; then
        # Discard partial outputs from failed attempts and clear _failed so
        # the verified fallback pair can preserve its last-known scan state.
        find "${dest}" -type f \
            \( -name "component-architecture.json" -o -name "security-findings.json" -o -name "_failed" -o -name "_scan-provenance.json" \) \
            -delete 2>/dev/null || true
        RESTORED_FALLBACKS+=("${key}")
    fi
    cp "${src}"/*.json "${dest}/"
    carried=$((carried + 1))
done

if [ ${#UNRESTORED_KEYS[@]} -gt 0 ]; then
    # A repo marked as skipped must be restorable: publishing without it
    # while advancing the scan state would strand it permanently.
    echo "::error::${#UNRESTORED_KEYS[@]} skipped repo(s) have no restorable snapshot data: ${UNRESTORED_KEYS[*]}"
    if [ -n "${GITHUB_STEP_SUMMARY:-}" ]; then
        {
            echo "### Skipped repositories without restorable snapshot data"
            printf -- '- %s\n' "${UNRESTORED_KEYS[@]}"
        } >> "${GITHUB_STEP_SUMMARY}"
    fi
    for key in "${UNRESTORED_KEYS[@]}"; do mark_failed "$key"; done
    exit 1
fi

report_unavailable_exclusions
if [ ${#RESTORED_FALLBACKS[@]} -gt 0 ]; then
    echo "::warning::reused last verified data for ${#RESTORED_FALLBACKS[@]} failed/unavailable repo(s): ${RESTORED_FALLBACKS[*]}"
    if [ -n "${GITHUB_STEP_SUMMARY:-}" ]; then
        {
            echo "### Repositories served from last verified data"
            printf -- '- %s\n' "${RESTORED_FALLBACKS[@]}"
        } >> "${GITHUB_STEP_SUMMARY}"
    fi
fi

echo "[*] Carried forward ${carried}/${#MISSING_KEYS[@]} repo(s)"
