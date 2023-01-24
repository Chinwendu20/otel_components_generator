package internal

import (
	"fmt"
	"github.com/Chinwendu20/otel_components_generator"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

// Command is the main entrypoint for this application
func Command() (*cobra.Command, error) {
	flagSet := flags()
	cmd := &cobra.Command{
		SilenceUsage:  true, // Don't print usage on Run error.
		SilenceErrors: true, // Don't print errors; main does it.
		Use:           "ocg",
		Long: fmt.Sprintf("OpenTelemetry Collector Builder (%s)", version) + `

ocg generates a custom OpenTelemetry Collector binary using the
options supplied by the commandline options. If no options are supplied
ocg requests for these interactively.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			checkEmptyConfigOptions()
			return
			otel_components_generator.GenerateComponent(Config)
		},
	}

	cmd.Flags().AddGoFlagSet(flagSet)
	return cmd, nil
}
