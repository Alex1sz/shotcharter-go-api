package models

import (
	"errors"
	"github.com/alex1sz/shotcharter-go-api/db"
	"time"
)

type NullTime struct {
	Time  time.Time `json:"time"`
	Valid bool      `json:"valid"`
}

type Game struct {
	ID        string   `db:"id" json:"id"`
	StartAt   NullTime `db:"start_at" json:"start_at,omitempty"`
	HomeScore uint8    `db:"home_score" json:"home_score"`
	AwayScore uint8    `db:"away_score" json:"away_score"`
	HomeTeam  Team     `db:"home_team" json:"home_team"`
	AwayTeam  Team     `db:"away_team" json:"away_team"`
	HomeShots []*Shot  `json:"home_shots,omitempty"`
	AwayShots []*Shot  `json:"away_shots,omitempty"`
}

type GameResp struct {
	*Game
	HomeShots []PublicShot `json:"home_shots,omitempty"`
	AwayShots []PublicShot `json:"away_shots,omitempty"`
}

func (game *Game) Create() (err error) {
	return db.Db.QueryRow("INSERT INTO games (home_team_id, away_team_id) VALUES ($1, $2) RETURNING id", game.HomeTeam.ID, game.AwayTeam.ID).Scan(&game.ID)
}

func (game *Game) GetShots() {
	shots := []*Shot{}
	db.Db.Select(&shots, `SELECT shots.id AS id, shots.player_id "player.id", shots.game_id "game.id", shots.team_id "team.id", pt_value, made, x_axis, y_axis, shots.created_at "created_at", shots.updated_at "updated_at" FROM shots WHERE shots.game_id = $1`, &game.ID)
	homeShots, awayShots := []*Shot{}, []*Shot{}

	for _, shot := range shots {
		if shot.Team.ID == game.HomeTeam.ID {
			homeShots = append(homeShots, shot)
		} else {
			awayShots = append(awayShots, shot)
		}
	}
	game.HomeShots, game.AwayShots = homeShots, awayShots
	return
}

func FindGameByID(id string) (game Game, err error) {
	if err = db.Db.Get(&game, `SELECT games.id as id, games.home_score, games.away_score, games.home_team_id "home_team.id", home_team.name "home_team.name", games.away_team_id "away_team.id", away_team.name "away_team.name" FROM games INNER JOIN teams AS home_team ON (games.home_team_id = home_team.id) INNER JOIN teams AS away_team ON (games.away_team_id = away_team.id) WHERE games.id = $1`, id); err != nil {
		return
	}
	game.GetShots()
	return
}

func (game Game) IsValid() (bool, error) {
	if game.HomeTeam.ID == game.AwayTeam.ID {
		err := errors.New("Invalid game HomeTeam.ID cannot be AwayTeam.ID")
		return false, err
	}
	return true, nil
}
