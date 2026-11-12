"""Regression tests for generated MkDocs platform navigation."""

import importlib.util
import os
import sys
import tempfile
import unittest
from unittest.mock import patch


SCRIPT = os.path.join(os.path.dirname(__file__), "generate-nav.py")
SPEC = importlib.util.spec_from_file_location("generate_nav", SCRIPT)
generate_nav = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(generate_nav)


class GenerateNavTest(unittest.TestCase):
    def test_empty_platform_list_clears_stale_nav_and_keeps_use_cases_index(self):
        with tempfile.TemporaryDirectory() as tmp:
            docs_dir = os.path.join(tmp, "docs")
            use_cases = os.path.join(docs_dir, "use-cases")
            os.makedirs(use_cases)
            with open(os.path.join(use_cases, "index.md"), "w") as f:
                f.write("# Use Cases\n")
            mkdocs_path = os.path.join(tmp, "mkdocs.yml")
            with open(mkdocs_path, "w") as f:
                f.write(
                    "nav:\n"
                    "  - Use Cases:\n"
                    "    - use-cases/index.md\n"
                    "    - ODH: use-cases/odh-platform/index.md\n"
                    "    - RHOAI: use-cases/rhoai-platform/index.md\n"
                    "  - Contributing: contributing.md\n"
                )

            with patch.object(sys, "argv", [SCRIPT, docs_dir, mkdocs_path]):
                generate_nav.main()

            with open(mkdocs_path) as f:
                nav = f.read()
            self.assertIn("- use-cases/index.md", nav)
            self.assertNotIn("odh-platform", nav)
            self.assertNotIn("rhoai-platform", nav)
            self.assertIn("- Contributing: contributing.md", nav)


if __name__ == "__main__":
    unittest.main()
