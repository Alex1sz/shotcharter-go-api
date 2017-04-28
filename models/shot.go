package models

import (
	"github.com/alex1sz/shotcharter-go/db"
)

type Shot struct {
	ID        string  `db:"id" json:"id"`
	Player    *Player `db:"player_id" json:"player,omitempty"`
	Game      *Game   `db:"game_id" json:"game,omitempty"`
	PtValue   uint64  `db:"pt_value" json:"pt_value"`
	Made      bool    `db:"made" json:"made"`
	XAxis     int     `db:"x_axis" json:"x_axis"`
	YAxis     int     `db:"y_axis" json:"y_axis"`
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedAt string  `db:"updated_at" json:"updated_at"`
}

func (shot *Shot) Create() (err error) {
	err = db.Db.QueryRow("insert into shots (player_id, game_id, pt_value, made, x_axis, y_axis) values ($1, $2, $3, $4, $5, $6) returning id", shot.Player.ID, shot.Game.ID, shot.PtValue, shot.Made, shot.XAxis, shot.YAxis).Scan(&shot.ID)

	return
}
