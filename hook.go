package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const hookScript = `#!/bin/sh
# redact pre-commit hook - scans staged files for secrets
files=$(git diff --cached --name-only --diff-filter=ACM)
if [ -z "$files" ]; then exit 0; fi
echo "redact: scanning staged files..."
for f in $files; do
  redact scan "$f" --quiet
  if [ $? -ne 0 ]; then
    echo "redact: secrets found in $f - commit blocked"
    exit 1
  fi
done
echo "redact: all clear"
`

func installHook() error {
	hookDir := filepath.Join(".git", "hooks")
	if _, err := os.Stat(hookDir); err != nil {
		return fmt.Errorf("not a git repository (no .git/hooks)")
	}
	hookPath := filepath.Join(hookDir, "pre-commit")
	if _, err := os.Stat(hookPath); err == nil {
		return fmt.Errorf("pre-commit hook already exists at %s", hookPath)
	}
	if err := os.WriteFile(hookPath, []byte(hookScript), 0755); err != nil {
		return err
	}
	fmt.Printf("Installed pre-commit hook at %s\n", hookPath)
	return nil
}
