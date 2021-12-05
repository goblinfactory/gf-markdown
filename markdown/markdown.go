package markdown

import (
	"errors"
	"path/filepath"
	"runtime"

	"github.com/goblinfactory/gf-markdown/markdown/internal/mystrings"
	"github.com/goblinfactory/gf-markdown/printer"
)

type params struct {
	printer *printer.Printer
	userParams
}

type userParams struct {
	verbose bool
	files   []string
}

// RunFromArgs runs gf-markdown app and returns it's exit status.
func RunFromArgs(args []string, printer *printer.Printer) Result {
	prms, err := parseParams(args)
	if err != nil {
		printer.PrintErrln(err.Error())
		printer.Flush()
		return Err2Arguments
	}
	retcode := Run(&params{printer, prms})
	return retcode
}

// Run gf-markdown app, print any reports to the buffered printer, flushes to writers, and returns it's exit status
func Run(p *params) Result {
	defer p.printer.Flush()
	reports, result := GetReports(p)
	if result == Success && !p.verbose {
		return Success
	}
	printReports(p.printer, reports, p.verbose)
	return result

}

func parseParams(args []string) (userParams, error) {
	v := mystrings.IsAny(args, "-v")
	files := mystrings.Remove(args, "-v")
	if len(files) == 0 {
		return userParams{}, errors.New("no files or glob path provided")
	}

	// if the os is windows, then you're only allowed 1 param other than -v and that's a single glob path.
	if runtime.GOOS == "windows" {
		if len(files) != 1 {
			return userParams{}, errors.New("only 1 glob path supported")
		}

		globbedfiles, err := filepath.Glob(files[0])
		if err != nil {
			return userParams{}, err
		}
		files = globbedfiles
	}
	return userParams{v, files}, nil
}

func usage(p *printer.Printer) {
	p.Println("usage: markdown [-v] [/a/b/filename1.txt] [readme.md] [...]")
}
