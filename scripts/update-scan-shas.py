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
from datetime import datetime, timezone


def find_arch_jsons(dirs, excluded_platforms=None):
    excluded_platforms = excluded_platforms or set()
    for d in dirs:
        # Results directories are named "<platform>-results"; the scan
        # state is keyed by platform and repo so a repository scanned
        # under multiple platforms keeps separate entries.
        platform = os.path.basename(d.rstrip("/")).removesuffix("-results")
        if platform in excluded_platforms:
            continue
        for root, _, files in os.walk(d):
            if "component-architecture.json" in files:
                # Both raw artifacts must exist for an analysis to count as
                # successful. A timeout can write architecture JSON before
                # failing to write security findings. Failed markers may be
                # beside versioned results or in an ancestor result dir.
                architecture_path = os.path.join(root, "component-architecture.json")
                security_path = os.path.join(root, "security-findings.json")
                if (
                    not valid_architecture(architecture_path)
                    or not valid_security_findings(security_path)
                ):
                    continue
                yield platform, architecture_path, result_run_metadata(root, d)


def load_json(path):
    try:
        with open(path) as f:
            return json.load(f)
    except (OSError, json.JSONDecodeError, UnicodeDecodeError):
        return None


def valid_architecture_data(data):
    return (
        isinstance(data, dict)
        and isinstance(data.get("repo"), str)
        and bool(data.get("repo"))
        and isinstance(data.get("commit_sha"), str)
        and bool(data.get("commit_sha"))
    )


def valid_architecture(path):
    """Require the identity/provenance fields needed to advance scan state."""
    return valid_architecture_data(load_json(path))


def valid_security_findings(path):
    """Security findings must be a valid JSON array, including an empty one."""
    return isinstance(load_json(path), list)


def parse_scan_time(value):
    """Parse RFC3339 timestamps to UTC datetimes for precision-safe ordering."""
    if not isinstance(value, str) or not value:
        return None
    try:
        parsed = datetime.fromisoformat(value.replace("Z", "+00:00"))
    except ValueError:
        return None
    if parsed.tzinfo is None or parsed.utcoffset() is None:
        return None
    return parsed.astimezone(timezone.utc)


def result_time(path, data=None):
    """Use result metadata or file mtime to timestamp an invalidation event."""
    if isinstance(data, dict):
        parsed = parse_scan_time(data.get("extracted_at"))
        if parsed is not None:
            return parsed
    try:
        return datetime.fromtimestamp(os.path.getmtime(path), timezone.utc)
    except OSError:
        return datetime.min.replace(tzinfo=timezone.utc)


def failure_marker_event(path):
    """Read timestamp and originating workflow provenance from a failure marker."""
    try:
        with open(path) as f:
            content = f.read().strip()
    except (OSError, UnicodeDecodeError):
        return result_time(path), {}

    try:
        marker = json.loads(content)
    except json.JSONDecodeError:
        marker = None
    if isinstance(marker, dict):
        parsed = parse_scan_time(marker.get("timestamp"))
        try:
            run_number = int(marker.get("workflow_run_number", 0))
            run_attempt = int(marker.get("workflow_run_attempt", 0))
        except (TypeError, ValueError):
            run_number = run_attempt = 0
        run_metadata = (
            {"workflow_run_number": run_number, "workflow_run_attempt": run_attempt}
            if run_number > 0 and run_attempt > 0
            else {}
        )
        return parsed if parsed is not None else result_time(path), run_metadata

    parsed = parse_scan_time(content)
    if parsed is not None:
        return parsed, {}
    return result_time(path), {}


def result_run_metadata(root, boundary):
    """Read the nearest artifact provenance sidecar up to the result root."""
    current = os.path.abspath(root)
    boundary = os.path.abspath(boundary)
    while current == boundary or current.startswith(boundary + os.sep):
        provenance = load_json(os.path.join(current, "_scan-provenance.json"))
        if isinstance(provenance, dict):
            try:
                run_number = int(provenance.get("workflow_run_number", 0))
                run_attempt = int(provenance.get("workflow_run_attempt", 0))
            except (TypeError, ValueError):
                run_number = run_attempt = 0
            if run_number > 0 and run_attempt > 0:
                return {
                    "workflow_run_number": run_number,
                    "workflow_run_attempt": run_attempt,
                }
        if current == boundary:
            break
        current = os.path.dirname(current)
    return {}


