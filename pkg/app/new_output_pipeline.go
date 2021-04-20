package app

import (
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/processors"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
)

func newOutputPipeline(config interface{}, outDir string) (processors.Processor, error) {
	filter := filters.NewNoOpFilter()
	outProcessor := scaffold.NewOutputFileProcessor(config, outDir)
	return processors.NewFilterProcessor(filter, outProcessor), nil
}
