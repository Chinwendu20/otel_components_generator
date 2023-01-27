package main

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

var (
	Config = config.NewConfig()
)

func flags() *flag.FlagSet {
	flagSet := new(flag.FlagSet)

	flagSet.StringVar(&Config.Component, componentTypeFlag, "", "The type of component to be generated")
	flagSet.StringVar(&Config.Module, goModuleNameFlag, "", "The name of the GO module")
	flagSet.StringVar(&Config.Output, outputDirectoryFlag, "", "The path to the directory for the generated source code")
	flagSet.StringVar(&Config.Signals, signalsFlag, "", "This could be of value, metrics, traces or logs")

	return flagSet
}

func checkEmptyConfigOptions() {
	if Config.Component == "" {
		obtainValueInteractively(componentTypeFlag, &Config.Component)
	}
	if Config.Module == "" {
		obtainValueInteractively(goModuleNameFlag, &Config.Module)
	}
	if Config.Output == "" {
		obtainValueInteractively(outputDirectoryFlag, &Config.Output)
	}
	if len(Config.Signals) == 0 && Config.Component != "extension" {
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
