package log_formatters

import "log_framework/models"

type LogFormatter interface {
	Format(log models.Log) string
}
