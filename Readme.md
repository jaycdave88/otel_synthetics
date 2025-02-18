# OpenTelemetry Synthetic Monitoring - Still Work In Progress (WIP)

## Overview
This project implements a **Synthetic Monitoring System** with **OpenTelemetry Collector** integration. It supports various synthetic checks such as **HTTP monitoring, SSL certificate validation, DNS resolution, browser automation**, and **TCP connectivity** checks.

The **OpenTelemetry Collector** is used to gather and forward synthetic test results to an observability backend such as **Observe Inc**.

## Features
- **SSL Certificate Monitoring**
- **HTTP Endpoint Monitoring**
- **DNS Resolution Checks**
- **Browser Automation for Full Page Load Metrics**
- **TCP/Port Connectivity Monitoring**
- **OpenTelemetry Collector Integration**
- **Configurable Synthetic Tests**
- **Batch Processing for Efficient Data Export**

## Prerequisites
Ensure you have the following installed on your machine:

### Required Dependencies
- [Go 1.24+](https://go.dev/doc/install)
- [Docker](https://www.docker.com/get-started)
- [OpenTelemetry Collector](https://opentelemetry.io/docs/collector/getting-started/)
- [Node.js](https://nodejs.org/en) *(Required for Browser Automation)*
- [Chromium & ChromeDriver](https://www.chromium.org/getting-involved/download-chromium/) *(Required for Browser Checks)*

## Installation
Clone the repository:
```sh
$ git clone https://github.com/jaycdave88/otel-synthetics.git
$ cd otel-synthetics
```

### Install Go Dependencies
```sh
$ go mod tidy
```

### Build OpenTelemetry Collector
Build a custom OpenTelemetry Collector with synthetic monitoring capabilities:
```sh
$ make build
```

Alternatively, you can run:
```sh
$ go build -o otel-synthetics cmd/otel-synthetics-collector/main.go
```

## Running the OpenTelemetry Collector Locally

### 1. Start the OpenTelemetry Collector
Run the collector with the provided configuration:
```sh
$ otelcol --config deploy/otel-collector-config.yaml
```

If using Docker:
```sh
$ docker run --rm -v $(pwd)/deploy/otel-collector-config.yaml:/etc/otel/config.yaml otel/opentelemetry-collector --config /etc/otel/config.yaml
```

### 2. Start Synthetic Monitoring
Once the collector is running, execute synthetic tests:
```sh
$ ./otel-synthetics
```

## Configuration
The **synthetic monitoring configuration** is defined in `deploy/otel-collector-config.yaml`. This file includes various test definitions, such as:

### Example HTTP Check:
```yaml
receivers:
  synthetics:
    collection_interval: 60s
    tests:
      - name: "api-health-check"
        type: "http"
        enabled: true
        http:
          url: "https://api.example.com/health"
          method: "GET"
          timeout: "10s"
          validation:
            status_code: 200
```

### Example SSL Certificate Check:
```yaml
      - name: "cert-expiry-check"
        type: "ssl"
        enabled: true
        ssl:
          host: "example.com"
          port: 443
          expiry_warning_days: 30
```


Run the following command to verify data is being sent to Observe:
```sh
$ otelcol --config deploy/otel-collector-config.yaml --log-level debug
```

## Running Tests
Run unit tests to validate synthetic monitoring components:
```sh
$ go test ./tests/...
```

## Deployment
For **containerized deployment**, use the provided Dockerfile:
```sh
$ docker build -t otel-synthetics .
$ docker run --rm -p 4317:4317 otel-synthetics
```

## Troubleshooting
- Check logs for errors:
```sh
$ tail -f otel-collector.log
```
- Verify collector is running:
```sh
$ curl http://localhost:4317/health
```
- Validate configuration:
```sh
$ otelcol --config deploy/otel-collector-config.yaml --dry-run
```

## Resources
- [OpenTelemetry Collector Docs](https://opentelemetry.io/docs/collector/custom-collector/)
- [Observe Inc. OpenTelemetry Integration](https://docs.observeinc.com/en/latest/integrations/opentelemetry/)

