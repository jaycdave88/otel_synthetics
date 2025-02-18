package exporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const typeStr = "synthetics_exporter"

// Type is the registered type for this exporter
var Type = component.MustNewType(typeStr)

func Factories() map[component.Type]exporter.Factory {
	return map[component.Type]exporter.Factory{
		Type: NewFactory(),
	}
}

// NewFactory creates a factory for synthetic monitoring exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		Type,
		createDefaultConfig,
		exporter.WithLogs(exporter.CreateLogsFunc(createLogsExporter), component.StabilityLevelDevelopment),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

type Config struct {
	component.Config `mapstructure:",squash"`
}

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {
	return newLogsExporter(set, cfg.(*Config))
}

func newLogsExporter(
	_ exporter.Settings,
	_ *Config,
) (exporter.Logs, error) {
	return nil, nil
}
