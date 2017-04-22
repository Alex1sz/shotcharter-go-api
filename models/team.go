package models

import (
	"github.com/alex1sz/shotcharter-go/db"
	"log"
)

type Team struct {
	ID      string `db:"id"`
	Name    string `db:"name"`
	Players []Player
}

func (team *Team) Create() (err error) {
	err = db.Db.QueryRow("insert into teams (name) values ($1) returning id", team.Name).Scan(&team.ID)

	if err != nil {
		log.Println(err)
	}
	return
}

func FindTeamById(id string) (team Team, err error) {
	team = Team{}

	err = db.Db.QueryRowx("select id, name from teams where id = $1", id).Scan(&team.ID, &team.Name)

	team.Players = []Player{}

	rows, err := db.Db.Queryx("select id, name, active, jersey_number from players where team_id = $1", id)

	if err != nil {
		return
	}

	for rows.Next() {
		player := Player{Team: &team}
		err = rows.Scan(&player.ID, &player.Name, &player.Active, &player.JerseyNumber)

		if err != nil {
			log.Fatalln(err)
		}
		team.Players = append(team.Players, player)
	}
	rows.Close()
	return
}
