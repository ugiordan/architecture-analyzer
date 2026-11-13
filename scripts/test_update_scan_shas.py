#!/usr/bin/env python3
"""Regression tests for scan-state generation from complete analysis results."""

import json
import itertools
import os
import subprocess
import sys
import tempfile
import unittest
from datetime import datetime, timezone


SCRIPT = os.path.join(os.path.dirname(__file__), "update-scan-shas.py")
MERGE_SCRIPT = os.path.join(os.path.dirname(__file__), "merge-scan-shas.py")


class UpdateScanShasTest(unittest.TestCase):
    def write_result(
        self,
        root,
        repo,
        *,
        security=True,
        failed=False,
        version=None,
        commit_sha="abc123",
        analyzer_version="v1",
        extracted_at="2026-11-11T05:00:00Z",
    ):
        result_dir = os.path.join(root, repo)
        if version:
            result_dir = os.path.join(result_dir, version)
        os.makedirs(result_dir, exist_ok=True)
        with open(os.path.join(result_dir, "component-architecture.json"), "w") as f:
            json.dump(
                {
                    "repo": repo.replace("/", "/", 1),
                    "commit_sha": commit_sha,
                    "analyzer_version": analyzer_version,
                    "extracted_at": extracted_at,
                },
                f,
            )
        if security:
            with open(os.path.join(result_dir, "security-findings.json"), "w") as f:
                json.dump([], f)
        if failed:
            marker_dir = os.path.join(root, repo)
            os.makedirs(marker_dir, exist_ok=True)
            with open(os.path.join(marker_dir, "_failed"), "w") as f:
                f.write("2026-11-12T00:00:00Z\n")

    def run_update(self, root, out, *extra_args, env=None):
        subprocess.run(
            [sys.executable, SCRIPT, root, *extra_args, "-o", out],
            check=True,
            capture_output=True,
            text=True,
            env=env,
        )
        with open(out) as f:
            return json.load(f)

    def test_only_complete_successful_results_advance_state(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            os.makedirs(results)
            self.write_result(results, "org/complete")
            self.write_result(results, "org/no-security", security=False)
            self.write_result(results, "org/failed-versioned", failed=True, version="v2")

            state = self.run_update(results, os.path.join(tmp, "shas.json"))

            self.assertEqual(set(state), {"odh/org/complete"})
            self.assertEqual(state["odh/org/complete"]["commit_sha"], "abc123")

    def test_failed_platform_is_invalidated_and_not_readded(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "rhoai-results")
            os.makedirs(results)
            self.write_result(results, "org/partial")
            existing = os.path.join(tmp, "shas.json")
            with open(existing, "w") as f:
                json.dump(
                    {
                        "rhoai/org/partial": {"commit_sha": "old"},
                        "odh/org/healthy": {"commit_sha": "keep"},
                    },
                    f,
                )

            state = self.run_update(
                results, existing, "--invalidate-platforms", "rhoai"
            )

            self.assertNotIn("rhoai/org/partial", state)
            self.assertEqual(state["odh/org/healthy"]["commit_sha"], "keep")

    def test_failed_marker_removes_old_scan_state(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            failed = os.path.join(results, "org", "repo")
            os.makedirs(failed)
            open(os.path.join(failed, "_failed"), "w").close()
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump(
                    {"odh/org/repo": {"commit_sha": "old", "analyzer_version": "v1"}},
                    f,
                )

            state = self.run_update(results, output)

            self.assertNotIn("odh/org/repo", state)

    def test_findings_only_directory_invalidates_old_scan_state(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            result_dir = os.path.join(results, "org", "repo")
            os.makedirs(result_dir)
            with open(os.path.join(result_dir, "security-findings.json"), "w") as f:
                json.dump([], f)
            with open(os.path.join(result_dir, "_scan-provenance.json"), "w") as f:
                json.dump({"workflow_run_number": 21, "workflow_run_attempt": 1}, f)
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump({"odh/org/repo": {
                    "commit_sha": "run-20",
                    "analyzer_version": "v1",
                    "scanned_at": "2026-11-11T00:00:00Z",
                    "workflow_run_number": 20,
                    "workflow_run_attempt": 1,
                }}, f)

            state = self.run_update(results, output)

            self.assertNotIn("odh/org/repo", state)

    def test_legacy_unqualified_state_key_is_not_retained(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            os.makedirs(results)
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump({"org/repo": {"commit_sha": "legacy"}}, f)

            state = self.run_update(results, output)

            self.assertNotIn("org/repo", state)

    def test_successful_platform_scan_replaces_legacy_key(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/repo", commit_sha="current")
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump({"org/repo": {"commit_sha": "legacy"}}, f)

            state = self.run_update(results, output)

            self.assertNotIn("org/repo", state)
            self.assertEqual(state["odh/org/repo"]["commit_sha"], "current")

    def test_invalid_security_findings_do_not_advance_state(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/repo", security=False)
            findings = os.path.join(results, "org", "repo", "security-findings.json")
            with open(findings, "w") as f:
                f.write("not-json")
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump({"odh/org/repo": {"commit_sha": "old"}}, f)

            state = self.run_update(results, output)

            self.assertNotIn("odh/org/repo", state)

    def test_missing_commit_sha_invalidates_existing_state(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/repo", commit_sha="")
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump({"odh/org/repo": {"commit_sha": "old"}}, f)

            state = self.run_update(results, output)

            self.assertNotIn("odh/org/repo", state)

    def test_undecodable_security_findings_invalidates_existing_state(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/repo", security=False)
            findings = os.path.join(results, "org", "repo", "security-findings.json")
            with open(findings, "wb") as f:
                f.write(b"\xff\xfe")
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump({"odh/org/repo": {"commit_sha": "old"}}, f)

            state = self.run_update(results, output)

            self.assertNotIn("odh/org/repo", state)

    def test_non_string_scan_timestamp_does_not_advance_state(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/repo", extracted_at=42)

            state = self.run_update(results, os.path.join(tmp, "shas.json"))

            self.assertEqual(state, {})

    def test_carried_forward_result_does_not_claim_current_run(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/repo", commit_sha="carried")
            output = os.path.join(tmp, "shas.json")
            env = dict(os.environ, GITHUB_RUN_NUMBER="21", GITHUB_RUN_ATTEMPT="1")

            state = self.run_update(results, output, env=env)

            self.assertEqual(state["odh/org/repo"]["commit_sha"], "carried")
            self.assertNotIn("workflow_run_number", state["odh/org/repo"])

    def test_result_uses_its_artifact_run_provenance(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/repo", commit_sha="analyzed")
            provenance = os.path.join(results, "org", "repo", "_scan-provenance.json")
            with open(provenance, "w") as f:
                json.dump({"workflow_run_number": 19, "workflow_run_attempt": 2}, f)
            output = os.path.join(tmp, "shas.json")
            env = dict(os.environ, GITHUB_RUN_NUMBER="21", GITHUB_RUN_ATTEMPT="1")

            state = self.run_update(results, output, env=env)

            self.assertEqual(state["odh/org/repo"]["workflow_run_number"], 19)
            self.assertEqual(state["odh/org/repo"]["workflow_run_attempt"], 2)

    def test_invalid_timestamp_without_run_provenance_invalidates_existing_state(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/repo", commit_sha="new", extracted_at=42)
            architecture = os.path.join(results, "org", "repo", "component-architecture.json")
            now = datetime.now(timezone.utc)
            os.utime(architecture, (now.timestamp() + 60, now.timestamp() + 60))
            output = os.path.join(tmp, "shas.json")
            scanned_at = now.isoformat(timespec="microseconds").replace("+00:00", "Z")
            with open(output, "w") as f:
                json.dump(
                    {"odh/org/repo": {
                        "commit_sha": "old",
                        "analyzer_version": "v1",
                        "scanned_at": scanned_at,
                    }},
                    f,
                )

            state = self.run_update(results, output)

            self.assertNotIn("odh/org/repo", state)

    def test_invalid_timestamp_with_newer_artifact_run_invalidates_old_state(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/repo", commit_sha="invalid-time", extracted_at=42)
            provenance = os.path.join(results, "org", "repo", "_scan-provenance.json")
            with open(provenance, "w") as f:
                json.dump({"workflow_run_number": 22, "workflow_run_attempt": 1}, f)
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump({"odh/org/repo": {
                    "commit_sha": "run-20",
                    "analyzer_version": "v1",
                    "scanned_at": "2026-11-12T12:00:00Z",
                    "workflow_run_number": 20,
                    "workflow_run_attempt": 1,
                }}, f)

            state = self.run_update(results, output)

            self.assertNotIn("odh/org/repo", state)

    def test_naive_extracted_at_is_rejected_and_invalidates_older_run(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(
                results,
                "org/repo",
                commit_sha="naive-time",
                extracted_at="2026-11-12T10:00:00",
            )
            provenance = os.path.join(results, "org", "repo", "_scan-provenance.json")
            with open(provenance, "w") as f:
                json.dump({"workflow_run_number": 22, "workflow_run_attempt": 1}, f)
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump({"odh/org/repo": {
                    "commit_sha": "run-20",
                    "analyzer_version": "v1",
                    "scanned_at": "2026-11-12T12:00:00Z",
                    "workflow_run_number": 20,
                    "workflow_run_attempt": 1,
                }}, f)

            state = self.run_update(results, output)

            self.assertNotIn("odh/org/repo", state)

    def test_failed_platform_whitespace_format_writes_invalidation_sidecar(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "rhoai-results")
            os.makedirs(results)
            output = os.path.join(tmp, "shas.json")
            sidecar = os.path.join(tmp, "invalidations.json")
            with open(output, "w") as f:
                json.dump(
                    {
                        "odh/org/one": {"commit_sha": "one"},
                        "rhoai/org/two": {"commit_sha": "two"},
                    },
                    f,
                )
            env = dict(os.environ, GITHUB_RUN_NUMBER="17", GITHUB_RUN_ATTEMPT="2")

            state = self.run_update(
                results,
                output,
                "--invalidate-platforms",
                "odh rhoai",
                "--invalidations-output",
                sidecar,
                env=env,
            )

            self.assertEqual(state, {})
            with open(sidecar) as f:
                invalidations = json.load(f)["invalidations"]
            self.assertEqual({event["prefix"] for event in invalidations}, {"odh/", "rhoai/"})
            self.assertTrue(all(event["workflow_run_number"] == 17 for event in invalidations))

    def test_invalidation_options_reject_following_flags_as_values(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            os.makedirs(results)
            output = os.path.join(tmp, "shas.json")
            invalidations = os.path.join(tmp, "invalidations.json")
            malformed = (
                ["--invalidate-platforms", "--invalidations-output", invalidations],
                ["--invalidations-output", "--invalidate-platforms", "odh"],
            )

            for args in malformed:
                with self.subTest(args=args):
                    result = subprocess.run(
                        [sys.executable, SCRIPT, results, *args, "-o", output],
                        capture_output=True,
                        text=True,
                    )
                    self.assertNotEqual(result.returncode, 0)

    def test_failure_marker_is_written_to_invalidation_sidecar(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            failed_dir = os.path.join(results, "org", "repo")
            os.makedirs(failed_dir)
            with open(os.path.join(failed_dir, "_failed"), "w") as f:
                f.write("2026-11-12T00:00:00Z\n")
            output = os.path.join(tmp, "shas.json")
            sidecar = os.path.join(tmp, "invalidations.json")
            with open(output, "w") as f:
                json.dump({"odh/org/repo": {"commit_sha": "old"}}, f)

            self.run_update(results, output, "--invalidations-output", sidecar)

            with open(sidecar) as f:
                invalidations = json.load(f)["invalidations"]
            self.assertIn(
                {"prefix": "odh/org/repo", "invalidated_at": "2026-11-12T00:00:00.000000Z"},
                invalidations,
            )

    def test_failure_marker_sidecar_preserves_originating_run(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            failed_dir = os.path.join(results, "org", "repo")
            os.makedirs(failed_dir)
            with open(os.path.join(failed_dir, "_failed"), "w") as f:
                json.dump(
                    {
                        "timestamp": "2026-11-12T00:00:00Z",
                        "workflow_run_number": 20,
                        "workflow_run_attempt": 1,
                    },
                    f,
                )
            output = os.path.join(tmp, "shas.json")
            sidecar = os.path.join(tmp, "invalidations.json")
            with open(output, "w") as f:
                json.dump({"odh/org/repo": {"commit_sha": "old"}}, f)
            env = dict(os.environ, GITHUB_RUN_NUMBER="21", GITHUB_RUN_ATTEMPT="1")

            self.run_update(results, output, "--invalidations-output", sidecar, env=env)

            with open(sidecar) as f:
                invalidations = json.load(f)["invalidations"]
            event = next(event for event in invalidations if event["prefix"] == "odh/org/repo")
            self.assertEqual(event["workflow_run_number"], 20)
            self.assertEqual(event["workflow_run_attempt"], 1)

    def test_invalid_marker_timestamp_keeps_valid_run_provenance(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            failed_dir = os.path.join(results, "org", "repo")
            os.makedirs(failed_dir)
            with open(os.path.join(failed_dir, "_failed"), "w") as f:
                json.dump(
                    {
                        "timestamp": "invalid",
                        "workflow_run_number": 20,
                        "workflow_run_attempt": 1,
                    },
                    f,
                )
            output = os.path.join(tmp, "shas.json")
            sidecar = os.path.join(tmp, "invalidations.json")
            with open(output, "w") as f:
                json.dump({"odh/org/repo": {"commit_sha": "old"}}, f)
            env = dict(os.environ, GITHUB_RUN_NUMBER="21", GITHUB_RUN_ATTEMPT="1")

            self.run_update(results, output, "--invalidations-output", sidecar, env=env)

            with open(sidecar) as f:
                invalidations = json.load(f)["invalidations"]
            event = next(event for event in invalidations if event["prefix"] == "odh/org/repo")
            self.assertEqual(event["workflow_run_number"], 20)

    def test_failed_marker_prevents_duplicate_incomplete_event_invalidating_newer_success(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            failed_dir = os.path.join(results, "org", "repo")
            incomplete_dir = os.path.join(failed_dir, "v2")
            os.makedirs(incomplete_dir)
            with open(os.path.join(failed_dir, "_failed"), "w") as f:
                json.dump(
                    {
                        "timestamp": "2026-11-12T00:00:00Z",
                        "workflow_run_number": 20,
                        "workflow_run_attempt": 1,
                    },
                    f,
                )
            architecture_path = os.path.join(incomplete_dir, "component-architecture.json")
            with open(architecture_path, "w") as f:
                f.write('{"repo":"org/repo"')
            future_mtime = datetime(2026, 11, 14, tzinfo=timezone.utc).timestamp()
            os.utime(architecture_path, (future_mtime, future_mtime))

            key = "odh/org/repo"
            output = os.path.join(tmp, "shas.json")
            sidecar = os.path.join(tmp, "invalidations.json")
            with open(output, "w") as f:
                json.dump(
                    {
                        key: {
                            "commit_sha": "run-21-success",
                            "analyzer_version": "v1",
                            "scanned_at": "2026-11-13T00:00:00Z",
                            "workflow_run_number": 21,
                            "workflow_run_attempt": 1,
                        }
                    },
                    f,
                )
            env = dict(os.environ, GITHUB_RUN_NUMBER="22", GITHUB_RUN_ATTEMPT="1")

            state = self.run_update(
                results, output, "--invalidations-output", sidecar, env=env
            )

            self.assertIn(key, state)
            self.assertEqual(state[key]["workflow_run_number"], 21)
            with open(sidecar) as f:
                invalidations = json.load(f)["invalidations"]
            self.assertEqual(len(invalidations), 1)
            self.assertEqual(invalidations[0]["workflow_run_number"], 20)

    def test_newer_child_failure_event_invalidates_success_after_older_marker(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            failed_dir = os.path.join(results, "org", "repo")
            incomplete_dir = os.path.join(failed_dir, "v2")
            os.makedirs(incomplete_dir)
            with open(os.path.join(failed_dir, "_failed"), "w") as f:
                json.dump(
                    {
                        "timestamp": "2026-11-12T00:00:00Z",
                        "workflow_run_number": 20,
                        "workflow_run_attempt": 1,
                    },
                    f,
                )
            with open(os.path.join(incomplete_dir, "_scan-provenance.json"), "w") as f:
                json.dump(
                    {"workflow_run_number": 22, "workflow_run_attempt": 1}, f
                )
            architecture_path = os.path.join(incomplete_dir, "component-architecture.json")
            with open(architecture_path, "w") as f:
                f.write('{"repo":"org/repo"')
            future_mtime = datetime(2026, 11, 14, tzinfo=timezone.utc).timestamp()
            os.utime(architecture_path, (future_mtime, future_mtime))

            key = "odh/org/repo"
            output = os.path.join(tmp, "shas.json")
            sidecar = os.path.join(tmp, "invalidations.json")
            with open(output, "w") as f:
                json.dump(
                    {
                        key: {
                            "commit_sha": "run-21-success",
                            "analyzer_version": "v1",
                            "scanned_at": "2026-11-13T00:00:00Z",
                            "workflow_run_number": 21,
                            "workflow_run_attempt": 1,
                        }
                    },
                    f,
                )

            state = self.run_update(results, output, "--invalidations-output", sidecar)

            self.assertNotIn(key, state)
            with open(sidecar) as f:
                invalidations = json.load(f)["invalidations"]
            self.assertEqual(
                {event["workflow_run_number"] for event in invalidations}, {20, 22}
            )

    def test_security_only_child_beneath_marker_does_not_duplicate_invalidation(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            failed_dir = os.path.join(results, "org", "repo")
            security_dir = os.path.join(failed_dir, "v2")
            os.makedirs(security_dir)
            with open(os.path.join(failed_dir, "_failed"), "w") as f:
                json.dump(
                    {
                        "timestamp": "2026-11-12T00:00:00Z",
                        "workflow_run_number": 20,
                        "workflow_run_attempt": 1,
                    },
                    f,
                )
            findings_path = os.path.join(security_dir, "security-findings.json")
            with open(findings_path, "w") as f:
                json.dump([], f)
            future_mtime = datetime(2026, 11, 14, tzinfo=timezone.utc).timestamp()
            os.utime(findings_path, (future_mtime, future_mtime))

            key = "odh/org/repo"
            output = os.path.join(tmp, "shas.json")
            sidecar = os.path.join(tmp, "invalidations.json")
            with open(output, "w") as f:
                json.dump(
                    {
                        key: {
                            "commit_sha": "run-21-success",
                            "analyzer_version": "v1",
                            "scanned_at": "2026-11-13T00:00:00Z",
                            "workflow_run_number": 21,
                            "workflow_run_attempt": 1,
                        }
                    },
                    f,
                )

            state = self.run_update(results, output, "--invalidations-output", sidecar)

            self.assertIn(key, state)
            with open(sidecar) as f:
                invalidations = json.load(f)["invalidations"]
            self.assertEqual(len(invalidations), 1)
            self.assertEqual(invalidations[0]["workflow_run_number"], 20)

    def test_historical_incomplete_result_does_not_inherit_current_run(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(
                results,
                "org/repo",
                version="older-incomplete",
                commit_sha="older-sha",
                security=False,
                extracted_at="2026-11-09T00:00:00Z",
            )
            self.write_result(
                results,
                "org/repo",
                version="valid-current",
                commit_sha="run-20-success",
                extracted_at="2026-11-11T00:00:00Z",
            )
            provenance_path = os.path.join(results, "org", "repo", "_scan-provenance.json")
            with open(provenance_path, "w") as f:
                json.dump(
                    {
                        "workflow_run_number": 19,
                        "workflow_run_attempt": 1,
                    },
                    f,
                )
            key = "odh/org/repo"
            output = os.path.join(tmp, "ours.json")
            main_state = os.path.join(tmp, "main.json")
            run_invalidations = os.path.join(tmp, "run-invalidations.json")
            main_invalidations = os.path.join(tmp, "main-invalidations.json")
            merged_path = os.path.join(tmp, "merged.json")
            output_invalidations = os.path.join(tmp, "merged-invalidations.json")
            initial_state = {key: {
                "commit_sha": "run-20-success",
                "analyzer_version": "v1",
                "scanned_at": "2026-11-11T00:00:00Z",
                "workflow_run_number": 20,
                "workflow_run_attempt": 1,
            }}
            for path in (output, main_state):
                with open(path, "w") as f:
                    json.dump(initial_state, f)
            with open(main_invalidations, "w") as f:
                json.dump({"invalidations": []}, f)
            env = dict(os.environ, GITHUB_RUN_NUMBER="21", GITHUB_RUN_ATTEMPT="1")

            self.run_update(
                results,
                output,
                "--invalidations-output",
                run_invalidations,
                env=env,
            )
            subprocess.run(
                [sys.executable, MERGE_SCRIPT, output, main_state, merged_path,
                 run_invalidations, main_invalidations, output_invalidations],
                check=True,
                capture_output=True,
                text=True,
            )

            with open(merged_path) as f:
                merged = json.load(f)
            with open(output_invalidations) as f:
                invalidations = json.load(f)["invalidations"]
            self.assertEqual(merged[key]["workflow_run_number"], 20)
            event = next(event for event in invalidations if event["prefix"] == key)
            self.assertEqual(event["workflow_run_number"], 19)

    def test_failed_version_excludes_sibling_successful_versions(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            os.makedirs(results)
            self.write_result(
                results, "org/repo", version="old", commit_sha="old-sha",
                extracted_at="2026-11-10T00:00:00Z",
            )
            failed_dir = os.path.join(results, "org", "repo", "new")
            os.makedirs(failed_dir)
            with open(os.path.join(failed_dir, "_failed"), "w") as f:
                f.write("2026-11-12T00:00:00Z\n")
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump({"odh/org/repo": {"commit_sha": "old-sha"}}, f)

            state = self.run_update(results, output)

            self.assertNotIn("odh/org/repo", state)

    def test_later_successful_version_restores_state_after_older_failure(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(
                results,
                "org/repo",
                version="old",
                commit_sha="old-sha",
                extracted_at="2026-11-09T00:00:00Z",
            )
            failed_dir = os.path.join(results, "org", "repo", "failed")
            os.makedirs(failed_dir)
            with open(os.path.join(failed_dir, "_failed"), "w") as f:
                f.write("2026-11-10T00:00:00Z\n")
            self.write_result(
                results,
                "org/repo",
                version="new",
                commit_sha="new-sha",
                extracted_at="2026-11-11T00:00:00Z",
            )
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump(
                    {"odh/org/repo": {
                        "commit_sha": "old-sha",
                        "analyzer_version": "v1",
                        "scanned_at": "2026-11-09T00:00:00Z",
                    }},
                    f,
                )

            state = self.run_update(results, output)

            self.assertEqual(state["odh/org/repo"]["commit_sha"], "new-sha")

    def test_org_failure_marker_invalidates_all_org_repos(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/one")
            self.write_result(results, "org/two")
            open(os.path.join(results, "org", "_failed"), "w").close()
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump(
                    {
                        "odh/org/one": {"commit_sha": "old1"},
                        "odh/org/two": {"commit_sha": "old2"},
                    },
                    f,
                )

            state = self.run_update(results, output)

            self.assertEqual(state, {})

    def test_failed_repo_prefix_does_not_match_similarly_named_repo(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/api-v2", commit_sha="keep")
            failed = os.path.join(results, "org", "api")
            os.makedirs(failed)
            open(os.path.join(failed, "_failed"), "w").close()
            output = os.path.join(tmp, "shas.json")

            state = self.run_update(results, output)

            self.assertEqual(state["odh/org/api-v2"]["commit_sha"], "keep")

    def test_newest_duplicate_versioned_result_wins_in_either_order(self):
        for first, second in (("old", "new"), ("new", "old")):
            with self.subTest(order=(first, second)), tempfile.TemporaryDirectory() as tmp:
                results = os.path.join(tmp, "odh-results")
                os.makedirs(results)
                values = {
                    "old": ("aaa", "2026-11-11T05:00:00Z"),
                    "new": ("bbb", "2026-11-11T06:00:00Z"),
                }
                for version in (first, second):
                    sha, timestamp = values[version]
                    self.write_result(
                        results,
                        "org/repo",
                        version=version,
                        commit_sha=sha,
                        extracted_at=timestamp,
                    )
                state = self.run_update(results, os.path.join(tmp, "shas.json"))
                self.assertEqual(state["odh/org/repo"]["commit_sha"], "bbb")

    def test_mixed_provenance_result_order_is_independent_of_directory_order(self):
        for order in (("tagged", "plain"), ("plain", "tagged")):
            with self.subTest(order=order), tempfile.TemporaryDirectory() as tmp:
                results = os.path.join(tmp, "odh-results")
                values = {
                    "tagged": ("tagged-sha", "2026-11-11T11:00:00Z"),
                    "plain": ("plain-sha", "2026-11-11T12:00:00Z"),
                }
                for version in order:
                    sha, timestamp = values[version]
                    self.write_result(
                        results,
                        "org/repo",
                        version=version,
                        commit_sha=sha,
                        extracted_at=timestamp,
                    )
                    if version == "tagged":
                        provenance = os.path.join(results, "org", "repo", version, "_scan-provenance.json")
                        with open(provenance, "w") as f:
                            json.dump({"workflow_run_number": 20, "workflow_run_attempt": 1}, f)

                state = self.run_update(results, os.path.join(tmp, "shas.json"))

                self.assertEqual(state["odh/org/repo"]["commit_sha"], "tagged-sha")

    def test_three_mixed_provenance_results_are_order_independent(self):
        results_data = {
            "run-20": ("run-20-sha", "2026-11-11T12:00:00Z", (20, 1)),
            "run-21": ("run-21-sha", "2026-11-11T10:00:00Z", (21, 1)),
            "untagged": ("untagged-sha", "2026-11-11T11:00:00Z", None),
        }
        for order in itertools.permutations(results_data):
            with self.subTest(order=order), tempfile.TemporaryDirectory() as tmp:
                results = os.path.join(tmp, "odh-results")
                for version in order:
                    sha, timestamp, run = results_data[version]
                    self.write_result(
                        results,
                        "org/repo",
                        version=version,
                        commit_sha=sha,
                        extracted_at=timestamp,
                    )
                    if run is not None:
                        provenance = os.path.join(results, "org", "repo", version, "_scan-provenance.json")
                        with open(provenance, "w") as f:
                            json.dump({
                                "workflow_run_number": run[0],
                                "workflow_run_attempt": run[1],
                            }, f)

                state = self.run_update(results, os.path.join(tmp, "shas.json"))

                self.assertEqual(state["odh/org/repo"]["commit_sha"], "run-21-sha")

    def test_newer_run_wins_over_later_extraction_timestamp(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/repo", commit_sha="run-21", extracted_at="2026-11-11T11:00:00Z")
            provenance = os.path.join(results, "org", "repo", "_scan-provenance.json")
            with open(provenance, "w") as f:
                json.dump({"workflow_run_number": 21, "workflow_run_attempt": 1}, f)
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump({"odh/org/repo": {
                    "commit_sha": "run-20",
                    "analyzer_version": "v1",
                    "scanned_at": "2026-11-11T12:00:00Z",
                    "workflow_run_number": 20,
                    "workflow_run_attempt": 1,
                }}, f)

            state = self.run_update(results, output)

            self.assertEqual(state["odh/org/repo"]["commit_sha"], "run-21")
            self.assertEqual(state["odh/org/repo"]["workflow_run_number"], 21)

    def test_newer_failure_run_invalidates_later_timestamped_old_state(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            failed_dir = os.path.join(results, "org", "repo")
            os.makedirs(failed_dir)
            with open(os.path.join(failed_dir, "_failed"), "w") as f:
                json.dump({
                    "timestamp": "2026-11-11T11:00:00Z",
                    "workflow_run_number": 21,
                    "workflow_run_attempt": 1,
                }, f)
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump({"odh/org/repo": {
                    "commit_sha": "run-20",
                    "analyzer_version": "v1",
                    "scanned_at": "2026-11-11T12:00:00Z",
                    "workflow_run_number": 20,
                    "workflow_run_attempt": 1,
                }}, f)

            state = self.run_update(results, output)

            self.assertNotIn("odh/org/repo", state)

    def test_equal_timestamps_use_deterministic_sha_tiebreaker(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            os.makedirs(results)
            timestamp = "2026-11-11T05:00:00Z"
            for version, sha in (("one", "aaa"), ("two", "bbb")):
                self.write_result(
                    results,
                    "org/repo",
                    version=version,
                    commit_sha=sha,
                    extracted_at=timestamp,
                )
            state = self.run_update(results, os.path.join(tmp, "shas.json"))
            self.assertEqual(state["odh/org/repo"]["commit_sha"], "bbb")

    def test_malformed_existing_rank_fields_do_not_abort_update(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            os.makedirs(results)
            timestamp = "2026-11-11T05:00:00Z"
            self.write_result(
                results, "org/repo", version="new", commit_sha="abc",
                analyzer_version="v1", extracted_at=timestamp,
            )
            output = os.path.join(tmp, "shas.json")
            with open(output, "w") as f:
                json.dump(
                    {"odh/org/repo": {
                        "scanned_at": timestamp,
                        "analyzer_version": None,
                        "commit_sha": None,
                    }},
                    f,
                )

            state = self.run_update(results, output)

            self.assertEqual(state["odh/org/repo"]["commit_sha"], "abc")

    def test_non_string_analyzer_version_uses_fallback_for_ranking(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            os.makedirs(results)
            self.write_result(
                results,
                "org/repo",
                version="null-version",
                commit_sha="aaa",
                analyzer_version=None,
                extracted_at="2026-11-11T05:00:00Z",
            )
            self.write_result(
                results,
                "org/repo",
                version="string-version",
                commit_sha="bbb",
                analyzer_version="v1",
                extracted_at="2026-11-11T05:00:00Z",
            )

            state = self.run_update(results, os.path.join(tmp, "shas.json"))

            self.assertEqual(state["odh/org/repo"]["analyzer_version"], "v1")

    def test_empty_analyzer_version_uses_dev_fallback(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            self.write_result(results, "org/repo", analyzer_version="")

            state = self.run_update(results, os.path.join(tmp, "shas.json"))

            self.assertEqual(state["odh/org/repo"]["analyzer_version"], "dev")

    def test_mixed_timestamp_precision_is_compared_as_instants(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            os.makedirs(results)
            self.write_result(
                results,
                "org/repo",
                version="whole-second",
                commit_sha="aaa",
                extracted_at="2026-11-11T05:00:00Z",
            )
            self.write_result(
                results,
                "org/repo",
                version="fractional-second",
                commit_sha="bbb",
                extracted_at="2026-11-11T05:00:00.1Z",
            )

            state = self.run_update(results, os.path.join(tmp, "shas.json"))

            self.assertEqual(state["odh/org/repo"]["commit_sha"], "bbb")


class MergeScanShasTest(unittest.TestCase):
    def merge(self, ours, main, run_invalidations=None, main_invalidations=None):
        with tempfile.TemporaryDirectory() as tmp:
            ours_path = os.path.join(tmp, "ours.json")
            main_path = os.path.join(tmp, "main.json")
            output_path = os.path.join(tmp, "merged.json")
            run_invalidations_path = os.path.join(tmp, "run-invalidations.json")
            main_invalidations_path = os.path.join(tmp, "main-invalidations.json")
            output_invalidations_path = os.path.join(tmp, "output-invalidations.json")
            for path, value in ((ours_path, ours), (main_path, main)):
                with open(path, "w") as f:
                    json.dump(value, f)
            command = [sys.executable, MERGE_SCRIPT, ours_path, main_path, output_path]
            if run_invalidations is not None or main_invalidations is not None:
                for path, events in (
                    (run_invalidations_path, run_invalidations or []),
                    (main_invalidations_path, main_invalidations or []),
                ):
                    with open(path, "w") as f:
                        json.dump({"invalidations": events}, f)
                command.extend(
                    [run_invalidations_path, main_invalidations_path, output_invalidations_path]
                )
            subprocess.run(command, check=True, capture_output=True, text=True)
            with open(output_path) as f:
                state = json.load(f)
            if os.path.exists(output_invalidations_path):
                with open(output_invalidations_path) as f:
                    events = json.load(f)["invalidations"]
                return state, events
            return state, None

    def test_merge_compares_fractional_and_offset_timestamps_as_instants(self):
        key = "odh/org/repo"
        fractional, _ = self.merge(
            {key: {"commit_sha": "fractional", "scanned_at": "2026-11-11T05:00:00.500Z"}},
            {key: {"commit_sha": "whole-second", "scanned_at": "2026-11-11T05:00:00Z"}},
        )
        self.assertEqual(fractional[key]["commit_sha"], "fractional")

        offset, _ = self.merge(
            {key: {"commit_sha": "older", "scanned_at": "2026-11-11T05:00:00+02:00"}},
            {key: {"commit_sha": "newer", "scanned_at": "2026-11-11T04:00:00Z"}},
        )
        self.assertEqual(offset[key]["commit_sha"], "newer")

    def test_merge_uses_analyzer_and_sha_tiebreakers(self):
        key = "odh/org/repo"
        current = {key: {
            "commit_sha": "bbb",
            "analyzer_version": "v1",
            "scanned_at": "2026-11-11T05:00:00Z",
        }}
        newer_analyzer = self.merge(
            {key: {
                "commit_sha": "aaa",
                "analyzer_version": "v2",
                "scanned_at": "2026-11-11T05:00:00Z",
            }},
            current,
        )
        self.assertEqual(newer_analyzer[0][key]["analyzer_version"], "v2")

        lower_sha = self.merge(
            {key: {
                "commit_sha": "aaa",
                "analyzer_version": "v1",
                "scanned_at": "2026-11-11T05:00:00Z",
            }},
            current,
        )
        self.assertEqual(lower_sha[0][key]["commit_sha"], "bbb")

    def test_invalidation_removes_only_entries_not_newer_than_event(self):
        state, _ = self.merge(
            {},
            {
                "odh/org/old": {"commit_sha": "old", "scanned_at": "2026-11-11T00:00:00Z"},
                "odh/org/new": {"commit_sha": "new", "scanned_at": "2026-11-13T00:00:00Z"},
            },
            run_invalidations=[{"prefix": "odh/org/", "invalidated_at": "2026-11-12T00:00:00Z"}],
        )
        self.assertNotIn("odh/org/old", state)
        self.assertEqual(state["odh/org/new"]["commit_sha"], "new")

    def test_later_workflow_run_survives_older_invalidation(self):
        key = "odh/org/repo"
        state, _ = self.merge(
            {},
            {key: {
                "commit_sha": "newer-run",
                "scanned_at": "2026-11-10T00:00:00Z",
                "workflow_run_number": 18,
                "workflow_run_attempt": 1,
            }},
            run_invalidations=[{
                "prefix": key,
                "invalidated_at": "2026-11-12T00:00:00Z",
                "workflow_run_number": 17,
                "workflow_run_attempt": 3,
            }],
        )
        self.assertEqual(state[key]["commit_sha"], "newer-run")

    def test_mixed_provenance_merge_prefers_tagged_entries(self):
        key = "odh/org/repo"
        tagged = {
            "commit_sha": "tagged",
            "analyzer_version": "v1",
            "scanned_at": "2026-11-11T11:00:00Z",
            "workflow_run_number": 20,
            "workflow_run_attempt": 1,
        }
        plain = {
            "commit_sha": "plain",
            "analyzer_version": "v1",
            "scanned_at": "2026-11-11T12:00:00Z",
        }
        newer_plain, _ = self.merge({key: plain}, {key: tagged})
        older_plain, _ = self.merge({key: tagged}, {key: plain})
        self.assertEqual(newer_plain[key]["commit_sha"], "tagged")
        self.assertEqual(older_plain[key]["commit_sha"], "tagged")

        tied_tagged, _ = self.merge(
            {key: {**tagged, "scanned_at": "2026-11-11T12:00:00Z"}},
            {key: plain},
        )
        tied_plain, _ = self.merge(
            {key: plain},
            {key: {**tagged, "scanned_at": "2026-11-11T12:00:00Z"}},
        )
        self.assertEqual(tied_tagged[key]["commit_sha"], "tagged")
        self.assertEqual(tied_plain[key]["commit_sha"], "tagged")

    def test_three_mixed_provenance_entries_merge_in_any_order(self):
        key = "odh/org/repo"
        entries = [
            {
                "commit_sha": "run-20-sha",
                "scanned_at": "2026-11-11T12:00:00Z",
                "workflow_run_number": 20,
                "workflow_run_attempt": 1,
            },
            {
                "commit_sha": "run-21-sha",
                "scanned_at": "2026-11-11T10:00:00Z",
                "workflow_run_number": 21,
                "workflow_run_attempt": 1,
            },
            {"commit_sha": "untagged-sha", "scanned_at": "2026-11-11T11:00:00Z"},
        ]
        for order in itertools.permutations(entries):
            with self.subTest(order=[entry["commit_sha"] for entry in order]):
                state = {}
                for entry in order:
                    state, _ = self.merge({key: entry}, state)
                self.assertEqual(state[key]["commit_sha"], "run-21-sha")

    def test_success_later_in_same_run_survives_failure_invalidation(self):
        key = "odh/org/repo"
        state, _ = self.merge(
            {},
            {key: {
                "commit_sha": "successful-retry",
                "scanned_at": "2026-11-12T10:30:00Z",
                "workflow_run_number": 17,
                "workflow_run_attempt": 1,
            }},
            run_invalidations=[{
                "prefix": key,
                "invalidated_at": "2026-11-12T10:00:00Z",
                "workflow_run_number": 17,
                "workflow_run_attempt": 1,
            }],
        )
        self.assertEqual(state[key]["commit_sha"], "successful-retry")

    def test_persisted_tombstone_blocks_older_run_restoration(self):
        key = "odh/org/repo"
        event = {
            "prefix": key,
            "invalidated_at": "2026-11-12T10:00:00Z",
            "workflow_run_number": 20,
            "workflow_run_attempt": 1,
        }
        state, tombstones = self.merge(
            {},
            {key: {
                "commit_sha": "stale",
                "scanned_at": "2026-11-11T00:00:00Z",
                "workflow_run_number": 19,
                "workflow_run_attempt": 1,
            }},
            run_invalidations=[event],
            main_invalidations=[],
        )
        self.assertNotIn(key, state)
        self.assertEqual(tombstones, [event])

        restored, persisted = self.merge(
            {key: {
                "commit_sha": "late-old-run",
                "scanned_at": "2026-11-11T00:00:00Z",
                "workflow_run_number": 19,
                "workflow_run_attempt": 2,
            }},
            state,
            run_invalidations=[],
            main_invalidations=tombstones,
        )
        self.assertNotIn(key, restored)
        self.assertEqual(persisted, tombstones)

    def test_historical_failure_is_not_reassigned_to_current_run(self):
        with tempfile.TemporaryDirectory() as tmp:
            results = os.path.join(tmp, "odh-results")
            failed_dir = os.path.join(results, "org", "repo")
            os.makedirs(failed_dir)
            with open(os.path.join(failed_dir, "_failed"), "w") as f:
                json.dump(
                    {
                        "timestamp": "2026-11-10T00:00:00Z",
                        "workflow_run_number": 10,
                        "workflow_run_attempt": 1,
                    },
                    f,
                )
            key = "odh/org/repo"
            success = {
                "commit_sha": "run-20-success",
                "analyzer_version": "v1",
                "scanned_at": "2026-11-11T00:00:00Z",
                "workflow_run_number": 20,
                "workflow_run_attempt": 1,
            }
            main_path = os.path.join(tmp, "main.json")
            ours_path = os.path.join(tmp, "ours.json")
            run_invalidations = os.path.join(tmp, "run-invalidations.json")
            main_invalidations = os.path.join(tmp, "main-invalidations.json")
            out_invalidations = os.path.join(tmp, "out-invalidations.json")
            merged_path = os.path.join(tmp, "merged.json")
            for path, value in (
                (main_path, {key: success}),
                (main_invalidations, {"invalidations": []}),
            ):
                with open(path, "w") as f:
                    json.dump(value, f)

            env = dict(os.environ, GITHUB_RUN_NUMBER="21", GITHUB_RUN_ATTEMPT="1")
            subprocess.run(
                [sys.executable, SCRIPT, results,
                 "--invalidations-output", run_invalidations, "-o", ours_path],
                check=True,
                capture_output=True,
                text=True,
                env=env,
            )
            subprocess.run(
                [sys.executable, MERGE_SCRIPT, ours_path, main_path, merged_path,
                 run_invalidations, main_invalidations, out_invalidations],
                check=True,
                capture_output=True,
                text=True,
            )

            with open(merged_path) as f:
                merged = json.load(f)
            with open(out_invalidations) as f:
                tombstones = json.load(f)["invalidations"]
            self.assertEqual(merged[key]["commit_sha"], "run-20-success")
            self.assertEqual(tombstones[0]["workflow_run_number"], 10)

    def test_newer_run_survives_persisted_tombstone(self):
        key = "odh/org/repo"
        state, tombstones = self.merge(
            {key: {
                "commit_sha": "newer-run",
                "scanned_at": "2026-11-10T00:00:00Z",
                "workflow_run_number": 21,
                "workflow_run_attempt": 1,
            }},
            {},
            run_invalidations=[{
                "prefix": key,
                "invalidated_at": "2026-11-12T00:00:00Z",
                "workflow_run_number": 20,
                "workflow_run_attempt": 1,
            }],
            main_invalidations=[],
        )
        self.assertEqual(state[key]["commit_sha"], "newer-run")
        self.assertEqual(len(tombstones), 1)

    def test_merge_drops_legacy_unqualified_state_keys(self):
        state, _ = self.merge(
            {},
            {
                "org/repo": {"commit_sha": "legacy"},
                "odh/org/repo": {"commit_sha": "scoped", "scanned_at": "2026-11-11T00:00:00Z"},
            },
        )
        self.assertNotIn("org/repo", state)
        self.assertEqual(state["odh/org/repo"]["commit_sha"], "scoped")

    def test_merge_rejects_invalidations_without_persistent_output(self):
        with tempfile.TemporaryDirectory() as tmp:
            ours = os.path.join(tmp, "ours.json")
            main = os.path.join(tmp, "main.json")
            output = os.path.join(tmp, "merged.json")
            run_invalidations = os.path.join(tmp, "run-invalidations.json")
            for path in (ours, main, run_invalidations):
                with open(path, "w") as f:
                    json.dump({"invalidations": []} if path == run_invalidations else {}, f)

            result = subprocess.run(
                [sys.executable, MERGE_SCRIPT, ours, main, output, run_invalidations],
                capture_output=True,
                text=True,
            )

            self.assertNotEqual(result.returncode, 0)

    def test_missing_run_invalidation_file_is_an_error(self):
        with tempfile.TemporaryDirectory() as tmp:
            ours, main, output, run_events, main_events, output_events = [
                os.path.join(tmp, name)
                for name in ("ours.json", "main.json", "merged.json", "run.json", "main-events.json", "out-events.json")
            ]
            for path, value in ((ours, {}), (main, {}), (main_events, {"invalidations": []})):
                with open(path, "w") as f:
                    json.dump(value, f)

            result = subprocess.run(
                [sys.executable, MERGE_SCRIPT, ours, main, output, run_events, main_events, output_events],
                capture_output=True,
                text=True,
            )

            self.assertNotEqual(result.returncode, 0)

    def test_timestamp_only_event_does_not_replace_run_provenance_event(self):
        key = "odh/org/repo"
        _, events = self.merge(
            {},
            {},
            run_invalidations=[{
                "prefix": key,
                "invalidated_at": "2026-11-12T12:00:00Z",
            }],
            main_invalidations=[{
                "prefix": key,
                "invalidated_at": "2026-11-12T10:00:00Z",
                "workflow_run_number": 20,
                "workflow_run_attempt": 1,
            }],
        )

        self.assertEqual(len(events), 2)
        self.assertIn(20, {event.get("workflow_run_number") for event in events})

    def test_invalidation_compaction_keeps_run_and_timestamp_only_events(self):
        key = "odh/org/repo"
        state, events = self.merge(
            {},
            {key: {
                "commit_sha": "run-21",
                "scanned_at": "2026-11-12T10:30:00Z",
                "workflow_run_number": 21,
                "workflow_run_attempt": 1,
            }},
            run_invalidations=[{
                "prefix": key,
                "invalidated_at": "2026-11-12T12:00:00Z",
                "workflow_run_number": 19,
                "workflow_run_attempt": 3,
            }],
            main_invalidations=[
                {
                    "prefix": key,
                    "invalidated_at": "2026-11-12T10:00:00Z",
                    "workflow_run_number": 20,
                    "workflow_run_attempt": 1,
                },
                {"prefix": key, "invalidated_at": "2026-11-12T11:00:00Z"},
            ],
        )

        self.assertNotIn(key, state)
        self.assertEqual(len(events), 3)
        self.assertEqual(
            {event["invalidated_at"] for event in events},
            {
                "2026-11-12T10:00:00Z",
                "2026-11-12T11:00:00Z",
                "2026-11-12T12:00:00Z",
            },
        )

    def test_merge_applies_invalidations_before_selecting_competing_entries(self):
        key = "odh/org/repo"
        state, _ = self.merge(
            {key: {
                "commit_sha": "run-21-earlier",
                "scanned_at": "2026-11-12T10:00:00Z",
                "workflow_run_number": 21,
                "workflow_run_attempt": 1,
            }},
            {key: {
                "commit_sha": "run-20-valid",
                "scanned_at": "2026-11-12T12:00:00Z",
                "workflow_run_number": 20,
                "workflow_run_attempt": 1,
            }},
            run_invalidations=[{
                "prefix": key,
                "invalidated_at": "2026-11-12T11:00:00Z",
            }],
            main_invalidations=[],
        )
        self.assertEqual(state[key]["commit_sha"], "run-20-valid")


if __name__ == "__main__":
    unittest.main()
