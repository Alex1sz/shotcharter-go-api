package models

import (
	// database/sql import needed to use sql.NullString for sqlx Scan() functionality
	// "database/sql"
	"github.com/alex1sz/shotcharter-go/db"
	"log"
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

	HomeTeam Team `db:"home_team" json:"home_team"`
	AwayTeam Team `db:"away_team" json:"away_team"`
}

func (game *Game) Create() (err error) {
	err = db.Db.QueryRow("insert into games (home_team_id, away_team_id) values ($1, $2) returning id", game.HomeTeam.ID, game.AwayTeam.ID).Scan(&game.ID)
	return
}

func FindGameByID(id string) (game Game, err error) {
	err = db.Db.Get(&game, `SELECT games.id as id, games.home_score, games.away_score, games.home_team_id "home_team.id", home_team.name "home_team.name", games.away_team_id "away_team.id", away_team.name "away_team.name" FROM games INNER JOIN teams AS home_team ON (games.home_team_id = home_team.id) INNER JOIN teams AS away_team ON (games.away_team_id = away_team.id) WHERE games.id = $1`, id)

	if err != nil {
		log.Println(err)
		return
	}
	return
}
