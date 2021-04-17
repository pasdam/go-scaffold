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
	assert.Equal(t, templateDir, string(options.TemplatePath))
}

func Test_ParseCLIOptions_success_shouldUseDefaultTemplatePath(t *testing.T) {
	outDir := "some-output-dir"
	oldArgs := mockArguments(false, "", outDir)
	defer func() { os.Args = oldArgs }()

	options, err := config.ParseCLIOptions()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, outDir, string(options.OutputPath))
	assert.Equal(t, "./", string(options.TemplatePath))
}

func Test_ParseCLIOptions_success_shouldParseOptionsWithLongFlags(t *testing.T) {
	templateDir := "some-template-dir"
	outDir := "some-output-dir"
	oldArgs := mockArguments(true, templateDir, outDir)
	defer func() { os.Args = oldArgs }()

	options, err := config.ParseCLIOptions()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, templateDir, string(options.TemplatePath))
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
	assert.Equal(t, templateDir, string(options.TemplatePath))
	assert.Equal(t, outDir, string(options.OutputPath))
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

func TestOptions_ConfigDirPath(t *testing.T) {
	type fields struct {
		TemplatePath flags.Filename
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Should return expected value",
			fields: fields{TemplatePath: "some-path-1"},
			want:   filepath.Join("some-path-1", ".go-scaffold"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &config.Options{
				TemplatePath: tt.fields.TemplatePath,
			}
			if got := o.ConfigDirPath(); got != tt.want {
				t.Errorf("Options.ConfigDirPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_InitScriptPath(t *testing.T) {
	type fields struct {
		TemplatePath flags.Filename
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Should return expected value",
			fields: fields{TemplatePath: "some-path-1"},
			want:   filepath.Join("some-path-1", ".go-scaffold", "initScript"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &config.Options{
				TemplatePath: tt.fields.TemplatePath,
			}
			if got := o.InitScriptPath(); got != tt.want {
				t.Errorf("Options.InitScriptPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_PromptsConfigPath(t *testing.T) {
	type fields struct {
		TemplatePath flags.Filename
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Should return expected value",
			fields: fields{TemplatePath: "some-path-1"},
			want:   filepath.Join("some-path-1", ".go-scaffold", "prompts.yaml"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &config.Options{
				TemplatePath: tt.fields.TemplatePath,
			}
			if got := o.PromptsConfigPath(); got != tt.want {
				t.Errorf("Options.PromptsConfigPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
