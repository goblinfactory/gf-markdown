package markdown

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/goblinfactory/gf-markdown/markdown/internal/ansi"
	"github.com/goblinfactory/gf-markdown/markdown/internal/mystrings"
	"github.com/goblinfactory/gf-markdown/markdown/internal/regexs"
)

// Result of running the markdown checks
type Result int

// Exit error codes
const (
	Success       = 0
	Err1Unhandled = 1
	Err2Arguments = 2
	Err3Links     = 3
)

// LinkCheck of checking a Link
type LinkCheck struct {
	Link
	Exists bool
	Error  string
}

// Link is a markdown link
type Link struct {
	Text    string
	RelPath string
}

// getReports main entry point for Markdown checker console app or integration test.
func getReports(p *params) ([]Report, Result) {
	reports := make([]Report, 0)
	for _, f := range p.files {
		r := GetReport(f)
		reports = append(reports, r)
	}
	if AllPassed(reports) {
		return reports, Success
	}
	return reports, Err3Links
}

// Report contains result of checking all the links in the file
type Report struct {
	File       string
	Pass       bool
	Links      []LinkCheck
	CntErrors  int
	FailReason Result
	Error      error
}

// GetReport checks that all the links in a markdown file are correct. Will switch to the directory and check the relative paths of any links in the file from 'it's' home directory.
func GetReport(filename string) Report {

	er := func(err error) Report {
		return Report{filename, false, make([]LinkCheck, 0), 0, Err1Unhandled, err}
	}
	dir := path.Dir(filename)
	name := filepath.Base(filename)
	cd, err := os.Getwd()
	if err != nil {
		return er(err)
	}

	defer os.Chdir(cd)

	err = os.Chdir(dir)
	if err != nil {
		return er(err)
	}

	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		return er(err)
	}

	links := FindLinks(bytes)

	lr := make([]LinkCheck, 0)

	allok := true
	errcnt := 0
	unhandled := false

	for _, link := range links {
		ok, err := checkDestFileExists(link.RelPath)
		if err != nil {
			lr = append(lr, LinkCheck{link, false, err.Error()})
			errcnt++
			allok = false
			unhandled = true
			continue
		}
		if !ok {
			lr = append(lr, LinkCheck{link, false, "Broken link"})
			errcnt++
			allok = false
			continue
		}
		if ok {
			lr = append(lr, LinkCheck{link, true, "(ok)"})
		}
	}

	// success, all links ok
	r := Report{filename, allok, lr, errcnt, Success, nil}
	if errcnt == 0 && unhandled == false {
		return r
	}

	// an unhandled error in a link
	if errcnt == 0 && unhandled == true {
		r.FailReason = Err1Unhandled
		return r
	}

	// at least 1 broken link
	r.FailReason = Err3Links
	return r
}

func printReports(p *Printer, reports []Report, verbose bool) {
	mw := maxWidth(reports)
	for _, r := range reports {
		printReport(p, mw, r, verbose)
	}
}

func printReport(p *Printer, maxWidth int, report Report, verbose bool) {

	if !report.Pass {
		p.Println("CheckLinks:%s has %s(%d) broken links%s", report.File, ansi.Red, report.CntErrors, ansi.Reset)
	}
	if report.Pass && verbose {
		p.Println("CheckLinks:%s has %sno broken links%s", report.File, ansi.Green, ansi.Reset)
	}
	for _, link := range report.Links {
		if link.Exists && verbose {
			p.Println("%s%s%s  ✓%s", ansi.Reset, mystrings.PadLeft(maxWidth+2, link.RelPath), ansi.Green, ansi.Reset)
		}

		if !link.Exists {
			p.Println("%s%s✗ %s%s", ansi.Red, mystrings.PadLeft(maxWidth+4, link.RelPath), link.Error, ansi.Reset)
		}
	}
}

func maxWidth(reports []Report) int {
	m := 0
	for _, rep := range reports {
		for _, res := range rep.Links {
			l := len(res.RelPath)
			if l > m {
				m = l
			}
		}
	}
	return m
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

func checkDestFileExists(relpath string) (bool, error) {
	cd, _ := os.Getwd()
	newpath := filepath.Join(cd, relpath)
	_, err := os.Stat(newpath)

	if err == nil {
		return true, nil
	}
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
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
