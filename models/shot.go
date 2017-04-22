package models

import (
	"github.com/alex1sz/shotcharter-go/db"
)

type Shot struct {
	ID      string `db:"id"`
	Player  *Player
	Game    *Game
	PtValue uint8 `db:"pt_value"`
	Made    bool  `db:"made"`
	XAxis   int   `db:"x_axis"`
	YAxis   int   `db:"y_axis"`
}

func (shot *Shot) Create() (err error) {
	err = db.Db.QueryRow("insert into shots (player_id, game_id, pt_value, made, x_axis, y_axis) values ($1, $2, $3, $4, $5, $6) returning id", shot.Player.ID, shot.Game.ID, shot.PtValue, shot.Made, shot.XAxis, shot.YAxis).Scan(&shot.ID)

	return
}
