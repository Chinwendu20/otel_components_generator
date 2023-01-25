package config

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

const DefaultOtelColVersion = "0.69.1"

type SignalSlice []string

var validSignals = SignalSlice{"metric", "trace", "log"}

func (sigs *SignalSlice) String() string {

	return fmt.Sprint(*sigs)
}

func (sigs *SignalSlice) Set(value string) error {
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

type ComponentString string

var validComponents = []ComponentString{"exporter", "receiver", "processor", "extension"}

func (compt *ComponentString) String() string {
	return fmt.Sprint(*compt)
}

func (compt *ComponentString) Set(s string) error {

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

type ConfigStruct struct {
	Logger         *zap.Logger
	SkipGetModules bool
	Component      ComponentString
	Module         string
	Output         string
	Signals        SignalSlice
	GoPath         string
}

func NewConfig() ConfigStruct {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("Experienced failure in obtaining logger instance: %v", err))
	}

	return ConfigStruct{
		Logger: log,
	}
}
