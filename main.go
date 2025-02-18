package main

import (
	"context"
	"log"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/otelcol"

	"github.com/jaycdave88/otel-synthetics/internal/exporter"
	"github.com/jaycdave88/otel-synthetics/internal/processor"
	receiver "github.com/jaycdave88/otel-synthetics/pkg/syntheticsreceiver"
)

func main() {
	ctx := context.Background()

	info := component.BuildInfo{
		Command:     "otel-synthetics",
		Description: "Custom OpenTelemetry Collector for Synthetic Monitoring",
		Version:     "1.0.0",
	}

	configProviderSettings := otelcol.ConfigProviderSettings{
		ResolverSettings: confmap.ResolverSettings{
			URIs: []string{"file:deploy/otel-collector-config.yaml"},
		},
	}

	settings := otelcol.CollectorSettings{
		BuildInfo:              info,
		Factories:              components, // Pass the function itself, not its result
		ConfigProviderSettings: configProviderSettings,
	}

	svc, err := otelcol.NewCollector(settings)
	if err != nil {
		log.Fatalf("Failed to create service: %v", err)
	}

	err = svc.Run(ctx)
	if err != nil {
		log.Fatalf("Service error: %v", err)
	}
}

func components() (otelcol.Factories, error) {
	factories := otelcol.Factories{}

	// Add custom receivers
	receiverFactories := receiver.Factories()
	factories.Receivers = receiverFactories

	// Add custom exporters
	exporterFactories := exporter.Factories()
	factories.Exporters = exporterFactories

	// Add custom processors
	processorFactories := processor.Factories()
	factories.Processors = processorFactories

	return factories, nil
}
