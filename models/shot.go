package models

import (
	"errors"
	"github.com/alex1sz/shotcharter-go-api/db"
)

type Shot struct {
	ID        string `db:"id" json:"id"`
	Player    Player `db:"player" json:"player,omitempty"`
	Game      Game   `db:"game" json:"game"`
	Team      Team   `db:"team" json:"team,omitempty"`
	PtValue   uint8  `db:"pt_value" json:"pt_value"`
	Made      bool   `db:"made" json:"made"`
	XAxis     uint64 `db:"x_axis" json:"x_axis"`
	YAxis     uint64 `db:"y_axis" json:"y_axis"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func (shot *Shot) Create() error {
	if !shot.IsValid() {
		return errors.New("Shot is invalid")
	}
	return db.Db.QueryRowx("INSERT INTO shots (player_id, game_id, team_id, pt_value, made, x_axis, y_axis) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", shot.Player.ID, shot.Game.ID, shot.Team.ID, shot.PtValue, shot.Made, shot.XAxis, shot.YAxis).Scan(&shot.ID)
}

func (shot *Shot) Update() (err error) {
	var result string
	err = db.Db.QueryRowx(`UPDATE shots SET
    player_id = ($2),
    team_id = ($3),
    pt_value = ($4),
    made = ($5),
    x_axis = ($6),
    y_axis = ($7)
    WHERE id = ($1) RETURNING id`, shot.ID, shot.Player.ID, shot.Team.ID, shot.PtValue, shot.Made, shot.XAxis, shot.YAxis).Scan(&result)

	if result != shot.ID && err == nil {
		return errors.New("Shot update failed: result not equal to shot.ID")
	}
	return
}

func (shot Shot) IsValid() (playerIsOnTeam bool) {
	playerIsOnTeam, err := RowExists(`
    SELECT 1 FROM players
    WHERE EXISTS(
      SELECT 1 FROM games
      WHERE players.team_id = games.home_team_id OR players.team_id = games.away_team_id AND games.id = $3)
    AND id=$1 AND team_id=$2`, shot.Player.ID, shot.Team.ID, shot.Game.ID)

	if err != nil || !playerIsOnTeam {
		return false
	}
	return true
}
