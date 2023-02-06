package main

import "C"
import (
	"errors"
	"flag"
	"fmt"
	"github.com/Chinwendu20/otel_components_generator/config"
	"strings"
)

const (
	componentTypeFlag   = "component"
	goModuleNameFlag    = "module"
	outputDirectoryFlag = "output"
	signalsFlag         = "signal"
)

func flags(cfg *config.ConfigStruct) *flag.FlagSet {
	flagSet := new(flag.FlagSet)

	flagSet.StringVar(&cfg.Component, componentTypeFlag, "", "The type of component to be generated")
	flagSet.StringVar(&cfg.Module, goModuleNameFlag, "", "The name of the GO module")
	flagSet.StringVar(&cfg.Output, outputDirectoryFlag, "", "The path to the directory for the generated source code")
	flagSet.StringVar(&cfg.Signals, signalsFlag, "", "This could be of value, metrics, traces or logs")

	return flagSet
}

func checkEmptyConfigOptions(cfg config.ConfigStruct) error {
	var emptyValues []string
	if cfg.Component == "" {
		emptyValues = append(emptyValues, fmt.Sprintf("- Value for %s required, please use the flag, --component\n", componentTypeFlag))
	}
	if cfg.Module == "" {
		emptyValues = append(emptyValues, fmt.Sprintf("- Value for %s required, please use the flag, --module\n", goModuleNameFlag))
	}
	if cfg.Output == "" {
		emptyValues = append(emptyValues, fmt.Sprintf("- Value for %s required, please use the flag, --output\n", outputDirectoryFlag))
	}
	if len(cfg.Signals) == 0 && cfg.Component != "extension" {
		emptyValues = append(emptyValues, fmt.Sprintf("- Value for %s required, please use the flag, --signal\n", signalsFlag))
	}
	if len(emptyValues) == 0 {
		return nil
	}

	emptyValues = append([]string{"\n"}, emptyValues...)
	return errors.New(strings.Join(emptyValues, " "))

}
