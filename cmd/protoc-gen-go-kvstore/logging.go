package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func setupLogger() {
	logFile := os.Getenv("PROTOC_GEN_GO_KVSTORE_LOGS")
	if logFile == "" {
		logFile = "/dev/null"
	}

	// open a file
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(f)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}
