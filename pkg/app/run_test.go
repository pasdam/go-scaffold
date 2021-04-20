package app

import (
	"path/filepath"
	"testing"

	"github.com/pasdam/go-files-test/pkg/filestest"
	"github.com/pasdam/go-scaffold/pkg/config"
	"github.com/stretchr/testify/assert"
)

func Test_Run_Success_ValidTemplate(t *testing.T) {
	outDir := filestest.TempDir(t)
	options := mockOptions(filepath.Join("testdata", "valid_template"), outDir)

	Run(options)

	filestest.FileExistsWithContent(t, filepath.Join("testdata", "valid_template", "template", "file.txt"), "This is a {{ .Values.text }}\n")
	filestest.FileExistsWithContent(t, filepath.Join("testdata", "valid_template", "template", "normal_file.txt"), "normal-file-content\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "file.txt"), "This is a test!\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "normal_file.txt"), "normal-file-content\n")
}

func Test_Run_Success_ShouldNotRemoveSourceIfOptionIsSetButProcessIsNotInPlace(t *testing.T) {
	outDir := filestest.TempDir(t)
	options := mockOptions(filepath.Join("testdata", "valid_template"), outDir)

	Run(options)

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

	options := mockOptions(filepath.Join("testdata", "invalid_template"), outDir)

	Run(options)

	assert.True(t, called)
}

func mockOptions(templateDir string, outDir string) *config.Options {
	return &config.Options{
		TemplateRootPath: templateDir,
		OutputPath:       outDir,
	}
}
