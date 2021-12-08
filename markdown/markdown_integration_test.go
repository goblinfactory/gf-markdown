package markdown

import (
	"testing"

	"github.com/goblinfactory/gf-markdown/markdown/internal/tests"
)

func TestPrintReportWhenNoBrokenLinksAndNotVerbosePrintsNothing(t *testing.T) {
	tests.RunIfEnvVarSet(t, "integration")
}

func TestPrintReportWhenNoBrokenLinksAndVerbosePrintsAllLinks(t *testing.T) {
	tests.RunIfEnvVarSet(t, "integration")
}

func TestPrintReportWhenBrokenLinksVerbose(t *testing.T) {
	tests.RunIfEnvVarSet(t, "integration")
}

func TestPrintReportWhenBrokenLinksAndNotVerboseOnlyPrintsTheBrokenLinks(t *testing.T) {
	tests.RunIfEnvVarSet(t, "integration")
}
