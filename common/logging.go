package common

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogging() {
	logLevel, err := log.ParseLevel(os.Getenv("AOC_LOG_LEVEL"))

	if err != nil {
		logLevel = log.WarnLevel
	}

	log.SetLevel(logLevel)
}
