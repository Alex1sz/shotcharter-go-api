package models

import (
	"errors"
	"github.com/alex1sz/shotcharter-go-api/db"
	// "log"
)

type Player struct {
	ID           string `db:"id" json:"id"`
	Name         string `db:"full_name" json:"name,omitempty"`
	Active       bool   `db:"active" json:"active,omitempty"`
	JerseyNumber uint8  `db:"jersey_number" json:"jersey_number,omitempty,string"`
	Team         Team   `db:"team" json:"team,omitempty"`

	CreatedAt string `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt string `db:"updated_at" json:"updated_at,omitempty"`
}

func (player *Player) Create() (err error) {
	if &player.Team == nil {
		return errors.New("Team not found")
	}
	playerValidBool, err := player.isValid()

	if !playerValidBool || err != nil {
		return errors.New("Invalid player: team not found")
	}

	return db.Db.QueryRow("INSERT INTO players (name, active, jersey_number, team_id) VALUES ($1, $2, $3, $4) RETURNING id", player.Name, player.Active, player.JerseyNumber, player.Team.ID).Scan(&player.ID)
}

func (player Player) isValid() (bool, error) {
	teamExistsBool, err := RowExists("SELECT 1 FROM teams WHERE id=$1", player.Team.ID)
	return teamExistsBool, err
}
