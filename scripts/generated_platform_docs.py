#!/usr/bin/env python3
"""Track and sync generated platform docs using published snapshot provenance."""

import json
import os
import subprocess
import sys
from datetime import datetime, timezone

PLATFORMS = ("odh", "rhoai")


class SnapshotMetadataError(ValueError):
    pass


def read_json(path):
    try:
        with open(path) as f:
            return json.load(f)
    except FileNotFoundError:
        return {}


def write_json(path, data):
    with open(path, "w") as f:
        json.dump(data, f, indent=2, sort_keys=True)
        f.write("\n")


def resolve_state_path(repo_root, state_path):
    repo_root = os.path.realpath(repo_root)
    absolute_path = os.path.realpath(os.path.join(repo_root, state_path))
    try:
        common_path = os.path.commonpath((repo_root, absolute_path))
    except ValueError as exc:
        raise ValueError("generated-doc provenance state must be inside the repository") from exc
    if common_path != repo_root:
        raise ValueError("generated-doc provenance state must be inside the repository")
    relative_path = os.path.relpath(absolute_path, repo_root)
    return absolute_path, relative_path


def parse_timestamp(value):
    if not isinstance(value, str) or not value:
        return None
    try:
        parsed = datetime.fromisoformat(value.replace("Z", "+00:00"))
    except ValueError:
        return None
    if parsed.tzinfo is None or parsed.utcoffset() is None:
        return None
    return parsed.astimezone(timezone.utc)


def run_git(repo_root, *args):
    return subprocess.run(
        ["git", *args], cwd=repo_root, capture_output=True, text=True, check=False
    )


def main_docs_match_snapshot(repo_root, platform, snapshot):
    if snapshot is None:
        return True
    docs_path = os.path.join(repo_root, "site", "docs", "use-cases", f"{platform}-platform")
    return os.path.isdir(docs_path) == (not snapshot.get("removed", False))


def overlay_platform_docs(repo_root, generated_ref, platform, entry):
    doc_path = f"site/docs/use-cases/{platform}-platform"
    docs_tree = os.path.join(repo_root, doc_path)
    if entry.get("removed", False):
        subprocess.run(["rm", "-rf", docs_tree], check=True)
        return True

    tree = run_git(repo_root, "ls-tree", "-d", "--name-only", generated_ref, doc_path)
    if tree.returncode != 0 or not tree.stdout.strip():
        return False

    subprocess.run(["rm", "-rf", docs_tree], check=True)
    checkout = run_git(repo_root, "checkout", generated_ref, "--", doc_path)
    if checkout.returncode != 0:
        raise RuntimeError(checkout.stderr)
    return True


def update_state(state_path, output_root, version, run_number, run_attempt,
                 enabled_platforms, failed_platforms=()):
    """Update provenance only for enabled, successful platform docs."""
    state = read_json(state_path)
    failed = set(failed_platforms)
    for platform in PLATFORMS:
        if platform not in enabled_platforms or platform in failed:
            continue
        metadata_path = os.path.join(output_root, platform, version, "snapshot-metadata.json")
        metadata = read_json(metadata_path)
        if metadata:
            state[platform] = {
                "version": metadata.get("version", version),
                "workflow_run_number": int(metadata.get("workflow_run_number", run_number)),
                "workflow_run_attempt": int(metadata.get("workflow_run_attempt", run_attempt)),
                "timestamp": metadata.get("timestamp", ""),
                "analyzer_version": metadata.get("analyzer_version", ""),
                "removed": bool(metadata.get("removed", False)),
            }
        else:
            state[platform] = {
                "version": version,
                "workflow_run_number": int(run_number),
                "workflow_run_attempt": int(run_attempt),
                "timestamp": datetime.now(timezone.utc).strftime("%Y-%m-%dT%H:%M:%SZ"),
                "analyzer_version": "",
                "removed": True,
            }
    write_json(state_path, state)
    return state


def matches_legacy_snapshot(candidate, snapshot):
    """Match pre-run-id metadata using its version and extraction time."""
    candidate_time = parse_timestamp(candidate.get("timestamp")) if candidate else None
    snapshot_time = parse_timestamp(snapshot.get("timestamp"))
    return bool(
        candidate
        and candidate.get("version") == snapshot.get("version")
        and candidate_time is not None
        and snapshot_time is not None
        and candidate_time == snapshot_time
        and candidate.get("removed", False) == snapshot.get("removed", False)
    )