def failure_marker_covers_result(root, boundary, result_metadata):
    """Check whether an ancestor failure marker covers this result's provenance."""
    current = os.path.abspath(root)
    boundary = os.path.abspath(boundary)
    marker_found = False
    marker_runs = []
    while current == boundary or current.startswith(boundary + os.sep):
        marker_path = os.path.join(current, "_failed")
        if os.path.isfile(marker_path):
            marker_found = True
            _, marker_metadata = failure_marker_event(marker_path)
            if not marker_metadata:
                marker_metadata = result_run_metadata(current, boundary)
            marker_run = run_key(marker_metadata)
            if marker_run is not None:
                marker_runs.append(marker_run)
        if current == boundary:
            break
        current = os.path.dirname(current)
    if not marker_found:
        return False
    result_run = run_key(result_metadata)
    if result_run is None:
        return True
    return bool(marker_runs) and max(marker_runs) >= result_run


def state_prefix_for_path(platform, relative):
    """Map a result/marker path to its platform-wide/org/repo scan-state prefix."""
    if relative in (".", ""):
        return f"{platform}/"
    parts = relative.split(os.sep)
    if len(parts) == 1:
        return f"{platform}/{parts[0]}/"
    return f"{platform}/{parts[0]}/{parts[1]}"


def current_workflow_metadata():
    try:
        run_number = int(os.environ.get("GITHUB_RUN_NUMBER", "0"))
        run_attempt = int(os.environ.get("GITHUB_RUN_ATTEMPT", "0"))
    except ValueError:
        return {}
    if run_number <= 0 or run_attempt <= 0:
        return {}
    return {
        "workflow_run_number": run_number,
        "workflow_run_attempt": run_attempt,
    }


def run_key(entry):
    try:
        number = int(entry.get("workflow_run_number", 0))
        attempt = int(entry.get("workflow_run_attempt", 0))
    except (AttributeError, TypeError, ValueError):
        return None
    return (number, attempt) if number > 0 and attempt > 0 else None


def scan_entry_rank(entry):
    run = run_key(entry)
    return (
        1 if run is not None else 0,
        run[0] if run is not None else 0,
        run[1] if run is not None else 0,
        parse_scan_time(entry.get("scanned_at")) or datetime.min.replace(tzinfo=timezone.utc),
        str(entry.get("analyzer_version") or ""),
        str(entry.get("commit_sha") or ""),
    )


def scan_entry_wins(incoming, current):
    return scan_entry_rank(incoming) > scan_entry_rank(current)


def find_state_invalidations(dirs):
    """Collect failure events from markers and incomplete results."""
    failure_events = []
    for d in dirs:
        platform = os.path.basename(d.rstrip("/")).removesuffix("-results")
        boundary = os.path.abspath(d)
        for root, _, files in os.walk(d):
            relative = os.path.relpath(os.path.abspath(root), boundary)
            if "_failed" in files:
                marker = os.path.join(root, "_failed")
                timestamp, run_metadata = failure_marker_event(marker)
                if not run_metadata:
                    run_metadata = result_run_metadata(root, boundary)
                failure_events.append(
                    (state_prefix_for_path(platform, relative), timestamp, run_metadata)
                )
            if "component-architecture.json" in files:
                architecture_path = os.path.join(root, "component-architecture.json")
                security_path = os.path.join(root, "security-findings.json")
                architecture = load_json(architecture_path)
                findings = load_json(security_path)
                arch_valid = valid_architecture_data(architecture)
                repo = architecture.get("repo") if isinstance(architecture, dict) else None
                repo_prefix = (
                    f"{platform}/{repo}"
                    if isinstance(repo, str) and repo
                    else state_prefix_for_path(platform, relative)
                )
                timestamp_valid = (
                    arch_valid
                    and isinstance(findings, list)
                    and parse_scan_time(architecture.get("extracted_at")) is not None
                )
                if not timestamp_valid:
                    artifact_run_metadata = result_run_metadata(root, boundary)
                    if not failure_marker_covers_result(root, boundary, artifact_run_metadata):
                        failure_events.append(
                            (
                                repo_prefix,
                                result_time(architecture_path, architecture),
                                artifact_run_metadata,
                            )
                        )
            elif "security-findings.json" in files and "_failed" not in files:
                findings_path = os.path.join(root, "security-findings.json")
                artifact_run_metadata = result_run_metadata(root, boundary)
                if not failure_marker_covers_result(root, boundary, artifact_run_metadata):
                    failure_events.append(
                        (
                            state_prefix_for_path(platform, relative),
                            result_time(findings_path, load_json(findings_path)),
                            artifact_run_metadata,
                        )
                    )
    return failure_events


def key_matches_prefix(key, prefix):
    if prefix.endswith("/"):
        return key.startswith(prefix)
    return key == prefix or key.startswith(prefix + "/")


