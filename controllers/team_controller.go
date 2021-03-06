package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go-api/models"
	"github.com/alex1sz/shotcharter-go-api/utilities"
	"github.com/gorilla/mux"
	"net/http"
)

// team specific decoder needed to avoid performance hit of using reflect
func decodeReqIntoTeam(w http.ResponseWriter, req *http.Request) (team models.Team) {
	if err := json.NewDecoder(req.Body).Decode(&team); err != nil {
		utils.RespondWithAppError(w, err, "Invalid team data", 500)
		return
	}
	return
}

// POST /teams
func CreateTeam(w http.ResponseWriter, req *http.Request) {
	team := decodeReqIntoTeam(w, req)

	if err := team.Create(); err != nil {
		utils.RespondWithAppError(w, err, "Unexpected error", 500)
		return
	}
	respondWithLeanTeamJSON(&team, w, 201)
}

// GET /teams/:id
func GetTeamByID(w http.ResponseWriter, req *http.Request) {
	team, err := models.FindTeamByID(mux.Vars(req)["id"])

	if err != nil {
		utils.HandleFindError(w, err)
		return
	}
	respondWithLeanTeamJSON(&team, w, 200)
}

// PATCH /teams/:id
func UpdateTeam(w http.ResponseWriter, req *http.Request) {
	team := decodeReqIntoTeam(w, req)

	if err := team.Update(); err != nil {
		utils.HandleFindError(w, err)
		return
	}
	respondWithLeanTeamJSON(&team, w, 200)
}

// json marshaler for lean Team response
func respondWithLeanTeamJSON(team *models.Team, w http.ResponseWriter, statusCode int) {
	var leanPlayers []models.LeanPlayer

	for _, player := range team.Players {
		leanPlayer := models.LeanPlayer{Player: player}
		leanPlayers = append(leanPlayers, leanPlayer)
	}
	json, err := json.Marshal(models.TeamResp{Team: team, Players: leanPlayers})

	if err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	utils.SetHeaders(w, statusCode)
	w.Write(json)
}
