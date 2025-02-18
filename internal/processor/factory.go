package processor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"
)

const typeStr = "synthetics_processor"

// Type is the registered type for this processor
var Type = component.MustNewType(typeStr)

// Factories initializes the processor factories
func Factories() map[component.Type]processor.Factory {
	return map[component.Type]processor.Factory{
		Type: NewFactory(),
	}
}

// NewFactory creates a factory for synthetic monitoring processor.
func NewFactory() processor.Factory {
	return processor.NewFactory(
		Type,
		createDefaultConfig,
		processor.WithLogs(createLogsProcessor, component.StabilityLevelDevelopment),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

type Config struct {
	component.Config `mapstructure:",squash"`
}

func createLogsProcessor(
	_ context.Context,
	params processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	return newLogsProcessor(params, cfg.(*Config), nextConsumer)
}

// logsProcessor implements the processor.Logs interface
type logsProcessor struct {
	logger       *zap.Logger
	config       *Config
	nextConsumer consumer.Logs
}

func newLogsProcessor(
	set processor.Settings,
	cfg *Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	return &logsProcessor{
		logger:       set.Logger,
		config:       cfg,
		nextConsumer: nextConsumer,
	}, nil
}

// Start implements the component.Component interface
func (p *logsProcessor) Start(_ context.Context, _ component.Host) error {
	return nil
}

// Shutdown implements the component.Component interface
func (p *logsProcessor) Shutdown(_ context.Context) error {
	return nil
}

// ConsumeLogs implements the consumer.Logs interface
func (p *logsProcessor) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	// Simply pass logs to the next consumer
	// In a real implementation, you'd do processing here
	return p.nextConsumer.ConsumeLogs(ctx, ld)
}

// Capabilities implements the processor.Logs interface
func (p *logsProcessor) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}