def trusted_generated_state(repo_root, generated_ref, snapshot_ref, state_path):
    """Return generated provenance entries verified by published snapshots."""
    _, git_path = resolve_state_path(repo_root, state_path)
    branch_state = load_ref_json(repo_root, generated_ref, git_path)
    trusted = {}
    for platform, candidate in branch_state.items():
        snapshot = latest_snapshot_state(repo_root, snapshot_ref, platform)
        if snapshot is None:
            continue
        if snapshot.get("workflow_run_number", 0) > 0:
            if matches_snapshot(candidate, snapshot):
                trusted[platform] = candidate
        elif matches_legacy_snapshot(candidate, snapshot):
            trusted[platform] = candidate
    return trusted


def prepare_state_from_generated_branch(
    repo_root, state_path, generated_ref, snapshot_ref, excluded_platforms=()
):
    """Merge only generated-branch state verified by published snapshots."""
    if not generated_ref:
        return
    state_file, git_path = resolve_state_path(repo_root, state_path)
    main_state = read_json(state_file)
    generated_state = trusted_generated_state(
        repo_root, generated_ref, snapshot_ref, git_path
    )
    write_json(
        state_file,
        merge_platform_state(main_state, generated_state, excluded_platforms),
    )


def merge_platform_state(main_state, generated_state, excluded_platforms=()):
    """Keep the newer per-platform provenance entry from either source."""
    merged = dict(main_state)
    excluded = set(excluded_platforms)
    for platform, entry in generated_state.items():
        if platform in excluded:
            continue
        current = merged.get(platform)
        if current is None or run_key(entry) > run_key(current):
            merged[platform] = entry
    return merged


def sync_disabled_platform_docs(repo_root, state_path, generated_ref,
                                snapshot_ref, enabled_platforms, failed_platforms=()):
    """Keep unscanned platform trees paired with published provenance.

    Returns platforms whose current main/generated trees do not match the
    latest published snapshot, allowing the workflow to withhold deployment.
    """
    _, git_state_path = resolve_state_path(repo_root, state_path)
    state = load_ref_json(repo_root, generated_ref, git_state_path)
    main_state = read_json(os.path.join(repo_root, state_path))
    excluded = set(enabled_platforms) | set(failed_platforms)
    stale = []
    selected_state = {}
    for platform in PLATFORMS:
        if platform in excluded:
            continue
        snapshot = latest_snapshot_state(repo_root, snapshot_ref, platform)
        source = choose_source(main_state, state, snapshot, platform)
        if source == "stale":
            stale.append(platform)
            continue
        if source == "main" and not main_docs_match_snapshot(repo_root, platform, snapshot):
            stale.append(platform)
            continue
        if source != "generated" or not generated_ref:
            continue
        if not overlay_platform_docs(repo_root, generated_ref, platform, state.get(platform, {})):
            stale.append(platform)
            continue
        selected_state[platform] = state.get(platform, snapshot or {})
    return stale, selected_state


def load_ref_json(repo_root, ref, path):
    if not ref:
        return {}
    result = run_git(repo_root, "show", f"{ref}:{path}")
    if result.returncode != 0:
        return {}
    try:
        return json.loads(result.stdout)
    except json.JSONDecodeError:
        return {}


def load_snapshot_metadata(repo_root, snapshot_ref, path):
    result = run_git(repo_root, "show", f"{snapshot_ref}:{path}")
    if result.returncode != 0:
        raise SnapshotMetadataError(f"could not read {path}: {result.stderr.strip()}")
    try:
        metadata = json.loads(result.stdout)
    except json.JSONDecodeError as exc:
        raise SnapshotMetadataError(f"invalid JSON in {path}") from exc
    if not isinstance(metadata, dict):
        raise SnapshotMetadataError(f"expected an object in {path}")
    return metadata


def latest_snapshot_state(repo_root, snapshot_ref, platform):
    """Return the newest published metadata by workflow run/attempt."""
    path = latest_snapshot_path(repo_root, snapshot_ref, platform)
    return load_snapshot_metadata(repo_root, snapshot_ref, path) if path else None


