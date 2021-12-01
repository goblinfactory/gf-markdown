# markdown link checker

cicd tool for validating markdown files contains no broken links, so that it can be added as a check as part of the build. If you rename a file and don't remember to update your readme, then `markdown` will stop the build until all the links are correct.

some text with [README1](README.md) an embedded link.

A line with 2 links [README2](README.md) and a broken link [README3](README2.md).

**Installing**

- tbd

**Usage standalone**

- tbd
- screenshot. (tbd)

**Adding to makefile**

- tbd

internal packages

- [ansi/ansi.go](internal/ansi/ansi.go) : *Ansi color printing*
- [mystrings/strings.go](../internal/strings.go) : *misc string utils*
- [regexs/pairmatcher.go](../internal/regexs/pairmatcher.go) : *regex 'pair' matching*


