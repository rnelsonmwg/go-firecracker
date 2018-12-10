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

type Logger struct {
	LoggerPipeName  string       `json:"log_fifo,omitempty"`
	MetricsPipeName string       `json:"metrics_fifo,omitempty"`
	Level           LoggingLevel `json:"level,omitempty"`
	ShowLevel       bool         `json:"show_level,omitempty"`
	ShowOrigin      bool         `json:"show_log_origin,omitempty"`
}

// InitLogger adds a logger with two named pipes for logging and metrics at a specific LoggingLevel.
// `showLevel` adds a logging level to the logger output, and `showOrigin` adds a file:line number.
func (cracker *Firecracker) InitLogger(logger *Logger) error {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetError(&apiError{}).
		SetBody(logger).
		Put("/logger")

	if err != nil {
		return err
	}

	return cracker.responseErrorStrict(resp, http.StatusNoContent)
}
