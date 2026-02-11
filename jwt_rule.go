package main

import "regexp"

var jwtPattern = regexp.MustCompile(`eyJ[A-Za-z0-9_-]{10,}\.eyJ[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}`)

func init() {
	rules = append(rules, struct{ Name string; Pat *regexp.Regexp; Sev string }{
		"jwt-token", jwtPattern, "HIGH",
	})
}
