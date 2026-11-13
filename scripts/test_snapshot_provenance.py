"""Tests for snapshot provenance stamping and freshness ordering."""

import json
import os
import subprocess
import sys
import tempfile
import unittest

from snapshot_provenance import __file__ as SNAPSHOT_PROVENANCE_SCRIPT
from snapshot_provenance import is_newer, stamp, validate_version_label, write_removal_marker


class SnapshotProvenanceTest(unittest.TestCase):
    def write_metadata(self, path, metadata):
        with open(path, "w") as f:
            json.dump(metadata, f)

    def test_stamp_preserves_snapshot_fields(self):
        with tempfile.TemporaryDirectory() as tmp:
            path = os.path.join(tmp, "snapshot-metadata.json")
            self.write_metadata(path, {"version": "2026-11-11", "timestamp": "t"})

            stamp(path, "17", "2")

            with open(path) as f:
                metadata = json.load(f)
            self.assertEqual(metadata["version"], "2026-11-11")
            self.assertEqual(metadata["workflow_run_number"], 17)
            self.assertEqual(metadata["workflow_run_attempt"], 2)

    def test_later_run_is_newer_even_when_old_run_is_retried_later(self):
        with tempfile.TemporaryDirectory() as tmp:
            path = os.path.join(tmp, "snapshot-metadata.json")
            self.write_metadata(
                path,
                {"workflow_run_number": 20, "workflow_run_attempt": 1},
            )

            self.assertTrue(is_newer(path, 19, 3))
            self.assertFalse(is_newer(path, 20, 2))
            self.assertFalse(is_newer(path, 21, 1))

    def test_legacy_metadata_is_replaceable(self):
        with tempfile.TemporaryDirectory() as tmp:
            path = os.path.join(tmp, "snapshot-metadata.json")
            self.write_metadata(path, {"version": "2026-11-11"})

            self.assertFalse(is_newer(path, 1, 1))

    def test_removal_marker_blocks_older_run(self):
        with tempfile.TemporaryDirectory() as tmp:
            path = os.path.join(tmp, "snapshot-metadata.json")

            write_removal_marker(path, "2026-11-12", "odh", 30, 1)

            self.assertTrue(is_newer(path, 29, 9))
            self.assertFalse(is_newer(path, 30, 2))

    def test_cli_stamp_and_freshness(self):
        with tempfile.TemporaryDirectory() as tmp:
            path = os.path.join(tmp, "snapshot-metadata.json")
            self.write_metadata(path, {"version": "2026-11-11"})

            subprocess.run(
                [sys.executable, SNAPSHOT_PROVENANCE_SCRIPT, "stamp", path, "30", "2"],
                check=True,
                capture_output=True,
                text=True,
            )
            result = subprocess.run(
                [sys.executable, SNAPSHOT_PROVENANCE_SCRIPT, "is-newer", path, "30", "1"],
                check=True,
                capture_output=True,
                text=True,
            )
            self.assertEqual(result.stdout.strip(), "yes")

    def test_cli_removed_marker(self):
        with tempfile.TemporaryDirectory() as tmp:
            path = os.path.join(tmp, "snapshot-metadata.json")
            subprocess.run(
                [
                    sys.executable,
                    SNAPSHOT_PROVENANCE_SCRIPT,
                    "removed",
                    path,
                    "2026-11-12",
                    "odh",
                    "30",
                    "1",
                ],
                check=True,
                capture_output=True,
                text=True,
            )
            with open(path) as f:
                metadata = json.load(f)
            self.assertTrue(metadata["removed"])
            self.assertEqual(metadata["workflow_run_number"], 30)

    def test_version_label_validation_rejects_path_traversal(self):
        for label in ("v1", "2026-11-12", "a" * 64):
            with self.subTest(label=label):
                validate_version_label(label)

        for label in ("", "../outside", "v1/../outside", ".hidden", "a" * 65):
            with self.subTest(label=label):
                with self.assertRaises(ValueError):
                    validate_version_label(label)

    def test_cli_validates_version_label(self):
        valid = subprocess.run(
            [sys.executable, SNAPSHOT_PROVENANCE_SCRIPT, "validate-version", "v1"],
            capture_output=True,
            text=True,
        )
        invalid = subprocess.run(
            [sys.executable, SNAPSHOT_PROVENANCE_SCRIPT, "validate-version", "../outside"],
            capture_output=True,
            text=True,
        )

        self.assertEqual(valid.returncode, 0, valid.stderr)
        self.assertNotEqual(invalid.returncode, 0)


if __name__ == "__main__":
    unittest.main()
