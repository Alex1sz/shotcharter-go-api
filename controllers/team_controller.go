package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go-api/models"
	"github.com/alex1sz/shotcharter-go-api/utilities"
	"github.com/gorilla/mux"
	"net/http"
)

// POST /teams
func CreateTeam(w http.ResponseWriter, req *http.Request) {
	var team models.Team
	err := json.NewDecoder(req.Body).Decode(&team)

	if err != nil {
		utils.RespondWithAppError(w, err, "Invalid team data", 500)
		return
	}
	team.Create()
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
	var team models.Team
	err := json.NewDecoder(req.Body).Decode(&team)

	if err != nil {
		utils.RespondWithAppError(w, err, "Invalid team data", 500)
		return
	}
	err = team.Update()

	if err != nil {
		utils.HandleFindError(w, err)
		return
	}
	utils.RespondWithJSON(w, team, 200)
}
