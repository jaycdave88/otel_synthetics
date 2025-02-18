package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jaycdave88/otel-synthetics/internal/exporter"
	"go.uber.org/zap"
)

func setupTestServer(t *testing.T) string {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><body>test page</body></html>"))
	}))
	t.Cleanup(func() { ts.Close() })
	return ts.URL
}

func TestCheckBrowserSuccess(t *testing.T) {
	t.Parallel()
	logger, _ := zap.NewDevelopment()
	checker := exporter.NewBrowserChecker(logger)

	testURL := setupTestServer(t)
	err := checker.CheckBrowser(testURL)
	if err != nil {
		t.Errorf("Expected successful page load, but got error: %v", err)
	}
}

func TestCheckBrowserFailure(t *testing.T) {
	t.Parallel()
	logger, _ := zap.NewDevelopment()
	checker := exporter.NewBrowserChecker(logger)

	err := checker.CheckBrowser("http://invalid-url-that-does-not-exist")
	if err == nil {
		t.Error("Expected an error for invalid URL but got none")
	}
}

func TestCheckBrowserLoadTimeout(t *testing.T) {
	t.Parallel()
	logger, _ := zap.NewDevelopment()
	checker := exporter.NewBrowserChecker(logger)

	// Create a server that delays response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(6 * time.Second) // Longer than our timeout
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	start := time.Now()
	err := checker.CheckBrowser(ts.URL)
	duration := time.Since(start)

	if err == nil {
		t.Error("Expected a timeout error but got none")
	}

	if !strings.Contains(err.Error(), "timeout") && !strings.Contains(err.Error(), "Timeout") {
		t.Errorf("Expected timeout error, got: %v", err)
	}

	// Check that it didn't take longer than expected
	if duration > 12*time.Second { // giving some buffer over the 10s timeout
		t.Errorf("Test took too long: %v", duration)
	}
}
