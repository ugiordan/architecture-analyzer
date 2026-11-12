"""Regression tests for failed and malformed carry-forward results."""

import json
import os
import pathlib
import shutil
import subprocess
import tempfile
import unittest
from datetime import datetime, timedelta, timezone


SCRIPT = pathlib.Path(__file__).with_name("carry-forward-snapshot-results.sh")


class CarryForwardSnapshotResultsTest(unittest.TestCase):
    def run_command(self, *command, cwd=None):
        return subprocess.run(
            list(command), cwd=cwd, check=True, capture_output=True, text=True
        )

    def test_invalid_architecture_json_is_removed_when_no_fallback_exists(self):
        with tempfile.TemporaryDirectory() as tmp:
            root = pathlib.Path(tmp)
            analyzer = root / "analyzer"
            scripts = analyzer / "scripts"
            scripts.mkdir(parents=True)
            shutil.copy2(SCRIPT, scripts / SCRIPT.name)
            (analyzer / "scan-config.yaml").write_text("platforms: {}\n")
            (analyzer / ".last-scan-shas.json").write_text("{}\n")

            remote = root / "remote.git"
            subprocess.run(["git", "init", "--bare", "-q", "-b", "main", str(remote)], check=True)
            subprocess.run(["git", "init", "-q", "-b", "main", str(analyzer)], check=True)
            subprocess.run(
                ["git", "-C", str(analyzer), "remote", "add", "origin", str(remote)],
                check=True,
            )

            fake_bin = root / "bin"
            fake_bin.mkdir()
            yq = fake_bin / "yq"
            yq.write_text('#!/bin/sh\nprintf "org/demo\\n"\n')
            yq.chmod(0o755)

            repo_dir = root / "odh-results" / "org" / "demo"
            repo_dir.mkdir(parents=True)
            (repo_dir / "component-architecture.json").write_text("not-json\n")
            (repo_dir / "security-findings.json").write_text("[]\n")

            env = dict(os.environ, PATH=f"{fake_bin}:{os.environ['PATH']}")
            result = subprocess.run(
                ["bash", str(scripts / SCRIPT.name), str(root / "odh-results")],
                env=env,
                capture_output=True,
                text=True,
            )

            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
            self.assertFalse((repo_dir / "component-architecture.json").exists())
            self.assertFalse((repo_dir / "security-findings.json").exists())
            self.assertTrue((repo_dir / "_failed").exists())

    def test_verified_skip_fetches_snapshot_through_existing_checkout(self):
        with tempfile.TemporaryDirectory() as tmp:
            root = pathlib.Path(tmp)
            remote = root / "remote.git"
            source = root / "snapshot-source"
            source.mkdir()
            self.run_command("git", "init", "-q", "-b", "snapshots", cwd=source)
            self.run_command("git", "config", "user.name", "test", cwd=source)
            self.run_command("git", "config", "user.email", "test@example.invalid", cwd=source)
            date_label = (datetime.now(timezone.utc).date() - timedelta(days=1)).isoformat()
            raw_dir = source / "output" / date_label / "odh" / date_label / "raw" / "demo"
            raw_dir.mkdir(parents=True)
            (raw_dir / "component-architecture.json").write_text(
                '{"repo":"org/demo","commit_sha":"abc123","analyzer_version":"v1"}\n'
            )
            (raw_dir / "security-findings.json").write_text("[]\n")
            snapshot_dir = source / "output" / date_label / "odh" / date_label
            snapshot_dir.mkdir(parents=True, exist_ok=True)
            metadata_path = snapshot_dir / "snapshot-metadata.json"
            metadata_path.write_text(json.dumps({
                "version": date_label,
                "platform": "odh",
                "analyzer_version": "v1",
                "workflow_run_number": 100,
                "workflow_run_attempt": 1,
                "repos_analyzed": {"org/demo": "abc123"},
            }) + "\n")
            self.run_command("git", "init", "--bare", "-q", "-b", "snapshots", str(remote))
            self.run_command("git", "add", ".", cwd=source)
            self.run_command("git", "commit", "-m", "snapshot", cwd=source)
            self.run_command("git", "remote", "add", "origin", str(remote), cwd=source)
            self.run_command("git", "push", "origin", "snapshots", cwd=source)

            analyzer = root / "analyzer"
            scripts = analyzer / "scripts"
            scripts.mkdir(parents=True)
            shutil.copy2(SCRIPT, scripts / SCRIPT.name)
            (analyzer / "scan-config.yaml").write_text("platforms: {}\n")
            (analyzer / ".last-scan-shas.json").write_text(
                '{"odh/org/demo":{"commit_sha":"abc123","analyzer_version":"v1"}}\n'
            )
            self.run_command("git", "init", "-q", "-b", "main", str(analyzer))
            self.run_command("git", "remote", "add", "origin", str(remote), cwd=analyzer)

            fake_bin = root / "bin"
            fake_bin.mkdir()
            yq = fake_bin / "yq"
            yq.write_text('#!/bin/sh\nprintf "org/demo\\n"\n')
            yq.chmod(0o755)
            results = root / "odh-results"
            skipped_repo = results / "org" / "demo"
            skipped_repo.mkdir(parents=True)
            (skipped_repo / "_skipped").touch()
            env = dict(os.environ, PATH=f"{fake_bin}:{os.environ['PATH']}")

            result = subprocess.run(
                ["bash", str(scripts / SCRIPT.name), str(results)],
                env=env,
                capture_output=True,
                text=True,
            )

            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
            self.assertEqual(
                (skipped_repo / "component-architecture.json").read_text(),
                (raw_dir / "component-architecture.json").read_text(),
            )
            self.assertEqual(
                (skipped_repo / "security-findings.json").read_text(),
                "[]\n",
            )

            newer_label = (datetime.now(timezone.utc).date() - timedelta(days=2)).isoformat()
            newer_raw = source / "output" / newer_label / "odh" / newer_label / "raw" / "demo"
            newer_raw.mkdir(parents=True)
            (newer_raw / "component-architecture.json").write_text(
                '{"repo":"org/demo","commit_sha":"abc123","analyzer_version":"v1","extracted_at":"2026-11-13T00:00:00Z"}\n'
            )
            (newer_raw / "security-findings.json").write_text("[]\n")
            newer_snapshot = newer_raw.parent.parent / "snapshot-metadata.json"
            newer_snapshot.write_text(json.dumps({
                "version": newer_label,
                "platform": "odh",
                "analyzer_version": "v2",
                "workflow_run_number": 101,
                "workflow_run_attempt": 1,
                "repos_analyzed": {"org/demo": "abc123"},
                "repo_analyzer_versions": {"org/demo": "v1"},
            }) + "\n")
            self.run_command("git", "add", ".", cwd=source)
            self.run_command("git", "commit", "-m", "newer snapshot with carried data", cwd=source)
            self.run_command("git", "push", "origin", "snapshots", cwd=source)

            (analyzer / ".last-scan-shas.json").write_text("{}\n")
            (skipped_repo / "component-architecture.json").write_text(
                '{"repo":"org/demo","commit_sha":"partial","analyzer_version":"v1"}\n'
            )
            (skipped_repo / "security-findings.json").write_text("[]\n")
            (skipped_repo / "_failed").write_text("failed\n")
            result = subprocess.run(
                ["bash", str(scripts / SCRIPT.name), str(results)],
                env=env,
                capture_output=True,
                text=True,
            )
            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
            self.assertIn("abc123", (skipped_repo / "component-architecture.json").read_text())
            self.assertIn("2026-11-13T00:00:00Z", (skipped_repo / "component-architecture.json").read_text())
            self.assertFalse((skipped_repo / "_failed").exists())

            metadata_path.write_text(json.dumps({
                "version": date_label,
                "platform": "odh",
                "analyzer_version": "v1",
                "repos_analyzed": {"org/demo": "different-sha"},
            }) + "\n")
            newer_metadata = json.loads(newer_snapshot.read_text())
            newer_metadata["repos_analyzed"]["org/demo"] = "different-sha"
            newer_snapshot.write_text(json.dumps(newer_metadata) + "\n")
            self.run_command("git", "add", ".", cwd=source)
            self.run_command("git", "commit", "-m", "mismatch snapshot provenance", cwd=source)
            self.run_command("git", "push", "origin", "snapshots", cwd=source)

            (skipped_repo / "component-architecture.json").write_text(
                '{"repo":"org/demo","commit_sha":"partial","analyzer_version":"v1"}\n'
            )
            (skipped_repo / "security-findings.json").write_text("[]\n")
            (skipped_repo / "_failed").write_text("failed\n")
            result = subprocess.run(
                ["bash", str(scripts / SCRIPT.name), str(results)],
                env=env,
                capture_output=True,
                text=True,
            )
            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
            self.assertFalse((skipped_repo / "component-architecture.json").exists())
            self.assertFalse((skipped_repo / "security-findings.json").exists())
            self.assertTrue((skipped_repo / "_failed").exists())


if __name__ == "__main__":
    unittest.main()
