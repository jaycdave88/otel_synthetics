package tests

import (
	"context"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/jaycdave88/otel-synthetics/config"
	"github.com/jaycdave88/otel-synthetics/internal/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

func TestExporterInitialization(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	cfg := &config.Config{
		ExporterEndpoint: "http://localhost:4317",
	}

	exp, err := exporter.NewExporter(cfg, logger)
	if err != nil {
		t.Fatalf("Failed to create exporter: %v", err)
	}

	if exp == nil {
		t.Fatal("Exporter instance is nil")
	}
}

func TestExporterConsumeLogs(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	cfg := &config.Config{
		ExporterEndpoint: "http://localhost:4317",
	}

	exp, _ := exporter.NewExporter(cfg, logger)
	ctx := context.Background()

	// Create test logs
	logs := plog.NewLogs()
	resourceLogs := logs.ResourceLogs().AppendEmpty()
	scopeLogs := resourceLogs.ScopeLogs().AppendEmpty()
	logRecord := scopeLogs.LogRecords().AppendEmpty()

	timestamp := pcommon.NewTimestampFromTime(time.Now())
	logRecord.SetObservedTimestamp(timestamp)
	logRecord.SetTimestamp(timestamp)
	logRecord.SetSeverityText("INFO")
	logRecord.Body().SetStr("Test log message")

	err := exp.ConsumeLogs(ctx, logs)
	if err != nil {
		t.Errorf("ConsumeLogs failed: %v", err)
	}
}
