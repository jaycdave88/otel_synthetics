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

	internal "github.com/jaycdave88/otel-synthetics/pkg/syntheticsreceiver"
)

func TestReceiverInitialization(t *testing.T) {
	factory := internal.NewFactory()
	assert.NotNil(t, factory, "Receiver factory should not return nil")

	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "Default config should not be nil")
}

func TestReceiverLifecycle(t *testing.T) {
	ctx := context.Background()
	factory := internal.NewFactory()
	cfg := factory.CreateDefaultConfig()

	// Create receiver with mock consumer
	sink := new(consumertest.LogsSink)
	id := component.NewIDWithName(internal.Type, "test")
	set := receiver.Settings{
		TelemetrySettings: component.TelemetrySettings{
			Logger: zap.NewNop(),
		},
		ID: id,
	}

	recv, err := factory.CreateLogs(ctx, set, cfg, sink)
	assert.NoError(t, err)
	assert.NotNil(t, recv)

	// Test Start
	err = recv.Start(ctx, nil)
	assert.NoError(t, err, "Receiver should start without error")

	time.Sleep(500 * time.Millisecond)

	// Test Shutdown
	err = recv.Shutdown(ctx)
	assert.NoError(t, err, "Receiver should shutdown without error")

	logCount := sink.LogRecordCount()
	assert.True(t, logCount > 0, "Receiver should process logs, but got %d logs", logCount)
}
