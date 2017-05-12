package models_test

import (
	"github.com/alex1sz/shotcharter-go/db"
	"github.com/alex1sz/shotcharter-go/models"
	"github.com/alex1sz/shotcharter-go/test/helpers/test_helper"
	// "log"
	"testing"
)

func TestSetupBeforeAndAfterCountsHelper(t *testing.T) {
	var pre_create_count, after_create_count, sql = test_helper.SetupBeforeAndAfterCounts("games")

	if after_create_count != 0 {
		t.Error("setupCountVariables failed, after_create_count expected to be 0" + sql)
	}

	if pre_create_count < 1 {
		t.Error("No games created!")
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

	player := models.Player{Name: "Alejandro Alejandro", Active: true, JerseyNumber: 24, Team: team}
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

	game := models.Game{HomeTeam: home_team, AwayTeam: away_team}
	game.Create()

	db.Db.Get(after_create_count, sql)

	if after_create_count > pre_create_count {
		t.Error("Game create failed!")
	}
}

func TestShotCreate(t *testing.T) {
	var pre_create_count, after_create_count, sql = test_helper.SetupBeforeAndAfterCounts("shots")

	player := test_helper.CreateTestPlayer()
	game := test_helper.CreateTestGameForHomeTeam(player.Team)

	shot := models.Shot{Player: player, Game: game, Team: player.Team, PtValue: 3, Made: true, XAxis: 312, YAxis: 250}
	shot.Create()

	db.Db.Get(after_create_count, sql)

	if after_create_count > pre_create_count {
		t.Error("Shot not created!")
	}
}

func TestFindTeamByID(t *testing.T) {
	team := test_helper.CreateTestTeam()
	test_helper.CreateTestPlayerForTeam(team)

	returnedTeam, err := models.FindTeamByID(team.ID)

	if len(returnedTeam.ID) < 1 {
		t.Error("FindTeamByID failed to return team")
	}

	if len(returnedTeam.Players) < 1 {
		t.Error("FindTeamByID failed to return players")
	}

	if err != nil {
		t.Error("FindTeamByID returns err!")
	}
}

func TestGameFindByID(t *testing.T) {
	game := test_helper.CreateTestGame()
	returnedGame, err := models.FindGameByID(game.ID)

	if len(returnedGame.ID) < 1 {
		t.Error("FindGameByID failed to return valid game")
	}

	if returnedGame.HomeTeam.ID != game.HomeTeam.ID {
		t.Error("FindGameByID failed!")
	}

	if err != nil {
		t.Error(err)
	}
}
