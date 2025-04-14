package main

import (
	"fmt"
	"os"

	"github.com/BagToad/gh-cobra-template/cmd"
)

func main() {
	// This is the main entry point for the gh-cobra-template extension.
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
