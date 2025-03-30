package arg

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a secret to a new version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
