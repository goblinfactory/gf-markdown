package main

import (
	"os"

	"github.com/goblinfactory/gf-markdown/markdown"
)

func main() {
	p := markdown.NewPrinter(os.Stdout, os.Stderr)
	result := markdown.RunFromArgs(os.Args[1:], p)
	os.Exit(int(result))
}
