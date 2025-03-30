package arg

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new secret",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
