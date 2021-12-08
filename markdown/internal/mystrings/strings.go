// Package mystrings contains miscalaneous string utilities for matching, padding, filtering and querying
package mystrings

import (
	"fmt"
	"strings"
)

// ContainsAny returns true if any of the strings are in the string
func ContainsAny(src string, matches ...string) bool {
	for _, s := range matches {
		if strings.Contains(src, s) {
			return true
		}
	}
	return false
}

// PadLeft return a string left padded. (left aligned) Text is not clipped
func PadLeft(width int, text string) string {
	wf := fmt.Sprintf("%%-%ss", tostring(width))
	ptext := fmt.Sprintf(wf, text)
	return ptext
}

// Remove an item from a slice
func Remove(src []string, item string) []string {
	item = strings.TrimSpace(item)
	items := make([]string, 0)
	for _, o := range src {
		if strings.TrimSpace(o) != item {
			items = append(items, o)
		}
	}
	return items
}

func tostring(number int) string {
	return fmt.Sprintf("%d", number)
}

// MatchAny returns true if any string matches f(x)
func MatchAny(src []string, match func(string) bool) bool {
	for _, i := range src {
		if match(i) {
			return true
		}
	}
	return false
}

// IsAny returns true if any string matches
func IsAny(src []string, match string) bool {
	for _, i := range src {
		if i == match {
			return true
		}
	}
	return false
}
