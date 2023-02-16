package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig()
	assert.Equal(t, cfg, Struct{Logger: cfg.Logger})

}

func TestConfigStruct_ValidateSignal(t *testing.T) {
	cfg := NewConfig()
	tests := []struct {
		name      string
		signals   string
		ExpectErr error
	}{
		{
			name:      "two valid signals",
			signals:   "trace,metric",
			ExpectErr: nil,
		},
		{
			name:      "one valid signal, one invaid signal",
			signals:   "trace,errsig",
			ExpectErr: validateSignalErr,
		},
		{
			name:      "three valid signals",
			signals:   "trace,metric,log",
			ExpectErr: nil,
		},
		{
			name:      "one valid signal, two invalid signal",
			signals:   "trace,errsig,errsig2",
			ExpectErr: validateSignalErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Signals = tt.signals

			assert.Equal(t, cfg.ValidateSignal(), tt.ExpectErr)
		})
	}
}

func TestConfigStruct_ValidateComponent(t *testing.T) {
	cfg := NewConfig()
	tests := []struct {
		name      string
		component string
		err       error
	}{
		{
			name:      "Exporter component",
			component: "exporter",
			err:       nil,
		},
		{
			name:      "Processor component",
			component: "processor",
			err:       nil,
		},
		{
			name:      "Receiver component",
			component: "receiver",
			err:       nil,
		},
		{
			name:      "Extension component",
			component: "extension",
			err:       nil,
		},
		{
			name:      "Fail component",
			component: "fail",
			err:       fmt.Errorf("Invalid input for component flag, accepted values are: %v", validComponents),
		},
		{
			name:      "Export component",
			component: "export",
			err:       fmt.Errorf("Invalid input for component flag, accepted values are: %v", validComponents),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Component = tt.component
			err := cfg.ValidateComponent()

			assert.Equal(t, err, tt.err)
		})
	}
}

func TestConfigStruct_ValidateModule(t *testing.T) {
	cfg := NewConfig()
	tests := []struct {
		name   string
		module string

		err error
	}{
		{

			module: "github.com/user13/myextensions",
			err:    nil,
		},
		{
			module: "github.com/user13/8myextensions",
			err:    fmt.Errorf("Invalid input for module flag, string must follow this pattern, github.com/<github username>/<package name>"),
		},
		{
			module: "github.com/user13/myextensions/",
			err:    fmt.Errorf("Invalid input for module flag, string must follow this pattern, github.com/<github username>/<package name>"),
		},
		{
			module: "github.com/user13/e-xt",
			err:    fmt.Errorf("Invalid input for module flag, string must follow this pattern, github.com/<github username>/<package name>"),
		},
		{
			module: "github.com/user13/my_extensions",
			err:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Module = tt.module
			err := cfg.ValidateModule()

			assert.Equal(t, err, tt.err)
		})
	}
}
