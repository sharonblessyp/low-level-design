package log_formatters

import "log_framework/models"

type TextFormatter struct {
}

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

func (tf *TextFormatter) Format(log models.Log) string {
	// [date] [level] message
	return "[" + log.TimeStamp.Format("2006-01-02T15:04:05 -070000") + "]" + "[" + models.LogLevelToString[log.Level] + "]" + log.Message
}
