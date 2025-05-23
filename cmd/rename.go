package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var validTypes = []string{"prefix", "suffix", "format", ""}

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename files in the specified directory based on a pattern",
	Long: `This command renames files in the specified directory. 
You can add a prefix, suffix, or replace a specific pattern in the filenames.`,
	Example: `rename-toolbox rename C:\folder a -t prefix ----> adds "a" prefix to all files in the folder
rename-toolbox rename C:\folder b -t suffix ----> adds "b" suffix to all files in the folder
rename-toolbox rename C:\folder txt -t format ----> changes all files format to .txt`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		pattern := args[1]
		renameType, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Println("Please enter a type")
			return
		}
		isValid := isValidType(renameType)
		if isValid {
			renameFiles(dir, pattern, renameType)
		}

	},
}

func isValidType(renameType string) bool {
	isValidType := false
	for _, validType := range validTypes {
		if renameType == validType {
			isValidType = true
			break
		}
	}

	if !isValidType {
		log.Printf("Error: Invalid type '%s'. Valid options are: prefix, suffix, format.\n", renameType)
		return false
	}
	return true
}

func init() {
	renameCmd.Flags().StringP("type", "t", "", "Type of renaming (prefix, suffix, format)")
	rootCmd.AddCommand(renameCmd)
}

func renameFiles(dir, pattern, renameType string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("Error reading directory: %v", err)
		return
	}
	for _, file := range files {
		oldName := file.Name()
		newName := applyPattern(oldName, pattern, renameType)
		err := os.Rename(filepath.Join(dir, oldName), filepath.Join(dir, newName))
		if err != nil {
			log.Printf("Failed to rename %s: %v\n", oldName, err)
			continue
		}
		log.Printf("Renamed: %s -> %s\n", oldName, newName)
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
	return replaceName(oldName, value)
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

func replaceName(oldName string, newName string) string {
	splitStr := strings.Split(oldName, ".")
	format := splitStr[1]
	return newName + "." + format
}
