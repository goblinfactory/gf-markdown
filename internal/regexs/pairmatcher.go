package regexs

import (
	"regexp"

	"github.com/goblinfactory/go-markdown/internal/mystrings"
)

// PatternMarkdownURI is a regex pattern that will return markdown uri's matching syntax [word](url)
const PatternMarkdownURI = `\[(.*?)\]\((.*?)\)`

// Pairs is a collection of Pair
type Pairs []Pair

// Pair is one of the search results from running a SearchForPairs search
type Pair struct {
	Match1 string
	Match2 string
}

// PairMatcher provides simple regex matching for pairs of texts
type PairMatcher struct {
	pattern       string
	match1ignores []string
	match2ignores []string
}

// NewPairMatcher creates a new regex pair matcher for running simple pair searches.
func NewPairMatcher(pattern string, match1ignores []string, match2ignores []string) PairMatcher {
	return PairMatcher{pattern, match1ignores, match2ignores}
}

// SearchForPairs returns all the "two matches" pairs in content. e.g. any regex search that will return 2 matches e.g. `\[(.*?)\]\((.*?)\)`
func (m *PairMatcher) SearchForPairs(content []byte) Pairs {
	matches := make([]Pair, 0)
	patternAll := regexp.MustCompile(m.pattern)
	allMatches := patternAll.FindAllStringSubmatch(string(content), -1)
	for _, match := range allMatches {
		// match 0 is the text that was matched on
		// todo, set breakpoint and manually check with a watch expression if this is true?
		p := Pair{match[1], match[2]}
		if mystrings.ContainsAny(p.Match1, m.match1ignores...) || mystrings.ContainsAny(p.Match2, m.match2ignores...) {
			continue
		}
		matches = append(matches, p)
	}
	return matches
}

// reference https://pkg.go.dev/regexp#Regexp.FindAllStringSubmatchIndex
