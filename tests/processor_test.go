package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	internal "github.com/jaycdave88/otel-synthetics/internal/processor"
)

func TestProcessorInitialization(t *testing.T) {
	proc := internal.Factory()
	assert.NotNil(t, proc, "Processor factory should not return nil")

	cfg := proc.CreateDefaultConfig()
	assert.NotNil(t, cfg, "Default config should not be nil")
}

func TestProcessorLifecycle(t *testing.T) {
	ctx := context.Background()
	proc := internal.Factory()
	cfg := proc.CreateDefaultConfig()

	// Create processor with mock consumer
	sink := new(consumertest.LogsSink)
	id := component.NewIDWithName(internal.Type, "test")
	set := processor.CreateSettings{
		TelemetrySettings: component.TelemetrySettings{
			Logger: zap.NewNop(),
		},
		ID: id,
	}

	p, err := proc.CreateLogsProcessor(ctx, set, cfg, sink)
	assert.NoError(t, err)
	assert.NotNil(t, p)

	// Test Start
	err = p.Start(ctx, nil)
	assert.NoError(t, err, "Processor should start without error")
	time.Sleep(100 * time.Millisecond)

	// Create test logs
	logs := plog.NewLogs()
	resourceLogs := logs.ResourceLogs().AppendEmpty()
	scopeLogs := resourceLogs.ScopeLogs().AppendEmpty()
	logRecord := scopeLogs.LogRecords().AppendEmpty()

	// Set timestamp
	now := time.Now().UnixNano()
	logRecord.SetTimestamp(plog.NewTimestampFromTime(time.Unix(0, now)))
	logRecord.SetSeverityText("INFO")
	logRecord.Body().SetStr("Test log message")

	// Test log processing
	err = p.ConsumeLogs(ctx, logs)
	assert.NoError(t, err, "Log processing should not error")

	// Test Shutdown
	err = p.Shutdown(ctx)
	assert.NoError(t, err, "Processor should shutdown without error")
}
