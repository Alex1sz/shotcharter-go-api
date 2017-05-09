package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go/utilities"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	// neccessary to catch sql.ErrNoRows
	"database/sql"

	"github.com/alex1sz/shotcharter-go/models"
)

// GET /games/:id
func GetGameByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var game models.Game
	game, err := models.FindGameByID(params["id"])

	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithAppError(w, err, "An unexpected error has occurred", 404)
		} else {
			utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		}
		return
	}
	jsonResp, err := json.Marshal(game)

	if err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
