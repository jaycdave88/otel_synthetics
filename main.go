package main

import (
	"context"
	"log"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/service"

	"github.com/jaycdave88/otel-synthetics/internal/exporter"
	"github.com/jaycdave88/otel-synthetics/internal/processor"
	"github.com/jaycdave88/otel-synthetics/internal/receiver"
)

func main() {
	ctx := context.Background()

	factories := []component.Factory{
		receiver.Factory(),
		processor.Factory(),
		exporter.Factory(),
	}

	info := component.BuildInfo{
		Command:     "otel-synthetics-collector",
		Description: "OpenTelemetry Collector for synthetic monitoring",
		Version:     "1.0.0",
	}

	configProvider, err := service.NewConfigProvider(
		service.ConfigProviderSettings{
			ResolverSettings: confmap.ResolverSettings{
				URIs:      []string{"file:deploy/otel-collector-config.yaml"},
				Providers: map[string]confmap.Provider{},
			},
		})
	if err != nil {
		log.Fatalf("Failed to create config provider: %v", err)
	}

	settings := service.Settings{
		BuildInfo:      info,
		Factories:      factories,
		ConfigProvider: configProvider,
	}

	svc, err := service.New(ctx, settings)
	if err != nil {
		log.Fatalf("Failed to create service: %v", err)
	}

	err = svc.Start(context.Background())
	if err != nil {
		log.Fatalf("Failed to start service: %v", err)
	}

	defer func() {
		err = svc.Shutdown(context.Background())
		if err != nil {
			log.Fatalf("Failed to shutdown service: %v", err)
		}
	}()

	// Block until the service is shutdown
	if err = svc.Run(ctx); err != nil {
		log.Fatalf("Service error: %v", err)
	}
}
