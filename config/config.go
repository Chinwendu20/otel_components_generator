package config

import (
	"fmt"
	"go.uber.org/zap"
	"strings"
)

const (
	DefaultOtelColVersion = "0.69.1"
)

var (
	validSignals    = []string{"metric", "trace", "log"}
	validComponents = []string{"exporter", "receiver", "processor", "extension"}
)

type ConfigStruct struct {
	Logger         *zap.Logger
	SkipGetModules bool
	Component      string
	Module         string
	Output         string
	Signals        string
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

func (cfg *ConfigStruct) ValidateSignal() error {
	for _, sig := range cfg.SetSignals() {
		valid := false
		for _, signal := range validSignals {
			if sig == signal {
				valid = true
				break

			}
		}
		if !valid {

			return fmt.Errorf("invalid input for signals flag, accepted values are: %v", validSignals)

		}

	}
	return nil
}

func (cfg *ConfigStruct) ValidateComponent() error {
	fmt.Println(cfg.Component)

	for _, component := range validComponents {

		if component == cfg.Component {
			return nil
		}
	}
	return fmt.Errorf("invalid input for component flag, accepted values are: %v", validComponents)

}

func (cfg *ConfigStruct) SetSignals() []string {

	return strings.Split(cfg.Signals, ",")

}
