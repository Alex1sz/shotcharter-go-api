package models

import (
	"errors"
	"github.com/alex1sz/shotcharter-go-api/db"
	// "log"
)

type Team struct {
	ID      string    `db:"id" json:"id"`
	Name    string    `db:"name" json:"name,omitempty"`
	Players []*Player `json:"players,omitempty"`
	Games   []Game    `json:"games,omitempty"`
}

func (team *Team) Create() (err error) {
	err = db.Db.QueryRow("INSERT INTO teams (name) VALUES ($1) RETURNING id", team.Name).Scan(&team.ID)
	return
}

func (team *Team) Update() (err error) {
	teamExistsBool, err := RowExists("SELECT 1 from teams WHERE id=$1", team.ID)

	if !teamExistsBool || err != nil {
		err = errors.New("resource not found")
		return
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
	err = db.Db.Get(&team, "SELECT id, name FROM teams WHERE id = $1", id)

	if err != nil {
		return
	}
	team.GetPlayers()
	return
}

func (team Team) PlayerIsOnTeam(player Player) bool {
	for _, player := range team.Players {
		if player.ID == player.ID {
			// player found
			return true
		}
	}
	return false
}
