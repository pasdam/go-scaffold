package processors

import (
	"io"

	"github.com/pasdam/go-scaffold/pkg/core"
	"github.com/pasdam/go-scaffold/pkg/templates"
)

type templateProcessor struct {
	data          interface{}
	nextProcessor core.Processor
}

// NewTemplateProcessor creates a new Processor that handles template files
func NewTemplateProcessor(data interface{}, nextProcessor core.Processor) core.Processor {
	return &templateProcessor{
		data:          data,
		nextProcessor: nextProcessor,
	}
}

func (p *templateProcessor) ProcessFile(filePath string, reader io.Reader) error {
	var err error
	reader, err = templates.ProcessTemplate(reader, p.data)
	if err != nil {
		return err
	}
	return p.nextProcessor.ProcessFile(filePath, reader)
}
