package arg

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
