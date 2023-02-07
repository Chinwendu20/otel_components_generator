package receivers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Chinwendu20/otel_components_generator/config"
)

const lengthOfDefaultTemplate = 5

func TestGenerateReceiver(t *testing.T) {
	cfg := config.NewConfig()
	tests := []struct {
		name                string
		signals             string
		diffExpectTemplates int
	}{
		{
			name:                "trace and metric signal",
			signals:             "trace,metric",
			diffExpectTemplates: 4,
		},
		{
			name:                "trace signal",
			signals:             "trace",
			diffExpectTemplates: 2,
		},
		{
			name:                "metric signal",
			signals:             "metric",
			diffExpectTemplates: 2,
		},
		{
			name:                "log signal",
			signals:             "log",
			diffExpectTemplates: 2,
		},
		{
			name:                "log and metric signal",
			signals:             "log,metric",
			diffExpectTemplates: 4,
		},
		{
			name:                "log and trace signal",
			signals:             "log,trace",
			diffExpectTemplates: 4,
		},
		{
			name:                "log,trace and metric signal",
			signals:             "log,trace,metric",
			diffExpectTemplates: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Signals = tt.signals
			templates := GenerateReceiver(cfg)
			assert.Equal(t, len(templates), tt.diffExpectTemplates+lengthOfDefaultTemplate)
		})
	}
}
