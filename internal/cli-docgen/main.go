package main

import (
	"log"
	"os"

	kv2 "git.huggins.io/kv2/cmd/kv2/arg"
	"github.com/spf13/cobra/doc"
)

func main() {
	os.RemoveAll("docs/cli-manual/")
	if err := os.MkdirAll("docs/cli-manual/", 0755); err != nil {
		log.Fatal(err)
	}

	err := doc.GenMarkdownTree(kv2.RootCmd, "docs/cli-manual/")

	if err != nil {
		log.Fatal(err)
	}
}
