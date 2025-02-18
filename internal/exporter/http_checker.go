package exporter

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// HTTPResult contains the result of an HTTP check
type HTTPResult struct {
	StatusCode     int
	ResponseTimeMS int64
	ErrorMessage   string
}

// HTTPChecker performs HTTP synthetic checks
type HTTPChecker struct {
	logger *zap.Logger
}

// NewHTTPChecker creates a new HTTP checker
func NewHTTPChecker(logger *zap.Logger) *HTTPChecker {
	return &HTTPChecker{
		logger: logger,
	}
}

// CheckHTTP performs an HTTP check on the given URL
func (c *HTTPChecker) CheckHTTP(url string) HTTPResult {
	start := time.Now()
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	elapsed := time.Since(start).Milliseconds()

	if err != nil {
		return HTTPResult{
			StatusCode:     0,
			ResponseTimeMS: elapsed,
			ErrorMessage:   err.Error(),
		}
	}
	defer resp.Body.Close()

	return HTTPResult{
		StatusCode:     resp.StatusCode,
		ResponseTimeMS: elapsed,
		ErrorMessage:   "",
	}
}
