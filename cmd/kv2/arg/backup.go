package arg

import (
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup the secrets database",
	Run: func(cmd *cobra.Command, args []string) {
		panic("not implemented")
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
