package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go-api/models"
	"github.com/alex1sz/shotcharter-go-api/utilities"
	"github.com/gorilla/mux"
	"net/http"
)

// GET /games/:id
func GetGameByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	game, err := models.FindGameByID(params["id"])

	if err != nil {
		utils.HandleFindError(w, err)
		return
	}
	jsonResp, err := marshalLeanGameJSON(&game)

	if err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	utils.SetHeaders(w, 200)
	w.Write(jsonResp)
}

// POST /games
func CreateGame(w http.ResponseWriter, req *http.Request) {
	var game models.Game
	if err := json.NewDecoder(req.Body).Decode(&game); err != nil {
		utils.RespondWithAppError(w, err, "Invalid team data", 500)
		return
	}
	if err := game.Create(); err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	utils.RespondWithJSON(w, game, 201)
}

// lean json marshaler for non redundant nested json for embedded types
func marshalLeanGameJSON(game *models.Game) ([]byte, error) {
	var leanAwayShots, leanHomeShots []models.PublicShot

	for _, shot := range append(game.HomeShots, game.AwayShots...) {
		publicShot := models.PublicShot{Shot: shot}

		if shot.Team.ID == game.HomeTeam.ID {
			leanHomeShots = append(leanHomeShots, publicShot)
		} else {
			leanAwayShots = append(leanAwayShots, publicShot)
		}
	}
	jsonResp, err := json.Marshal(models.GameResp{
		Game:      game,
		HomeShots: leanHomeShots,
		AwayShots: leanAwayShots,
	})
	return jsonResp, err
}
