package scaffold_test

import (
	"errors"
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pasdam/go-io-utilx/pkg/ioutilx"
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_NewFileSystemProvider_Fail_FolderDoesNotExist(t *testing.T) {
	var filter filters.Filter
	processor := newMockFileProcessor()

	provider := scaffold.NewFileSystemProvider("some-non-existing-folder")
	err := provider.ProvideFiles(filter, processor)
	assert.Equal(t, "open some-non-existing-folder: no such file or directory", err.Error())
}

func TestFileSystemProvider_ProvideFiles_Fail_ShouldProcessAllFileIfNoFilterIsSpecified(t *testing.T) {
	var filter filters.Filter
	processor := newMockFileProcessor()
	expectedErr := errors.New("some-error")
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(expectedErr)

	provider := scaffold.NewFileSystemProvider(filepath.Join("testdata", "file_system_provider"))
	actualErr := provider.ProvideFiles(filter, processor)

	assert.Equal(t, expectedErr, actualErr)
	assert.Equal(t, 1, len(processor.ReadersMap))
}

func TestFileSystemProvider_ProvideFiles_Success_ShouldProcessAllFileIfNoFilterIsSpecified(t *testing.T) {
	var filter filters.Filter
	processor := newMockFileProcessor()
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(nil)

	provider := scaffold.NewFileSystemProvider(filepath.Join("testdata", "file_system_provider"))
	err := provider.ProvideFiles(filter, processor)
	assert.Nil(t, err)

	verifyProcessedFile(t, processor, "file0", "file0-content\n")
	verifyProcessedFile(t, processor, "file1", "file1-content\n")
	verifyProcessedFile(t, processor, filepath.Join("test_folder", "fileA"), "fileA-content\n")
	assert.Equal(t, 0, len(processor.ReadersMap))
}

func TestFileSystemProvider_ProvideFiles_Success_ShouldProcessAllFileIfFilterAcceptsAll(t *testing.T) {
	filter := &mockFilter{File: "no-file-will-match"}
	processor := newMockFileProcessor()
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(nil)

	provider := scaffold.NewFileSystemProvider(filepath.Join("testdata", "file_system_provider"))
	err := provider.ProvideFiles(filter, processor)
	assert.Nil(t, err)

	verifyProcessedFile(t, processor, "file0", "file0-content\n")
	verifyProcessedFile(t, processor, "file1", "file1-content\n")
	verifyProcessedFile(t, processor, filepath.Join("test_folder", "fileA"), "fileA-content\n")
	assert.Equal(t, 0, len(processor.ReadersMap))
}

func TestFileSystemProvider_ProvideFiles_Success_ShouldNotProcessFilesIgnoredByTheFilter(t *testing.T) {
	filter := &mockFilter{File: "file0"}
	processor := newMockFileProcessor()
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(nil)

	provider := scaffold.NewFileSystemProvider(filepath.Join("testdata", "file_system_provider"))
	err := provider.ProvideFiles(filter, processor)
	assert.Nil(t, err)

	verifyProcessedFile(t, processor, "file1", "file1-content\n")
	verifyProcessedFile(t, processor, filepath.Join("test_folder", "fileA"), "fileA-content\n")
	assert.Equal(t, 0, len(processor.ReadersMap))
}

func verifyProcessedFile(t *testing.T, processor *mockFileProcessor, filePath string, content string) {
	processor.AssertCalled(t, "ProcessFile", filePath, mock.Anything)
	assert.Equal(t, content, processor.ReadersMap[filePath])
	delete(processor.ReadersMap, filePath)
}

type mockFilter struct {
	File    string
	Include bool
}

func (m *mockFilter) Accept(filePath string) bool {
	if m.Include {
		return strings.Contains(filePath, m.File)

	} else {
		return !strings.Contains(filePath, m.File)
	}
}

type mockFileProcessor struct {
	mock.Mock

	ReadersMap map[string]string
}

func newMockFileProcessor() *mockFileProcessor {
	return &mockFileProcessor{
		ReadersMap: make(map[string]string),
	}
}

func (p *mockFileProcessor) ProcessFile(filePath string, reader io.Reader) error {
	p.ReadersMap[filePath] = ioutilx.ReaderToString(reader)
	args := p.Called(filePath, reader)
	return args.Error(0)
}
