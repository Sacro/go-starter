package logging

import (
	"fmt"
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/logfmt"
	"github.com/apex/log/handlers/text"
)

// Configure configures the logger with the specified handler and level
func Configure(logHandler, logLevel *string) error {
	if *logHandler != "" {
		handler, ok := logOutputs[*logHandler]
		if !ok {
			return fmt.Errorf("Level: %v not found", *logHandler)
		}
		log.SetHandler(handler)
	}

	level, err := log.ParseLevel(*logLevel)

	if err != nil {
		return fmt.Errorf("Unable to parse log level: %w", err)
	}

	log.SetLevel(level)
	return nil
}

// GetLogLevels returns the log levels available as a comma seperated string
func GetLogLevels() string {
	return strings.Join([]string{
		log.DebugLevel.String(),
		log.InfoLevel.String(),
		log.WarnLevel.String(),
		log.ErrorLevel.String(),
		log.FatalLevel.String(),
	}, ", ")
}

// logOutputs is a map of strings to log handlers
var logOutputs = map[string]log.Handler{
	"cli":    cli.Default,
	"json":   json.Default,
	"logfmt": logfmt.Default,
	"text":   text.Default,
}

// GetLogOutputs returns the log handlers as a comma seperated string
func GetLogOutputs() string {
	return strings.Join([]string{
		"cli",
		"json",
		"logfmt",
		"text",
	}, ", ")
}
