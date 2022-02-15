package customlogger

import (
	"github.com/rs/zerolog"
)

var loggerInstance *zerolog.Logger

func SetLoggerInstance(Logger *zerolog.Logger) {
	loggerInstance = Logger
}

func GetLoggerInstance() *zerolog.Logger {
	return loggerInstance
}
