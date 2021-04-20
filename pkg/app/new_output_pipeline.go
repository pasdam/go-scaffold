package app

import (
	"github.com/pasdam/go-scaffold/pkg/core"
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/processors"
)

func newOutputPipeline(config interface{}, outDir string) (core.Processor, error) {
	filter := filters.NewNoOpFilter()
	outProcessor := processors.NewOutputFileProcessor(config, outDir)
	return processors.NewFilterProcessor(filter, outProcessor), nil
}
