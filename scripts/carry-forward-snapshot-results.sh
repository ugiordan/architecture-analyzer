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
FAILED_KEYS=()
for results_dir in "${RESULTS_DIRS[@]}"; do
    platform="$(basename "${results_dir}" | sed 's/-results$//')"
    while IFS= read -r repo; do
        [ -n "$repo" ] || continue
        short="${repo##*/}"
        key="${platform}/${short}"
        if [ -f "${results_dir}/${repo}/component-architecture.json" ]; then
            # Partial results from a failed analysis do not count: the
            # run may be missing security-findings.json and must not be
            # published or recorded as current.
            if [ -f "${results_dir}/${repo}/_failed" ]; then
                FAILED_KEYS+=("${key}")
            fi
            continue
        fi
        if [ ! -f "${results_dir}/${repo}/_skipped" ]; then
            FAILED_KEYS+=("${key}")
            continue
        fi
        if [ -z "${DEST_FOR_KEY[${key}]:-}" ]; then
            MISSING_KEYS+=("${key}")
        fi
        DEST_FOR_KEY["${key}"]="${results_dir}/${repo}"
        REPO_FOR_KEY["${key}"]="${repo}"
    done < <(yq -r ".platforms.${platform}.orgs | to_entries[] | .key as \$org | .value[] | \"\(\$org)/\(.)\"" "${SCAN_CONFIG}" 2>/dev/null)
done

if [ ${#FAILED_KEYS[@]} -gt 0 ]; then
    echo "::error::${#FAILED_KEYS[@]} repo(s) have no results and were not marked skipped (analysis failed); their previous data is NOT restored: ${FAILED_KEYS[*]}"
    # Stop publication: aggregating would emit a snapshot permanently
    # missing these repos, and the scan state must not advance past them.
    exit 1
fi

if [ ${#MISSING_KEYS[@]} -eq 0 ]; then
    echo "[*] No missing repos to carry forward"
    exit 0
fi

echo "[*] Carrying forward ${#MISSING_KEYS[@]} repo(s) from the latest snapshot"

# Sparse-clone only the raw directories of the missing repos from the
# snapshots branch (blobless clone keeps this cheap).
SNAP_DIR="$(mktemp -d)"
trap 'rm -rf "${SNAP_DIR}"' EXIT
SNAP_URL="$(git -C "${ANALYZER_DIR}" remote get-url origin 2>/dev/null || echo "")"
if [ -z "${SNAP_URL}" ]; then
    echo "::error::cannot determine origin URL; cannot carry forward ${#MISSING_KEYS[@]} marked repo(s)"
    exit 1
fi
if ! git clone --depth 1 --branch snapshots --filter=blob:none --sparse \
        "${SNAP_URL}" "${SNAP_DIR}" 2>/dev/null; then
    echo "::error::snapshots branch unavailable; cannot carry forward ${#MISSING_KEYS[@]} marked repo(s)"
    exit 1
fi

# Restrict both the sparse patterns and the source lookup to each repo's
# own platform so identical names across platforms cannot cross-contaminate.
PATTERNS=()
for key in "${MISSING_KEYS[@]}"; do
    PATTERNS+=("output/*/${key%%/*}/*/raw/${key##*/}/")
done
(
    cd "${SNAP_DIR}"
    git sparse-checkout set --no-cone "${PATTERNS[@]}" 2>/dev/null || true
)

# No snapshot data at all is fatal when repos must be restored: publishing
# without them while advancing the scan state would strand them.
if [ ! -d "${SNAP_DIR}/output" ]; then
    echo "::error::no snapshot output directory; cannot carry forward ${#MISSING_KEYS[@]} marked repo(s)"
    exit 1
fi

carried=0
UNRESTORED_KEYS=()
STATE_FILE="${ANALYZER_DIR}/.last-scan-shas.json"
for key in "${MISSING_KEYS[@]}"; do
    platform="${key%%/*}"
    short="${key##*/}"
    # Newest dated snapshot carrying this repo for its platform. Custom
    # version labels are excluded so they cannot override dated snapshots,
    # and future-dated labels are excluded too: a manual run with a
    # future date would otherwise always outrank valid current data.
    today="$(date -u +%Y-%m-%d)"
    src="$(find "${SNAP_DIR}/output" -type d -path "*/${platform}/*/raw/${short}" 2>/dev/null \
        | grep -E "/output/[0-9]{4}-[0-9]{2}-[0-9]{2}/" \
        | awk -v today="$today" -v prefix="${SNAP_DIR}/output/" \
            'index($0, prefix) == 1 {
                 date = substr($0, length(prefix) + 1, 10)
                 if (date <= today) print
             }' \
        | sort | tail -1 || true)"
    if [ -z "${src}" ] || [ -z "$(ls "${src}"/*.json 2>/dev/null)" ]; then
        UNRESTORED_KEYS+=("${key}")
        continue
    fi
    # Provenance: the restored data must match the scan-state entry that
    # justified the skip (same commit and analyzer version). A mismatch
    # means the snapshot holds different data than the state describes,
    # for example a newer run that published but failed to commit state;
    # restoring it would publish stale data, so stop instead.
    if [ -f "${STATE_FILE}" ]; then
        state_key="${platform}/${REPO_FOR_KEY[${key}]}"
        MATCHES=$(python3 -c "
import json, sys
state = json.load(open(sys.argv[1])).get(sys.argv[3], {})
data = json.load(open(sys.argv[2]))
ok = state.get('commit_sha') and \
    state.get('commit_sha') == data.get('commit_sha') and \
    state.get('analyzer_version') == data.get('analyzer_version')
print('yes' if ok else 'no')
" "${STATE_FILE}" "${src}/component-architecture.json" "${state_key}")
        if [ "${MATCHES}" != "yes" ]; then
            echo "::warning::${key}: snapshot data does not match the scan state entry"
            UNRESTORED_KEYS+=("${key} (provenance mismatch)")
            continue
        fi
    fi
    dest="${DEST_FOR_KEY[${key}]}"
    mkdir -p "${dest}"
    cp "${src}"/*.json "${dest}/"
    carried=$((carried + 1))
done

if [ ${#UNRESTORED_KEYS[@]} -gt 0 ]; then
    # A repo marked as skipped must be restorable: publishing without it
    # while advancing the scan state would strand it permanently.
    echo "::error::${#UNRESTORED_KEYS[@]} skipped repo(s) have no restorable snapshot data: ${UNRESTORED_KEYS[*]}"
    exit 1
fi

echo "[*] Carried forward ${carried}/${#MISSING_KEYS[@]} repo(s)"
