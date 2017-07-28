package models

import (
	"errors"
	"github.com/alex1sz/shotcharter-go-api/db"
)

type Team struct {
	ID      string    `db:"id" json:"id"`
	Name    string    `db:"name" json:"name,omitempty"`
	Players []*Player `json:"players,omitempty"`
	Games   []Game    `json:"games,omitempty"`
}

// for lean team json without redundant nested types
type TeamResp struct {
	*Team
	Players []LeanPlayer `json:"players,omitempty"`
}

func (team *Team) Create() error {
	if !team.IsValid() {
		return errors.New("Team is invalid: name must be present")
	}
	return db.Db.QueryRow("INSERT INTO teams (name) VALUES ($1) RETURNING id", team.Name).Scan(&team.ID)
}

func (team *Team) Update() (err error) {
	teamExistsBool, err := RowExists("SELECT 1 from teams WHERE id=$1", team.ID)

	if !teamExistsBool || err != nil {
		return errors.New("resource not found")
	}
	_, err = db.Db.Exec(`UPDATE teams SET name = ($1) WHERE id = ($2)`, team.Name, team.ID)
	return
}

func (team *Team) GetPlayers() {
	players := []*Player{}
	db.Db.Select(&players, "SELECT id, name AS full_name, active, jersey_number, created_at, updated_at FROM players WHERE team_id = $1", &team.ID)

	team.Players = players
	return
}

func FindTeamByID(id string) (team Team, err error) {
	if err = db.Db.Get(&team, "SELECT id, name FROM teams WHERE id = $1", id); err != nil {
		return
	}
	team.GetPlayers()
	return
}

func (team Team) PlayerIsOnTeam(player Player) bool {
	for _, player := range team.Players {
		if player.ID == player.ID {
			return true
		}
	}
	return false
}

func (team Team) IsValid() bool {
	if team.Name != "" {
		return true
	}
	return false
}
