package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	date    = "unknown"
)

func versionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Version of ocg",
		Long:  "Prints the version of the ocg binary",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(fmt.Sprintf("%s version %s", cmd.Parent().Name(), version))
		},
	}
}
