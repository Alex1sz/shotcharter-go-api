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
	utils.RespondWithJSON(w, player)
}
