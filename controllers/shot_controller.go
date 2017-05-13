package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go/utilities"
	// "log"
	"github.com/alex1sz/shotcharter-go/models"
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
	utils.RespondWithJSON(w, shot)
}
