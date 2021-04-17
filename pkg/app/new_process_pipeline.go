package app

import (
	"github.com/pasdam/go-scaffold/pkg/processors"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
)

func newProcessPipeline(config interface{}, srcDir string, outDir string, templateHelper *scaffold.TemplateHelper, errHandler func(v ...interface{})) processors.Processor {
	outPipeline, err := newOutputPipeline(config, outDir, templateHelper)
	if err != nil {
		errHandler("Error while creating the processing pipeline. ", err)
		return nil
	}
	return outPipeline
}
