package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go/utilities"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/alex1sz/shotcharter-go/models"
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
	params := mux.Vars(req)
	// log.Println("GET /teams/" + params["id"])

	t := models.Team{ID: params["id"]}
	log.Println(t)

	team, err := models.FindTeamByID(t.ID)

	log.Println(team)
	log.Print(&team)

	jsonResp, err := json.Marshal(&team)

	if err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

	// json.NewEncoder(w).Encode(team)
}
