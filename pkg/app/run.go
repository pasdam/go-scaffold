package app

import (
	"log"

	"github.com/pasdam/go-scaffold/pkg/scaffold"
)

var errHandler = log.Fatal

// Run starts the app
func Run() {
	options := readOptions(errHandler)
	data := parseAndRunPrompts(options.PromptsConfigPath(), errHandler)

	if options.TemplatePath == options.OutputPath {
		log.Fatal("Can't generate file in the input folder, please specify an output directory")
		return
	}

	templateHelper := &scaffold.TemplateHelper{}

	fileProcessor := newProcessPipeline(
		data,
		string(options.TemplatePath),
		string(options.OutputPath),
		templateHelper,
		errHandler,
	)

	provider := scaffold.NewFileSystemProvider(string(options.TemplatePath))
	err := provider.ProvideFiles(nil, fileProcessor)
	if err != nil {
		errHandler("Error while processing files. ", err)
		return
	}

	runInitScript(options.InitScriptPath(), string(options.OutputPath), errHandler)
}
