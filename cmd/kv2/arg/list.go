package arg

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available secrets and versions",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := kv2.List()
		if err != nil {
			panic(err)
		}

		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
