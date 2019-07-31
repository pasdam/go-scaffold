package scaffold_test

import (
	"os"
	"testing"

	"github.com/pasdam/go-project-template/pkg/scaffold"
	"github.com/pasdam/go-project-template/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func Test_ProcessFile_Fail_ApplyTemplateFails(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	file, err := os.Open("test/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	err = scaffold.ProcessFile(
		file,
		"invalid-data",
		outDir,
		"template_file.tpl",
	)

	assert.NotNil(t, err)
	verifyOutputFileDoesNotExist(t, outDir, "template_file.tpl")
}

func Test_ProcessFile_Fail_CantWriteOutputFile(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	file, err := os.Open("test/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	err = scaffold.ProcessFile(
		file,
		validData,
		outDir,
		"some-non-existent-folder/test/template_file.tpl",
	)

	assert.NotNil(t, err)
	verifyOutputFileDoesNotExist(t, outDir, "test/template_file.tpl")
}

func Test_ProcessFile_Success_FileIsATemplate(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	file, err := os.Open("test/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	err = scaffold.ProcessFile(
		file,
		validData,
		outDir,
		"template_file.tpl",
	)

	assert.Nil(t, err)
	testutils.FileExists(t, outDir+"template_file", "This is a *test*\n")
}

func Test_ProcessFile_Success_FileIsNotATemplate(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	file, err := os.Open("test/regular_file.txt")
	assert.Nil(t, err)
	defer file.Close()

	err = scaffold.ProcessFile(
		file,
		nil,
		outDir,
		"test/regular_file.txt",
	)

	assert.Nil(t, err)
	testutils.FileExists(t, outDir+"test/regular_file.txt", "regular-file-content\n")
}

func verifyOutputFileDoesNotExist(t *testing.T, outDir string, filePath string) {
	_, err := os.Stat(outDir + filePath)
	assert.NotNil(t, err)
}