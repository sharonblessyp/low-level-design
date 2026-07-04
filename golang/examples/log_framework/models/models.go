package models

import "time"

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

var LogLevelToString = map[LogLevel]string{
	Debug: "DEBUG",
	Info:  "INFO",
	Warn:  "WARN",
	Error: "ERROR",
}

type Log struct {
	Message   string
	TimeStamp time.Time
	Level     LogLevel
}

/*
1. builder struct
2. methods to set all args
3. build method to return the final object
*/
type LogBuilder struct {
	msg       string
	timeStamp time.Time
	level     LogLevel
}

func NewLogBuilder(msg string) *LogBuilder {
	return &LogBuilder{
		// consider default level to be info
		level: Info,
	}
}

func (lb *LogBuilder) SetMessage(msg string) *LogBuilder {
	lb.msg = msg
	return lb
}

func (lb *LogBuilder) SetTimeStamp(timeStamp time.Time) *LogBuilder {
	lb.timeStamp = timeStamp
	return lb
}

func (lb *LogBuilder) SetLogLevel(level LogLevel) *LogBuilder {
	lb.level = level
	return lb
}

func (lb *LogBuilder) Build() Log {
	return Log{
		Message:   lb.msg,
		TimeStamp: lb.timeStamp,
		Level:     lb.level,
	}
}
