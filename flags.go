package main

import "C"
import (
	"flag"
	"fmt"
	"github.com/Chinwendu20/otel_components_generator/config"
	"log"
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

func checkEmptyConfigOptions(cfg config.ConfigStruct) {
	if cfg.Component == "" {
		obtainValueInteractively(componentTypeFlag, &Config.Component)
	}
	if cfg.Module == "" {
		obtainValueInteractively(goModuleNameFlag, &Config.Module)
	}
	if cfg.Output == "" {
		obtainValueInteractively(outputDirectoryFlag, &Config.Output)
	}
	if len(cfg.Signals) == 0 && cfg.Component != "extension" {
		obtainValueInteractively(signalsFlag, &Config.Signals)
	}

}

func obtainValueInteractively(value string, valstore *string) {
	fmt.Printf("Input value for %s, no default setting:", value)
	_, err := fmt.Scanln(valstore)
	if err != nil {
		log.Fatal(err)
	}

}
