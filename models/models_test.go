package models_test

import (
	// "github.com/alex1sz/shotcharter-go/db"
	"github.com/alex1sz/shotcharter-go/models"
	"github.com/alex1sz/shotcharter-go/test/helpers/rand"
	"github.com/alex1sz/shotcharter-go/test/helpers/test_helper"
	// "log"
	"testing"
)

// test helper for checking presence of attribute
func isPresent(attribute interface{}) bool {
	if attribute != nil {
		return true
	}
	return false
}

func TestTeamCreate(t *testing.T) {
	team := models.Team{Name: rand.String(10)}
	err := team.Create()

	if !isPresent(team.ID) {
		t.Error("team Create() failed. Expected team.ID to be present")
	}
	if err != nil {
		t.Error("Team Create() returns error")
	}
}

func TestPlayerCreate(t *testing.T) {
	team := test_helper.CreateTestTeam()

	player := models.Player{Name: "Alejandro Alejandro", Active: true, JerseyNumber: 24, Team: team}
	err := player.Create()

	if !isPresent(player.ID) {
		t.Error("Player Create() failed. Expected player.ID to be present")
	}
	if err != nil {
		t.Error("Player Create() returns error")
	}
}

func TestGameCreate(t *testing.T) {
	home_team := test_helper.CreateTestTeam()
	away_team := test_helper.CreateTestTeam()

	game := models.Game{HomeTeam: home_team, AwayTeam: away_team}
	err := game.Create()

	if !isPresent(game.ID) {
		t.Error("Game not created: game.ID not present")
	}

	if err != nil {
		t.Error("game Create() returns error")
	}
}

func TestShotCreate(t *testing.T) {
	player := test_helper.CreateTestPlayer()
	game := test_helper.CreateTestGameForHomeTeam(player.Team)

	shot := models.Shot{Player: player, Game: game, Team: player.Team, PtValue: 3, Made: true, XAxis: 312, YAxis: 250}
	err := shot.Create()

	if !isPresent(shot.ID) {
		t.Error("Shot Create() failed: shot.ID not present")
	}

	if err != nil {
		t.Error("shot Create() returns err")
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

func TestGameIsValid(t *testing.T) {
	team := test_helper.CreateTestTeam()
	game := models.Game{HomeTeam: team, AwayTeam: team}

	gameValidBool, err := game.IsValid()

	if gameValidBool {
		t.Error("game.IsValid() failed! expected bool to be false")
	}
	if err == nil {
		t.Error("Expected: 'Invalid game HomeTeam.ID cannot be AwayTeam.ID'")
	}
}
