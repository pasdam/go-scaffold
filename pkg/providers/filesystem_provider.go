package providers

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/processors"
)

type fileSystemProvider struct {
	filesPath []string
	filesInfo []os.FileInfo
	inputDir  string
}

// NewFileSystemProvider creates a new instance of a FileProvider that reads
// file from the filesystem.
func NewFileSystemProvider(inputDir string) FileProvider {
	return &fileSystemProvider{
		inputDir: inputDir,
	}
}

func (p *fileSystemProvider) ProvideFiles(filesFilter filters.Filter, processor processors.Processor) error {
	err := p.indexDir(p.inputDir)
	if err != nil {
		return err
	}

	for len(p.filesPath) > 0 {
		filePath, reader := p.nextFile()
		if reader != nil {

			relativePath, _ := filepath.Rel(p.inputDir, filePath)
			if filesFilter == nil || filesFilter.Accept(relativePath) {
				err = processor.ProcessFile(relativePath, reader)
				if err != nil {
					// TODO: clean output folder
					reader.Close()
					log.Printf("Error while processing file %s\n", relativePath)
					return err
				}
			}

			reader.Close()
		}
	}
	return nil
}

func (p *fileSystemProvider) nextFile() (string, io.ReadCloser) {
	nextFileInfo := p.filesInfo[0]
	nextFilePath := p.filesPath[0]

	listSize := len(p.filesPath)
	if listSize > 1 {
		p.filesPath = p.filesPath[1:len(p.filesPath)]
		p.filesInfo = p.filesInfo[1:len(p.filesInfo)]

	} else {
		p.filesPath = nil
		p.filesInfo = nil
	}

	var reader io.ReadCloser

	if nextFileInfo.IsDir() {
		p.indexDir(nextFilePath)
		return nextFilePath, nil
	}

	reader, _ = os.Open(nextFilePath)
	return nextFilePath, reader
}

func (p *fileSystemProvider) indexDir(dirPath string) error {
	filesInfo, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	acceptedInfo := make([]os.FileInfo, 0, len(filesInfo))
	acceptedPaths := make([]string, 0, len(filesInfo))
	for i := 0; i < len(filesInfo); i++ {
		filePath := filepath.Join(dirPath, filesInfo[i].Name())
		acceptedInfo = append(acceptedInfo, filesInfo[i])
		acceptedPaths = append(acceptedPaths, filePath)
	}
	p.filesInfo = append(acceptedInfo, p.filesInfo...)
	p.filesPath = append(acceptedPaths, p.filesPath...)

	return nil
}
