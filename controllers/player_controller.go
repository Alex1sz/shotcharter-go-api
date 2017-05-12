package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go/models"
	"github.com/alex1sz/shotcharter-go/utilities"
	// "log"
	"net/http"
)

// POST /players
func CreatePlayer(w http.ResponseWriter, req *http.Request) {
	var player models.Player
	err := json.NewDecoder(req.Body).Decode(&player)

	if err != nil {
		utils.RespondWithAppError(w, err, "Invalid player data", 500)
		return
	}
	playerIsValid, err := player.IsValid()

	if !playerIsValid {
		utils.RespondWithAppError(w, err, "Invalid player data", 500)
		return
	}
	player.Create()
	jsonResp, err := json.Marshal(player)

	if err != nil {
		utils.RespondWithAppError(w, err, "Unexpected error occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonResp)
}
