package models

import (
	"errors"
	"github.com/alex1sz/shotcharter-go/db"
	"log"
)

type Player struct {
	ID           string `db:"id"`
	Name         string `db:"name"`
	Active       bool   `db:"active"`
	JerseyNumber uint8  `db:"jersey_number"`
	Team         *Team
	Shots        []Shot
}

func (player *Player) Create() (err error) {
	if player.Team == nil {
		err = errors.New("Team not found")
		return
	}
	err = db.Db.QueryRow("insert into players (name, active, jersey_number, team_id) values ($1, $2, $3, $4) returning id", player.Name, player.Active, player.JerseyNumber, player.Team.ID).Scan(&player.ID)

	if err != nil {
		log.Println(err)
	}
	return
}
