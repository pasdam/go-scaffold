package config

import (
	"path/filepath"

	"github.com/jessevdk/go-flags"
)

// Options contains the app run configuration
type Options struct {
	OutputPath   flags.Filename `short:"o" long:"output" description:"Path of the output dir, if not specified the template will be generated in place" default:"./"`
	TemplatePath flags.Filename `short:"t" long:"template" description:"Path of the template folder" default:"./"`
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

// ConfigDirPath returns the path of the template's configuration dir
func (o *Options) ConfigDirPath() string {
	return filepath.Join(string(o.TemplatePath), ".go-scaffold")
}

// InitScriptPath returns the path of the template's init script
func (o *Options) InitScriptPath() string {
	return filepath.Join(o.ConfigDirPath(), "initScript")
}

// PromptsConfigPath returns the path of the template's  prompts configuration
func (o *Options) PromptsConfigPath() string {
	return filepath.Join(o.ConfigDirPath(), "prompts.yaml")
}
