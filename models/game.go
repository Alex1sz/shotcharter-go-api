package models

import (
	"github.com/alex1sz/shotcharter-go/db"
	"log"
	"time"
)

type Game struct {
	ID        string    `db:"id" json:"id"`
	StartAt   time.Time `db:"start_at" json:"start_at,omitempty"`
	HomeTeam  *Team     `json:"home_team,omitempty"`
	AwayTeam  *Team     `json:"away_team, omitempty"`
	HomeScore uint8     `db:"home_score" json:"home_score"`
	AwayScore uint8     `db:"away_score" json:"away_score"`
	CreatedAt string    `db:"created_at" json:"created_at"`
	UpdatedAt string    `db:"updated_at" json:"updated_at"`
}

func (game *Game) Create() (err error) {
	err = db.Db.QueryRow("insert into games (home_team_id, away_team_id) values ($1, $2) returning id", game.HomeTeam.ID, game.AwayTeam.ID).Scan(&game.ID)
	return
}

func FindGameByID(id string) (game Game, err error) {
	game = Game{}

	// err = db.Db.QueryRowx("select (*) from games where id = $1", id).Scan(&game.ID, &game.StartAt, &game.HomeTeam, &game.AwayTeam, &game.HomeScore, &game.AwayScore, &game.CreatedAt, &game.UpdatedAt)
	err = db.Db.Get(&game, "select (*) from games where id = $1", id)

	if err != nil {
		log.Println(err)
		return
	}

	log.Print(game)
	log.Println(game)

	return game, err
}
