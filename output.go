package main

import ("encoding/json";"fmt";"os";"strings")

func PrintFindings(findings []Finding, format string) {
	switch format {
	case "json": json.NewEncoder(os.Stdout).Encode(findings)
	default:
		fmt.Printf("\nredact: %d findings\n%s\n", len(findings), strings.Repeat("-",40))
		for _, f := range findings {
			fmt.Printf("  [%s] %s:%d %s -> %s\n", f.Severity, f.File, f.Line, f.Rule, f.Match)
		}
	}
}
