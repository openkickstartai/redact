package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInstallHookNoGit(t *testing.T) {
	tmp := t.TempDir()
	old, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(old)
	if err := installHook(); err == nil {
		t.Error("should fail without .git")
	}
}

func TestInstallHookSuccess(t *testing.T) {
	tmp := t.TempDir()
	old, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(old)
	os.MkdirAll(filepath.Join(tmp, ".git", "hooks"), 0755)
	if err := installHook(); err != nil {
		t.Fatal(err)
	}
	hook := filepath.Join(tmp, ".git", "hooks", "pre-commit")
	data, _ := os.ReadFile(hook)
	if len(data) == 0 { t.Error("hook file empty") }
}

func TestInstallHookExists(t *testing.T) {
	tmp := t.TempDir()
	old, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(old)
	hookDir := filepath.Join(tmp, ".git", "hooks")
	os.MkdirAll(hookDir, 0755)
	os.WriteFile(filepath.Join(hookDir, "pre-commit"), []byte("existing"), 0755)
	if err := installHook(); err == nil {
		t.Error("should fail if hook exists")
	}
}
