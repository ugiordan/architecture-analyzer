#!/bin/bash
# Usage: analyze-repo.sh <org/repo> <results-base-dir> [version-label]
#
# Clones a repository, runs the architecture analyzer, and cleans up.
# When version-label is provided, output goes to <results-base>/<repo>/<version>/.
set -euo pipefail

# Security hardening for Go toolchain on untrusted repos
export CGO_ENABLED=0
export GOFLAGS="-mod=readonly"
export GONOSUMCHECK=""
export GONOSUMDB=""
export GOMAXPROCS=2

# Memory limit for the analyzer binary (prevents OOM kills on CI runners)
# GitHub-hosted runners have 7GB RAM. Cap Go heap at 5.5GB to leave room for OS + other processes.
export GOMEMLIMIT="${GOMEMLIMIT:-5500MiB}"

REPO="${1:?Usage: analyze-repo.sh <org/repo> <results-base-dir> [version-label]}"
RESULTS_BASE="${2:?Usage: analyze-repo.sh <org/repo> <results-base-dir> [version-label]}"
VERSION_LABEL="${3:-}"
SCAN_PLATFORM="${SCAN_PLATFORM:-}"

# SIGTERM/SIGINT trap: when GitHub Actions timeout kills the process,
# clean up and exit 0 so the job shows a warning, not a red error.
trap 'echo "::warning::${REPO}: terminated (timeout or signal)"; \
      chmod -R u+w "${CLONE_DIR:-/dev/null}" 2>/dev/null || true; \
      rm -rf "${CLONE_DIR:-/dev/null}" 2>/dev/null || true; exit 0' TERM INT

# Validate repo name to prevent command injection
if [[ ! "$REPO" =~ ^[a-zA-Z0-9_.-]+/[a-zA-Z0-9_.-]+$ ]]; then
    echo "[!] Invalid repo format: $REPO (expected: org/repo)" >&2
    exit 1
fi
if [ -n "${VERSION_LABEL}" ] && [[ ! "${VERSION_LABEL}" =~ ^[a-zA-Z0-9][a-zA-Z0-9._-]{0,63}$ ]]; then
    echo "::error::Invalid version label: ${VERSION_LABEL}" >&2
    exit 1
fi

ORG="${REPO%%/*}"
SHORT="${REPO##*/}"
OUTDIR="${RESULTS_BASE}/${ORG}/${SHORT}"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ANALYZER_DIR="$(dirname "${SCRIPT_DIR}")"

# mark_skipped records that this repo intentionally produced no results
# while remaining verifiably unchanged or frozen (unchanged skip, or
# archived), so the aggregate job may restore the previous snapshot for
# it. Clone failures and symlink-rejected repos are NOT marked: their
# data is unverified and must not be published as current; a repo with
# no results and no marker counts as a failed analysis.
mark_skipped() {
    mkdir -p "${OUTDIR}"
    touch "${OUTDIR}/_skipped"
}

# mark_unavailable records a repo that could not be cloned or was rejected
# by the symlink boundary. Carry-forward may use its last verified snapshot,
# but does not treat it as an analysis failure for the whole platform.
mark_unavailable() {
    mkdir -p "${OUTDIR}"
    touch "${OUTDIR}/_unavailable"
}

# Binary location: CI downloads pre-built binary, local dev builds on demand
if [ -x "${GITHUB_WORKSPACE:-}/arch-analyzer" ]; then
    ANALYZER_BIN="${GITHUB_WORKSPACE}/arch-analyzer"
elif [ -x "${ANALYZER_DIR}/arch-analyzer" ]; then
    ANALYZER_BIN="${ANALYZER_DIR}/arch-analyzer"
else
    echo "[*] Building arch-analyzer..."
    (cd "${ANALYZER_DIR}" && go build -o arch-analyzer ./cmd/arch-analyzer/)
    ANALYZER_BIN="${ANALYZER_DIR}/arch-analyzer"
fi

REPO_URL="https://github.com/${REPO}.git"
SHAS_FILE="${ANALYZER_DIR}/.last-scan-shas.json"

