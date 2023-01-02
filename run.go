package main

import (
	"fmt"
	"kubectl-categories/pkg/resources"

	"github.com/spf13/cobra"
)

func run(command *cobra.Command, args []string) error {

	discoveryClient, err := getDiscoveryClient()
	if err != nil {
		return err
	}

	categories, err := resources.GetResourceCategories(discoveryClient)
	if err != nil {
		return err
	}

	fmt.Print(categories)
	return nil
}
