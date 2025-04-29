package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var numberCmd = &cobra.Command{
	Use:   "number",
	Short: "Labels all files in a directory with numbers",
	Long: `Labels all files in a directory with numbers.
Example: a.txt b.txt c.txt => 1a.txt 2b.txt 3c.txt`,
	Example: "rename-toolbox number C:/folder",
	DisableFlagsInUseLine: true,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		numberFiles(dir)
	},
}

func init() {
	rootCmd.AddCommand(numberCmd)
}

func numberFiles(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("Error reading directory: %v", err)
		return
	}
	for i, file := range files {
		oldName := file.Name()
		parts := strings.Split(oldName, ".")
		fileName, format := parts[0], parts[1]
		newName := strconv.Itoa(i) + fileName + "." + format
		err := os.Rename(filepath.Join(dir, oldName), filepath.Join(dir, newName))
		if err != nil {
			log.Printf("Failed to rename %s: %v\n", oldName, err)
			continue
		}
		log.Printf("Renamed: %s -> %s\n", oldName, newName)

	}

}
