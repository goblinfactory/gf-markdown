package regexs_test

import (
	"testing"

	"github.com/goblinfactory/markdown/internal/regexs"
	"github.com/google/go-cmp/cmp"
)

func TestPairMatcher(t *testing.T) {
	content := []byte(`
	-   [one two](p/e/a.go)
	-   [a(b)](http://exclude.go)
    -   [a(b)](pkg/error/b.go)
    -   [word](pkg/c.go)
	-   [cats](pkg/cat.gif)
    -   [1 word 1](d.go)
    -   external links
	-   [routing](https://www.honeybadger.io)
	`)

	m := regexs.NewPairMatcher(regexs.PatternMarkdownURI, nil, []string{"://"})
	links := m.SearchForPairs(content)

	expected := regexs.Pairs{
		{"one two", "p/e/a.go"},
		{"a(b)", "pkg/error/b.go"},
		{"word", "pkg/c.go"},
		{"cats", "pkg/cat.gif"},
		{"1 word 1", "d.go"},
	}
	diff := cmp.Diff(expected, links)
	if diff != "" {
		t.Fatal(diff)
	}
}
