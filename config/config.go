package config

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"go.uber.org/zap"
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
	errValidateSignal = fmt.Errorf("Invalid input for signals flag, accepted values are: %v", validSignals)
)

type Struct struct {
	Logger         *zap.Logger
	SkipGetModules bool
	Component      string
	Module         string
	Output         string
	Signals        string
	GoPath         string
}

func NewConfig() Struct {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("Experienced failure in obtaining logger instance: %v", err))
	}

	return Struct{
		Logger: log,
	}
}

func (cfg *Struct) ValidateSignal() error {
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

			return errValidateSignal

		}

	}
	return nil
}

func (cfg *Struct) ValidateComponent() error {

	for _, component := range validComponents {

		if component == cfg.Component {
			return nil
		}
	}
	return fmt.Errorf("Invalid input for component flag, accepted values are: %v", validComponents)

}

func (cfg *Struct) ValidateModule() error {

	match, err := regexp.MatchString(`^github\.com/\w+/[A-Za-z]\w+[A-Za-z]$`, cfg.Module)
	if err != nil {

		log.Fatal(err)

	}
	if match {

		return nil

	}

	return fmt.Errorf("Invalid input for module flag, string must follow this pattern, github.com/<github username>/<package name>")

}

func (cfg *Struct) SetSignals() []string {

	return strings.Split(cfg.Signals, ",")

}
