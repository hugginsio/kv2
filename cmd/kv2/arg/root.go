package arg

import (
	"os"

	"github.com/spf13/cobra"
)

var output string
var jsonOutput bool
var rootCmd = &cobra.Command{
	Use:   "kv2",
	Short: "kv2 provides an interface for your secrets manager",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "JSON output")
}
