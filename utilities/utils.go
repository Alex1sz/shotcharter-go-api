package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
)

func RespondWithAppError(w http.ResponseWriter, handlerError error, message string, statusCode int) {
	errorObject := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: statusCode,
	}

	log.Printf("AppError]: %s\n", handlerError)
	// Error.Printf("AppError]: %s\n", handlerError)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if json, err := json.Marshal(errorResource{Data: errorObject}); err == nil {
		w.Write(json)
	}
}
