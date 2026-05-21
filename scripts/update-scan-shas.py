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
        for root, _, files in os.walk(d):
            if "component-architecture.json" in files:
                yield os.path.join(root, "component-architecture.json")


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

    for path in find_arch_jsons(dirs):
        with open(path) as f:
            data = json.load(f)
        repo = data.get("repo", "")
        sha = data.get("commit_sha", "")
        ver = data.get("analyzer_version", "")
        if repo and sha:
            existing[repo] = {"commit_sha": sha, "analyzer_version": ver}

    with open(output, "w") as f:
        json.dump(existing, f, indent=2, sort_keys=True)
        f.write("\n")

    print(f"Updated {output}: {len(existing)} repos tracked")


if __name__ == "__main__":
    main()
