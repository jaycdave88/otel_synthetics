// pkg/receiver/factory.go
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

// Factories returns all receiver factories
func Factories() map[component.Type]receiver.Factory {
	return map[component.Type]receiver.Factory{
		Type: NewFactory(),
	}
}

// NewFactory creates a factory for synthetic monitoring receiver
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		Type,
		createDefaultConfig,
		receiver.WithLogs(createLogsReceiver, component.StabilityLevelDevelopment),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

type Config struct {
	component.Config   `mapstructure:",squash"`
	CollectionInterval string `mapstructure:"collection_interval"`
}

func createLogsReceiver(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (receiver.Logs, error) {
	return &logsReceiver{
		logger:       params.Logger,
		config:       cfg.(*Config),
		nextConsumer: nextConsumer,
	}, nil
}
