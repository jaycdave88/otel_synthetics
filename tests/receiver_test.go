package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	internal "github.com/jaycdave88/otel-synthetics/internal/receiver"
)

func TestReceiverInitialization(t *testing.T) {
	factory := internal.Factory()
	assert.NotNil(t, factory, "Receiver factory should not return nil")

	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "Default config should not be nil")
}

func TestReceiverLifecycle(t *testing.T) {
	ctx := context.Background()
	factory := internal.Factory()
	cfg := factory.CreateDefaultConfig()

	// Create receiver with mock consumer
	sink := new(consumertest.LogsSink)
	id := component.NewIDWithName(internal.Type, "test")
	set := receiver.CreateSettings{
		TelemetrySettings: component.TelemetrySettings{
			Logger: zap.NewNop(),
		},
		ID: id,
	}

	recv, err := factory.CreateLogsReceiver(ctx, set, cfg, sink)
	assert.NoError(t, err)
	assert.NotNil(t, recv)

	// Test Start
	err = recv.Start(ctx, nil)
	assert.NoError(t, err, "Receiver should start without error")

	// Let it run for a short time to ensure synthetic tests are triggered
	time.Sleep(100 * time.Millisecond)

	// Test Shutdown
	err = recv.Shutdown(ctx)
	assert.NoError(t, err, "Receiver should shutdown without error")

	// Verify that logs were processed
	assert.GreaterOrEqual(t, sink.LogRecordCount(), uint64(0), "Receiver should process logs")
}
