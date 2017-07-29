package controllers

import (
	"encoding/json"
	"github.com/alex1sz/shotcharter-go-api/models"
	"github.com/alex1sz/shotcharter-go-api/utilities"
	"github.com/gorilla/mux"
	"net/http"
)

// game request decoder, game specific to avoid performance hit of using reflect
func decodeReqIntoGame(w http.ResponseWriter, req *http.Request) (game models.Game) {
	if err := json.NewDecoder(req.Body).Decode(&game); err != nil {
		utils.RespondWithAppError(w, err, "Invalid game data in request", 500)
		return
	}
	return
}

// GET /games/:id
func GetGameByID(w http.ResponseWriter, req *http.Request) {
	game, err := models.FindGameByID(mux.Vars(req)["id"])

	if err != nil {
		utils.HandleFindError(w, err)
		return
	}
	respondWithLeanGameJSON(&game, w)
}

// POST /games
func CreateGame(w http.ResponseWriter, req *http.Request) {
	game := decodeReqIntoGame(w, req)

	if err := game.Create(); err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	respondWithLeanGameJSON(&game, w)
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

// respondWithLeanJSON takes a game, marshals to lean JSON and responds
func respondWithLeanGameJSON(game *models.Game, w http.ResponseWriter) {
	json, err := marshalLeanGameJSON(game)

	if err != nil {
		utils.RespondWithAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	utils.SetHeaders(w, 200)
	w.Write(json)
}
