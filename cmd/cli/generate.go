package main

import (
	"fmt"

	"github.com/pasdam/go-scaffold/pkg/app"
	"github.com/pasdam/go-scaffold/pkg/config"
	"github.com/spf13/cobra"
)

func generate(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		cmd.PrintErr("Invalid number of command arguments")
		cmd.Help()
		return
	}

	values, err := cmd.Flags().GetStringArray("values")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	options := &config.Options{
		OutputPath:       args[1],
		TemplateRootPath: args[0],
		Values:           values,
	}

	app.Run(options)
}