def latest_snapshot_path(repo_root, snapshot_ref, platform):
    """Return the path of the newest snapshot metadata for one platform."""
    if not snapshot_ref:
        return None
    result = run_git(repo_root, "ls-tree", "-r", "--name-only", snapshot_ref, "output")
    if result.returncode != 0:
        raise SnapshotMetadataError(
            f"could not list snapshot metadata: {result.stderr.strip()}"
        )

    candidates = []
    for path in result.stdout.splitlines():
        parts = path.split("/")
        if len(parts) < 5 or parts[0] != "output" or parts[-3] != platform or parts[-1] != "snapshot-metadata.json":
            continue
        metadata = load_snapshot_metadata(repo_root, snapshot_ref, path)
        if metadata.get("platform", platform) != platform:
            raise SnapshotMetadataError(f"platform does not match metadata path {path}")
        try:
            run_number = int(metadata.get("workflow_run_number", 0))
            run_attempt = int(metadata.get("workflow_run_attempt", 0))
        except (TypeError, ValueError):
            run_number = run_attempt = 0
        timestamp = parse_timestamp(metadata.get("timestamp"))
        candidates.append((
            run_number,
            run_attempt,
            timestamp or datetime.min.replace(tzinfo=timezone.utc),
            path,
        ))
    if not candidates:
        return None
    return max(candidates, key=lambda item: item[:3])[3]


def platform_snapshot_is_newer(repo_root, snapshot_ref, platform, run_number, run_attempt):
    """Check whether any published label for a platform came from a newer run."""
    latest = latest_snapshot_state(repo_root, snapshot_ref, platform)
    if latest is None:
        return False
    return run_key(latest) > (int(run_number), int(run_attempt))


def run_key(state):
    return (
        int(state.get("workflow_run_number", 0)),
        int(state.get("workflow_run_attempt", 0)),
    )


def matches_snapshot(candidate, snapshot):
    required = all(
        candidate.get(key) == snapshot.get(key)
        for key in ("version", "workflow_run_number", "workflow_run_attempt")
    )
    if not candidate or not required:
        return False
    if candidate.get("removed", False) != snapshot.get("removed", False):
        return False
    if snapshot.get("analyzer_version") and candidate.get("analyzer_version") != snapshot["analyzer_version"]:
        return False
    return True


def choose_source(main_state, generated_state, snapshot_state, platform):
    """Choose main/generated docs or report stale/unknown snapshot provenance."""
    main = main_state.get(platform, {})
    generated = generated_state.get(platform, {})

    # A generated-docs branch can be pushed before its snapshot. Without a
    # published snapshot there is no proof that its newer run marker became
    # current data; keep main rather than deploying uncommitted provenance.
    if snapshot_state is None:
        return "main"

    if snapshot_state:
        if snapshot_state.get("workflow_run_number", 0) > 0:
            if matches_snapshot(main, snapshot_state):
                return "main"
            if matches_snapshot(generated, snapshot_state):
                return "generated"
            return "stale"

        # Legacy snapshot metadata has no workflow run id. Require a docs
        # candidate for the same label whose extraction predates the
        # snapshot timestamp. A newer marker for the same label means docs
        # were generated by a run whose snapshot was never published.
        if matches_legacy_snapshot(main, snapshot_state):
            return "main"
        if matches_legacy_snapshot(generated, snapshot_state):
            return "generated"
        if not main and not generated:
            return "main"
        return "stale"

    return "main"


def sync_docs(repo_root, generated_ref, snapshot_ref, state_path):
    """Overlay only platform docs whose provenance matches published data."""
    state_file, git_path = resolve_state_path(repo_root, state_path)
    main_state = read_json(state_file)
    generated_state = load_ref_json(repo_root, generated_ref, git_path)
    choices = {}
    for platform in PLATFORMS:
        try:
            snapshot = latest_snapshot_state(repo_root, snapshot_ref, platform)
        except SnapshotMetadataError as exc:
            print(f"::warning::{platform} snapshot metadata is unreadable: {exc}")
            return False
        choice = choose_source(main_state, generated_state, snapshot, platform)
        if choice == "main" and not main_docs_match_snapshot(repo_root, platform, snapshot):
            choice = "stale"
        choices[platform] = choice
        if choices[platform] == "stale":
            print(f"::warning::{platform} docs do not match the latest published snapshot provenance")
            return False

    for platform, source in choices.items():
        if source != "generated" or not generated_ref:
            continue
        entry = generated_state.get(platform, {})
        if not overlay_platform_docs(repo_root, generated_ref, platform, entry):
            print(f"::warning::{platform} generated docs tree is missing or unreadable")
            return False
    return True


