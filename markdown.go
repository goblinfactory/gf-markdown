package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/goblinfactory/markdown/internal/ansi"
	"github.com/goblinfactory/markdown/internal/mystrings"
	"github.com/goblinfactory/markdown/internal/regexs"
	"github.com/goblinfactory/markdown/printer"
)

func main() {
	p := &printer.Printer{}
	pass := Main(p, os.Args[1:])
	if !pass {
		log.Fatalf("broken links:%v", os.Args[1:])
	}
}

// Main entry point for Markdown checker console app or integration test.
func Main(p *printer.Printer, args []string) bool {
	fmt.Printf("%v\n", args)
	verbose := mystrings.IsAny(os.Args, "-v")
	args = mystrings.Remove(args, "-v")
	defer p.Flush()
	ac := len(args)
	if ac == 0 {
		usage(p)
		return false
	}

	var reports []Report

	if runtime.GOOS == "windows" {
		reports = checkAllWindowsExpandGlob(args[0], verbose)
	} else {
		reports = checkAll(args, verbose)
	}

	printReports(p, reports, verbose)
	return allPassed(reports)
}

func usage(p *printer.Printer) {
	p.Println("usage: markdown [-v] [/a/b/filename1.txt] [readme.md] [...]")
}

func checkAll(matches []string, verbose bool) []Report {
	// need to test this on windows
	reports := make([]Report, 0)

	for _, f := range matches {
		r := CheckOne(f, verbose)
		reports = append(reports, r)
	}
	return reports
}

func checkAllWindowsExpandGlob(globpath string, verbose bool) []Report {
	// need to test this on windows
	reports := make([]Report, 0)

	matches, err := filepath.Glob(globpath)

	check(err)

	for _, f := range matches {
		r := CheckOne(f, verbose)
		reports = append(reports, r)
	}
	return reports
}

// Report contains result of the link check
type Report struct {
	Filename      string
	Pass          bool
	MarkdownFiles []Result
	CntErrors     int
}

// CheckOne checks that all the links in a markdown file are correct. verbose will include non broken links in report.
func CheckOne(path string, verbose bool) Report {

	bytes, err := ioutil.ReadFile(path)
	errcnt := 0
	check(err)
	links := FindLinks(bytes)
	results := make([]Result, 0)
	for _, link := range links {
		ok, _ := checkGoSourceFileExists(link.RelPath)
		if !ok {
			errcnt++
		}
		result := Result{link, ok}
		if ok && verbose {
			results = append(results, result)
		}
		if !ok {
			results = append(results, result)
		}
	}

	return Report{
		path,
		errcnt == 0,
		results,
		errcnt,
	}
}

func allPassed(reports []Report) bool {
	allValid := true
	for _, r := range reports {
		if !r.Pass {
			allValid = false
		}
	}
	return allValid
}

func printReports(p *printer.Printer, reports []Report, verbose bool) {
	mw := maxWidth(reports)
	for _, r := range reports {
		printReport(p, mw, r, verbose)
	}
}

func printReport(p *printer.Printer, maxWidth int, report Report, verbose bool) {

	if verbose {
		p.Println("%s%s", mystrings.PadLeft(maxWidth+4, "link"), "ok")
		p.Println("%s%s", mystrings.PadLeft(maxWidth+4, "----"), "--")
	}

	if !report.Pass {
		p.Println("CheckLinks:%s has %s(%d) broken links%s", report.Filename, ansi.Red, report.CntErrors, ansi.Reset)
	}

	for _, link := range report.MarkdownFiles {
		if link.Exists {
			p.Println("%s%s%s  ✓%s", ansi.Reset, mystrings.PadLeft(maxWidth+2, link.RelPath), ansi.Green, ansi.Reset)
		} else {
			p.Println(" - %s%s(broken)%s", mystrings.PadLeft(maxWidth, link.RelPath), ansi.Red, ansi.Reset)
		}
	}
}

func maxWidth(reports []Report) int {
	m := 0
	for _, rep := range reports {
		for _, res := range rep.MarkdownFiles {
			l := len(res.RelPath)
			if l > m {
				m = l
			}
		}
	}
	return m
}

// Result ...
type Result struct {
	Link
	Exists bool
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Link is a markdown link parsed from a markdown file with format [text](uri)
type Link struct {
	Text    string
	RelPath string
}

func parseLinks(pairs regexs.Pairs) []Link {
	m := make([]Link, len(pairs))
	for i, p := range pairs {
		m[i] = Link{p.Match1, p.Match2}
	}
	return m
}

// FindLinks finds all the internal hyperlinks in a markdown file (in this case a sequence of bytes) using markdown [text](hyperlink) format. Ignores any external links
func FindLinks(content []byte) []Link {
	pm := regexs.NewPairMatcher(regexs.PatternMarkdownURI, nil, []string{"://"})
	pairs := pm.SearchForPairs(content)
	links := parseLinks(pairs)
	return links
}

func checkGoSourceFileExists(relpath string) (bool, error) {
	_, err := os.Stat(relpath)
	if err == nil {
		return true, nil
	}
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	panic(err)
}

// Random idea
// -----------
// use git history to check for git files with old names, and see if we can work out if they have been renamed, find files with same content?
// would need to be a new git tool, "movedWhere ?" that can find out where git file was moved to.
// then can have option to show broken links and have parameter to allow for automatic repair broken links, and preview repairs.
