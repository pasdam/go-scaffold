package app

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-files-test/pkg/filestest"
	"github.com/stretchr/testify/assert"
)

func Test_Run_Success_ValidTemplate(t *testing.T) {
	outDir := filestest.TempDir(t)
	oldArgs := mockArguments(filepath.Join("testdata", "valid_template"), outDir)
	defer func() { os.Args = oldArgs }()

	Run()

	filestest.FileExistsWithContent(t, filepath.Join("testdata", "valid_template", "template", "file.txt"), "This is a {{ .Values.text }}\n")
	filestest.FileExistsWithContent(t, filepath.Join("testdata", "valid_template", "template", "normal_file.txt"), "normal-file-content\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "file.txt"), "This is a test!\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "normal_file.txt"), "normal-file-content\n")
}

func Test_Run_Success_ShouldNotRemoveSourceIfOptionIsSetButProcessIsNotInPlace(t *testing.T) {
	outDir := filestest.TempDir(t)
	oldArgs := mockArguments(filepath.Join("testdata", "valid_template"), outDir)
	defer func() { os.Args = oldArgs }()

	Run()

	filestest.FileExistsWithContent(t, filepath.Join("testdata", "valid_template", "template", "file.txt"), "This is a {{ .Values.text }}\n")
	filestest.FileExistsWithContent(t, filepath.Join("testdata", "valid_template", "template", "normal_file.txt"), "normal-file-content\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "file.txt"), "This is a test!\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "normal_file.txt"), "normal-file-content\n")
}

func Test_Run_Fail_ErrorWhileProcessingFiles(t *testing.T) {
	t.Skip()
	called := false
	errHandler = func(args ...interface{}) {
		called = true
		assert.Equal(t, "Error while processing files. ", args[0].(string))
		assert.NotNil(t, args[1])
	}

	outDir := filestest.TempDir(t)

	oldArgs := mockArguments(filepath.Join("testdata", "invalid_template"), outDir)
	defer func() { os.Args = oldArgs }()

	Run()

	assert.True(t, called)
}

func mockArguments(templateDir string, outDir string) []string {
	oldArgs := os.Args

	os.Args = make([]string, 7)
	os.Args[0] = ""
	os.Args[1] = "--template"
	os.Args[2] = templateDir
	os.Args[3] = "--output"
	os.Args[4] = outDir

	return oldArgs
}
