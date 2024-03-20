package errorHandling

import (
	// Standard
	"log/slog"
)

// if err, fatal
func PanicCheck(err error, message string) {
	if err != nil {
		slog.Error(message, err)
		panic(err)
	}
}

// if err, log
func LogCheck(err error, message string) {
	if err != nil {
		slog.Error(message, err)
	}
}
