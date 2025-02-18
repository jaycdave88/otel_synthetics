package tests

import (
	"sync"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/jaycdave88/otel-synthetics/internal/exporter"
)

func TestSyntheticLoadHandling(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	httpChecker := exporter.NewHTTPChecker(logger)
	sslChecker := exporter.NewSSLChecker(logger)

	var wg sync.WaitGroup
	numRequests := 100 // High-frequency test

	start := time.Now()

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Run HTTP check
			httpResult := httpChecker.CheckHTTP("https://example.com")
			if httpResult.StatusCode != 200 {
				t.Errorf("HTTP Check failed with status %d", httpResult.StatusCode)
			}

			// Run SSL check
			sslResult := sslChecker.CheckSSL("google.com")
			if !sslResult.Valid {
				t.Errorf("SSL Check failed for google.com")
			}
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)

	t.Logf("Completed %d synthetic tests in %v", numRequests, elapsed)
}
