#!/usr/bin/env python3
"""Merge a run's scan SHAs into the current state on main.

Usage: merge-scan-shas.py <ours.json> <main.json> <output.json> [run-invalidations.json main-invalidations.json output-invalidations.json]

Concurrent runs can finish out of order; an older run must not replace a
newer run's entries. Per repo, the entry with the newer scanned_at wins;
entries only present on one side are kept.
"""
import json
import sys
from datetime import datetime, timezone


def parse_scan_time(value):
    if not isinstance(value, str) or not value:
        return None
    try:
        parsed = datetime.fromisoformat(value.replace("Z", "+00:00"))
    except ValueError:
        return None
    if parsed.tzinfo is None:
        parsed = parsed.replace(tzinfo=timezone.utc)
    return parsed.astimezone(timezone.utc)


def entry_time(entry):
    return parse_scan_time(entry.get("scanned_at")) or datetime.min.replace(tzinfo=timezone.utc)


def run_key(value):
    try:
        number = int(value.get("workflow_run_number", 0))
        attempt = int(value.get("workflow_run_attempt", 0))
    except (AttributeError, TypeError, ValueError):
        return None
    return (number, attempt) if number > 0 and attempt > 0 else None


def entry_rank(entry):
    run = run_key(entry)
    return (
        1 if run is not None else 0,
        run[0] if run is not None else 0,
        run[1] if run is not None else 0,
        entry_time(entry),
        str(entry.get("analyzer_version") or ""),
        str(entry.get("commit_sha") or ""),
    )


def incoming_wins(incoming, current):
    return entry_rank(incoming) > entry_rank(current)


def key_matches_prefix(key, prefix):
    if prefix.endswith("/"):
        return key.startswith(prefix)
    return key == prefix or key.startswith(prefix + "/")


def load_invalidation_events(path, *, required=False):
    if not path:
        return []
    try:
        with open(path) as f:
            return json.load(f).get("invalidations", [])
    except FileNotFoundError:
        if required:
            raise
        return []


def merge_invalidation_events(main_events, run_events):
    latest_by_run = {}
    latest_by_time = {}
    latest_by_timestamp_only = {}
    for event in [*main_events, *run_events]:
        prefix = event.get("prefix")
        if not isinstance(prefix, str) or parse_scan_time(event.get("invalidated_at")) is None:
            continue
        event_run = run_key(event)
        incoming_time = parse_scan_time(event["invalidated_at"])
        if event_run is not None:
            current = latest_by_run.get(prefix)
            current_run = run_key(current) if current else None
            current_time = parse_scan_time(current.get("invalidated_at")) if current else None
            if (
                current is None
                or event_run > current_run
                or (event_run == current_run and (current_time is None or incoming_time >= current_time))
            ):
                latest_by_run[prefix] = event

        else:
            current = latest_by_timestamp_only.get(prefix)
            current_time = parse_scan_time(current.get("invalidated_at")) if current else None
            if current is None or current_time is None or incoming_time >= current_time:
                latest_by_timestamp_only[prefix] = event

        current = latest_by_time.get(prefix)
        current_time = parse_scan_time(current.get("invalidated_at")) if current else None
        current_run = run_key(current) if current else None
        if (
            current is None
            or incoming_time > current_time
            or (
                incoming_time == current_time
                and event_run is not None
                and (current_run is None or event_run > current_run)
            )
        ):
            latest_by_time[prefix] = event

    merged = []
    prefixes = set(latest_by_run) | set(latest_by_time) | set(latest_by_timestamp_only)
    for prefix in sorted(prefixes):
        for event in (
            latest_by_run.get(prefix),
            latest_by_timestamp_only.get(prefix),
            latest_by_time.get(prefix),
        ):
            if event is not None and event not in merged:
                merged.append(event)
    return merged


def apply_invalidations(merged, events):
    removed = 0
    for event in events:
        prefix = event["prefix"]
        invalidated_at = parse_scan_time(event.get("invalidated_at"))
        event_run = run_key(event)
        for key, entry in list(merged.items()):
            if not key_matches_prefix(key, prefix):
                continue
            current_run = run_key(entry)
            if event_run is not None and current_run is not None and event_run != current_run:
                should_remove = current_run < event_run
            else:
                should_remove = entry_time(entry) <= invalidated_at
            if should_remove:
                del merged[key]
                removed += 1
    return removed


def main():
    if len(sys.argv) not in (4, 7):
        print(
            f"Usage: {sys.argv[0]} <ours.json> <main.json> <output.json> "
            "[run-invalidations.json main-invalidations.json output-invalidations.json]",
            file=sys.stderr,
        )
        sys.exit(1)

    ours_path, main_path, out_path = sys.argv[1:4]
    run_invalidations_path = sys.argv[4] if len(sys.argv) == 7 else None
    main_invalidations_path = sys.argv[5] if len(sys.argv) == 7 else None
    out_invalidations_path = sys.argv[6] if len(sys.argv) == 7 else None

    with open(ours_path) as f:
        ours = json.load(f)
    try:
        with open(main_path) as f:
            merged = json.load(f)
    except FileNotFoundError:
        merged = {}

    for key in [key for key in merged if key.count("/") == 1]:
        del merged[key]

    try:
        invalidations = merge_invalidation_events(
            load_invalidation_events(main_invalidations_path),
            load_invalidation_events(run_invalidations_path, required=run_invalidations_path is not None),
        )
    except FileNotFoundError as exc:
        print(f"Required run invalidation file not found: {exc.filename}", file=sys.stderr)
        sys.exit(1)
    removed = apply_invalidations(merged, invalidations)
    removed += apply_invalidations(ours, invalidations)

    for repo, entry in ours.items():
        current = merged.get(repo)
        if current is None or incoming_wins(entry, current):
            merged[repo] = entry

    with open(out_path, "w") as f:
        json.dump(merged, f, indent=2, sort_keys=True)
        f.write("\n")

    if out_invalidations_path:
        with open(out_invalidations_path, "w") as f:
            json.dump({"invalidations": invalidations}, f, indent=2, sort_keys=True)
            f.write("\n")

    print(f"Merged {len(ours)} entries into {len(merged)} tracked repos; applied {removed} invalidations")


if __name__ == "__main__":
    main()
