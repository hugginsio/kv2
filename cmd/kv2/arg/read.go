package arg

import (
	"fmt"

	"git.huggins.io/kv2/api"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read <secret>",
	Short: "Read the latest version of a secret",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res, err := kv2.Read(api.ReadSecretRequest{Key: args[0]})
		if err != nil {
			panic(err)
		}

		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
