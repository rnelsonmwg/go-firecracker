package firecracker

import (
	"net/http"
)

type LoggingLevel string

const (
	ErrorLoggingLevel   LoggingLevel = "Error"
	WarningLoggingLevel LoggingLevel = "Warning"
	InfoLoggingLevel    LoggingLevel = "Info"
	DebugLoggingLevel   LoggingLevel = "Debug"
)

// InitLogger adds a logger with two named pipes for logging and metrics at a specific LoggingLevel.
// `showLevel` adds a logging level to the logger output, and `showOrigin` adds a file:line number.
func (cracker *Firecracker) InitLogger(
	loggerPipeName string,
	metricsPipeName string,
	level LoggingLevel,
	showLevel bool,
	showOrigin bool,
) error {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetError(&apiError{}).
		SetBody(map[string]interface{}{
			"log_fifo":        loggerPipeName,
			"metrics_fifo":    metricsPipeName,
			"level":           level,
			"show_level":      showLevel,
			"show_log_origin": showOrigin, // file path and line number
		}).
		Put("/logger")

	if err != nil {
		return err
	}

	return cracker.responseErrorStrict(resp, http.StatusNoContent)
}