def is_invalidated(key, entry, platform_prefixes, failure_events):
    if any(key_matches_prefix(key, prefix) for prefix in platform_prefixes):
        return True
    success_run = run_key(entry)
    success_time = parse_scan_time(entry.get("scanned_at"))
    for prefix, failure_time, failure_metadata in failure_events:
        if not key_matches_prefix(key, prefix):
            continue
        failure_run = run_key(failure_metadata)
        if success_run is not None and failure_run is not None:
            if success_run > failure_run:
                continue
            if success_run < failure_run:
                return True
        if success_time is None or success_time <= failure_time:
            return True
    return False


def main():
    if "-o" not in sys.argv:
        print(f"Usage: {sys.argv[0]} <dir> [<dir>...] -o <output>", file=sys.stderr)
        sys.exit(1)

    out_idx = sys.argv.index("-o")
    args = sys.argv[1:out_idx]
    output = sys.argv[out_idx + 1]
    dirs = []
    invalidate_platforms = set()
    invalidations_output = None
    i = 0
    while i < len(args):
        if args[i] == "--invalidate-platforms":
            if i + 1 >= len(args) or not args[i + 1] or args[i + 1].startswith("--"):
                print("--invalidate-platforms requires a value", file=sys.stderr)
                sys.exit(1)
            platforms = {p for p in args[i + 1].replace(",", " ").split() if p}
            if not platforms:
                print("--invalidate-platforms requires a value", file=sys.stderr)
                sys.exit(1)
            invalidate_platforms.update(platforms)
            i += 2
        elif args[i] == "--invalidations-output":
            if i + 1 >= len(args) or not args[i + 1] or args[i + 1].startswith("--"):
                print("--invalidations-output requires a path", file=sys.stderr)
                sys.exit(1)
            invalidations_output = args[i + 1]
            i += 2
        else:
            dirs.append(args[i])
            i += 1

    existing = {}
    if os.path.exists(output):
        with open(output) as f:
            existing = json.load(f)

    # The previous schema keyed only by org/repo. It cannot safely be
    # migrated because a repo may be scanned under multiple platforms;
    # discard legacy entries once so each platform is rescanned separately.
    for key in [key for key in existing if key.count("/") == 1]:
        del existing[key]

    invalidated_prefixes = {f"{platform}/" for platform in invalidate_platforms}
    run_metadata = current_workflow_metadata()
    failure_events = find_state_invalidations(dirs)

    if invalidations_output:
        invalidations = []
        now = datetime.now(timezone.utc).isoformat(timespec="microseconds").replace("+00:00", "Z")
        for platform in sorted(invalidate_platforms):
            invalidations.append({"prefix": f"{platform}/", "invalidated_at": now, **run_metadata})
        for prefix, timestamp, failure_run in failure_events:
            invalidation = {
                "prefix": prefix,
                "invalidated_at": timestamp.isoformat(timespec="microseconds").replace("+00:00", "Z"),
                **failure_run,
            }
            invalidations.append(invalidation)
        with open(invalidations_output, "w") as f:
            json.dump({"invalidations": invalidations}, f, indent=2, sort_keys=True)
            f.write("\n")

    # Failed or incomplete results must invalidate their previous SHA,
    # otherwise the next run could skip the same repo again forever.
    for key in [
        key for key in existing
        if is_invalidated(key, existing[key], invalidated_prefixes, failure_events)
    ]:
        del existing[key]

    for platform, path, artifact_run_metadata in find_arch_jsons(dirs, invalidate_platforms):
        with open(path) as f:
            data = json.load(f)
        repo = data.get("repo", "")
        sha = data.get("commit_sha", "")
        ver = data.get("analyzer_version", "")
        if not isinstance(ver, str) or not ver:
            # Match the CLI's default analyzer version when malformed or
            # legacy architecture JSON omits the string value.
            ver = "dev"
        # extracted_at (UTC RFC3339) records when the data was produced;
        # the state commit merges concurrent runs by keeping the newer
        # entry per repo.
        scanned_at = data.get("extracted_at", "")
        candidate_time = parse_scan_time(scanned_at)
        if candidate_time is None:
            print(f"::warning::{platform}/{repo}: missing or invalid extracted_at; scan state not advanced")
            continue
        if repo and sha:
            key = f"{platform}/{repo}"
            candidate = {
                "commit_sha": sha,
                "analyzer_version": ver,
                "scanned_at": scanned_at,
                **artifact_run_metadata,
            }
            if is_invalidated(key, candidate, invalidated_prefixes, failure_events):
                continue
            current = existing.get(key)
            if current is None or scan_entry_wins(candidate, current):
                existing[key] = candidate

    with open(output, "w") as f:
        json.dump(existing, f, indent=2, sort_keys=True)
        f.write("\n")

    print(f"Updated {output}: {len(existing)} repos tracked")


if __name__ == "__main__":
    main()
