package project

import (
	"context"
	"go.opentelemetry.io/collector/component/componenttest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/receiver/receivertest"
)

func TestLogsReceiverNoErrors(t *testing.T) {
	f := NewFactory()
	lle, err := f.CreateLogsReceiver(context.Background(), receivertest.NewNopCreateSettings(), f.CreateDefaultConfig(), nil)
	require.NotNil(t, lle)
	assert.NoError(t, err)

	assert.NoError(t, lle.Start(context.Background(), componenttest.NewNopHost()))

	assert.NoError(t, lle.Shutdown(context.Background()))
}
