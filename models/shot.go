package models

import (
	"github.com/alex1sz/shotcharter-go/db"
	// "log"
)

type Shot struct {
	ID        string `db:"id" json:"id"`
	Player    Player `db:"player_id" json:"player,omitempty"`
	Game      Game   `db:"game_id" json:"-"`
	Team      Team   `db:"team_id" json:"team"`
	PtValue   uint64 `db:"pt_value" json:"pt_value"`
	Made      bool   `db:"made" json:"made"`
	XAxis     int    `db:"x_axis" json:"x_axis"`
	YAxis     int    `db:"y_axis" json:"y_axis"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func (shot *Shot) Create() (s Shot, err error) {
	err = db.Db.QueryRow("insert into shots (player_id, game_id, team_id, pt_value, made, x_axis, y_axis) values ($1, $2, $3, $4, $5, $6) returning id", shot.Player.ID, shot.Game.ID, shot.Team.ID, shot.PtValue, shot.Made, shot.XAxis, shot.YAxis).Scan(&shot.ID)

	if err != nil {
		return
	}
	return
}

func (shot Shot) IsValid() (playerIsOnTeam bool, err error) {
	playerIsOnTeam, err = RowExists("select 1 from players where id=$1 AND team_id=$2", shot.Player.ID, shot.Team.ID)

	if err != nil || !playerIsOnTeam {
		return
	}
	return true, err
}
