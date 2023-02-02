package main

import (
	"github.com/Chinwendu20/otel_components_generator/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlags(t *testing.T) {

	flgs := flags()
	err := flgs.Parse([]string{"--component=exporter", "--module=pop", "--output=./pop", "--signal=trace"})
	require.NoError(t, err)
	assert.Equal(t, Config, config.ConfigStruct{
		Logger:         Config.Logger,
		SkipGetModules: false,
		Component:      "exporter",
		Module:         "pop",
		Output:         "./pop",
		Signals:        "trace",
	})
}
