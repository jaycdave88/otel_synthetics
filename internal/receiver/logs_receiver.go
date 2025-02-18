// internal/receiver/logs_receiver.go
package receiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver"
)

type logsReceiver struct {
	params   receiver.CreateSettings
	config   *Config
	consumer consumer.Logs
	cancel   context.CancelFunc
}

func newLogsReceiver(params receiver.CreateSettings, config *Config, consumer consumer.Logs) (receiver.Logs, error) {
	return &logsReceiver{
		params:   params,
		config:   config,
		consumer: consumer,
	}, nil
}

func (r *logsReceiver) Start(ctx context.Context, _ component.Host) error {
	ctx, r.cancel = context.WithCancel(ctx)

	// Start a goroutine that generates logs periodically
	go func() {
		ticker := time.NewTicker(10 * time.Second) // Configurable interval
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				r.generateAndSendLogs(ctx)
			}
		}
	}()

	return nil
}

func (r *logsReceiver) Shutdown(ctx context.Context) error {
	if r.cancel != nil {
		r.cancel()
	}
	return nil
}

func (r *logsReceiver) generateAndSendLogs(ctx context.Context) {
	logs := plog.NewLogs()
	rl := logs.ResourceLogs().AppendEmpty()

	// Set resource attributes if needed
	// rl.Resource().Attributes().PutStr("service.name", "synthetic-monitoring")

	sl := rl.ScopeLogs().AppendEmpty()
	logRecord := sl.LogRecords().AppendEmpty()

	// Fill in log record data
	logRecord.SetTimestamp(plog.NewTimestampFromTime(time.Now()))
	logRecord.SetSeverityText("INFO")
	logRecord.Body().SetStr("Synthetic log message")

	// Send logs to the next consumer
	if err := r.consumer.ConsumeLogs(ctx, logs); err != nil {
		r.params.Logger.Error("Failed to send logs", err)
	}
}
