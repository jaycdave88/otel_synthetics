package exporter

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/jaycdave88/otel-synthetics/config"
)

// Exporter exports synthetic monitoring data
type Exporter struct {
	cfg    *config.Config
	logger *zap.Logger
}

// NewExporter creates a new exporter
func NewExporter(cfg *config.Config, logger *zap.Logger) (*Exporter, error) {
	return &Exporter{
		cfg:    cfg,
		logger: logger,
	}, nil
}

// ConsumeLogs consumes logs
func (e *Exporter) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	e.logger.Debug("Consuming logs", zap.Int("count", logs.LogRecordCount()))
	return nil
}

func (e *Exporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}
