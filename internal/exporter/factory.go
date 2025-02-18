package exporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

const typeStr = "synthetics_exporter"

// Type is the registered type for this exporter
var Type = component.MustNewType(typeStr)

// NewFactory creates a factory for synthetic monitoring exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		Type,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, component.StabilityLevelDevelopment),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

// Config defines configuration for synthetic monitoring exporter.
type Config struct {
	component.Config `mapstructure:",squash"`
}

func createTracesExporter(
	_ context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
	nextConsumer ptrace.Consumer,
) (exporter.Traces, error) {
	return newTraceExporter(set, cfg)
}

// Implement a basic trace exporter
type traceExporter struct {
	config   Config
	settings exporter.CreateSettings
}

func newTraceExporter(set exporter.CreateSettings, cfg component.Config) (exporter.Traces, error) {
	pCfg := cfg.(*Config)
	return &traceExporter{
		config:   *pCfg,
		settings: set,
	}, nil
}

func (e *traceExporter) Start(_ context.Context, _ component.Host) error {
	return nil
}

func (e *traceExporter) Shutdown(_ context.Context) error {
	return nil
}

func (e *traceExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	// Implement trace export logic here
	return nil
}

// Capabilities returns the capabilities of the exporter
func (e *traceExporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}
