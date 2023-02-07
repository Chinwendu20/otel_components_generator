package main

import (
	"github.com/spf13/cobra"

	"github.com/Chinwendu20/otel_components_generator/config"
)

var (
	Config = config.NewConfig()
)

func main() {
	cmd := command(Config)
	cobra.CheckErr(cmd.Execute())
}
