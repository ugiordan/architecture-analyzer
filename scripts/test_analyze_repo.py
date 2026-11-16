"""Regression tests for retry cleanup in analyze-repo.sh."""

import os
import pathlib
import shutil
import subprocess
import tempfile
import unittest


SCRIPT = pathlib.Path(__file__).with_name("analyze-repo.sh")


class AnalyzeRepoTest(unittest.TestCase):
    def write_executable(self, path, contents):
        path.write_text(contents)
        path.chmod(0o755)

    def test_versioned_scans_preserve_other_versions_and_retries_clean_current_outputs(self):
        with tempfile.TemporaryDirectory() as tmp:
            root = pathlib.Path(tmp)
            fake_bin = root / "bin"
            fake_bin.mkdir()
            real_git = shutil.which("git")
            self.assertIsNotNone(real_git)
            self.write_executable(
                fake_bin / "git",
                f'''#!/bin/bash
if [ "$1" = "clone" ]; then
    destination=""
    for arg in "$@"; do destination="$arg"; done
    mkdir -p "$destination"
    exit 0
fi
if [ "$1" = "-C" ] && [ "$3" = "rev-parse" ]; then
    printf 'abc123\\n'
    exit 0
fi
exec "{real_git}" "$@"
''',
            )
            self.write_executable(fake_bin / "yq", "#!/bin/sh\nprintf '\\n'\n")
            self.write_executable(fake_bin / "gh", "#!/bin/sh\nexit 1\n")
            self.write_executable(fake_bin / "timeout", "#!/bin/sh\nshift\nexec \"$@\"\n")
            self.write_executable(
                root / "arch-analyzer",
                '''#!/bin/bash
if [ "$1" = "version" ]; then
    printf 'arch-analyzer v1\\n'
    exit 0
fi
[ "$1" = "full-analysis" ] || exit 2
shift
output=""
org=""
version=""
repo_path=""
while [ $# -gt 0 ]; do
    case "$1" in
        -output-dir) output="$2"; shift 2 ;;
        -org) org="$2"; shift 2 ;;
        -version) version="$2"; shift 2 ;;
        *) repo_path="$1"; shift ;;
    esac
done
if [ -n "$version" ]; then output="${output}/${version}"; fi
mkdir -p "$output"
repo_name="${repo_path##*/}"
printf '{"repo":"%s/%s","commit_sha":"abc123","analyzer_version":"v1","extracted_at":"2026-11-12T12:00:00Z"}\\n' "$org" "$repo_name" > "$output/component-architecture.json"
if [ "$FAKE_ANALYSIS_MODE" = "complete" ]; then
    printf '[]\\n' > "$output/security-findings.json"
    exit 0
fi
exit 7
''',
            )

            env = dict(
                os.environ,
                PATH=f"{fake_bin}:{os.environ['PATH']}",
                GITHUB_WORKSPACE=str(root),
                RUNNER_TEMP=str(root),
                FORCE_RESCAN="1",
                SCAN_PLATFORM="rhoai",
                ANALYSIS_TIMEOUT="20",
                FAKE_ANALYSIS_MODE="complete",
            )
            first = subprocess.run(
                ["bash", str(SCRIPT), "org/demo", "results", "v1"],
                cwd=root,
                env=env,
                capture_output=True,
                text=True,
            )
            self.assertEqual(first.returncode, 0, first.stdout + first.stderr)
            v1_findings = root / "results" / "org" / "demo" / "v1" / "security-findings.json"
            self.assertTrue(v1_findings.is_file())

            unversioned = subprocess.run(
                ["bash", str(SCRIPT), "org/demo", "results"],
                cwd=root,
                env=env,
                capture_output=True,
                text=True,
            )
            self.assertEqual(unversioned.returncode, 0, unversioned.stdout + unversioned.stderr)
            self.assertTrue(v1_findings.is_file())

            second_version = subprocess.run(
                ["bash", str(SCRIPT), "org/demo", "results", "v2"],
                cwd=root,
                env=env,
                capture_output=True,
                text=True,
            )
            self.assertEqual(second_version.returncode, 0, second_version.stdout + second_version.stderr)
            v2_findings = root / "results" / "org" / "demo" / "v2" / "security-findings.json"
            self.assertTrue(v1_findings.is_file())
            self.assertTrue(v2_findings.is_file())

            invalid_version = subprocess.run(
                ["bash", str(SCRIPT), "org/demo", "results", "../outside"],
                cwd=root,
                env=env,
                capture_output=True,
                text=True,
            )
            self.assertNotEqual(invalid_version.returncode, 0)
            self.assertTrue(v1_findings.is_file())
            self.assertTrue(v2_findings.is_file())

            env["FAKE_ANALYSIS_MODE"] = "partial"
            second = subprocess.run(
                ["bash", str(SCRIPT), "org/demo", "results", "v2"],
                cwd=root,
                env=env,
                capture_output=True,
                text=True,
            )
            self.assertEqual(second.returncode, 7, second.stdout + second.stderr)
            self.assertFalse(v2_findings.exists())
            self.assertTrue(v1_findings.is_file())
            self.assertTrue((root / "results" / "org" / "demo" / "_failed").is_file())


if __name__ == "__main__":
    unittest.main()
