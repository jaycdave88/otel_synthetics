FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache \
    build-base \
    git \
    nodejs \
    npm \
    chromium \
    chromium-chromedriver

# Copy go.mod & download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build application
RUN CGO_ENABLED=0 GOOS=linux go build -o /otel-synthetics cmd/otel-synthetics-collector/main.go

# Final runtime image
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache \
    ca-certificates \
    chromium \
    nodejs \
    npm \
    font-noto

# Copy compiled binary
COPY --from=builder /otel-synthetics /usr/local/bin/

# Copy OpenTelemetry Collector config
COPY deploy/otel-collector-config.yaml /etc/otel/config.yaml

# Set environment variables
ENV OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317

ENTRYPOINT ["/usr/local/bin/otel-synthetics"]