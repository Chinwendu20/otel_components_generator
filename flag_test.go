package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Chinwendu20/otel_components_generator/config"
)

func TestFlags(t *testing.T) {
	cfg := config.NewConfig()
	flgs := flags(&cfg)
	err := flgs.Parse([]string{"--component=exporter", "--module=github.com/user13/myexporter", "--signal=trace", "--output=./pop"})
	require.NoError(t, err)
	assert.Equal(t, cfg, config.Struct{
		Logger:         cfg.Logger,
		SkipGetModules: false,
		Component:      "exporter",
		Module:         "github.com/user13/myexporter",
		Output:         "./pop",
		Signals:        "trace",
	})
}
