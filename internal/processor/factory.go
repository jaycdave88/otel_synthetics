package processor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
)

const typeStr = "synthetics_processor"

// Type is the registered type for this processor
var Type = component.MustNewType(typeStr)

// NewFactory creates a factory for synthetic monitoring processor.
func NewFactory() processor.Factory {
	return processor.NewFactory(
		Type,
		createDefaultConfig,
		processor.WithTraces(createTracesProcessor, component.StabilityLevelDevelopment),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

// Config defines configuration for synthetic monitoring processor.
type Config struct {
	component.Config `mapstructure:",squash"`
}

func createTracesProcessor(
	_ context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer ptrace.Consumer,
) (processor.Traces, error) {
	return newTraceProcessor(set, cfg, nextConsumer)
}

// Implement a basic trace processor
type traceProcessor struct {
	nextConsumer ptrace.Consumer
	config       Config
	settings     processor.CreateSettings
}

func newTraceProcessor(set processor.CreateSettings, cfg component.Config, nextConsumer ptrace.Consumer) (processor.Traces, error) {
	pCfg := cfg.(*Config)
	return &traceProcessor{
		nextConsumer: nextConsumer,
		config:       *pCfg,
		settings:     set,
	}, nil
}

func (p *traceProcessor) Start(_ context.Context, _ component.Host) error {
	return nil
}

func (p *traceProcessor) Shutdown(_ context.Context) error {
	return nil
}

func (p *traceProcessor) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	// Implement trace processing logic here

	// Forward to the next consumer
	return p.nextConsumer.ConsumeTraces(ctx, td)
}

// Capabilities returns the capabilities of the processor
func (p *traceProcessor) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}
