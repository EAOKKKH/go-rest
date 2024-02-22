package logger

import (
	"go-rest/internal/config"
	"log"

	"github.com/sirupsen/logrus"
)

func NewLogger(cfg *config.Config) *logrus.Logger {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	logLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Panicf("error parsing log level: %v", err)
	}

	logger.SetLevel(logLevel)

	return logger
}
