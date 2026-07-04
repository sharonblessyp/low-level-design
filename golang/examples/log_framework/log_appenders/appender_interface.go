package appender

import (
	"log_framework/log_formatters"
	"log_framework/models"
)

/*
1. function must implement to output logs
- files
- console
2. Formats the log
*/

type Appender interface {
	SetFormatters(formatter log_formatters.LogFormatter)
	Append(log models.Log) error
}
