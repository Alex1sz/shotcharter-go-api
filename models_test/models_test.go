package models_test

import (
	"github.com/alex1sz/shotcharter-go/db"
	"github.com/alex1sz/shotcharter-go/models"
	"testing"
)

// Create test team for usage in tests
func createTestTeam() (team models.Team) {
	team = models.Team{Name: "Test team..."}
	team.Create()
	return team
}

func TestTeamCreate(t *testing.T) {
	sql := "SELECT COUNT(*) from teams"
	var pre_create_count, after_create_count int

	db.Db.Get(pre_create_count, sql)

	createTestTeam()

	db.Db.Get(after_create_count, sql)

	if after_create_count > pre_create_count {
		t.Error("Team count not incremented by 1")
	}
}

func TestPlayerCreate(t *testing.T) {
	var pre_create_count, after_create_count int
	sql := "SELECT COUNT(*) FROM players"

	db.Db.Get(pre_create_count, sql)

	// create team to associate to player
	team := createTestTeam()

	player := models.Player{Name: "Alejandro Alejandro", Active: true, JerseyNumber: 24, Team: &team}
	player.Create()

	db.Db.Get(after_create_count, sql)

	if after_create_count > pre_create_count {
		t.Error("Player create failed! Count not incremented.")
	}
}
