package tests

import (
	"sync"
	"testing"

	"github.com/jaycdave88/otel-synthetics/internal/exporter"
	"go.uber.org/zap"
)

func TestLoadHandling(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	checker := exporter.NewHTTPChecker(logger)

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ { // Simulating 50 concurrent tests
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := checker.CheckHTTP("https://example.com")
			if result.StatusCode != 200 && result.ErrorMessage == "" {
				t.Errorf("Unexpected response: status=%d, error=%s", result.StatusCode, result.ErrorMessage)
			}
		}()
	}
	wg.Wait()
}