# ODH forks are scanned from their main/master branch, not inferred upstream
# module paths. Repositories without either branch can set a repo_branches
# override in scan-config.yaml.
SOURCE_BRANCH=""
if [ "${SCAN_PLATFORM}" = "odh" ] && [ "${ORG}" = "opendatahub-io" ]; then
    if command -v yq &>/dev/null; then
        SOURCE_BRANCH=$(yq -r ".platforms.odh.repo_branches.\"${REPO}\" // \"\"" "${ANALYZER_DIR}/scan-config.yaml" 2>/dev/null || true)
    fi
    if [ -z "${SOURCE_BRANCH}" ]; then
        DEFAULT_BRANCH=$(git ls-remote --symref "${REPO_URL}" HEAD 2>/dev/null | awk '$1 == "ref:" && $2 ~ /^refs\/heads\// { sub("refs/heads/", "", $2); print $2; exit }')
        if [ "${DEFAULT_BRANCH}" = "main" ] || [ "${DEFAULT_BRANCH}" = "master" ]; then
            SOURCE_BRANCH="${DEFAULT_BRANCH}"
        elif git ls-remote --exit-code --heads "${REPO_URL}" refs/heads/main >/dev/null 2>&1; then
            SOURCE_BRANCH="main"
        elif git ls-remote --exit-code --heads "${REPO_URL}" refs/heads/master >/dev/null 2>&1; then
            SOURCE_BRANCH="master"
        else
            echo "::error::${REPO} has no main/master branch and no repo_branches override" >&2
            exit 1
        fi
    fi
fi

# Skip archived repos whose HEAD is unchanged since the last scan and
# whose results match the current analyzer version. An archived repo that
# moved since the last scan is scanned once more so its final state is
# captured, and FORCE_RESCAN bypasses this skip too.
if command -v gh &>/dev/null; then
    if gh api "repos/${REPO}" --jq '.archived' 2>/dev/null | grep -q "true"; then
        if [ -n "${SOURCE_BRANCH}" ]; then
            ARCHIVED_SHA=$(git ls-remote "${REPO_URL}" "refs/heads/${SOURCE_BRANCH}" 2>/dev/null | awk '{print $1}')
        else
            ARCHIVED_SHA=$(git ls-remote "${REPO_URL}" HEAD 2>/dev/null | awk '{print $1}')
        fi
        STATE_KEY="${SCAN_PLATFORM}/${REPO}"
        PREV_SHA=""
        PREV_VER=""
        if [ -f "${SHAS_FILE}" ]; then
            PREV_SHA=$(python3 -c "import json,sys; d=json.load(open(sys.argv[1])); print(d.get(sys.argv[2],{}).get('commit_sha',''))" "${SHAS_FILE}" "${STATE_KEY}" 2>/dev/null || echo "")
            PREV_VER=$(python3 -c "import json,sys; d=json.load(open(sys.argv[1])); print(d.get(sys.argv[2],{}).get('analyzer_version',''))" "${SHAS_FILE}" "${STATE_KEY}" 2>/dev/null || echo "")
        fi
        ANALYZER_VER=$("${ANALYZER_BIN}" version 2>/dev/null | awk '{print $2}' || echo "")
        if [ "${FORCE_RESCAN:-}" != "1" ] && [ -n "${ARCHIVED_SHA}" ] && [ "${ARCHIVED_SHA}" = "${PREV_SHA}" ] \
            && [ -n "${PREV_VER}" ] && [ "${ANALYZER_VER}" = "${PREV_VER}" ]; then
            echo "::warning::Skipping ${REPO}: repository is archived and unchanged since the last scan"
            mark_skipped
            exit 0
        fi
        echo "[*] ${REPO} is archived but changed since the last scan (or the analyzer changed); scanning its final state"
    fi
fi

mkdir -p "${OUTDIR}"

CLONE_BASE="${RUNNER_TEMP:-/tmp}/arch-analyzer-repos"
CLONE_DIR="${CLONE_BASE}/${SHORT}"

export GOTMPDIR="${RUNNER_TEMP:-/tmp}/gotmp"
mkdir -p "${GOTMPDIR}"

