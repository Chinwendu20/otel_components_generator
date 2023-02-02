package exporters

import (
	"github.com/Chinwendu20/otel_components_generator/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

const lengthOfDefaultTemplate = 3

func TestGenerateExporter(t *testing.T) {
	cfg := config.NewConfig()
	tests := []struct {
		name                string
		signals             string
		diffExpectTemplates int
	}{
		{
			name:                "trace and metric signal",
			signals:             "trace,metric",
			diffExpectTemplates: 2,
		},
		{
			name:                "trace signal",
			signals:             "trace",
			diffExpectTemplates: 1,
		},
		{
			name:                "metric signal",
			signals:             "metric",
			diffExpectTemplates: 1,
		},
		{
			name:                "log signal",
			signals:             "log",
			diffExpectTemplates: 1,
		},
		{
			name:                "log and metric signal",
			signals:             "log,metric",
			diffExpectTemplates: 2,
		},
		{
			name:                "log and trace signal",
			signals:             "log,trace",
			diffExpectTemplates: 2,
		},
		{
			name:                "log,trace and metric signal",
			signals:             "log,trace,metric",
			diffExpectTemplates: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Signals = tt.signals
			templates := GenerateExporter(cfg)
			assert.Equal(t, len(templates), tt.diffExpectTemplates+lengthOfDefaultTemplate)
		})
	}
}
