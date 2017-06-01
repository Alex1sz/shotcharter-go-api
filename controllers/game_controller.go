package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go-api/models"
	"github.com/alex1sz/shotcharter-go-api/utilities"
	"github.com/gorilla/mux"
	"net/http"
)

// GET /games/:id
func GetGameByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	game, err := models.FindGameByID(params["id"])

	if err != nil {
		utils.HandleFindError(w, err)
		return
	}
	utils.RespondWithJSON(w, game, 200)
}

// POST /games
func CreateGame(w http.ResponseWriter, req *http.Request) {
	var game models.Game
	if err := json.NewDecoder(req.Body).Decode(&game); err != nil {
		utils.RespondWithAppError(w, err, "Invalid team data", 500)
		return
	}
	if err := game.Create(); err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	utils.RespondWithJSON(w, game, 201)
}
