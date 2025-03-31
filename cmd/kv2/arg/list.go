package arg

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

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

		if jsonOutput {
			json.NewEncoder(os.Stdout).Encode(res)
			return
		}

		data := [][]string{
			{"KEY", "VERSION"},
		}

		for _, s := range res {
			data = append(data, []string{s.Key, fmt.Sprintf("%d", len(s.Versions))})
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)

		for _, row := range data {
			fmt.Fprintln(w, row[0]+"\t"+row[1])
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
