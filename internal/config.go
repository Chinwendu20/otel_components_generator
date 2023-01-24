package internal

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/exec"
)

const defaultOtelColVersion = "0.69.1"

type signalSlice []string

type componentString string

type ConfigStruct struct {
	Logger         *zap.Logger
	SkipGetModules bool
	Component      componentString
	Module         string
	Output         string
	Signals        signalSlice
	GoPath         string
}

func newConfig() ConfigStruct {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("Experienced failure in obtaining logger instance: %v", err))
	}

	return ConfigStruct{
		Logger: log,
	}
}

func (cfg *ConfigStruct) ProcessOutputPath() error {
	if _, err := os.Stat(cfg.Output); os.IsNotExist(err) {
		cfg.Logger.Info("Output path not found, creating directory")
		if err = os.Mkdir(cfg.Output, 0750); err != nil {
			return fmt.Errorf("failed to create output path: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf(" %w", err)
	}
	cfg.Logger.Info("Output path created")
	return nil

}

func (cfg *ConfigStruct) SetGoPath() error {
	if !cfg.SkipGetModules {
		path, err := exec.LookPath("go")
		if err != nil {
			return errors.New("GO binary not found")
		}
		cfg.GoPath = path
		cfg.Logger.Info("Using go", zap.String("go-executable", cfg.GoPath))
	}
	return nil
}
