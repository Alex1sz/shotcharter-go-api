package utils

import (
	"database/sql"
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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if json, err := json.Marshal(errorResource{Data: errorObject}); err == nil {
		w.Write(json)
	}
}

// HandleFindError responds when error occurs in FindByID methods
func HandleFindError(w http.ResponseWriter, err error) {
	if err == sql.ErrNoRows {
		RespondWithAppError(w, err, "Error not found", 404)
	} else {
		RespondWithAppError(w, err, "An unexpected error has occurred", 500)
	}
	return
}

// abstracts out shared marshaling of responseJSON, err handling, setting of headers
func RespondWithJSON(w http.ResponseWriter, modelObj interface{}) {
	jsonResp, err := json.Marshal(modelObj)

	if err != nil {
		RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonResp)
}
