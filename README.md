# markdown link checker

cicd tool for validating markdown files contains no broken links, so that it can be added as a check as part of the build. If you rename a file and don't remember to update your readme, then `markdown` will stop the build until all the links are correct.

**Installing**

- tbd

**Usage standalone**

Run from the root of your solution using either glob patterns e.g. `markdown testdata/**/*.md` in nix or osx, and when running in windows specify each of the filenames separated with spaces.

> markdown testdata/**/*.md -v

*Running with verbose*

![markdown testdata/**/*.md -v](markdown1.png)

> markdown testdata/**/*.md

*Running without verbose*

![markdown testdata/**/*.md](markdown2.png)

**Adding to makefile**

- tbd

internal packages

- [ansi/ansi.go](internal/ansi/ansi.go) : *Ansi color printing*
- [mystrings/strings.go](internal/mystrings/strings.go) : *misc string utils*
- [regexs/pairmatcher.go](internal/regexs/pairmatcher.go) : *regex 'pair' matching*


