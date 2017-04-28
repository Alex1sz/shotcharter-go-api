package models

import (
	"errors"
	"github.com/alex1sz/shotcharter-go/db"
	"log"
)

type Player struct {
	ID           string `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Active       bool   `db:"active" json:"active,omitempty"`
	JerseyNumber int64  `db:"jersey_number" json:"jersey_number"`
	Team         Team   `db:"team_id" json:"-"`
	// Shots        []Shot `db:"shots" json:"shots,omitempty"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func (player *Player) Create() (p Player, err error) {
	if &player.Team == nil {
		err = errors.New("Team not found")
		return
	}
	err = db.Db.QueryRow("insert into players (name, active, jersey_number, team_id) values ($1, $2, $3, $4) returning id", player.Name, player.Active, player.JerseyNumber, player.Team.ID).Scan(&player.ID)

	if err != nil {
		log.Println(err)
		return
	}
	return
}
