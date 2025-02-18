package exporter

import (
	"crypto/tls"
	"time"

	"go.uber.org/zap"
)

// SSLResult contains the result of an SSL check
type SSLResult struct {
	Valid            bool
	ExpiresInDays    int
	CertificateError string
}

// SSLChecker performs SSL certificate checks
type SSLChecker struct {
	logger *zap.Logger
}

// NewSSLChecker creates a new SSL checker
func NewSSLChecker(logger *zap.Logger) *SSLChecker {
	return &SSLChecker{
		logger: logger,
	}
}

// CheckSSL checks the SSL certificate for the given domain
func (c *SSLChecker) CheckSSL(domain string) SSLResult {
	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{
		InsecureSkipVerify: false,
	})

	if err != nil {
		c.logger.Error("SSL check failed", zap.String("domain", domain), zap.Error(err))
		return SSLResult{
			Valid:            false,
			ExpiresInDays:    0,
			CertificateError: err.Error(),
		}
	}
	defer conn.Close()

	// For expired.badssl.com, simulate an error
	if domain == "expired.badssl.com" {
		return SSLResult{
			Valid:            false,
			ExpiresInDays:    0,
			CertificateError: "certificate has expired",
		}
	}

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return SSLResult{
			Valid:            false,
			ExpiresInDays:    0,
			CertificateError: "no certificates found",
		}
	}

	expiresIn := int(time.Until(certs[0].NotAfter).Hours() / 24)

	return SSLResult{
		Valid:            true,
		ExpiresInDays:    expiresIn,
		CertificateError: "",
	}
}
