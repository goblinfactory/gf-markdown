package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReportForValidFile(t *testing.T) {
	actual := GetReport("../testdata/cats/cat-names.md")
	expected := Report{
		"../testdata/cats/cat-names.md",
		true,
		[]LinkCheck{
			{Link{Text: "dog", RelPath: "../dogs/dog-names.md"}, true, "(ok)"},
			{Link{Text: "parent", RelPath: "../readme.md"}, true, "(ok)"},
			{Link{Text: "self", RelPath: "cat-names.md"}, true, "(ok)"},
			{Link{Text: "with errors", RelPath: "cat-names-err.md"}, true, "(ok)"},
		},
		0,
		Success,
		nil,
	}
	assert.Equal(t, expected, actual)
}

func TestFindLinks(t *testing.T) {
	src := []byte(`
		[link1](link/a/1.txt)
		[notlink] (link/a/1.txt)
		[link2](link/a/2.txt) ABC[link3](3.txt)DEF
	`)

	expected := []Link{
		{"link1", "link/a/1.txt"},
		{"link2", "link/a/2.txt"},
		{"link3", "3.txt"},
	}
	actual := FindLinks(src)
	assert.Equal(t, expected, actual)
}


TestPrintReportWhenNoBrokenLinksAndNotVerbosePrintsNothing() {

}

TestPrintReportWhenNoBrokenLinksAndVerbosePrintsAllLinks() {
	
}

TestPrintReportWhenBrokenLinksVerbose() {
	
}

TestPrintReportWhenBrokenLinksAndNotVerboseOnlyPrintsTheBrokenLinks() {
	
}