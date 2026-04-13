package handlers

import (
	"net/http"
	logger "serv-test/log"
)

// HTTPError handles HTTP errors and writes appropriate response
func HTTPError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	logger.Logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// LogError logs an error using the centralized logger
func LogError(err error) {
	logger.Logger.Error(err.Error())
}

// LogInfo logs an info message using the centralized logger
func LogInfo(msg string) {
	logger.Logger.Info(msg)
}

// LogWarn logs a warning message using the centralized logger
func LogWarn(msg string) {
	logger.Logger.Warn(msg)
}

// LogDebug logs a debug message using the centralized logger
func LogDebug(msg string) {
	logger.Logger.Debug(msg)
}
