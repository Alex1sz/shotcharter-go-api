package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go-api/models"
	"github.com/alex1sz/shotcharter-go-api/utilities"
	"github.com/gorilla/mux"
	// "log"
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
	utils.RespondWithJSON(w, team, 201)
}

// GET /teams/:id
func GetTeamByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	team, err := models.FindTeamByID(params["id"])

	if err != nil {
		utils.HandleFindError(w, err)
		return
	}
	utils.RespondWithJSON(w, team, 200)
}

// PATCH /teams/:id
func UpdateTeam(w http.ResponseWriter, req *http.Request) {
	team := decodeReqIntoTeam(w, req)

	if err := team.Update(); err != nil {
		utils.HandleFindError(w, err)
		return
	}
	utils.RespondWithJSON(w, team, 200)
}