CLONE_ARGS=(--depth 1)
if [ -n "${SOURCE_BRANCH}" ]; then
    CLONE_ARGS+=(--branch "${SOURCE_BRANCH}")
    echo "[*] Cloning ${REPO} at ${SOURCE_BRANCH}..."
else
    echo "[*] Cloning ${REPO} at its default branch..."
fi
rm -rf "${CLONE_DIR}"
git clone "${CLONE_ARGS[@]}" "${REPO_URL}" "${CLONE_DIR}" 2>/dev/null || {
    echo "::warning::Skipping ${REPO}: clone failed (repo may be private or inaccessible)"
    mark_unavailable
    exit 0
}

# Symlink boundary check helper
check_symlinks() {
    local resolved_clone
    resolved_clone=$(readlink -f "${CLONE_DIR}" 2>/dev/null || echo "${CLONE_DIR}")
    local escaped
    escaped=$(find "${CLONE_DIR}" -type l -exec readlink -f {} \; 2>/dev/null | grep -v "^${resolved_clone}" || true)
    if [ -n "$escaped" ]; then
        echo "::warning::Skipping ${REPO}: symlinks escape clone boundary"
        rm -rf "${CLONE_DIR}"
        mark_unavailable
        exit 0
    fi
}

# Check symlinks after clone
check_symlinks

# Skip analysis if repo hasn't changed since last scan and analyzer version is the same.
# Uses .last-scan-shas.json (committed to repo) to track previous scan state.
# Set FORCE_RESCAN=1 to bypass this check.
CLONE_SHA=$(git -C "${CLONE_DIR}" rev-parse HEAD 2>/dev/null || echo "")
if [ -n "${CLONE_SHA}" ] && [ -f "${SHAS_FILE}" ] && [ "${FORCE_RESCAN:-}" != "1" ]; then
    STATE_KEY="${SCAN_PLATFORM}/${REPO}"
    PREV_SHA=$(python3 -c "import json,sys; d=json.load(open(sys.argv[1])); print(d.get(sys.argv[2],{}).get('commit_sha',''))" "${SHAS_FILE}" "${STATE_KEY}" 2>/dev/null || echo "")
    PREV_VER=$(python3 -c "import json,sys; d=json.load(open(sys.argv[1])); print(d.get(sys.argv[2],{}).get('analyzer_version',''))" "${SHAS_FILE}" "${STATE_KEY}" 2>/dev/null || echo "")
    ANALYZER_VER=$("${ANALYZER_BIN}" version 2>/dev/null | awk '{print $2}' || echo "")
    if [ "${CLONE_SHA}" = "${PREV_SHA}" ] && [ -n "${PREV_VER}" ] && [ "${ANALYZER_VER}" = "${PREV_VER}" ]; then
        echo "[*] Skipping ${SHORT}: HEAD ${CLONE_SHA:0:8} unchanged, analyzer ${ANALYZER_VER} unchanged"
        chmod -R u+w "${CLONE_DIR}" 2>/dev/null || true
        rm -rf "${CLONE_DIR}"
        mark_skipped
        exit 0
    fi
    echo "[*] Change detected for ${SHORT}: HEAD ${PREV_SHA:0:8} -> ${CLONE_SHA:0:8}, analyzer ${PREV_VER} -> ${ANALYZER_VER}"
fi

# Download Go dependencies for go/packages analysis (isolated cache)
if [ -f "${CLONE_DIR}/go.mod" ]; then
    echo "[*] Downloading Go modules for ${SHORT}..."
    (
        export GOMODCACHE="${CLONE_DIR}/.gomod-cache"
        export GOCACHE="${CLONE_DIR}/.gobuild-cache"
        export GOFLAGS="-modcacherw"
        export GONOSUMCHECK=""
        export GONOSUMDB=""
        cd "${CLONE_DIR}" && timeout 120 go mod download 2>&1
    ) || {
        echo "::warning::go mod download failed for ${REPO}, Go AST extraction will use fallback"
    }
    # Re-check symlinks after go mod download (TOCTOU mitigation)
    check_symlinks
fi

