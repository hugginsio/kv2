package arg

import (
	"os"

	"git.huggins.io/kv2/client"
	"github.com/spf13/cobra"
)

var kv2 *client.Client
var output string
var jsonOutput bool
var rootCmd = &cobra.Command{
	Use:   "kv2",
	Short: "kv2 provides an interface for your secrets manager",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		kv2 = client.NewClient("http://kv2")
	},
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
