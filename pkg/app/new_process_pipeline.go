package app

import (
	"github.com/pasdam/go-scaffold/pkg/processors"
)

func newProcessPipeline(config interface{}, srcDir string, outDir string, errHandler func(v ...interface{})) processors.Processor {
	outPipeline, err := newOutputPipeline(config, outDir)
	if err != nil {
		errHandler("Error while creating the processing pipeline. ", err)
		return nil
	}
	return outPipeline
}
