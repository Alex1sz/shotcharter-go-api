package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go/models"
	"github.com/alex1sz/shotcharter-go/utilities"
	"github.com/gorilla/mux"
	"log"
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

	jsonResp, err := json.Marshal(team)
	if err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonResp)
}

// GET /teams/:id
func GetTeamByID(w http.ResponseWriter, req *http.Request) {
	log.Println("GET request /teams/:id")
	params := mux.Vars(req)
	team, err := models.FindTeamByID(params["id"])

	if err != nil {
		utils.HandleFindError(w, err)
		return
	}
	jsonResp, err := json.Marshal(team)

	if err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
