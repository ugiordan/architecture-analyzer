#!/usr/bin/env python3
"""Stamp and compare GitHub Actions run provenance in snapshot metadata."""

import json
import re
import sys
from datetime import datetime, timezone

VERSION_LABEL_RE = re.compile(r"[a-zA-Z0-9][a-zA-Z0-9._-]{0,63}")


def validate_version_label(version):
    if not isinstance(version, str) or not VERSION_LABEL_RE.fullmatch(version):
        raise ValueError("invalid version label: must match [a-zA-Z0-9][a-zA-Z0-9._-]{0,63}")


def stamp(path, run_number, run_attempt):
    """Record the workflow run and attempt that produced a snapshot."""
    with open(path) as f:
        metadata = json.load(f)
    metadata["workflow_run_number"] = int(run_number)
    metadata["workflow_run_attempt"] = int(run_attempt)
    with open(path, "w") as f:
        json.dump(metadata, f, indent=2, sort_keys=True)
        f.write("\n")


def write_removal_marker(path, version, platform, run_number, run_attempt):
    """Write tombstone metadata for a platform intentionally removed."""
    validate_version_label(version)
    metadata = {
        "version": version,
        "timestamp": datetime.now(timezone.utc).strftime("%Y-%m-%dT%H:%M:%SZ"),
        "platform": platform,
        "repos_analyzed": {},
        "removed": True,
        "workflow_run_number": int(run_number),
        "workflow_run_attempt": int(run_attempt),
    }
    with open(path, "w") as f:
        json.dump(metadata, f, indent=2, sort_keys=True)
        f.write("\n")


def is_newer(path, run_number, run_attempt):
    """Return true when metadata belongs to a later workflow run/attempt.

    Metadata without run fields is from before run provenance was added and
    is treated as older so the first run using this mechanism can replace it.
    """
    with open(path) as f:
        metadata = json.load(f)
    existing = (
        int(metadata.get("workflow_run_number", 0)),
        int(metadata.get("workflow_run_attempt", 0)),
    )
    current = (int(run_number), int(run_attempt))
    return existing > current


def main():
    action = sys.argv[1] if len(sys.argv) > 1 else ""
    valid_args = (action in {"stamp", "is-newer"} and len(sys.argv) == 5) or (
        action == "removed" and len(sys.argv) == 7
    ) or (
        action == "validate-version" and len(sys.argv) == 3
    )
    if not valid_args:
        print(
            f"Usage: {sys.argv[0]} stamp <metadata.json> <run-number> <run-attempt> | "
            f"is-newer <metadata.json> <run-number> <run-attempt> | "
            f"removed <metadata.json> <version> <platform> <run-number> <run-attempt> | "
            f"validate-version <version>",
            file=sys.stderr,
        )
        sys.exit(2)

    if action == "validate-version":
        try:
            validate_version_label(sys.argv[2])
        except ValueError as exc:
            print(str(exc), file=sys.stderr)
            sys.exit(1)
        return

    if action == "removed":
        path, version, platform, run_number, run_attempt = sys.argv[2:]
        try:
            write_removal_marker(path, version, platform, run_number, run_attempt)
        except ValueError as exc:
            print(str(exc), file=sys.stderr)
            sys.exit(1)
        return

    path, run_number, run_attempt = sys.argv[2:]
    if action == "stamp":
        stamp(path, run_number, run_attempt)
    else:
        print("yes" if is_newer(path, run_number, run_attempt) else "no")


if __name__ == "__main__":
    main()
