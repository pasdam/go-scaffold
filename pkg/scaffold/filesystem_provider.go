package scaffold

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type fileSystemProvider struct {
	filesPath   []string
	filesInfo   []os.FileInfo
	templateDir string
}

func NewFileSystemProvider(templateDir string) (FileProvider, error) {
	provider := &fileSystemProvider{
		templateDir: formatPath(templateDir, ""),
	}
	err := provider.Reset()
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func (self *fileSystemProvider) Reset() error {
	return self.indexDir(self.templateDir)
}

func (self *fileSystemProvider) HasMoreFiles() bool {
	return len(self.filesPath) > 0
}

func (self *fileSystemProvider) NextFile() (string, io.ReadCloser, error) {
	nextFilePath := self.filesPath[0]
	nextFileInfo := self.filesInfo[0]

	listSize := len(self.filesPath)
	if listSize > 1 {
		self.filesPath = self.filesPath[1:len(self.filesPath)]
		self.filesInfo = self.filesInfo[1:len(self.filesInfo)]

	} else {
		self.filesPath = nil
		self.filesInfo = nil
	}

	if nextFileInfo.IsDir() {
		self.indexDir(nextFilePath)
		return self.NextFile()
	}

	reader, err := os.Open(nextFilePath)
	return strings.TrimPrefix(nextFilePath, self.templateDir), reader, err
}

func (self *fileSystemProvider) indexDir(dirPath string) error {
	filesInfo, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	filesPath := make([]string, len(filesInfo))
	for i := 0; i < len(filesInfo); i++ {
		filesPath[i] = formatPath(dirPath, filesInfo[i].Name())
	}

	// prepend slices
	self.filesInfo = append(filesInfo, self.filesInfo...)
	self.filesPath = append(filesPath, self.filesPath...)

	return nil
}

func formatPath(parent, child string) string {
	var separator string
	if !strings.HasSuffix(parent, "/") {
		separator = "/"
	}
	return fmt.Sprintf("%s%s%s", parent, separator, child)
}
