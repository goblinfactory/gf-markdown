package tests

import (
	"fmt"
	"os"
	"testing"
)

// RunIfAnyEnvVarSet skips the test if the environment variable is not set.
func RunIfAnyEnvVarSet(t *testing.T, envVars ...string) {
	for _, e := range envVars {
		v := os.Getenv(e)
		if v != "" {
			return
		}
	}
	t.Skip(fmt.Sprintf("set any of these test env vars %v to run this test", envVars))
}

// RunIfEnvVarSet skips the test if the environment variable is not set.
func RunIfEnvVarSet(t *testing.T, envvar string) {
	v := os.Getenv(envvar)
	if v != "" {
		return
	}
	t.Skip(fmt.Sprintf("set env var [%s] to run this test", envvar))
}

// notes
// to debug an integration test from within visual studio code add the following
