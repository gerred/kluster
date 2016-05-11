package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the base command for the kluster CLI
var RootCmd = &cobra.Command{
	Use:   "kluster",
	Short: "kluster is a tool for managing Kubernetes clusters",
	Long:  `Manage Kubernetes clusters the right way. Create, upgrade, and scale clusters.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}
