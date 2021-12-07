package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Expected results
// ----------------
// Success
// Err1Unhandled
// Err2Arguments
// Err3Links

func TestPassingNoArgsShouldExitWithErr2(t *testing.T) {

	args := []string{}
	p := NewTestWriter()
	retCode := RunFromArgs(args, p)
	assert.Equal(t, Err2Arguments, retCode)
}

func TestIgnoringFiles(t *testing.T) {

}

// func TestIgnoreFolders(t *testing.T) {
// 	panic("can you see me")
// 	// p := &Printer{}
// 	// args := []string{"testdata/**/*.md", ""}
// 	// pass := Main(p, args)

// }
