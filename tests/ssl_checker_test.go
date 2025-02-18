package tests

import (
	"testing"

	"github.com/jaycdave88/otel-synthetics/internal/exporter"
	"go.uber.org/zap"
)

func TestCheckSSLValid(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	checker := exporter.NewSSLChecker(logger)

	result := checker.CheckSSL("google.com")
	if !result.Valid {
		t.Errorf("Expected valid SSL, got invalid")
	}
}

func TestCheckSSLExpired(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	checker := exporter.NewSSLChecker(logger)

	result := checker.CheckSSL("expired.badssl.com")
	if result.Valid {
		t.Errorf("Expected expired SSL, got valid")
	}
}
