package models

import (
	"errors"
	"github.com/alex1sz/shotcharter-go/db"
	// "log"
)

type Player struct {
	ID           string `db:"id" json:"id"`
	Name         string `db:"full_name" json:"name,omitempty"`
	Active       bool   `db:"active" json:"active,omitempty"`
	JerseyNumber int64  `db:"jersey_number" json:"jersey_number,omitempty"`
	Team         Team   `db:"team" json:"team"`
	// Shots        []Shot `db:"shots" json:"shots,omitempty"`
	CreatedAt string `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt string `db:"updated_at" json:"updated_at,omitempty"`
}

func (player *Player) Create() (err error) {
	if &player.Team == nil {
		err = errors.New("Team not found")
		return
	}
	err = db.Db.QueryRow("insert into players (name, active, jersey_number, team_id) values ($1, $2, $3, $4) returning id", player.Name, player.Active, player.JerseyNumber, player.Team.ID).Scan(&player.ID)

	if err != nil {
		return
	}
	return
}

func (player Player) IsValid() (bool, error) {
	teamExistsBool, err := RowExists("select 1 from teams WHERE id=$1", player.Team.ID)
	return teamExistsBool, err
}
