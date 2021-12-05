package tests

import (
	"os"
	"testing"
)

// RunIfEnvVarSet skips the test if the environment variable is not set.
func RunIfEnvVarSet(t *testing.T, envVars ...string) {
	for _, e := range envVars {
		v := os.Getenv(e)
		if v != "" {
			return
		}
	}
	t.Skip("Set '%w' to run this test.", envVars)
}
