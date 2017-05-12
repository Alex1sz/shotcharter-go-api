package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go/utilities"
	"github.com/gorilla/mux"
	// "log"
	"net/http"

	"github.com/alex1sz/shotcharter-go/models"
)

// POST /teams/:id/players
func CreatePlayer(w http.ResponseWriter, req *http.Request) {
	var player models.Player
	params := mux.Vars(req)
	team, err := models.FindTeamByID(params["id"])

	if err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error occurred req not valid", 500)
	}
	player.Team = team
	err = json.NewDecoder(req.Body).Decode(&player)

	if err != nil {
		utils.RespondWithAppError(w, err, "Invalid player data", 500)
	}
	player.Create()
	jsonResp, err := json.Marshal(player)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonResp)
}