# Resolve aliases from scan-config.yaml (if present)
ALIASES_ARGS=""
SCAN_CONFIG="${ANALYZER_DIR}/scan-config.yaml"
if [ -f "${SCAN_CONFIG}" ] && command -v yq &>/dev/null; then
    # Search all platforms for this repo's aliases
    ALIASES=$(yq -r "
        .platforms[].repo_overrides.\"${SHORT}\".aliases // [] | join(\",\")
    " "${SCAN_CONFIG}" 2>/dev/null | grep -v '^$' | head -1 || true)
    if [ -n "${ALIASES}" ]; then
        ALIASES_ARGS="-aliases"
        ALIASES_VAL="${ALIASES}"
    fi
fi

# Full analysis (architecture + optional code graph)
echo "[*] Analyzing ${SHORT}..."
VERSION_ARGS=""
if [ -n "${VERSION_LABEL}" ]; then
    VERSION_ARGS="-version ${VERSION_LABEL}"
fi
ANALYSIS_TIMEOUT="${ANALYSIS_TIMEOUT:-2400}"
FULL_ARGS=(-output-dir "${OUTDIR}" -org "${ORG}")
if [ -n "${VERSION_ARGS}" ]; then
    read -r -a _version_args <<< "${VERSION_ARGS}"
    FULL_ARGS+=("${_version_args[@]}")
fi
if [ -n "${ALIASES_ARGS}" ]; then
    FULL_ARGS+=("${ALIASES_ARGS}" "${ALIASES_VAL}")
fi
FULL_ARGS+=("${CLONE_DIR}")
CLEANUP_DIR="${OUTDIR}"
if [ -n "${VERSION_LABEL}" ]; then
    CLEANUP_DIR="${OUTDIR}/${VERSION_LABEL}"
fi
if [ -d "${CLEANUP_DIR}" ]; then
    find "${CLEANUP_DIR}" -maxdepth 1 -type f \
        \( -name "component-architecture.json" -o -name "security-findings.json" -o -name "_failed" \) \
        -delete 2>/dev/null || true
fi
rm -f "${OUTDIR}/_failed"
python3 -c 'import json,os; from datetime import datetime,timezone; print(json.dumps({"timestamp":datetime.now(timezone.utc).isoformat(timespec="microseconds").replace("+00:00","Z"),"workflow_run_number":int(os.getenv("GITHUB_RUN_NUMBER","0")),"workflow_run_attempt":int(os.getenv("GITHUB_RUN_ATTEMPT","0"))}))' \
    > "${OUTDIR}/_scan-provenance.json" || echo "::warning::Could not write scan provenance for ${REPO}"

timeout "${ANALYSIS_TIMEOUT}" "${ANALYZER_BIN}" full-analysis "${FULL_ARGS[@]}" || {
    rc=$?
    if [ "$rc" -eq 124 ]; then
        echo "::warning::Analysis timed out for ${REPO} after ${ANALYSIS_TIMEOUT}s, partial results may be available"
    else
        echo "::warning::Analysis failed for ${REPO} (exit ${rc}), partial results may be available"
    fi
    # Mark the run as failed so partial results are neither published as
    # current nor recorded in the scan SHAs; the next run rescans. Exit
    # with the analyzer's status so the workflow retry loop runs.
    python3 -c 'import json,os; from datetime import datetime,timezone; print(json.dumps({"timestamp":datetime.now(timezone.utc).isoformat(timespec="microseconds").replace("+00:00","Z"),"workflow_run_number":int(os.getenv("GITHUB_RUN_NUMBER","0")),"workflow_run_attempt":int(os.getenv("GITHUB_RUN_ATTEMPT","0"))}))' \
        > "${OUTDIR}/_failed" || touch "${OUTDIR}/_failed"
    if [ "$rc" -eq 124 ]; then
        # ANALYSIS_TIMEOUT is 40 minutes while the matrix job is capped at
        # 60 minutes; a retry cannot finish and only consumes the remaining
        # runner time. Keep the failed marker and stop retries cleanly.
        exit 0
    fi
    exit "$rc"
}

# Cleanup (chmod needed because Go module cache sets files read-only)
chmod -R u+w "${CLONE_DIR}" 2>/dev/null || true
rm -rf "${CLONE_DIR}"
echo "[*] Done: ${OUTDIR}"
