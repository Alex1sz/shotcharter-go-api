package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go-api/models"
	"github.com/alex1sz/shotcharter-go-api/utilities"
	// "log"
	"net/http"
)

func decodeReqIntoShot(w http.ResponseWriter, req *http.Request) (shot models.Shot) {
	if err := json.NewDecoder(req.Body).Decode(&shot); err != nil {
		utils.RespondWithAppError(w, err, "Invalid shot data", 500)
		return
	}
	return
}

// POST /shots
func CreateShot(w http.ResponseWriter, req *http.Request) {
	shot := decodeReqIntoShot(w, req)

	if err := shot.Create(); err != nil {
		utils.RespondWithAppError(w, err, "Unexpected error on create", 500)
		return
	}
	utils.RespondWithJSON(w, shot, 201)
}

// PATCH /shots/:id
func UpdateShot(w http.ResponseWriter, req *http.Request) {
	shot := decodeReqIntoShot(w, req)

	if err := shot.Update(); err != nil {
		utils.HandleFindError(w, err)
		return
	}
	utils.RespondWithJSON(w, shot, 200)
}
