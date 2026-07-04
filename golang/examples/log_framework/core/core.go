package core

import (
	"fmt"
	appender "log_framework/log_appenders"
	"log_framework/models"
	"sync"
)

type Logger struct {
	Config LoggerConfig
	mu     sync.Mutex
}

var (
	once     sync.Once
	Instance *Logger //The global instance variable is what all goroutines share.
)

func NewLoggerInstance(logLevel models.LogLevel, logAppender appender.Appender) *Logger {
	once.Do(func() {
		Instance = &Logger{
			Config: NewLoggerConfig(logLevel, logAppender),
			mu:     sync.Mutex{},
		}
	})
	return Instance
}

// config struct to hold the configuration for the logger
type LoggerConfig struct {
	LogLevel    models.LogLevel
	LogAppender appender.Appender
}

func NewLoggerConfig(logLevel models.LogLevel, logAppender appender.Appender) LoggerConfig {
	return LoggerConfig{
		LogLevel:    logLevel,
		LogAppender: logAppender,
	}
}

func (l *Logger) Log(log models.Log, configuredLogLevel models.LogLevel, logLevelToLog models.LogLevel) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if configuredLogLevel <= logLevelToLog {
		err := l.Config.LogAppender.Append(log)
		if err != nil {
			fmt.Println("Error appending log:", err)
			return err
		}
	}
	return nil
}

func (l *Logger) Info(log models.Log) error {
	return l.Log(log, l.Config.LogLevel, models.Info)
}

func (l *Logger) Debug(log models.Log) error {
	return l.Log(log, l.Config.LogLevel, models.Debug)
}

func (l *Logger) Warn(log models.Log) error {
	return l.Log(log, l.Config.LogLevel, models.Warn)
}

func (l *Logger) Error(log models.Log) error {
	return l.Log(log, l.Config.LogLevel, models.Error)
}
