package logging

import (
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/logfmt"
	"github.com/apex/log/handlers/text"
)

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

// LogOutputs is a map of strings to log handlers
var LogOutputs = map[string]log.Handler{
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
