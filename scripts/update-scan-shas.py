#!/usr/bin/env python3
"""Build .last-scan-shas.json from component-architecture.json files.

Usage: update-scan-shas.py <results-dir> [<results-dir2> ...] -o <output-path>

Scans result directories for component-architecture.json files and extracts
commit_sha + analyzer_version from each. Merges with any existing SHAs file
so repos not in this run keep their previous entry.
"""
import json
import os
import sys


def find_arch_jsons(dirs):
    for d in dirs:
        # Results directories are named "<platform>-results"; the scan
        # state is keyed by platform and repo so a repository scanned
        # under multiple platforms keeps separate entries.
        platform = os.path.basename(d.rstrip("/")).removesuffix("-results")
        for root, _, files in os.walk(d):
            if "component-architecture.json" in files:
                # Partial results from a failed analysis (marked by
                # analyze-repo.sh) must not advance the skip state. The
                # marker sits beside the result, or in the parent when
                # full-analysis wrote a versioned subdirectory.
                if "_failed" in files or os.path.exists(
                    os.path.join(os.path.dirname(root), "_failed")
                ):
                    continue
                yield platform, os.path.join(root, "component-architecture.json")


def main():
    if "-o" not in sys.argv:
        print(f"Usage: {sys.argv[0]} <dir> [<dir>...] -o <output>", file=sys.stderr)
        sys.exit(1)

    out_idx = sys.argv.index("-o")
    dirs = sys.argv[1:out_idx]
    output = sys.argv[out_idx + 1]

    existing = {}
    if os.path.exists(output):
        with open(output) as f:
            existing = json.load(f)

    for platform, path in find_arch_jsons(dirs):
        with open(path) as f:
            data = json.load(f)
        repo = data.get("repo", "")
        sha = data.get("commit_sha", "")
        ver = data.get("analyzer_version", "")
        # extracted_at (UTC RFC3339) records when the data was produced;
        # the state commit merges concurrent runs by keeping the newer
        # entry per repo.
        scanned_at = data.get("extracted_at", "")
        if repo and sha:
            existing[f"{platform}/{repo}"] = {
                "commit_sha": sha,
                "analyzer_version": ver,
                "scanned_at": scanned_at,
            }

    with open(output, "w") as f:
        json.dump(existing, f, indent=2, sort_keys=True)
        f.write("\n")

    print(f"Updated {output}: {len(existing)} repos tracked")


if __name__ == "__main__":
    main()
