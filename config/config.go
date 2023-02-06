package config

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"regexp"
	"strings"
)

const (
	DefaultOtelColVersion = "0.69.1"
	ConfigFileName        = "config.go"
	FactoryFileName       = "factory.go"
	GoModFileName         = "go.mod"
	LogFileName           = "log.go"
	TraceFileName         = "trace.go"
	MetricFileName        = "metric.go"
	ConfigTestFileName    = "config_test.go"
	FactoryTestFileName   = "factory_test.go"
	LogTestFileName       = "log_test.go"
	MetricTestFileName    = "metric_test.go"
	TraceTestFileName     = "trace_test.go"
)

var (
	validSignals      = []string{"metric", "trace", "log"}
	validComponents   = []string{"exporter", "receiver", "processor", "extension"}
	validateSignalErr = fmt.Errorf("Invalid input for signals flag, accepted values are: %v", validSignals)
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
	if cfg.Component == "extension" {

		return nil
	}
	for _, sig := range cfg.SetSignals() {
		valid := false
		for _, signal := range validSignals {
			if sig == signal {
				valid = true
				break

			}
		}
		if !valid {

			return validateSignalErr

		}

	}
	return nil
}

func (cfg *ConfigStruct) ValidateComponent() error {

	for _, component := range validComponents {

		if component == cfg.Component {
			return nil
		}
	}
	return fmt.Errorf("Invalid input for component flag, accepted values are: %v", validComponents)

}

func (cfg *ConfigStruct) ValidateModule() error {

	match, err := regexp.MatchString(`^github.com/\w+/[A-Za-z]\w+[A-Za-z]$`, cfg.Module)
	if err == nil {
		if match {

			return nil

		} else {

			return fmt.Errorf("Invalid input for module flag, string must follow this pattern, github.com/<github username>/<package name>")
		}
	}
	log.Fatal(err)
	return err
}

func (cfg *ConfigStruct) SetSignals() []string {

	return strings.Split(cfg.Signals, ",")

}
