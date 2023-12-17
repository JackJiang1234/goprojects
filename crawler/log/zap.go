package log

import (
	"time"

	"go.uber.org/zap"
)

func zapUsage() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	url := "www.google.com"
	logger.Info("failed to fetch url", zap.String("url", url), zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
}