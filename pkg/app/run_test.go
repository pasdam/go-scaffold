package app

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/otiai10/copy"
	"github.com/pasdam/go-files-test/pkg/filestest"
	"github.com/pasdam/mockit/matchers/argument"
	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func Test_Run_Success_ValidTemplate(t *testing.T) {
	outDir := filestest.TempDir(t)

	mockPrompt(t)
	oldArgs := mockArguments(filepath.Join("testdata", "valid_template"), outDir)
	defer func() { os.Args = oldArgs }()

	mockit.MockFunc(t, runInitScript).With(filepath.Join("testdata", "valid_template", ".go-scaffold", "initScript"), outDir, log.Fatal).Return()

	Run()

	filestest.FileExistsWithContent(t, filepath.Join("testdata", "valid_template", "file.txt.tpl"), "This is a {{ .text }}\n")
	filestest.FileExistsWithContent(t, filepath.Join("testdata", "valid_template", "normal_file.txt"), "normal-file-content\n")
	filestest.FileExistsWithContent(t, filepath.Join("testdata", "valid_template", ".go-scaffold", "prompts.yaml"), "prompts:\n  - name: text\n    type: string\n    default: default-text\n    message: Enter text value\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "file.txt"), "This is a test!\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "normal_file.txt"), "normal-file-content\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, ".go-scaffold", "prompts.yaml"), "prompts:\n  - name: text\n    type: string\n    default: default-text\n    message: Enter text value\n")
}

func Test_Run_Success_ShouldNotRemoveSourceIfOptionIsSetButProcessIsNotInPlace(t *testing.T) {
	outDir := filestest.TempDir(t)

	mockPrompt(t)
	oldArgs := mockArguments(filepath.Join("testdata", "valid_template"), outDir)
	defer func() { os.Args = oldArgs }()

	mockit.MockFunc(t, runInitScript).With(filepath.Join("testdata", "valid_template", ".go-scaffold", "initScript"), outDir, log.Fatal).Return()

	Run()

	filestest.FileExistsWithContent(t, filepath.Join("testdata", "valid_template", "file.txt.tpl"), "This is a {{ .text }}\n")
	filestest.FileExistsWithContent(t, filepath.Join("testdata", "valid_template", "normal_file.txt"), "normal-file-content\n")
	filestest.FileExistsWithContent(t, filepath.Join("testdata", "valid_template", ".go-scaffold", "prompts.yaml"), "prompts:\n  - name: text\n    type: string\n    default: default-text\n    message: Enter text value\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "file.txt"), "This is a test!\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "normal_file.txt"), "normal-file-content\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, ".go-scaffold", "prompts.yaml"), "prompts:\n  - name: text\n    type: string\n    default: default-text\n    message: Enter text value\n")
}

func Test_Run_Fail_ErrorWhileProcessingFiles(t *testing.T) {
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

func Test_Run_Fail_TemplateDirIsSameAsOutDir(t *testing.T) {
	outDir := filestest.TempDir(t)

	oldArgs := mockArguments(outDir, outDir)
	defer func() { os.Args = oldArgs }()

	copy.Copy(filepath.Join("testdata", "valid_template"), outDir)

	captor := argument.Captor{}
	mockit.MockFunc(t, log.Fatal).With(captor.Capture).Return()

	Run()

	assert.Equal(t, "Can't generate file in the input folder, please specify an output directory", captor.Value.([]interface{})[0].(string))
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "file.txt.tpl"), "This is a {{ .text }}\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, "normal_file.txt"), "normal-file-content\n")
	filestest.FileExistsWithContent(t, filepath.Join(outDir, ".go-scaffold", "prompts.yaml"), "prompts:\n  - name: text\n    type: string\n    default: default-text\n    message: Enter text value\n")
}

func mockPrompt(t *testing.T) {
	data := make(map[string]interface{})
	data["text"] = "test!"
	mockit.MockFunc(t, parseAndRunPrompts).With(argument.Any, log.Fatal).Return(data)
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
