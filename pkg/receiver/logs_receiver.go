package receiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

// ✅ Ensure only one declaration of `logsReceiver`
type logsReceiver struct {
	logger       *zap.Logger
	config       *Config
	nextConsumer consumer.Logs
	cancel       context.CancelFunc
}

func (r *logsReceiver) Start(ctx context.Context, _ component.Host) error {
	ctx, r.cancel = context.WithCancel(ctx)

	go func() {
		ticker := time.NewTicker(50 * time.Millisecond)
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

	sl := rl.ScopeLogs().AppendEmpty()
	logRecord := sl.LogRecords().AppendEmpty()

	logRecord.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	logRecord.SetSeverityText("INFO")
	logRecord.Body().SetStr("Synthetic log message")

	r.logger.Info("Sending logs", zap.Int("count", logs.LogRecordCount()))

	if err := r.nextConsumer.ConsumeLogs(ctx, logs); err != nil {
		r.logger.Error("Failed to send logs", zap.Error(err))
	}
}
