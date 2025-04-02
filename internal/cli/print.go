// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"git.huggins.io/kv2/api"
)

func PrintTable(headers []string, rows [][]string) {
	// TODO
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)

	for _, col := range headers {
		fmt.Fprintf(w, "%s\t", col)
	}

	fmt.Fprintf(w, "\n")

	for _, row := range rows {
		for _, col := range row {
			fmt.Fprintf(w, "%s\t", col)
		}

		fmt.Fprintf(w, "\n")
	}

	w.Flush()
}

func PrintErrorOutput(asJson bool, err error) {
	if asJson {
		json.NewEncoder(os.Stdout).Encode(api.ErrorResponse{Message: err.Error()})
	} else {
		fmt.Println(err)
	}
}
