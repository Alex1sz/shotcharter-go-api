package models_test

import (
	"github.com/alex1sz/shotcharter-go/db"
	"github.com/alex1sz/shotcharter-go/models"
	"github.com/alex1sz/shotcharter-go/test/helpers/test_helper"
	"testing"
)

func TestSetupBeforeAndAfterCountsHelper(t *testing.T) {
	var pre_create_count, after_create_count, sql = test_helper.SetupBeforeAndAfterCounts("games")

	if after_create_count != 0 {
		t.Error("setupCountVariables failed, after_create_count expected to be 0")
	}

	if pre_create_count < 1 {
		t.Error("No games created!")
	}

	if sql != "SELECT COUNT(*) from games" {
		t.Error("setupCountVariables failed wrong sql query")
	}
}

func TestTeamCreate(t *testing.T) {
	var pre_create_count, after_create_count, sql = test_helper.SetupBeforeAndAfterCounts("teams")

	test_helper.CreateTestTeam()

	db.Db.Get(after_create_count, sql)

	if after_create_count > pre_create_count {
		t.Error("Team create failed!")
	}
}

func TestPlayerCreate(t *testing.T) {
	var pre_create_count, after_create_count, sql = test_helper.SetupBeforeAndAfterCounts("players")

	team := test_helper.CreateTestTeam()

	player := models.Player{Name: "Alejandro Alejandro", Active: true, JerseyNumber: 24, Team: &team}
	player.Create()

	db.Db.Get(after_create_count, sql)

	if after_create_count > pre_create_count {
		t.Error("Player create failed!")
	}
}

func TestGameCreate(t *testing.T) {
	var pre_create_count, after_create_count, sql = test_helper.SetupBeforeAndAfterCounts("games")

	home_team := test_helper.CreateTestTeam()
	away_team := test_helper.CreateTestTeam()

	game := models.Game{HomeTeam: &home_team, AwayTeam: &away_team}
	game.Create()

	db.Db.Get(after_create_count, sql)

	if after_create_count > pre_create_count {
		t.Error("Game create failed!")
	}
}

// helper method creates test player w/ team
// func createTestPlayer() (player models.Player) {
// 	team := createTestTeam()
//
// 	player = models.Player{Name: "Some player name", Active: true, JerseyNumber: 23, Team: &team}
// 	player.Create()
//
// 	return player
// }

// helper method creates game w/ away team
// func createTestGameForHomeTeam(homeTeam *models.Team) (game models.Game) {
// 	away_team := createTestTeam()
// 	game = models.Game{HomeTeam: homeTeam, AwayTeam: &away_team}
//
// 	return game
// }

func TestShotCreate(t *testing.T) {
	var pre_create_count, after_create_count, sql = test_helper.SetupBeforeAndAfterCounts("shots")

	player := test_helper.CreateTestPlayer()
	game := test_helper.CreateTestGameForHomeTeam(player.Team)

	shot := models.Shot{Player: &player, Game: &game, PtValue: 3, Made: true, XAxis: 312, YAxis: 250}
	shot.Create()

	db.Db.Get(after_create_count, sql)

	if after_create_count > pre_create_count {
		t.Error("Shot not created!")
	}
}
