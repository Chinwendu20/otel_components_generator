package otel_components_generator

import (
	"github.com/Chinwendu20/otel_components_generator/internal"
	"github.com/spf13/cobra"
)

func main() {
	cmd, err := internal.Command()
	cobra.CheckErr(err)
	cobra.CheckErr(cmd.Execute())
}
