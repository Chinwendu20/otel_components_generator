package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig()
	assert.Equal(t, cfg, ConfigStruct{Logger: cfg.Logger})

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
