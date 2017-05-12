package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go/utilities"
	// gorilla/mux used for req params
	// "github.com/gorilla/mux"
	// "log"
	"net/http"
	// neccessary to catch sql.ErrNoRows
	// "database/sql"

	"github.com/alex1sz/shotcharter-go/models"
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
		utils.RespondWithAppError(w, err, "Shot associations are not valid", 500)
	}
	shot.Create()
	jsonResp, err := json.Marshal(shot)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonResp)
}
