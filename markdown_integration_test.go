package main

import (
	"os"
	"testing"

	"github.com/goblinfactory/markdown/internal/tests"
	"github.com/goblinfactory/markdown/printer"
)

func TestPassingNoArgsShouldExitWith1(t *testing.T) {
	tests.RunIfEnvVarSet(t, "integration")

	args := []string{"markdown_integration_test.go"}

	p := printer.New(os.Stdout)
	retCode := run(os.Args, p)
	os.Exit(retCode)
}
