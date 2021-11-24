package common

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogging() {
	logLevel, err := log.ParseLevel(os.Getenv("AoC_LogLevel"))

	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
}
