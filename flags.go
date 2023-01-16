package Otel_Component_Generator

import (
	"flag"
	"fmt"
	"log"
)

type config struct {
	component string
	module    string
	output    string
	signal    string
}

const (
	componentTypeFlag   = "component"
	goModuleNameFlag    = "module"
	outputDirectoryFlag = "output"
	signalsFlag         = "signal"
)

var cfgs = new(config)

func flags() *flag.FlagSet {
	flagSet := new(flag.FlagSet)

	flagSet.StringVar(&cfgs.component, componentTypeFlag, "", "The type of component to be generated")
	flagSet.StringVar(&cfgs.module, goModuleNameFlag, "", "The name of the GO module")
	flagSet.StringVar(&cfgs.output, outputDirectoryFlag, "", "The path to the directory for the generated source code")
	flagSet.StringVar(&cfgs.signal, signalsFlag, "", "This could be of value, metrics, traces or logs")

	return flagSet
}

func checkConfigOptions() {
	if cfgs.component == "" {
		obtainValueInteractively(componentTypeFlag)
	}
	if cfgs.module == "" {
		obtainValueInteractively(goModuleNameFlag)
	}
	if cfgs.module == "" {
		obtainValueInteractively(outputDirectoryFlag)
	}
	if cfgs.module == "" {
		obtainValueInteractively(signalsFlag)
	}

}

func obtainValueInteractively(value string) {
	fmt.Println("Input value for {}, no default setting:", value)
	_, err := fmt.Scanln(&cfgs.component)
	if err != nil {
		log.Fatal(err)
	}

}
