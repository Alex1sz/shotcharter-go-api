package models_test

import (
	"github.com/alex1sz/shotcharter-go/db"
	"github.com/alex1sz/shotcharter-go/models"
	"testing"
)

func TestTeamCreate(t *testing.T) {
	var team = models.Team{Name: "Test team..."}
	team.Create()

	after_teams := []models.Team{}
	db.Db.Select(&after_teams, "SELECT * FROM teams")

	if len(after_teams) > 0 {
		t.Error("Team count not incremented by 1")
	}
}

func TestPlayerCreate(t *testing.T) {
	team := models.Team{}
	err := db.Db.QueryRow("SELECT id, name FROM teams").Scan(&team.ID, &team.Name)

	if err != nil {
		t.Error("Team not found...")
	}
	sql := "SELECT * FROM players"

	pre_create := []models.Player{}
	db.Db.Select(&pre_create, sql)

	player := models.Player{Name: "Alejandro Alejandro", Active: true, JerseyNumber: 24, Team: &team}
	player.Create()

	after_create := []models.Player{}
	db.Db.Select(&after_create, sql)

	if len(after_create) != len(pre_create)+1 {
		t.Error("Player create failed! Count not incremented.")
	}
}
