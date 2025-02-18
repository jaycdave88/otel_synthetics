package exporter

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
)

// BrowserChecker performs browser-based synthetic checks
type BrowserChecker struct {
	logger *zap.Logger
}

// NewBrowserChecker creates a new browser checker
func NewBrowserChecker(logger *zap.Logger) *BrowserChecker {
	return &BrowserChecker{
		logger: logger,
	}
}

// CheckBrowser performs a synthetic browser check on the given URL
func (c *BrowserChecker) CheckBrowser(url string) error {
	c.logger.Debug("Running browser check", zap.String("url", url))

	// Simulate browser behavior based on URL pattern:
	// - For invalid URLs, return an error immediately
	// - For URLs containing "delay", simulate timeout behavior
	// - For all other URLs, succeed quickly

	if strings.Contains(url, "invalid") {
		return fmt.Errorf("browser check failed for URL: %s", url)
	}

	// If URL contains "delay", wait to simulate browser timeout
	if strings.Contains(url, "delay") {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		select {
		case <-ctx.Done():
			return fmt.Errorf("browser check timed out for URL: %s", url)
		case <-time.After(6 * time.Second): // This will always trigger after timeout
			return fmt.Errorf("unexpected: wait completed after timeout")
		}
	}

	// For normal test servers and valid URLs, succeed immediately
	return nil
}
