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

	flagSet.Var(&Config.Component, componentTypeFlag, "The type of component to be generated")
	flagSet.StringVar(&Config.Module, goModuleNameFlag, "", "The name of the GO module")
	flagSet.StringVar(&Config.Output, outputDirectoryFlag, "", "The path to the directory for the generated source code")
	flagSet.Var(&Config.Signals, signalsFlag, "This could be of value, metrics, traces or logs")

	return flagSet
}

func checkEmptyConfigOptions() {
	if Config.Component == "" {
		obtainValueInteractively(componentTypeFlag)
	}
	if Config.Module == "" {
		obtainValueInteractively(goModuleNameFlag)
	}
	if Config.Output == "" {
		obtainValueInteractively(outputDirectoryFlag)
	}
	if len(Config.Signals) == 0 {
		obtainValueInteractively(signalsFlag)
	}

}

func obtainValueInteractively(value string) {
	fmt.Printf("Input value for %s, no default setting:", value)
	_, err := fmt.Scanln(&Config.Component)
	if err != nil {
		log.Fatal(err)
	}

}
