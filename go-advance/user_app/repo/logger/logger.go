package logger

import (
	"log"
	"os"

	"github.com/rs/zerolog"
)

var instance *zerolog.Logger

func GetLogger() *zerolog.Logger {
	var logpath = "zerologs.log"
	tempFile, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error finding log file")
	}
	if instance == nil {
		logLevel := zerolog.InfoLevel
		zerolog.SetGlobalLevel(logLevel)
		logger := zerolog.New(tempFile).With().Timestamp().Logger()
		instance = &logger
	}
	return instance
}
