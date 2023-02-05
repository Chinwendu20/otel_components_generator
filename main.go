package main

import (
	"github.com/Chinwendu20/otel_components_generator/config"
	"github.com/spf13/cobra"
)

var (
	Config = config.NewConfig()
)

func main() {
	cmd, err := command(Config)
	cobra.CheckErr(err)
	cobra.CheckErr(cmd.Execute())
}
