package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jessevdk/go-flags"
	"github.com/pasdam/go-scaffold/pkg/config"
	"github.com/stretchr/testify/assert"
)

func Test_ParseCLIOptions_fail_shouldReturnErrorIfAnInvalidParameterIsSpecifiedSpecified(t *testing.T) {
	templateDir := "some-template-dir"
	oldArgs := mockArguments(false, templateDir, "")
	defer func() { os.Args = oldArgs }()
	os.Args = append(os.Args, "--invalid-param")

	options, err := config.ParseCLIOptions()

	assert.NotNil(t, err)
	assert.Equal(t, "unknown flag `invalid-param'", err.Error())
	assert.Nil(t, options)
}

func Test_ParseCLIOptions_success_shouldUseDefaultOutputPath(t *testing.T) {
	templateDir := "some-template-dir"
	oldArgs := mockArguments(false, templateDir, "")
	defer func() { os.Args = oldArgs }()

	options, err := config.ParseCLIOptions()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, "./", string(options.OutputPath))
	assert.Equal(t, templateDir, string(options.TemplateRootPath))
}

func Test_ParseCLIOptions_success_shouldUseDefaultTemplatePath(t *testing.T) {
	outDir := "some-output-dir"
	oldArgs := mockArguments(false, "", outDir)
	defer func() { os.Args = oldArgs }()

	options, err := config.ParseCLIOptions()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, outDir, string(options.OutputPath))
	assert.Equal(t, "./", string(options.TemplateRootPath))
}

func Test_ParseCLIOptions_success_shouldParseOptionsWithLongFlags(t *testing.T) {
	templateDir := "some-template-dir"
	outDir := "some-output-dir"
	oldArgs := mockArguments(true, templateDir, outDir)
	defer func() { os.Args = oldArgs }()

	options, err := config.ParseCLIOptions()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, templateDir, string(options.TemplateRootPath))
	assert.Equal(t, outDir, string(options.OutputPath))
}

func Test_ParseCLIOptions_success_shouldParseOptionsWithShortFlags(t *testing.T) {
	templateDir := "some-template-dir"
	outDir := "some-output-dir"
	oldArgs := mockArguments(false, templateDir, outDir)
	defer func() { os.Args = oldArgs }()

	options, err := config.ParseCLIOptions()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, templateDir, string(options.TemplateRootPath))
	assert.Equal(t, outDir, string(options.OutputPath))
}

func TestOptions_ManifestPath(t *testing.T) {
	type fields struct {
		OutputPath       flags.Filename
		TemplateRootPath flags.Filename
		Values           []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Should return expected value",
			fields: fields{TemplateRootPath: "manifest-test"},
			want:   filepath.Join("manifest-test", "Manifest.yaml"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &config.Options{
				OutputPath:       tt.fields.OutputPath,
				TemplateRootPath: tt.fields.TemplateRootPath,
				Values:           tt.fields.Values,
			}
			if got := o.ManifestPath(); got != tt.want {
				t.Errorf("Options.ManifestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_TemplateDirPath(t *testing.T) {
	type fields struct {
		OutputPath       flags.Filename
		TemplateRootPath flags.Filename
		Values           []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Should return expected value",
			fields: fields{TemplateRootPath: "template-test"},
			want:   filepath.Join("template-test", "template"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &config.Options{
				OutputPath:       tt.fields.OutputPath,
				TemplateRootPath: tt.fields.TemplateRootPath,
				Values:           tt.fields.Values,
			}
			if got := o.TemplateDirPath(); got != tt.want {
				t.Errorf("Options.TemplateDirPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_ValuesPath(t *testing.T) {
	type fields struct {
		OutputPath       flags.Filename
		TemplateRootPath flags.Filename
		Values           []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Should return expected value",
			fields: fields{TemplateRootPath: "values-test"},
			want:   filepath.Join("values-test", "values.yaml"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &config.Options{
				OutputPath:       tt.fields.OutputPath,
				TemplateRootPath: tt.fields.TemplateRootPath,
				Values:           tt.fields.Values,
			}
			if got := o.ValuesPath(); got != tt.want {
				t.Errorf("Options.ValuesPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockArguments(useLongFlags bool, templateDir string, outDir string) []string {
	oldArgs := os.Args

	os.Args = make([]string, 7)
	os.Args[0] = ""
	if templateDir != "" {
		if useLongFlags {
			os.Args[1] = "--template"
		} else {
			os.Args[1] = "-t"
		}
		os.Args[2] = templateDir
	}
	if outDir != "" {
		if useLongFlags {
			os.Args[3] = "--output"
		} else {
			os.Args[3] = "-o"
		}
		os.Args[4] = outDir
	}

	return oldArgs
}
