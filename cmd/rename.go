package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename files in the specified directory based on a pattern",
	Long: `This command renames files in the specified directory. 
    You can add a prefix, suffix, or replace a specific pattern in the filenames.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		pattern := args[1]
		renameType, err := cmd.Flags().GetString("type")
		if err != nil {
			fmt.Println("Please enter a type")
		}
		renameFiles(dir, pattern, renameType)
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
	renameCmd.Flags().StringP("type", "t", "prefix", "Type of renaming (prefix, suffix, format)")

}

func renameFiles(dir, pattern, renameType string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, file := range files {
		oldName := file.Name()
		newName := applyPattern(oldName, pattern, renameType)
		err := os.Rename(filepath.Join(dir, oldName), filepath.Join(dir, newName))
		if err != nil {
			fmt.Printf("Failed to rename %s: %v\n", oldName, err)
			continue
		}
		fmt.Printf("Renamed: %s -> %s\n", oldName, newName)
	}
}

func applyPattern(oldName string, value string, renameType string) string {
	switch renameType {
	case "prefix":
		return addPrefix(oldName, value)
	case "suffix":
		return addSuffix(oldName, value)
	case "format":
		return changeFormat(oldName, value)
	}
	return addPrefix(oldName, value)
}
func addPrefix(oldName string, prefix string) string {
	return prefix + oldName
}

func addSuffix(oldName string, suffix string) string {
	splitStr := strings.Split(oldName, ".")
	fileName, format := splitStr[0], splitStr[1]
	return fileName + suffix + format
}

func changeFormat(oldName string, newFormat string) string {
	fileName := strings.Split(oldName, ".")[0]
	return fileName + "." + newFormat
}
