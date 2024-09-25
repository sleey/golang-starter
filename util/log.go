package util

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitializeLog() {
	// Determine log level
	logLevel := zerolog.InfoLevel
	if IsLocalDev() {
		logLevel = zerolog.DebugLevel
	}

	// Set up io.Writer
	ioWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime}
	if !IsLocalDev() {
		ioWriter = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	}

	// Initialize logger
	log.Logger = zerolog.New(ioWriter).
		Level(logLevel).
		With().
		Timestamp().
		Caller().
		Str("environment", Environment).
		Str("service", "backend-api").
		Logger()
}
