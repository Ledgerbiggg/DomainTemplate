package test

import (
	"testing"
)

// TestLogs test logs
func TestLogs(t *testing.T) {
	logger.Debug("test debug logs success")
	logger.Info("test info logs success")
	logger.Warn("test error logs success")
	logger.Error("test fatal logs success")
}
