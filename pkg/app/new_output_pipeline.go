package app

import (
	"errors"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/pasdam/go-scaffold/pkg/core"
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/processors"
)

func newOutputPipeline(config interface{}, outDir string) (core.Processor, error) {
	filter := filters.NewNoOpFilter()
	funcMap := template.FuncMap{
		"camelcase": strcase.ToCamel,
		"cat": func(values ...string) string {
			return strings.Join(values, " ")
		},
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		"replace": func(target string, replacement string, source string) string {
			return strings.ReplaceAll(source, target, replacement)
		},
	}
	outProcessor := processors.NewOutputFileProcessor(config, outDir, funcMap)
	return processors.NewFilterProcessor(filter, outProcessor), nil
}
