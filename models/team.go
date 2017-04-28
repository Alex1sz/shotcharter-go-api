package models

import (
	"github.com/alex1sz/shotcharter-go/db"
	"log"
)

type Team struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt string    `db:"created_at" json:"created_at"`
	UpdatedAt string    `db:"updated_at" json:"updated_at"`
	Players   []*Player `db:""`
	Games     []*Game   `db:""`
}

func (team *Team) Create() (err error) {
	err = db.Db.QueryRow("insert into teams (name) values ($1) returning id", team.Name).Scan(&team.ID)

	if err != nil {
		log.Println(err)
		return
	}
	return
}

func GetTeamPlayers(team Team) (Players []*Player) {
	team.Players = []*Player{}

	rows, err := db.Db.Queryx("select id, name, active, jersey_number, created_at, updated_at from players where team_id = $1", team.ID)

	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		player := Player{Team: &team}
		err = rows.Scan(&player.ID, &player.Name, &player.Active, &player.JerseyNumber, &player.CreatedAt, &player.UpdatedAt)

		if err != nil {
			log.Println(err)
			return
		}
		team.Players = append(team.Players, &player)
	}
	rows.Close()
	return
}

func FindTeamByID(id string) (team Team, err error) {
	err = db.Db.Get(&team, "select id, name from teams where id = $1", id)

	if err != nil {
		log.Println(err)
		return
	}
	team.Players = GetTeamPlayers(team)

	return
}
