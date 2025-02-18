// internal/receiver/factory.go
package receiver

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const typeStr = "synthetics"

// Type is the registered type for this receiver
var Type = component.MustNewType(typeStr)

// NewFactory creates a factory for synthetic monitoring receiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		Type,
		createDefaultConfig,
		receiver.WithLogs(createLogsReceiver, component.StabilityLevelDevelopment),
		receiver.WithMetrics(createMetricsReceiver, component.StabilityLevelDevelopment),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

// Config defines configuration for synthetic monitoring receiver.
type Config struct {
	component.Config `mapstructure:",squash"`
	// Add any necessary config options here, like:
	// Endpoint string `mapstructure:"endpoint"`
	// Interval time.Duration `mapstructure:"interval"`
}

func createLogsReceiver(
	_ context.Context,
	params receiver.CreateSettings,
	cfg component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	return newLogsReceiver(params, cfg.(*Config), consumer)
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.CreateSettings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	return newMetricsReceiver(params, cfg.(*Config), consumer)
}
