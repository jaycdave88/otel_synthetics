package tests

import (
	"context"
	"testing"
	"time"

	"github.com/jaycdave88/otel-synthetics/config"
	"github.com/jaycdave88/otel-synthetics/internal/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

func TestIntegration(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	ctx := context.Background()

	// Run an SSL check
	sslChecker := exporter.NewSSLChecker(logger)
	sslResult := sslChecker.CheckSSL("google.com")
	if !sslResult.Valid {
		t.Errorf("Integration Test: SSL test failed")
	}

	// Run an HTTP check
	httpChecker := exporter.NewHTTPChecker(logger)
	httpResult := httpChecker.CheckHTTP("https://example.com")
	if httpResult.StatusCode != 200 {
		t.Errorf("Integration Test: Expected HTTP 200, got %d", httpResult.StatusCode)
	}

	// Simulate exporter push logs
	cfg := &config.Config{
		ExporterEndpoint: "http://localhost:4317",
	}
	exp, _ := exporter.NewExporter(cfg, logger)

	// Create test logs
	logs := plog.NewLogs()
	resourceLogs := logs.ResourceLogs().AppendEmpty()
	scopeLogs := resourceLogs.ScopeLogs().AppendEmpty()
	logRecord := scopeLogs.LogRecords().AppendEmpty()

	// Set timestamp using pcommon.Timestamp
	timestamp := pcommon.Timestamp(time.Now().UnixNano())
	logRecord.SetObservedTimestamp(timestamp)
	logRecord.SetTimestamp(timestamp)
	logRecord.SetSeverityText("INFO")
	logRecord.Body().SetStr("Integration test log message")

	err := exp.ConsumeLogs(ctx, logs)
	if err != nil {
		t.Errorf("Exporter failed to push logs: %v", err)
	}
}
