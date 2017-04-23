package models

import (
	"github.com/alex1sz/shotcharter-go/db"
	"time"
)

type Game struct {
	ID        string
	StartAt   time.Time
	HomeTeam  *Team
	AwayTeam  *Team
	HomeScore uint8
	AwayScore uint8
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (game *Game) Create() (err error) {
	err = db.Db.QueryRow("insert into games (home_team_id, away_team_id) values ($1, $2) returning id", game.HomeTeam.ID, game.AwayTeam.ID).Scan(&game.ID)

	return
}
