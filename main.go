package main

import ("flag";"fmt";"os")

var version = "0.1.0"

func main() {
	if len(os.Args) < 2 { fmt.Println("usage: redact <scan|version>"); os.Exit(1) }
	switch os.Args[1] {
	case "scan":
		cmd := flag.NewFlagSet("scan", flag.ExitOnError)
		fmt := cmd.String("format","table","output format")
		cmd.Parse(os.Args[2:])
		path := "."
		if cmd.NArg() > 0 { path = cmd.Arg(0) }
		f := ScanPath(path)
		PrintFindings(f, *fmt)
		if len(f) > 0 { os.Exit(1) }
	case "version": fmt.Printf("redact %s\n", version)
	}
}
