package main

import "C"
import (
	"flag"

	"github.com/Chinwendu20/otel_components_generator/config"
)

const (
	componentTypeFlag   = "component"
	goModuleNameFlag    = "module"
	outputDirectoryFlag = "output"
	signalsFlag         = "signal"
)

func flags(cfg *config.Struct) *flag.FlagSet {
	flagSet := new(flag.FlagSet)

	flagSet.StringVar(&cfg.Component, componentTypeFlag, "", "The type of component to be generated")
	flagSet.StringVar(&cfg.Module, goModuleNameFlag, "", "The name of the GO module")
	flagSet.StringVar(&cfg.Output, outputDirectoryFlag, "", "The path to the directory for the generated source code")
	flagSet.StringVar(&cfg.Signals, signalsFlag, "", "This could be of value, metrics, traces or logs")

	return flagSet
}
