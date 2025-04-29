package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "rename-toolbox",
    Short: "A CLI tool to batch rename files in a directory",
    Long: `File Renamer is a simple tool to rename multiple files in a directory
    based on the specified pattern.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func init() {
    rootCmd.AddCommand(renameCmd)
}


