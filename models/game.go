package models

import (
	// database/sql import needed to use sql.NullString for sqlx Scan() functionality
	"database/sql"
	"github.com/alex1sz/shotcharter-go/db"
	"log"
	// "time"
)

type Game struct {
	ID        string         `db:"id" json:"id"`
	StartAt   sql.NullString `db:"start_at" json:"start_at,omitempty"`
	HomeScore uint64         `db:"home_score" json:"home_score"`
	AwayScore uint64         `db:"away_score" json:"away_score"`
	CreatedAt string         `db:"created_at" json:"created_at"`
	UpdatedAt string         `db:"updated_at" json:"updated_at"`
	HomeTeam  *Team          `db:"home_team_id" json:"home_team"`
	AwayTeam  *Team          `db:"away_team_id" json:"away_team"`
}

func (game *Game) Create() (err error) {
	err = db.Db.QueryRow("insert into games (home_team_id, away_team_id) values ($1, $2) returning id", game.HomeTeam.ID, game.AwayTeam.ID).Scan(&game.ID)
	return
}

func RetrieveTeams(homeTeamID string, awayTeamID string) (homeTeam Team, awayTeam Team) {
	teams := []Team{}
	db.Db.Select(teams, "SELECT * from teams where id in ($1, $2)", homeTeamID, awayTeamID)

	return teams[0], teams[1]
}

func FindGameByID(id string) (game Game, err error) {
	err = db.Db.Get(&game, "SELECT * from games where id = $1", id)

	if err != nil {
		log.Println(err)
		return
	}
	teams := []Team{}
	db.Db.Select(teams, "SELECT * from teams where id in ($1, $2)", game.HomeTeam.ID, game.AwayTeam.ID)

	var homeTeam, awayTeam Team = teams[0], teams[1]
	game.HomeTeam = &homeTeam
	game.AwayTeam = &awayTeam

	return game, err
}
