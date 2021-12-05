package main

import (
	"os"

	"github.com/goblinfactory/gf-markdown/markdown"
	"github.com/goblinfactory/gf-markdown/printer"
)

func main() {
	p := printer.New(os.Stdout, os.Stderr)
	result := markdown.RunFromArgs(os.Args[1:], p)
	os.Exit(int(result))
}
