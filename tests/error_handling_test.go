package tests

import (
	"testing"

	"github.com/jaycdave88/otel-synthetics/internal/exporter"
	"go.uber.org/zap"
)

func TestHandleInvalidURL(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	checker := exporter.NewHTTPChecker(logger)

	result := checker.CheckHTTP("invalid-url")
	if result.ErrorMessage == "" {
		t.Errorf("Expected error for invalid URL")
	}
}

func TestHandleNetworkFailure(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	checker := exporter.NewHTTPChecker(logger)

	result := checker.CheckHTTP("https://unknown.example.com")
	if result.ErrorMessage == "" {
		t.Errorf("Expected network failure error")
	}
}
