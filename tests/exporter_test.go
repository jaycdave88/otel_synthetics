package tests

import (
	"context"
	"testing"

	"go.uber.org/zap"

	"github.com/jaycdave88/otel-synthetics/config"
	"github.com/jaycdave88/otel-synthetics/internal/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
)

func TestExporterInitialization(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	cfg := &config.Config{
		ExporterEndpoint: "http://localhost:4317",
	}

	exporter, err := exporter.NewExporter(cfg, logger)
	if err != nil {
		t.Fatalf("Failed to create exporter: %v", err)
	}

	if exporter == nil {
		t.Fatal("Exporter is nil")
	}
}

func TestExporterConsumeLogs(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	cfg := &config.Config{
		ExporterEndpoint: "http://localhost:4317",
	}

	exp, _ := exporter.NewExporter(cfg, logger)
	ctx := context.Background()

	// Create empty logs
	logs := plog.NewLogs()

	err := exp.ConsumeLogs(ctx, logs)
	if err != nil {
		t.Errorf("ConsumeLogs failed: %v", err)
	}
}
