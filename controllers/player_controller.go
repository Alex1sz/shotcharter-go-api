package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go-api/models"
	"github.com/alex1sz/shotcharter-go-api/utilities"
	"net/http"
)

// POST /players
func CreatePlayer(w http.ResponseWriter, req *http.Request) {
	var player models.Player
	if err := json.NewDecoder(req.Body).Decode(&player); err != nil {
		utils.RespondWithAppError(w, err, "Invalid player data", 500)
		return
	}

	if err := player.Create(); err != nil {
		utils.RespondWithAppError(w, err, err.Error(), 404)
		return
	}
	utils.RespondWithJSON(w, player, 201)
}