def main():
    if len(sys.argv) < 2:
        print("Usage: generated_platform_docs.py <update|sync> ...", file=sys.stderr)
        sys.exit(2)
    action = sys.argv[1]
    if action == "update" and len(sys.argv) == 12:
        _, _, repo_root, state_path, output_root, version, run_number, run_attempt, enabled, failed, generated_ref, snapshot_ref = sys.argv
        repo_root = os.path.realpath(repo_root)
        state_path, _ = resolve_state_path(repo_root, state_path)
        if not os.path.isabs(output_root):
            output_root = os.path.join(repo_root, output_root)
        enabled_platforms = {p for p in enabled.replace(",", " ").split() if p}
        failed_platforms = {p for p in failed.replace(",", " ").split() if p}
        excluded_platforms = enabled_platforms | failed_platforms
        docs_stale = []
        selected_docs_state = {}
        try:
            docs_stale, selected_docs_state = sync_disabled_platform_docs(
                repo_root,
                state_path,
                generated_ref or None,
                snapshot_ref or None,
                enabled_platforms,
                failed_platforms,
            )
            prepare_state_from_generated_branch(
                repo_root,
                state_path,
                generated_ref or None,
                snapshot_ref or None,
                failed_platforms | set(docs_stale),
            )
        except SnapshotMetadataError as exc:
            print(f"::warning::snapshot metadata is unreadable: {exc}")
            docs_stale = sorted(set(docs_stale) | (set(PLATFORMS) - excluded_platforms))
            selected_docs_state = {}
        state = update_state(
            state_path,
            output_root,
            version,
            run_number,
            run_attempt,
            enabled_platforms,
            failed_platforms,
        )
        # Keep the provenance paired with the disabled-platform tree chosen
        # above; a newer unpublished main marker must not describe older
        # generated-branch files.
        state.update(selected_docs_state)
        write_json(state_path, state)
        output = os.environ.get("GITHUB_OUTPUT")
        if output:
            with open(output, "a") as f:
                f.write(f"docs-stale-platforms={' '.join(docs_stale)}\n")
        if docs_stale:
            print(f"Unverified docs provenance for: {', '.join(docs_stale)}")
        return
    if action == "sync" and len(sys.argv) == 6:
        _, _, repo_root, generated_ref, snapshot_ref, state_path = sys.argv
        should_deploy = sync_docs(repo_root, generated_ref or None, snapshot_ref or None, state_path)
        output = os.environ.get("GITHUB_OUTPUT")
        if output:
            with open(output, "a") as f:
                f.write(f"should-deploy={'true' if should_deploy else 'false'}\n")
        print(f"should-deploy={'true' if should_deploy else 'false'}")
        return
    if action == "platform-newer" and len(sys.argv) == 7:
        _, _, repo_root, snapshot_ref, platform, run_number, run_attempt = sys.argv
        newer = platform_snapshot_is_newer(
            repo_root, snapshot_ref or None, platform, run_number, run_attempt
        )
        print("yes" if newer else "no")
        return
    if action == "latest-snapshot-path" and len(sys.argv) == 5:
        _, _, repo_root, snapshot_ref, platform = sys.argv
        print(latest_snapshot_path(repo_root, snapshot_ref or None, platform) or "")
        return
    print("Usage: generated_platform_docs.py update <repo-root> <state> <output-root> <version> <run> <attempt> <enabled-csv> <failed-csv> <generated-ref> <snapshot-ref>")
    print("   or: generated_platform_docs.py sync <repo-root> <generated-ref> <snapshot-ref> <state-path>")
    print("   or: generated_platform_docs.py platform-newer <repo-root> <snapshot-ref> <platform> <run> <attempt>")
    print("   or: generated_platform_docs.py latest-snapshot-path <repo-root> <snapshot-ref> <platform>")
    sys.exit(2)


if __name__ == "__main__":
    main()
