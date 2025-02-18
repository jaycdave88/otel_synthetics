package config

// Config defines settings for synthetic monitoring
type Config struct {
	ExporterEndpoint string `yaml:"exporter_endpoint"`
}
