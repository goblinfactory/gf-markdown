package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotProvidingArgsShouldReturnWithArgumentErrorExitCode(t *testing.T) {
	args := []string{}
	p := NewTestWriter()
	retCode := RunFromArgs(args, p)
	assert.Equal(t, Err2Arguments, retCode)
}
