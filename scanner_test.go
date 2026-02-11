package main

import ("os";"path/filepath";"testing")

func TestDetectAWS(t *testing.T) {
	tmp := t.TempDir()
	os.WriteFile(filepath.Join(tmp,"c.py"), []byte(`KEY="AKIAIOSFODNN7EXAMPLE"`), 0644)
	findings := ScanPath(tmp)
	ok := false
	for _, f := range findings { if f.Rule=="aws-key" { ok=true } }
	if !ok { t.Error("missed AWS key") }
}

func TestCleanDir(t *testing.T) {
	tmp := t.TempDir()
	os.WriteFile(filepath.Join(tmp,"ok.txt"), []byte("hello world"), 0644)
	for _, f := range ScanPath(tmp) {
		if f.Severity=="HIGH"||f.Severity=="CRITICAL" { t.Errorf("false positive: %+v",f) }
	}
}

func TestEntropy(t *testing.T) {
	if entropy("aB3$xK9!mP2@") <= entropy("aaaaaa") { t.Error("entropy calc wrong") }
}

func TestMask(t *testing.T) {
	m := mask("AKIAIOSFODNN7EXAMPLE")
	if m == "AKIAIOSFODNN7EXAMPLE" { t.Error("not masked") }
}
