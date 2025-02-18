package tests

import (
	"testing"

	"github.com/jaycdave88/otel-synthetics/internal/exporter"
	"go.uber.org/zap"
)

func TestCheckHTTPSuccess(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	checker := exporter.NewHTTPChecker(logger)

	result := checker.CheckHTTP("https://example.com")
	if result.StatusCode != 200 {
		t.Errorf("Expected HTTP 200, got %d", result.StatusCode)
	}
}

func TestCheckHTTPTimeout(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	checker := exporter.NewHTTPChecker(logger)

	// Test with a non-routable IP address that should timeout quickly
	result := checker.CheckHTTP("http://192.0.2.1:12345") // Using TEST-NET-1 IP address

	if result.ErrorMessage == "" {
		t.Error("Expected timeout error, got no error")
	}

	if result.StatusCode != 0 {
		t.Errorf("Expected status code 0 for timeout, got %d", result.StatusCode)
	}

	// Verify it completed within the timeout period
	if result.ResponseTimeMS > 6000 { // Should complete within 6 seconds (5s timeout + some overhead)
		t.Errorf("Request took too long: %d ms", result.ResponseTimeMS)
	}
}
