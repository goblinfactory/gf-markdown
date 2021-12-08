// Package tests enables us to run (or bypass) the integration tests by checking if an environment variable has been set or not.
// VSCode is currently unable to debug Integration tests using the go convention, named {package}_integration_test.go
// hence this approach.
// In order to run the integration tests from within vscode, you need to set appropriate environment variables.
// add the following to your settings.json
// ```json
//   "go.testEnvVars": {
// 	"integration": "yes"
// },
// ```
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
