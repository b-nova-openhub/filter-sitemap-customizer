package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:           "fisicus",
		Short:         "fisicus – command-line tool to manage {PLEASE DESCRIBE}",
		Long:          ``,
		Version:       "0.0.0",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
