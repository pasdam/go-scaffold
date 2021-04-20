package processors

import (
	"io"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/core"
)

type outputFileProcessor struct {
	outDir             string
	tempplateProcessor core.Processor
}

// NewOutputFileProcessor creates a new instance of a FileProcessor that process
// templates and creates the output files.
// THe variables to use for the template are in config.
func NewOutputFileProcessor(config interface{}, outDir string) core.Processor {
	writeProcessor := NewWriteProcessor()
	templateProcessor := NewTemplateProcessor(config, writeProcessor)

	return &outputFileProcessor{
		outDir:             outDir,
		tempplateProcessor: templateProcessor,
	}
}

func (p *outputFileProcessor) ProcessFile(filePath string, reader io.Reader) error {
	outPath := filepath.Join(p.outDir, filePath)

	return p.tempplateProcessor.ProcessFile(outPath, reader)
}
