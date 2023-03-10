package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Chinwendu20/otel_components_generator/config"
)

// This returns a cobra command and it is the main entrypoint for this application
func command(cfg config.Struct) *cobra.Command {
	flagSet := flags(&cfg)
	cmd := &cobra.Command{
		SilenceUsage:  true, // Don't print usage on Run error.
		SilenceErrors: true, // Don't print errors; main does it.
		Use:           "ocg",
		Long: fmt.Sprintf("OpenTelemetry Collector Generator (%s)", version) + `

ocg generates a custom OpenTelemetry Collector binary using the
options supplied by the commandline options. 
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := checkEmptyConfigOptions(cfg); err != nil {
				return err
			}
			if err := validateComponent(cfg); err != nil {

				return err
			}
			return generateComponent(cfg)
		},
	}

	cmd.Flags().AddGoFlagSet(flagSet)
	cmd.AddCommand(versionCommand())
	return cmd
}
