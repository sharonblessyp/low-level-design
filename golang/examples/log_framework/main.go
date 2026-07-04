package main

import (
	"fmt"
	"log_framework/core"
	appender "log_framework/log_appenders"
	formatter "log_framework/log_formatters"
	"log_framework/models"
	"time"
)

func main() {
	// get logger instance
	Logger := core.NewLoggerInstance(models.Debug, appender.NewFileAppender(formatter.NewTextFormatter()))

	// create a log
	logbuilder := models.NewLogBuilder("This is an info log")
	log := logbuilder.SetMessage("This is an info log").SetTimeStamp(time.Now()).SetLogLevel(models.Info).Build()
	err := Logger.Info(log)
	if err != nil {
		fmt.Printf("Error logging info message: %v\n", err)
	}

	log = logbuilder.SetMessage("This is a debug log").SetTimeStamp(time.Now()).SetLogLevel(models.Debug).Build()
	err = Logger.Debug(log)
	if err != nil {
		fmt.Printf("Error logging debug message: %v\n", err)
	}

	log = logbuilder.SetMessage("This is a warning log").SetTimeStamp(time.Now()).SetLogLevel(models.Warn).Build()
	err = Logger.Warn(log)
	if err != nil {
		fmt.Printf("Error logging warning message: %v\n", err)
	}

	log = logbuilder.SetMessage("This is an error log").SetTimeStamp(time.Now()).SetLogLevel(models.Error).Build()
	err = Logger.Error(log)
	if err != nil {
		fmt.Printf("Error logging error message: %v\n", err)
	}
}
