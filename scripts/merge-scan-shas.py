#!/usr/bin/env python3
"""Merge a run's scan SHAs into the current state on main.

Usage: merge-scan-shas.py <ours.json> <main.json> <output.json>

Concurrent runs can finish out of order; an older run must not replace a
newer run's entries. Per repo, the entry with the newer scanned_at wins;
entries only present on one side are kept.
"""
import json
import sys


def main():
    if len(sys.argv) != 4:
        print(f"Usage: {sys.argv[0]} <ours.json> <main.json> <output.json>", file=sys.stderr)
        sys.exit(1)

    ours_path, main_path, out_path = sys.argv[1:4]

    with open(ours_path) as f:
        ours = json.load(f)
    try:
        with open(main_path) as f:
            merged = json.load(f)
    except FileNotFoundError:
        merged = {}

    for repo, entry in ours.items():
        current = merged.get(repo)
        if current is None or entry.get("scanned_at", "") >= current.get("scanned_at", ""):
            merged[repo] = entry

    with open(out_path, "w") as f:
        json.dump(merged, f, indent=2, sort_keys=True)
        f.write("\n")

    print(f"Merged {len(ours)} entries into {len(merged)} tracked repos")


if __name__ == "__main__":
    main()
