package cmd

import (
	"log"
	"os"
	"path"

	"github.com/gerred/kluster/vendor/github.com/spf13/cobra"
)

var driver string

var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new Kubernetes cluster on the specified provider",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Specify one name argument.")
		}

		home := os.Getenv("HOME")
		name := args[0]

		log.Printf("Creating cluster %s using driver %s.", name, driver)
		os.MkdirAll(path.Join(home, ".kluster", name), 0755)
	},
}

func init() {
	createCmd.Flags().StringVarP(&driver, "driver", "d", "virtualbox", "backend driver to use")
}
