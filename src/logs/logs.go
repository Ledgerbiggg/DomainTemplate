package logs

import (
	"DomainTemplate/src/config"
	"fmt"
	"go.uber.org/fx"
	"log"
	"os"
	"time"
)

// ConsoleLogger LOGGER
type ConsoleLogger struct {
	logLevel int // log level
}

// LOG_LEVEL
const (
	DebugLevel = iota // 0
	InfoLevel         // 1
	WarnLevel         // 2
	ErrorLevel        // 3
)

// Params injector
type Params struct {
	fx.In
	Config *config.GConfig
}

// NewConsoleLogger gets a new ConsoleLogger
func NewConsoleLogger(p Params) *ConsoleLogger {
	return &ConsoleLogger{
		logLevel: p.Config.LogLevel,
	}
}

func (l *ConsoleLogger) log(level int, levelStr, message string) {
	if level >= l.logLevel {
		timestamp := time.Now().Format("2006-01-02 15:04:05")

		logMessage := fmt.Sprintf("[%s] [%s]: %s\n", timestamp, levelStr, message)
		log.Println(logMessage)

		file, err := os.OpenFile(
			fmt.Sprintf("./logs/%s.txt", time.Now().Format("2006-01-02")),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening log file: %v\n", err)
			return
		}
		defer file.Close()

		if _, err := file.WriteString(logMessage); err != nil {
			fmt.Printf("Error writing to log file: %v\n", err)
		}
	}
}

// Debug DEBUG
func (l *ConsoleLogger) Debug(message string) {
	l.log(DebugLevel, "DEBUG", message)
}

// Info INFO
func (l *ConsoleLogger) Info(message string) {
	l.log(InfoLevel, "INFO", message)
}

// Warn WARN
func (l *ConsoleLogger) Warn(message string) {
	l.log(WarnLevel, "WARN", message)
}

// Error ERROR
func (l *ConsoleLogger) Error(message string) {
	l.log(ErrorLevel, "ERROR", message)
}
