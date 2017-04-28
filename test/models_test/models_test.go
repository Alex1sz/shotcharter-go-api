package models_test

import (
	"github.com/alex1sz/shotcharter-go/db"
	"github.com/alex1sz/shotcharter-go/models"
	"github.com/alex1sz/shotcharter-go/test/helpers/test_helper"
	"log"
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

func TestFindTeamByID(t *testing.T) {
	team := test_helper.CreateTestTeam()

	returnedTeam, err := models.FindTeamByID(team.ID)

	if len(returnedTeam.ID) < 1 {
		t.Error("FindTeamByID failed to return team")
	}

	if err != nil {
		t.Error("FindTeamByID returns err!")
	}
}

func TestGameFindByIDReturnsGame(t *testing.T) {
	game := test_helper.CreateTestGame()

	var returnedGame, err = models.FindGameByID(game.ID)

	if len(returnedGame.ID) < 1 {
		t.Error("FindGameByID failed to return valid game")
	}

	if err != nil {
		log.Println(err)
		t.Error("FindGameByID returns err along with game")
	}
}

func TestGameFindByIDReturnsGameWithAssociatedTeams(t *testing.T) {
	game := test_helper.CreateTestGame()

	if len(game.HomeTeam.ID) < 1 {
		t.Error("CreateTestGame helper failed ")
	}
	var returnedGame, err = models.FindGameByID(game.ID)

	if returnedGame.HomeTeam.ID != game.HomeTeam.ID {
		log.Println(returnedGame.HomeTeam.ID)
		t.Error("FindGameByID failed!")
	}

	if err != nil {
		t.Error(err)
	}
}

func testRetrieveTeams(t *testing.T) {
	game := test_helper.CreateTestGame()
	var teamH, teamA models.Team = models.RetrieveTeams(game.HomeTeam.ID, game.AwayTeam.ID)

	if teamH.ID != game.HomeTeam.ID {
		t.Error("retrieveTeams fails")
	}

	if teamA.ID != game.HomeTeam.ID {
		t.Error("retreiveTeams fails ID's mismatched")
	}
}

func testGetTeamPlayers(t *testing.T) {
	team := test_helper.CreateTestTeam()
	test_helper.CreateTestPlayerForTeam(&team)

	players := models.GetTeamPlayers(team)

	if len(players) != 1 {
		t.Error("GetTeamPlayers failed len of returned players is off")
	}

	if team.Players[0].ID != team.Players[0].ID {
		t.Error("GetTeamPlayers failed")
	}
}
