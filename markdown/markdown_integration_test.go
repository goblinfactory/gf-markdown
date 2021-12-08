package markdown

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
)

func TestPrintReportWhenNoBrokenLinksAndNotVerbosePrintsNothing(t *testing.T) {
	args := []string{"testdata/nobroken/README.md"}
	p := NewTestWriter()
	retCode := RunFromArgs(args, p)
	p.Flush()
	assert.Equal(t, Success, retCode)
	assert.True(t, len(p.history) == 0)
}

func TestPrintReportWhenNoBrokenLinksAndVerbosePrintsAllLinks(t *testing.T) {
	args := []string{"testdata/nobroken/README.md", "-v"}
	p := NewTestWriter()
	retCode := RunFromArgs(args, p)
	p.Flush()
	assert.Equal(t, Success, retCode)
	cupaloy.SnapshotT(t, p.history)
}

func TestPrintReportWhenBrokenLinksVerbose(t *testing.T) {
	args := []string{"testdata/brokenlinks/README.md", "-v"}
	p := NewTestWriter()
	retCode := RunFromArgs(args, p)
	p.Flush()
	assert.Equal(t, Err3Links, retCode)
	cupaloy.SnapshotT(t, p.history)
}

func TestPrintReportWhenBrokenLinksAndNotVerboseOnlyPrintsTheBrokenLinks(t *testing.T) {
	args := []string{"testdata/brokenlinks/README.md"}
	p := NewTestWriter()
	retCode := RunFromArgs(args, p)
	p.Flush()
	assert.Equal(t, Err3Links, retCode)
	cupaloy.SnapshotT(t, p.history)
}
