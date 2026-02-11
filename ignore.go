package main

import ("bufio";"os";"path/filepath";"strings")

type IgnoreRules struct { Patterns []string }

func LoadIgnoreFile(path string) *IgnoreRules {
	r := &IgnoreRules{}
	f, err := os.Open(filepath.Join(path, ".redactignore"))
	if err != nil { return r }
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		l := strings.TrimSpace(sc.Text())
		if l != "" && !strings.HasPrefix(l, "#") { r.Patterns = append(r.Patterns, l) }
	}
	return r
}

func (r *IgnoreRules) ShouldIgnore(path string) bool {
	for _, p := range r.Patterns {
		if m, _ := filepath.Match(p, filepath.Base(path)); m { return true }
		if strings.Contains(path, p) { return true }
	}
	return false
}
