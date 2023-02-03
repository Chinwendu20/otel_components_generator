package project

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/pdata/plog"
)

func TestLogsExporterNoErrors(t *testing.T) {
	f := NewFactory()
	lle, err := f.CreateLogsExporter(context.Background(), exportertest.NewNopCreateSettings(), f.CreateDefaultConfig())
	require.NotNil(t, lle)
	assert.NoError(t, err)

	assert.NoError(t, lle.ConsumeLogs(context.Background(), plog.NewLogs()))

	assert.NoError(t, lle.Shutdown(context.Background()))
}
