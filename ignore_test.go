package main

import ("os";"path/filepath";"testing")

func TestLoadIgnore(t *testing.T) {
	tmp := t.TempDir()
	os.WriteFile(filepath.Join(tmp,".redactignore"), []byte("*.test\nfixtures/\n"), 0644)
	r := LoadIgnoreFile(tmp)
	if len(r.Patterns) != 2 { t.Errorf("want 2, got %d", len(r.Patterns)) }
}

func TestShouldIgnore(t *testing.T) {
	r := &IgnoreRules{Patterns: []string{"*.test","fixtures/"}}
	if !r.ShouldIgnore("x.test") { t.Error("should ignore .test") }
	if r.ShouldIgnore("main.go") { t.Error("should not ignore main.go") }
}
