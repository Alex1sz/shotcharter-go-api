package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go-api/models"
	"github.com/alex1sz/shotcharter-go-api/utilities"
	// "log"
	"net/http"
)

// POST /shots
func CreateShot(w http.ResponseWriter, req *http.Request) {
	var shot models.Shot
	err := json.NewDecoder(req.Body).Decode(&shot)

	if err != nil {
		utils.RespondWithAppError(w, err, "Invalid shot data", 500)
		return
	}
	shotIsValid, err := shot.IsValid()

	if !shotIsValid {
		utils.RespondWithAppError(w, err, "Invalid shot data: see associations", 500)
		return
	}
	shot.Create()
	utils.RespondWithJSON(w, shot, 201)
}

// PATCH /shots/:id
func UpdateShot(w http.ResponseWriter, req *http.Request) {
	var shot models.Shot
	err := json.NewDecoder(req.Body).Decode(&shot)

	if err != nil {
		utils.RespondWithAppError(w, err, "Invalid shot data", 500)
		return
	}
	err = shot.Update()

	if err != nil {
		utils.HandleFindError(w, err)
		return
	}
	utils.RespondWithJSON(w, shot, 200)
}
