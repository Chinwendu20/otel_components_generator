package internal

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"
)

const (
	componentTypeFlag   = "component"
	goModuleNameFlag    = "module"
	outputDirectoryFlag = "output"
	signalsFlag         = "signal"
)

var (
	Config          = newConfig()
	validSignals    = signalSlice{"metric", "trace", "log"}
	validComponents = []componentString{"exporter", "receiver", "processor", "extension"}
)

func (compt *componentString) String() string {
	return fmt.Sprint(*compt)
}

func (compt *componentString) Set(s string) error {

	if len(*compt) > 0 {
		return errors.New("component flag already set")
	}

	for _, component := range validComponents {

		if component == *compt {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("invalid component. Accepted values are: %v", validComponents))

}

func (sigs *signalSlice) String() string {

	return fmt.Sprint(*sigs)
}

func (sigs *signalSlice) Set(value string) error {
	if len(*sigs) > 0 {
		return errors.New("interval flag already set")
	}
	for _, signal := range strings.Split(value, ",") {

		*sigs = append(*sigs, signal)
	}
	for _, sig := range *sigs {
		valid := false
		for _, signal := range validSignals {
			if sig == signal {
				valid = true
				break

			}
		}
		if !valid {

			return errors.New(fmt.Sprintf("invalid signal.Accepted values are: %v", validSignals))

		}

	}
	return nil
}

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
	fmt.Println("Input value for {}, no default setting:", value)
	_, err := fmt.Scanln(&Config.Component)
	if err != nil {
		log.Fatal(err)
	}

}
