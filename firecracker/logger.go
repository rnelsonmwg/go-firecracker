package firecracker

import (
	"errors"
	"net/http"
)

type LoggingLevel string

const (
	ErrorLoggingLevel   LoggingLevel = "Error"
	WarningLoggingLevel              = "Warning"
	InfoLoggingLevel                 = "Info"
	DebugLoggingLevel                = "Debug"
)

func (cracker *Firecracker) InitLogger(loggerPipeName string, metricsPipeName string, level LoggingLevel, showLevel bool, showOrigin bool) error {
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

	if resp.StatusCode() != http.StatusNoContent {
		if e, ok := resp.Error().(*apiError); ok {
			return errors.New(e.Message)
		}

		return errInvalidServerError
	}

	return errInvalidServerResponse
}
