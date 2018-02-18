package main

import (
	"log"

	"github.com/spf13/cobra/doc"
	"github.com/trusch/stellarctl/cmd"
)

func main() {
	root := cmd.RootCmd
	err := doc.GenMarkdownTree(root, "./")
	if err != nil {
		log.Fatal(err)
	}
}
