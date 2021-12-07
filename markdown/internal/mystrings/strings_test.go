package mystrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsAny(t *testing.T) {
	src := "apple banana carrot bmw ford"
	assert.True(t, ContainsAny(src, "ferarri", "apple"))
	assert.True(t, ContainsAny(src, "carrot", "ford"))
	assert.False(t, ContainsAny(src, "ferarri"))
	assert.False(t, ContainsAny(src, "ferarri", "monza"))
}

func TestPadLeft(t *testing.T) {
	assert.Equal(t, "12345     ", PadLeft(10, "12345"))
	assert.Equal(t, "123456789012", PadLeft(10, "123456789012"))
	assert.Equal(t, "   ", PadLeft(3, ""))
}

func TestRemove(t *testing.T) {
	src := []string{"one", "two", "three", "four"}
	empty := []string{}
	assert.Equal(t, []string{"one", "three", "four"}, Remove(src, "two"))
	assert.Equal(t, []string{"one", "two", "three"}, Remove(src, "four"))
	assert.Equal(t, empty, Remove(empty, "four"))
}

func TestMatchAny(t *testing.T) {
	src := []string{"one", "two", "three", "four"}
	assert.True(t, MatchAny(src, func(s string) bool {
		return s == "four"
	}))
	assert.False(t, MatchAny(src, func(s string) bool {
		return s == "five"
	}))
}

func TestIsAny(t *testing.T) {
	src := []string{"one", "two", "three", "four"}
	assert.True(t, IsAny(src, "four"))
	assert.False(t, IsAny(src, "FOUR"))
	assert.False(t, IsAny(src, "five"))
}
