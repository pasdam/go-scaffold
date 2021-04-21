package app

import (
	"text/template"

	"github.com/pasdam/go-scaffold/pkg/core"
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/processors"
)

func newOutputPipeline(config interface{}, outDir string) (core.Processor, error) {
	filter := filters.NewNoOpFilter()
	funcMap := template.FuncMap{}
	outProcessor := processors.NewOutputFileProcessor(config, outDir, funcMap)
	return processors.NewFilterProcessor(filter, outProcessor), nil
}
