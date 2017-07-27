package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go-api/models"
	"github.com/alex1sz/shotcharter-go-api/utilities"
	"net/http"
)

func decodeReqIntoPlayer(w http.ResponseWriter, req *http.Request) (player models.Player) {
	if err := json.NewDecoder(req.Body).Decode(&player); err != nil {
		utils.RespondWithAppError(w, err, "Invalid player data in request", 500)
		return
	}
	return
}

// POST /players
func CreatePlayer(w http.ResponseWriter, req *http.Request) {
	player := decodeReqIntoPlayer(w, req)

	if err := player.Create(); err != nil {
		utils.RespondWithAppError(w, err, err.Error(), 404)
		return
	}
	utils.RespondWithJSON(w, player, 201)
}

// PATCH /players/:id
// TO DO add generic decodeReq method that takes and returns interface{}
func UpdatePlayer(w http.ResponseWriter, req *http.Request) {
	player := decodeReqIntoPlayer(w, req)

	if err := player.Update(); err != nil {
		utils.RespondWithAppError(w, err, "Unexpected internal error", 500)
		return
	}
	utils.RespondWithJSON(w, player, 201)
}
