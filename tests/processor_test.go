package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	internal "github.com/jaycdave88/otel-synthetics/internal/processor"
)

func TestProcessorInitialization(t *testing.T) {
	proc := internal.NewFactory()
	assert.NotNil(t, proc, "Processor factory should not return nil")

	cfg := proc.CreateDefaultConfig()
	assert.NotNil(t, cfg, "Default config should not be nil")
}

func TestProcessorLifecycle(t *testing.T) {
	ctx := context.Background()
	proc := internal.NewFactory()
	cfg := proc.CreateDefaultConfig()

	// Create processor with mock consumer
	sink := new(consumertest.LogsSink)
	id := component.NewIDWithName(internal.Type, "test")
	set := processor.Settings{
		TelemetrySettings: component.TelemetrySettings{
			Logger: zap.NewNop(),
		},
		ID: id,
	}

	p, err := proc.CreateLogs(ctx, set, cfg, sink)
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

	timestamp := pcommon.NewTimestampFromTime(time.Now())
	logRecord.SetObservedTimestamp(timestamp)
	logRecord.SetTimestamp(timestamp)
	logRecord.SetSeverityText("INFO")
	logRecord.Body().SetStr("Test processor log")

	// Test log processing
	err = p.ConsumeLogs(ctx, logs)
	assert.NoError(t, err, "Log processing should not fail")

	// Test Shutdown
	err = p.Shutdown(ctx)
	assert.NoError(t, err, "Processor should shutdown without error")
}
