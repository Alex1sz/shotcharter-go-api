package models

import (
	"github.com/alex1sz/shotcharter-go/db"
	"log"
)

type Team struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	// CreatedAt string `db:"created_at" json:"created_at"`
	// UpdatedAt string `db:"updated_at" json:"updated_at"`
	Players []*Player //`json:"players,omitempty"`
	Games   []Game    `json:"games,omitempty"`
}

func (team *Team) Create() (err error) {
	err = db.Db.QueryRow("insert into teams (name) values ($1) returning id", team.Name).Scan(&team.ID)

	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (team *Team) GetPlayers() {
	players := []*Player{}
	db.Db.Select(&players, "SELECT id, name AS full_name, active, jersey_number, created_at, updated_at from players where team_id = $1", &team.ID)

	team.Players = players
	return
}

func FindTeamByID(id string) (team Team, err error) {
	err = db.Db.Get(&team, "select id, name from teams where id = $1", id)

	if err != nil {
		log.Println(err)
		return team, err
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
