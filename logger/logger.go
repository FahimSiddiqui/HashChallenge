package logger

import (
	"log"
)

// Naive Logging demonstration

// TODO: In-production I would use logrus go-package to log requests.

type LoggingLevel struct {
	Trace int
	Debug int
	Info  int
	Error int
	Fatal int
	Panic int
}

var Level = &LoggingLevel{
	Trace: 0,
	Debug: 1,
	Info:  2,
	Error: 3,
	Fatal: 4,
	Panic: 5,
}

func PushLogs(message interface{}, level int) {
	switch level {
	case Level.Info, Level.Trace, Level.Debug: // Debug level logs
		log.Println(message) // Temporarily printing statements
	default:
		log.Panic("Invalid log implementation")
	}
	// TODO: Same log levels can be added for fatal, info, trace levels too.
}

func PushErrLogs(message interface{}, level int, err error) {
	switch level {
	case Level.Error: // Error level logs, temporarily printing statements
		log.Println(message, err.Error())
	}
}
