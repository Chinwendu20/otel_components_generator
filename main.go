package main

import (
	"github.com/spf13/cobra"
)

func main() {
	cmd, err := command(Config)
	cobra.CheckErr(err)
	cobra.CheckErr(cmd.Execute())
}
