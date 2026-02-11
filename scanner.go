package main

import ("bufio";"math";"os";"path/filepath";"regexp";"strings")

type Finding struct {
	File string `json:"file"`
	Line int `json:"line"`
	Rule string `json:"rule"`
	Match string `json:"match"`
	Severity string `json:"severity"`
}

var rules = []struct{ Name string; Pat *regexp.Regexp; Sev string }{
	{"aws-key", regexp.MustCompile(`AKIA[0-9A-Z]{16}`), "HIGH"},
	{"gh-token", regexp.MustCompile(`gh[ps]_[A-Za-z0-9_]{36,}`), "HIGH"},
	{"private-key", regexp.MustCompile(`-----BEGIN (RSA |EC )?PRIVATE KEY-----`), "CRITICAL"},
	{"generic-secret", regexp.MustCompile(`(?i)(password|secret|token)\s*[:=]\s*['"][^'"{\s]{8,}`), "MEDIUM"},
}

func ScanPath(root string) []Finding {
	var out []Finding
	filepath.Walk(root, func(p string, i os.FileInfo, e error) error {
		if e != nil || i.IsDir() || i.Size() > 1<<20 { return nil }
		if strings.HasPrefix(i.Name(),".") { return nil }
		f, _ := os.Open(p)
		if f == nil { return nil }
		defer f.Close()
		sc := bufio.NewScanner(f); ln := 0
		for sc.Scan() {
			ln++; line := sc.Text()
			for _, r := range rules {
				if r.Pat.MatchString(line) {
					out = append(out, Finding{p,ln,r.Name,mask(r.Pat.FindString(line)),r.Sev})
				}
			}
		}
		return nil
	})
	return out
}

func entropy(s string) float64 {
	f := map[rune]float64{}; for _,c := range s { f[c]++ }
	var e float64; l := float64(len(s))
	for _,v := range f { p := v/l; if p>0 { e -= p*math.Log2(p) } }
	return e
}

func mask(s string) string {
	if len(s)<=8 { return "****" }
	return s[:4]+"****"+s[len(s)-4:]
}
