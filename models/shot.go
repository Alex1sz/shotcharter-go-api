package models

import (
	"github.com/alex1sz/shotcharter-go-api/db"
	// "log"
)

type Shot struct {
	ID        string `db:"id" json:"id"`
	Player    Player `db:"player" json:"player,omitempty"`
	Game      Game   `db:"game" json:"-"`
	Team      Team   `db:"team" json:"team,omitempty"`
	PtValue   uint64 `db:"pt_value" json:"pt_value"`
	Made      bool   `db:"made" json:"made"`
	XAxis     uint64 `db:"x_axis" json:"x_axis"`
	YAxis     uint64 `db:"y_axis" json:"y_axis"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func (shot *Shot) Create() (err error) {
	err = db.Db.QueryRow("INSERT INTO shots (player_id, game_id, team_id, pt_value, made, x_axis, y_axis) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", shot.Player.ID, shot.Game.ID, shot.Team.ID, shot.PtValue, shot.Made, shot.XAxis, shot.YAxis).Scan(&shot.ID)

	return
}

func (shot Shot) IsValid() (playerIsOnTeam bool, err error) {
	playerIsOnTeam, err = RowExists("SELECT 1 FROM players WHERE id=$1 AND team_id=$2", shot.Player.ID, shot.Team.ID)

	if err != nil || !playerIsOnTeam {
		return
	}
	return true, err
}
