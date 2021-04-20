package scaffold_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-files-test/pkg/filestest"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/stretchr/testify/assert"
)

func TestOutputFileProcessor_ProcessFile_Fail_ApplyTemplateFails(t *testing.T) {
	outDir := filestest.TempDir(t)

	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	processor := scaffold.NewOutputFileProcessor(
		"invalid-data",
		outDir,
	)
	err = processor.ProcessFile("template_file.tpl", file)

	assert.NotNil(t, err)
	filestest.PathDoesNotExist(t, filepath.Join(outDir, "template_file.tpl"))
}

func TestOutputFileProcessor_ProcessFile_Success_FileIsATemplate(t *testing.T) {
	outDir := filestest.TempDir(t)

	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	processor := scaffold.NewOutputFileProcessor(
		struct{ Text string }{Text: "*test*"},
		outDir,
	)
	err = processor.ProcessFile("template_file.tpl", file)

	assert.Nil(t, err)
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "template_file.tpl"), "This is a *test*\n")
}

func TestOutputFileProcessor_ProcessFile_Success_FileIsNotATemplate(t *testing.T) {
	outDir := filestest.TempDir(t)

	file, err := os.Open("testdata/regular_file.txt")
	assert.Nil(t, err)
	defer file.Close()

	processor := scaffold.NewOutputFileProcessor(
		nil,
		outDir,
	)
	err = processor.ProcessFile("regular_file.txt", file)

	assert.Nil(t, err)
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "regular_file.txt"), "regular-file-content\n")
}
