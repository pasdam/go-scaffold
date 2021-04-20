package app

import (
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/processors"
)

func newOutputPipeline(config interface{}, outDir string) (processors.Processor, error) {
	filter := filters.NewNoOpFilter()
	outProcessor := processors.NewOutputFileProcessor(config, outDir)
	return processors.NewFilterProcessor(filter, outProcessor), nil
}
