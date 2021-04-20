package config

import (
	"path/filepath"
)

// Options contains the app run configuration
type Options struct {
	OutputPath       string
	TemplateRootPath string
	Values           []string
	ManifestName     string
}

// ManifestPath returns the path of the template's manifest
func (o *Options) ManifestPath() string {
	if len(o.ManifestName) == 0 {
		return filepath.Join(string(o.TemplateRootPath), "Manifest.yaml")
	}
	return filepath.Join(string(o.TemplateRootPath), o.ManifestName)
}

// TemplateDirPath returns the path of the template dir
func (o *Options) TemplateDirPath() string {
	return filepath.Join(string(o.TemplateRootPath), "template")
}

// ValuesPath returns the path of the template values definition
func (o *Options) ValuesPath() string {
	return filepath.Join(string(o.TemplateRootPath), "values.yaml")
}
