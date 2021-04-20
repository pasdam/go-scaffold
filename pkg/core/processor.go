package core

import (
	"io"
)

// Processor represents an object that can process files
type Processor interface {

	// ProcessFile processes the specified file, using the reader to read the
	// content.
	ProcessFile(path string, reader io.Reader) error
}
