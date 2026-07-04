package appender

import (
	"log_framework/log_formatters"
	"log_framework/models"
	"os"
)

type FileAppender struct {
	Formatter log_formatters.LogFormatter
}

func NewFileAppender(formatter log_formatters.LogFormatter) *FileAppender {
	return &FileAppender{
		Formatter: formatter,
	}
}

func (fa *FileAppender) SetFormatters(formatter log_formatters.LogFormatter) {
	fa.Formatter = formatter
	return
}

func (fa *FileAppender) Append(log models.Log) error {
	formattedLog := fa.Formatter.Format(log)

	// create a file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = file.WriteString(formattedLog)
	if err != nil {
		return err
	}

	// move the cursor to next line
	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}

	return nil
}
