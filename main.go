package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "kubectl categories",
		Short:   "Show resource categories to use with kubectl get <category>",
		Example: "kubectl categories",
		Args:    cobra.ExactArgs(0),
		RunE:    run,
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
