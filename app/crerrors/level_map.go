package crerrors

import "net/http"

const (
	// LevelInfo : info level error
	LevelInfo = "info"

	// LevelError : error level error
	LevelError = "error"

	// LevelWarn : warn level error
	LevelWarn = "warn"

	// LevelFatal : fatal level error
	LevelFatal = "fatal"
)

var levelMap = map[int]string{
	http.StatusLocked:              LevelWarn,
	http.StatusBadRequest:          LevelError,
	http.StatusNotFound:            LevelError,
	http.StatusUnauthorized:        LevelError,
	http.StatusInternalServerError: LevelError,
	http.StatusGatewayTimeout:      LevelError,
}

func getLevel(requestStatusCode int) string {
	level, ok := levelMap[requestStatusCode]

	if !ok {
		level = LevelError
	}

	return level
}
