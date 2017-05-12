package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go/utilities"
	"github.com/gorilla/mux"
	// "log"
	"net/http"
	// neccessary to catch sql.ErrNoRows
	// "database/sql"

	"github.com/alex1sz/shotcharter-go/models"
)

// GET /games/:id
func GetGameByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var game models.Game
	game, err := models.FindGameByID(params["id"])

	if err != nil {
		utils.HandleFindError(w, err)
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

// POST /games
func CreateGame(w http.ResponseWriter, req *http.Request) {
	var game models.Game
	err := json.NewDecoder(req.Body).Decode(&game)

	if err != nil {
		utils.RespondWithAppError(w, err, "Invalid team data", 500)
		return
	}
	game, err = game.Create()

	if err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	jsonResp, err := json.Marshal(game)

	if err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonResp)
}
