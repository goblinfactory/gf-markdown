package main

import (
	"testing"

	"github.com/goblinfactory/markdown/internal/tests"
	"github.com/goblinfactory/markdown/printer"
	"github.com/stretchr/testify/assert"
)

func TestPassingNoArgsShouldExitWith1(t *testing.T) {
	tests.RunIfEnvVarSet(t, "integration")

	args := []string{"markdown_integration_test.go"}

	p := printer.NewTestWriter()
	retCode := run(args, p)

	assert.Equal()
}
