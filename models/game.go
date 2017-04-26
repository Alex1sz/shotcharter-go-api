package models

import (
	// database/sql import needed to use sql.NullString for sqlx Scan() functionality
	"database/sql"
	"github.com/alex1sz/shotcharter-go/db"
	"log"
)

type Game struct {
	ID        string         `db:"id" json:"id"`
	StartAt   sql.NullString `db:"start_at" json:"start_at,omitempty"`
	HomeScore uint8          `db:"home_score" json:"home_score"`
	AwayScore uint8          `db:"away_score" json:"away_score"`
	CreatedAt string         `db:"created_at" json:"created_at"`
	UpdatedAt string         `db:"updated_at" json:"updated_at"`
	HomeTeam  *Team          `db:"home_team_id" json:"home_team,omitempty"`
	AwayTeam  *Team          `db:"away_team_id" json:"away_team,omitempty"`
}

func (game *Game) Create() (err error) {
	err = db.Db.QueryRow("insert into games (home_team_id, away_team_id) values ($1, $2) returning id", game.HomeTeam.ID, game.AwayTeam.ID).Scan(&game.ID)
	return
}

func FindGameByID(id string) (game Game, err error) {
	err = db.Db.Get(&game, "select * from games where id = $1", id)

	if err != nil {
		log.Println(err)
		return
	}
	return game, err
}
