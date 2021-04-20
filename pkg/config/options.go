package config

import (
	"path/filepath"

	"github.com/jessevdk/go-flags"
)

// Options contains the app run configuration
type Options struct {
	OutputPath       flags.Filename `short:"o" long:"output" description:"path of the output dir, if not specified the template will be generated in place" default:"./"`
	TemplateRootPath flags.Filename `short:"t" long:"template" description:"path of the template root folder" default:"./"`
	Values           []string       `short:"f" long:"values" description:"specify values in a YAML file"`
}

// ParseCLIOptions parses the options from the CLI
func ParseCLIOptions() (*Options, error) {
	var options Options
	_, err := flags.Parse(&options)
	if err != nil {
		return nil, err
	}
	return &options, nil
}

// ManifestPath returns the path of the template's manifest
func (o *Options) ManifestPath() string {
	return filepath.Join(string(o.TemplateRootPath), "Manifest.yaml")
}

// TemplateDirPath returns the path of the template dir
func (o *Options) TemplateDirPath() string {
	return filepath.Join(string(o.TemplateRootPath), "template")
}

// ValuesPath returns the path of the template values definition
func (o *Options) ValuesPath() string {
	return filepath.Join(string(o.TemplateRootPath), "values.yaml")
}
