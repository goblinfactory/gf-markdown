package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/goblinfactory/go-markdown/internal/ansi"
	"github.com/goblinfactory/go-markdown/internal/mystrings"
	"github.com/goblinfactory/go-markdown/internal/regexs"
	"github.com/goblinfactory/go-markdown/printer"
)

// Exit error codes
const (
	ErrUnhandled   = 1
	ErrNoArguments = 2
)

func main() {
	p := printer.New(os.Stdout)
	retCode := run(os.Args, p)
	os.Exit(retCode)
}

func run(args []string, p *printer.Printer) (retcode int) {

	defer p.Flush()

	defer func() {
		if err := recover(); err != nil {
			retcode = 1
			log.Println("panic:", err)
		}
	}()

	reports, err := GetReports(p, os.Args[1:])
	check(err)
	printReports(p, reports)
	if !AllPassed(reports) {
		retcode = 1
	}
	return retcode
}

// GetReports main entry point for Markdown checker console app or integration test.
func GetReports(p *printer.Printer, args []string) ([]Report, error) {

	verbose := mystrings.IsAny(os.Args, "-v")
	args = mystrings.Remove(args, "-v")

	ac := len(args)
	if ac == 0 {
		usage(p)
		return nil, errors.New("no arguments provided")
	}

	var reports []Report

	if runtime.GOOS == "windows" {
		reports = checkAllWindowsExpandGlob(args[0], verbose)
	} else {
		reports = checkAll(args, verbose)
	}

	return reports, nil
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
	Verbose       bool
	Filename      string
	Pass          bool
	MarkdownFiles []LinkResult
	CntErrors     int
}

// CheckOne checks that all the links in a markdown file are correct.
// verbose will include non broken links in report.
func CheckOne(fpath string, verbose bool) Report {

	dir := path.Dir(fpath)
	name := filepath.Base(fpath)
	cd, err := os.Getwd()
	check(err)
	defer os.Chdir(cd)
	os.Chdir(dir)

	bytes, err := ioutil.ReadFile(name)
	errcnt := 0
	check(err)
	links := FindLinks(bytes)
	results := make([]LinkResult, 0)
	for _, link := range links {
		ok, _ := checkGoSourceFileExists(link.RelPath)
		if !ok {
			errcnt++
		}
		result := LinkResult{link, ok}
		if ok && verbose {
			results = append(results, result)
		}
		if !ok {
			results = append(results, result)
		}
	}

	return Report{
		verbose,
		fpath,
		errcnt == 0,
		results,
		errcnt,
	}
}

// AllPassed returns false if any of the markdown files have broken links
func AllPassed(reports []Report) bool {
	allValid := true
	for _, r := range reports {
		if !r.Pass {
			allValid = false
		}
	}
	return allValid
}

func printReports(p *printer.Printer, reports []Report) {
	mw := maxWidth(reports)
	for _, r := range reports {
		printReport(p, mw, r)
	}
}

func printReport(p *printer.Printer, maxWidth int, report Report) {

	if !report.Pass {
		p.Println("CheckLinks:%s has %s(%d) broken links%s", report.Filename, ansi.Red, report.CntErrors, ansi.Reset)
	}
	if report.Pass && report.Verbose {
		p.Println("CheckLinks:%s has %sno broken links%s", report.Filename, ansi.Green, ansi.Reset)
	}
	for _, link := range report.MarkdownFiles {
		if link.Exists {
			p.Println("%s%s%s  ✓%s", ansi.Reset, mystrings.PadLeft(maxWidth+2, link.RelPath), ansi.Green, ansi.Reset)
		} else {
			p.Println("%s%s✗%s", ansi.Red, mystrings.PadLeft(maxWidth+4, link.RelPath), ansi.Reset)
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

// LinkResult of checking a Link
type LinkResult struct {
	Link
	Exists bool
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Link is a markdown link
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

// FindLinks finds all the internal hyperlinks in a markdown file (in this case a sequence of bytes)
// using markdown [text](hyperlink) format. Ignores any external links
func FindLinks(content []byte) []Link {
	pm := regexs.NewPairMatcher(regexs.PatternMarkdownURI, nil, []string{"://"})
	pairs := pm.SearchForPairs(content)
	links := parseLinks(pairs)
	return links
}

func checkGoSourceFileExists(relpath string) (bool, error) {
	cd, _ := os.Getwd()
	newpath := filepath.Join(cd, relpath)
	_, err := os.Stat(newpath)

	if err == nil {
		return true, nil
	}
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	panic(err)
}
