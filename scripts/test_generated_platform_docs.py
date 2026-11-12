"""Tests for generated platform-doc provenance and deployment selection."""

import json
import os
import shutil
import subprocess
import sys
import tempfile
import unittest

from generated_platform_docs import (
    choose_source,
    latest_snapshot_path,
    latest_snapshot_state,
    matches_legacy_snapshot,
    merge_platform_state,
    prepare_state_from_generated_branch,
    resolve_state_path,
    sync_disabled_platform_docs,
    sync_docs,
    update_state,
)

SCRIPT = os.path.join(os.path.dirname(__file__), "generated_platform_docs.py")


class GeneratedPlatformDocsTest(unittest.TestCase):
    def write_json(self, path, value):
        os.makedirs(os.path.dirname(path), exist_ok=True)
        with open(path, "w") as f:
            json.dump(value, f)

    def read_json(self, path):
        with open(path) as f:
            return json.load(f)

    def git(self, repo, *args):
        return subprocess.run(
            ["git", *args], cwd=repo, check=True, capture_output=True, text=True
        )

    def init_repo(self, path):
        os.makedirs(path)
        self.git(path, "init", "-b", "main")
        self.git(path, "config", "user.name", "test")
        self.git(path, "config", "user.email", "test@example.invalid")

    def commit_all(self, repo, message):
        self.git(repo, "add", "-A")
        self.git(repo, "commit", "-m", message)

    def test_update_state_preserves_disabled_and_failed_platforms(self):
        with tempfile.TemporaryDirectory() as tmp:
            output = os.path.join(tmp, "platform-output")
            state_path = os.path.join(tmp, ".generated-platform-provenance.json")
            self.write_json(state_path, {"rhoai": {"version": "old", "workflow_run_number": 3}})
            self.write_json(
                os.path.join(output, "odh", "v1", "snapshot-metadata.json"),
                {
                    "version": "v1",
                    "workflow_run_number": 10,
                    "workflow_run_attempt": 2,
                    "removed": False,
                },
            )

            state = update_state(state_path, output, "v1", 10, 2, {"odh", "rhoai"}, {"rhoai"})

            self.assertEqual(state["odh"]["workflow_run_number"], 10)
            self.assertEqual(state["rhoai"]["version"], "old")

    def test_update_cli_resolves_relative_state_path_under_repo_root(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo_root = os.path.join(tmp, "repo")
            caller_dir = os.path.join(tmp, "caller")
            os.makedirs(repo_root)
            os.makedirs(caller_dir)

            subprocess.run(
                [
                    sys.executable,
                    SCRIPT,
                    "update",
                    repo_root,
                    ".generated-platform-provenance.json",
                    "platform-output",
                    "v1",
                    "10",
                    "1",
                    "",
                    "",
                    "",
                    "",
                ],
                cwd=caller_dir,
                check=True,
                capture_output=True,
                text=True,
            )

            self.assertTrue(os.path.isfile(os.path.join(repo_root, ".generated-platform-provenance.json")))
            self.assertFalse(os.path.exists(os.path.join(caller_dir, ".generated-platform-provenance.json")))

    def test_resolve_state_path_rejects_symlink_escape(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo_root = os.path.join(tmp, "repo")
            outside = os.path.join(tmp, "outside.json")
            os.makedirs(repo_root)
            with open(outside, "w") as f:
                json.dump({}, f)
            os.symlink(outside, os.path.join(repo_root, ".generated-platform-provenance.json"))

            with self.assertRaises(ValueError):
                resolve_state_path(repo_root, ".generated-platform-provenance.json")

    def test_prepare_state_uses_repo_relative_path_for_git_show(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            state_path = ".generated-platform-provenance.json"
            self.write_json(os.path.join(repo, state_path), {
                "odh": {
                    "version": "v1",
                    "workflow_run_number": 19,
                    "workflow_run_attempt": 1,
                    "analyzer_version": "v2",
                    "removed": False,
                }
            })
            self.commit_all(repo, "main provenance")

            self.git(repo, "checkout", "-b", "snapshots")
            self.write_json(
                os.path.join(repo, "output", "v1", "odh", "v1", "snapshot-metadata.json"),
                {
                    "platform": "odh",
                    "version": "v1",
                    "workflow_run_number": 20,
                    "workflow_run_attempt": 1,
                    "analyzer_version": "v2",
                    "removed": False,
                },
            )
            self.commit_all(repo, "published snapshot")
            self.git(repo, "checkout", "main")

            self.git(repo, "checkout", "-b", "generated")
            self.write_json(os.path.join(repo, state_path), {
                "odh": {
                    "version": "v1",
                    "workflow_run_number": 20,
                    "workflow_run_attempt": 1,
                    "analyzer_version": "v2",
                    "removed": False,
                }
            })
            self.commit_all(repo, "generated provenance")
            self.git(repo, "checkout", "main")

            prepare_state_from_generated_branch(
                repo, os.path.join(repo, state_path), "generated", "snapshots"
            )

            state = self.read_json(os.path.join(repo, state_path))
            self.assertEqual(state["odh"]["workflow_run_number"], 20)

    def test_generated_state_does_not_replace_newer_main_state(self):
        main = {"rhoai": {"workflow_run_number": 20, "workflow_run_attempt": 1}}
        generated = {"rhoai": {"workflow_run_number": 18, "workflow_run_attempt": 4}}
        self.assertEqual(merge_platform_state(main, generated), main)

    def test_failed_platform_generated_state_is_not_merged(self):
        main = {
            "rhoai": {"workflow_run_number": 19, "workflow_run_attempt": 1},
            "odh": {"workflow_run_number": 19, "workflow_run_attempt": 1},
        }
        generated = {
            "rhoai": {"workflow_run_number": 20, "workflow_run_attempt": 1},
            "odh": {"workflow_run_number": 20, "workflow_run_attempt": 1},
        }

        merged = merge_platform_state(main, generated, {"rhoai"})

        self.assertEqual(merged["rhoai"]["workflow_run_number"], 19)
        self.assertEqual(merged["odh"]["workflow_run_number"], 20)

    def test_update_cli_accepts_space_separated_platforms(self):
        with tempfile.TemporaryDirectory() as tmp:
            state_path = os.path.join(tmp, ".generated-platform-provenance.json")
            output = os.path.join(tmp, "platform-output")
            self.write_json(os.path.join(output, "odh", "v1", "snapshot-metadata.json"), {
                "version": "v1", "workflow_run_number": 10,
                "workflow_run_attempt": 1, "removed": False,
            })
            subprocess.run(
                [sys.executable, SCRIPT, "update", tmp, state_path, output,
                 "v1", "10", "1", "odh rhoai", "", "", ""],
                check=True,
                capture_output=True,
                text=True,
            )
            state = self.read_json(state_path)
            self.assertEqual(state["odh"]["workflow_run_number"], 10)
            self.assertTrue(state["rhoai"]["removed"])

    def test_disabled_platform_docs_follow_their_provenance(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            state_path = ".generated-platform-provenance.json"
            doc_path = os.path.join(
                repo, "site", "docs", "use-cases", "rhoai-platform", "index.md"
            )
            os.makedirs(os.path.dirname(doc_path))
            with open(doc_path, "w") as f:
                f.write("main-stale\n")
            self.write_json(
                os.path.join(repo, state_path),
                {"rhoai": {"version": "v1", "workflow_run_number": 1, "removed": False}},
            )
            self.commit_all(repo, "main docs")

            self.git(repo, "checkout", "-b", "snapshots")
            self.write_json(
                os.path.join(repo, "output", "v2", "rhoai", "v2", "snapshot-metadata.json"),
                {"platform": "rhoai", "version": "v2", "workflow_run_number": 2,
                 "workflow_run_attempt": 1, "removed": False},
            )
            self.commit_all(repo, "published rhoai snapshot")
            self.git(repo, "checkout", "main")

            self.git(repo, "checkout", "-b", "generated")
            with open(doc_path, "w") as f:
                f.write("generated-current\n")
            self.write_json(
                os.path.join(repo, state_path),
                {"rhoai": {"version": "v2", "workflow_run_number": 2,
                           "workflow_run_attempt": 1, "removed": False}},
            )
            self.commit_all(repo, "generated rhoai docs")
            self.git(repo, "checkout", "main")

            stale, selected = sync_disabled_platform_docs(
                repo, state_path, "generated", "snapshots", {"odh"}
            )

            self.assertEqual(stale, [])
            self.assertEqual(selected["rhoai"]["workflow_run_number"], 2)
            with open(doc_path) as f:
                self.assertEqual(f.read(), "generated-current\n")

    def test_legacy_snapshot_without_provenance_preserves_main_docs(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            doc_path = os.path.join(
                repo, "site", "docs", "use-cases", "rhoai-platform", "index.md"
            )
            os.makedirs(os.path.dirname(doc_path))
            with open(doc_path, "w") as f:
                f.write("main-legacy\n")
            self.commit_all(repo, "main legacy docs")

            self.git(repo, "checkout", "-b", "snapshots")
            self.write_json(
                os.path.join(repo, "output", "v1", "rhoai", "v1", "snapshot-metadata.json"),
                {"platform": "rhoai", "version": "v1", "timestamp": "2026-11-10T00:00:00Z"},
            )
            self.commit_all(repo, "legacy snapshot")
            self.git(repo, "checkout", "main")

            self.git(repo, "checkout", "-b", "generated")
            with open(doc_path, "w") as f:
                f.write("generated-legacy-current\n")
            self.commit_all(repo, "maintained legacy generated docs")
            self.git(repo, "checkout", "main")

            stale, selected = sync_disabled_platform_docs(
                repo, ".generated-platform-provenance.json", "generated", "snapshots", {"odh"}
            )

            self.assertEqual(stale, [])
            self.assertEqual(selected, {})
            with open(doc_path) as f:
                self.assertEqual(f.read(), "main-legacy\n")

    def test_disabled_sync_reports_mismatched_generated_legacy_candidate(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            state_path = ".generated-platform-provenance.json"
            doc_path = os.path.join(
                repo, "site", "docs", "use-cases", "rhoai-platform", "index.md"
            )
            os.makedirs(os.path.dirname(doc_path))
            with open(doc_path, "w") as f:
                f.write("main-docs\n")
            self.write_json(os.path.join(repo, state_path), {})
            self.commit_all(repo, "main without platform provenance")

            self.git(repo, "checkout", "-b", "snapshots")
            self.write_json(
                os.path.join(repo, "output", "v1", "rhoai", "v1", "snapshot-metadata.json"),
                {"platform": "rhoai", "version": "v1", "timestamp": "2026-11-12T05:00:00Z"},
            )
            self.commit_all(repo, "legacy snapshot")
            self.git(repo, "checkout", "main")

            self.git(repo, "checkout", "-b", "generated")
            with open(doc_path, "w") as f:
                f.write("unverified-generated-docs\n")
            self.write_json(
                os.path.join(repo, state_path),
                {"rhoai": {"version": "v1", "timestamp": "2026-11-12T04:59:00Z"}},
            )
            self.commit_all(repo, "generated docs with mismatched legacy provenance")
            self.git(repo, "checkout", "main")

            stale, selected = sync_disabled_platform_docs(
                repo, state_path, "generated", "snapshots", {"odh"}
            )

            self.assertEqual(stale, ["rhoai"])
            self.assertEqual(selected, {})
            with open(doc_path) as f:
                self.assertEqual(f.read(), "main-docs\n")

            self.assertFalse(
                sync_docs(
                    repo,
                    "generated",
                    "snapshots",
                    os.path.join(repo, state_path),
                )
            )
            with open(doc_path) as f:
                self.assertEqual(f.read(), "main-docs\n")

    def test_disabled_platform_does_not_replace_newer_main_docs(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            state_path = ".generated-platform-provenance.json"
            doc_path = os.path.join(
                repo, "site", "docs", "use-cases", "rhoai-platform", "index.md"
            )
            os.makedirs(os.path.dirname(doc_path))
            with open(doc_path, "w") as f:
                f.write("main-newer\n")
            self.write_json(
                os.path.join(repo, state_path),
                {"rhoai": {"version": "v2", "workflow_run_number": 20,
                           "workflow_run_attempt": 1, "removed": False}},
            )
            self.commit_all(repo, "newer main docs")

            self.git(repo, "checkout", "-b", "snapshots")
            self.write_json(
                os.path.join(repo, "output", "v2", "rhoai", "v2", "snapshot-metadata.json"),
                {"platform": "rhoai", "version": "v2", "workflow_run_number": 20,
                 "workflow_run_attempt": 1, "removed": False},
            )
            self.commit_all(repo, "published newer rhoai snapshot")
            self.git(repo, "checkout", "main")

            self.git(repo, "checkout", "-b", "generated")
            with open(doc_path, "w") as f:
                f.write("generated-older\n")
            self.write_json(
                os.path.join(repo, state_path),
                {"rhoai": {"version": "v1", "workflow_run_number": 18,
                           "workflow_run_attempt": 4, "removed": False}},
            )
            self.commit_all(repo, "older generated docs")
            self.git(repo, "checkout", "main")

            stale, selected = sync_disabled_platform_docs(
                repo, state_path, "generated", "snapshots", {"odh"}
            )

            self.assertEqual(stale, [])
            self.assertNotIn("rhoai", selected)
            with open(doc_path) as f:
                self.assertEqual(f.read(), "main-newer\n")

    def test_disabled_sync_rejects_missing_main_docs_for_nonremoved_snapshot(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            state_path = ".generated-platform-provenance.json"
            provenance = {
                "version": "v1",
                "workflow_run_number": 20,
                "workflow_run_attempt": 1,
                "removed": False,
            }
            self.write_json(os.path.join(repo, state_path), {"odh": provenance})
            self.commit_all(repo, "main provenance without docs")

            self.git(repo, "checkout", "-b", "snapshots")
            self.write_json(
                os.path.join(repo, "output", "v1", "odh", "v1", "snapshot-metadata.json"),
                {"platform": "odh", **provenance},
            )
            self.commit_all(repo, "published snapshot")
            self.git(repo, "checkout", "main")

            stale, selected = sync_disabled_platform_docs(
                repo, state_path, None, "snapshots", {"rhoai"}
            )

            self.assertEqual(stale, ["odh"])
            self.assertEqual(selected, {})

    def test_choose_only_provenance_matching_latest_snapshot(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            snapshot_path = os.path.join(
                repo, "output", "2026-11-12", "odh", "2026-11-12", "snapshot-metadata.json"
            )
            self.write_json(
                snapshot_path,
                {
                    "platform": "odh",
                    "version": "2026-11-12",
                    "workflow_run_number": 20,
                    "workflow_run_attempt": 1,
                    "removed": False,
                },
            )
            self.commit_all(repo, "snapshot")

            snapshot = latest_snapshot_state(repo, "HEAD", "odh")
            main = {"odh": {"version": "2026-11-12", "workflow_run_number": 19, "workflow_run_attempt": 1, "removed": False}}
            generated = {"odh": {"version": "2026-11-12", "workflow_run_number": 20, "workflow_run_attempt": 1, "removed": False}}
            self.assertEqual(choose_source(main, generated, snapshot, "odh"), "generated")

            main["odh"] = generated["odh"].copy()
            generated["odh"]["workflow_run_number"] = 18
            self.assertEqual(choose_source(main, generated, snapshot, "odh"), "main")

            main["odh"]["workflow_run_number"] = 17
            self.assertEqual(choose_source(main, generated, snapshot, "odh"), "stale")

    def test_latest_snapshot_lookup_ignores_other_platforms(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            for platform, run in (("odh", 5), ("rhoai", 8)):
                path = os.path.join(repo, "output", "2026-11-12", platform, "2026-11-12", "snapshot-metadata.json")
                self.write_json(path, {"platform": platform, "version": "2026-11-12", "workflow_run_number": run})
            self.commit_all(repo, "snapshots")

            odh = latest_snapshot_state(repo, "HEAD", "odh")
            rhoai = latest_snapshot_state(repo, "HEAD", "rhoai")
            self.assertEqual(odh["workflow_run_number"], 5)
            self.assertEqual(rhoai["workflow_run_number"], 8)

    def test_latest_snapshot_uses_run_provenance_not_label_order(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            older_label = os.path.join(
                repo, "output", "2026-11-10", "odh", "2026-11-10", "snapshot-metadata.json"
            )
            newer_label = os.path.join(
                repo, "output", "2026-11-11", "odh", "2026-11-11", "snapshot-metadata.json"
            )
            self.write_json(older_label, {
                "platform": "odh", "version": "2026-11-10",
                "workflow_run_number": 30, "workflow_run_attempt": 1,
            })
            self.write_json(newer_label, {
                "platform": "odh", "version": "2026-11-11",
                "workflow_run_number": 29, "workflow_run_attempt": 1,
            })
            self.commit_all(repo, "snapshots out of label order")

            selected = latest_snapshot_path(repo, "HEAD", "odh")
            self.assertIn("output/2026-11-10/odh/2026-11-10/snapshot-metadata.json", selected)

    def test_latest_legacy_snapshot_uses_timestamp_instants(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            whole_second = os.path.join(
                repo, "output", "v1", "odh", "v1", "snapshot-metadata.json"
            )
            fractional_second = os.path.join(
                repo, "output", "v2", "odh", "v2", "snapshot-metadata.json"
            )
            self.write_json(whole_second, {
                "platform": "odh", "version": "v1", "timestamp": "2026-11-12T05:00:00Z"
            })
            self.write_json(fractional_second, {
                "platform": "odh", "version": "v2", "timestamp": "2026-11-12T05:00:00.500Z"
            })
            self.commit_all(repo, "legacy snapshots with fractional timestamps")

            selected = latest_snapshot_path(repo, "HEAD", "odh")

            self.assertIn("output/v2/odh/v2/snapshot-metadata.json", selected)

    def test_legacy_doc_approval_compares_timestamp_instants(self):
        snapshot = {"version": "v1", "timestamp": "2026-11-12T05:00:00Z"}
        generated_later = {
            "version": "v1",
            "timestamp": "2026-11-12T05:00:00.500Z",
            "removed": False,
        }
        self.assertFalse(matches_legacy_snapshot(generated_later, snapshot))
        generated_earlier = {
            "version": "v1",
            "timestamp": "2026-11-12T04:59:00Z",
            "removed": False,
        }
        self.assertFalse(matches_legacy_snapshot(generated_earlier, snapshot))
        exact = {"version": "v1", "timestamp": snapshot["timestamp"], "removed": False}
        self.assertTrue(matches_legacy_snapshot(exact, snapshot))

    def test_sync_overlays_generated_docs_matching_snapshot_provenance(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            docs = os.path.join(repo, "site", "docs", "use-cases", "odh-platform")
            os.makedirs(docs)
            with open(os.path.join(docs, "index.md"), "w") as f:
                f.write("main-old\n")
            self.write_json(
                os.path.join(repo, ".generated-platform-provenance.json"),
                {"odh": {"version": "2026-11-12", "workflow_run_number": 19,
                         "workflow_run_attempt": 1, "removed": False}},
            )
            self.commit_all(repo, "main docs")

            self.git(repo, "checkout", "-b", "snapshots")
            self.write_json(
                os.path.join(repo, "output", "2026-11-12", "odh", "2026-11-12", "snapshot-metadata.json"),
                {"platform": "odh", "version": "2026-11-12", "workflow_run_number": 20,
                 "workflow_run_attempt": 1, "removed": False},
            )
            self.commit_all(repo, "published snapshot")

            self.git(repo, "checkout", "main")
            self.git(repo, "checkout", "-b", "generated")
            with open(os.path.join(docs, "index.md"), "w") as f:
                f.write("generated-current\n")
            self.write_json(
                os.path.join(repo, ".generated-platform-provenance.json"),
                {"odh": {"version": "2026-11-12", "workflow_run_number": 20,
                         "workflow_run_attempt": 1, "removed": False}},
            )
            self.commit_all(repo, "generated docs")

            self.git(repo, "checkout", "main")
            self.assertTrue(
                sync_docs(
                    repo,
                    "generated",
                    "snapshots",
                    ".generated-platform-provenance.json",
                )
            )
            with open(os.path.join(docs, "index.md")) as f:
                self.assertEqual(f.read(), "generated-current\n")

    def test_sync_rejects_missing_generated_tree_before_deleting_main_docs(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            state_path = ".generated-platform-provenance.json"
            docs = os.path.join(repo, "site", "docs", "use-cases", "odh-platform")
            os.makedirs(docs)
            with open(os.path.join(docs, "index.md"), "w") as f:
                f.write("main-current\n")
            self.write_json(
                os.path.join(repo, state_path),
                {"odh": {"version": "v1", "workflow_run_number": 19,
                         "workflow_run_attempt": 1, "removed": False}},
            )
            self.commit_all(repo, "main docs")

            self.git(repo, "checkout", "-b", "snapshots")
            self.write_json(
                os.path.join(repo, "output", "v1", "odh", "v1", "snapshot-metadata.json"),
                {"platform": "odh", "version": "v1", "workflow_run_number": 20,
                 "workflow_run_attempt": 1, "removed": False},
            )
            self.commit_all(repo, "published snapshot")
            self.git(repo, "checkout", "main")

            self.git(repo, "checkout", "-b", "generated")
            shutil.rmtree(docs)
            self.write_json(
                os.path.join(repo, state_path),
                {"odh": {"version": "v1", "workflow_run_number": 20,
                         "workflow_run_attempt": 1, "removed": False}},
            )
            self.commit_all(repo, "generated provenance without docs tree")
            self.git(repo, "checkout", "main")

            self.assertFalse(
                sync_docs(
                    repo,
                    "generated",
                    "snapshots",
                    os.path.join(repo, state_path),
                )
            )
            stale, selected = sync_disabled_platform_docs(
                repo, state_path, "generated", "snapshots", {"rhoai"}
            )
            self.assertEqual(stale, ["odh"])
            self.assertEqual(selected, {})
            with open(os.path.join(docs, "index.md")) as f:
                self.assertEqual(f.read(), "main-current\n")

            output_file = os.path.join(tmp, "github-output")
            update = subprocess.run(
                [
                    sys.executable,
                    SCRIPT,
                    "update",
                    repo,
                    state_path,
                    "platform-output",
                    "v2",
                    "21",
                    "1",
                    "rhoai",
                    "",
                    "generated",
                    "snapshots",
                ],
                cwd=repo,
                env=dict(os.environ, GITHUB_OUTPUT=output_file),
                capture_output=True,
                text=True,
            )
            self.assertEqual(update.returncode, 0, update.stdout + update.stderr)
            state = self.read_json(os.path.join(repo, state_path))
            self.assertEqual(state["odh"]["workflow_run_number"], 19)
            with open(output_file) as f:
                self.assertIn("docs-stale-platforms=odh", f.read())

    def test_sync_rejects_missing_main_docs_for_nonremoved_snapshot(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            state_path = ".generated-platform-provenance.json"
            provenance = {
                "version": "v1",
                "workflow_run_number": 20,
                "workflow_run_attempt": 1,
                "removed": False,
            }
            self.write_json(os.path.join(repo, state_path), {"odh": provenance})
            self.commit_all(repo, "main provenance without docs")

            self.git(repo, "checkout", "-b", "snapshots")
            self.write_json(
                os.path.join(repo, "output", "v1", "odh", "v1", "snapshot-metadata.json"),
                {"platform": "odh", **provenance},
            )
            self.commit_all(repo, "published snapshot")
            self.git(repo, "checkout", "main")

            self.assertFalse(sync_docs(repo, None, "snapshots", state_path))

    def test_sync_allows_missing_main_docs_for_removed_snapshot(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            state_path = ".generated-platform-provenance.json"
            provenance = {
                "version": "v1",
                "workflow_run_number": 20,
                "workflow_run_attempt": 1,
                "removed": True,
            }
            self.write_json(os.path.join(repo, state_path), {"odh": provenance})
            self.commit_all(repo, "main removal provenance")

            self.git(repo, "checkout", "-b", "snapshots")
            self.write_json(
                os.path.join(repo, "output", "v1", "odh", "v1", "snapshot-metadata.json"),
                {"platform": "odh", **provenance},
            )
            self.commit_all(repo, "published removal snapshot")
            self.git(repo, "checkout", "main")

            self.assertTrue(sync_docs(repo, None, "snapshots", state_path))

    def test_sync_rejects_existing_main_docs_for_removed_snapshot(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            state_path = ".generated-platform-provenance.json"
            docs = os.path.join(repo, "site", "docs", "use-cases", "odh-platform")
            os.makedirs(docs)
            with open(os.path.join(docs, "index.md"), "w") as f:
                f.write("stale docs\n")
            provenance = {
                "version": "v1",
                "workflow_run_number": 20,
                "workflow_run_attempt": 1,
                "removed": True,
            }
            self.write_json(os.path.join(repo, state_path), {"odh": provenance})
            self.commit_all(repo, "main removal provenance with stale docs")

            self.git(repo, "checkout", "-b", "snapshots")
            self.write_json(
                os.path.join(repo, "output", "v1", "odh", "v1", "snapshot-metadata.json"),
                {"platform": "odh", **provenance},
            )
            self.commit_all(repo, "published removal snapshot")
            self.git(repo, "checkout", "main")

            self.assertFalse(sync_docs(repo, None, "snapshots", state_path))

    def test_sync_rejects_unreadable_snapshot_metadata(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            state_path = ".generated-platform-provenance.json"
            docs = os.path.join(repo, "site", "docs", "use-cases", "odh-platform")
            os.makedirs(docs)
            with open(os.path.join(docs, "index.md"), "w") as f:
                f.write("main docs\n")
            old_provenance = {
                "version": "v1",
                "workflow_run_number": 19,
                "workflow_run_attempt": 1,
                "removed": False,
            }
            self.write_json(os.path.join(repo, state_path), {"odh": old_provenance})
            self.commit_all(repo, "main docs with old provenance")

            self.git(repo, "checkout", "-b", "snapshots")
            self.write_json(
                os.path.join(repo, "output", "v1", "odh", "v1", "snapshot-metadata.json"),
                {"platform": "odh", **old_provenance},
            )
            corrupt_path = os.path.join(
                repo, "output", "v2", "odh", "v2", "snapshot-metadata.json"
            )
            os.makedirs(os.path.dirname(corrupt_path))
            with open(corrupt_path, "w") as f:
                f.write('{"platform":"odh"')
            self.commit_all(repo, "corrupt latest snapshot metadata")
            self.git(repo, "checkout", "main")

            self.assertFalse(sync_docs(repo, None, "snapshots", state_path))

    def test_update_records_enabled_platform_when_snapshot_metadata_is_unreadable(self):
        with tempfile.TemporaryDirectory() as tmp:
            repo = os.path.join(tmp, "repo")
            self.init_repo(repo)
            state_path = ".generated-platform-provenance.json"
            self.write_json(
                os.path.join(repo, state_path),
                {"odh": {"version": "v1", "workflow_run_number": 18,
                          "workflow_run_attempt": 1, "removed": False}},
            )
            self.commit_all(repo, "main provenance")

            self.git(repo, "checkout", "-b", "snapshots")
            corrupt = os.path.join(
                repo, "output", "v2", "rhoai", "v2", "snapshot-metadata.json"
            )
            os.makedirs(os.path.dirname(corrupt))
            with open(corrupt, "w") as f:
                f.write('{"platform":"rhoai"')
            self.commit_all(repo, "corrupt rhoai snapshot metadata")
            self.git(repo, "checkout", "main")

            self.git(repo, "checkout", "-b", "generated")
            self.write_json(
                os.path.join(repo, state_path),
                {"rhoai": {"version": "v2", "workflow_run_number": 20,
                           "workflow_run_attempt": 1, "removed": False}},
            )
            self.commit_all(repo, "generated rhoai provenance")
            self.git(repo, "checkout", "main")

            self.write_json(
                os.path.join(repo, "platform-output", "rhoai", "v2", "snapshot-metadata.json"),
                {"version": "v2", "workflow_run_number": 22,
                 "workflow_run_attempt": 1, "removed": False},
            )
            output_file = os.path.join(tmp, "github-output")
            result = subprocess.run(
                [
                    sys.executable,
                    SCRIPT,
                    "update",
                    repo,
                    state_path,
                    "platform-output",
                    "v2",
                    "22",
                    "1",
                    "rhoai",
                    "",
                    "generated",
                    "snapshots",
                ],
                cwd=repo,
                env=dict(os.environ, GITHUB_OUTPUT=output_file),
                capture_output=True,
                text=True,
            )

            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
            state = self.read_json(os.path.join(repo, state_path))
            self.assertEqual(state["odh"]["workflow_run_number"], 18)
            self.assertEqual(state["rhoai"]["workflow_run_number"], 22)
            with open(output_file) as f:
                self.assertIn("docs-stale-platforms=odh", f.read())

    def test_legacy_snapshot_rejects_docs_generated_after_publication(self):
        snapshot = {
            "version": "2026-11-12",
            "workflow_run_number": 0,
            "workflow_run_attempt": 0,
            "timestamp": "2026-11-12T05:00:00Z",
            "analyzer_version": "legacy",
            "removed": False,
        }
        main = {
            "odh": {
                "version": "2026-11-12",
                "workflow_run_number": 0,
                "workflow_run_attempt": 0,
                "timestamp": "2026-11-12T05:00:00Z",
                "removed": False,
            }
        }
        generated = {
            "odh": {
                "version": "2026-11-12",
                "workflow_run_number": 0,
                "workflow_run_attempt": 0,
                "timestamp": "2026-11-12T06:00:00Z",
                "removed": False,
            }
        }
        self.assertEqual(choose_source(main, generated, snapshot, "odh"), "main")

    def test_legacy_snapshot_rejects_older_main_docs(self):
        snapshot = {
            "version": "2026-11-12",
            "workflow_run_number": 0,
            "workflow_run_attempt": 0,
            "timestamp": "2026-11-12T05:00:00Z",
            "removed": False,
        }
        main = {
            "odh": {
                "version": "2026-11-12",
                "timestamp": "2026-11-12T04:59:00Z",
                "removed": False,
            }
        }
        self.assertEqual(choose_source(main, {}, snapshot, "odh"), "stale")

    def test_legacy_snapshot_with_only_unpublished_docs_is_stale(self):
        snapshot = {
            "version": "2026-11-12",
            "workflow_run_number": 0,
            "workflow_run_attempt": 0,
            "timestamp": "2026-11-12T05:00:00Z",
            "removed": False,
        }
        generated = {
            "odh": {
                "version": "2026-11-12",
                "workflow_run_number": 0,
                "workflow_run_attempt": 0,
                "timestamp": "2026-11-12T06:00:00Z",
                "removed": False,
            }
        }
        self.assertEqual(choose_source({}, generated, snapshot, "odh"), "stale")

    def test_generated_docs_are_not_approved_without_a_published_snapshot(self):
        main = {"odh": {"version": "v1", "workflow_run_number": 10}}
        generated = {"odh": {"version": "v2", "workflow_run_number": 11}}
        self.assertEqual(choose_source(main, generated, None, "odh"), "main")


if __name__ == "__main__":
    unittest.main()
