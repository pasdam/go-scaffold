package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "go-scaffold",
		Short: "go-scaffold generates files and projects from templates",
	}

	generateCmd := &cobra.Command{
		Use:   "generate [flags] template-dir output-dir",
		Short: "generates files and projects from templates",
		Run:   generate,
		Args:  cobra.ExactArgs(2),
	}
	generateCmd.Flags().StringArrayP("values", "f", nil, "overriding values file")
	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
