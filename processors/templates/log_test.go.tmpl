package project

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/processor/processortest"
)

func TestLogsProcessorNoErrors(t *testing.T) {
	f := NewFactory()
	lle, err := f.CreateLogsProcessor(context.Background(), processortest.NewNopCreateSettings(), f.CreateDefaultConfig(), nil)
	require.NotNil(t, lle)
	assert.NoError(t, err)

	assert.NoError(t, lle.ConsumeLogs(context.Background(), plog.NewLogs()))

	assert.NoError(t, lle.Shutdown(context.Background()))
}
