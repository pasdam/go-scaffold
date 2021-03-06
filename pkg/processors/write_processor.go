package processors

import (
	"io"

	"github.com/pasdam/go-io-utilx/pkg/ioutilx"
)

type writeProcessor struct{}

// NewWriteProcessor creates a new instance of a Processor that write each the
// content from the reader to the specified path
func NewWriteProcessor() Processor {
	return &writeProcessor{}
}

func (p *writeProcessor) ProcessFile(filePath string, reader io.Reader) error {
	return ioutilx.ReaderToFile(reader, filePath)
}
