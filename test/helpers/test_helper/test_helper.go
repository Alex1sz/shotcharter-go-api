package test_helper

import (
	"github.com/alex1sz/shotcharter-go-api/models"
	"github.com/alex1sz/shotcharter-go-api/test/helpers/rand"
)

// Create test team for usage in tests
func CreateTestTeam() (team models.Team) {
	team = models.Team{Name: rand.String(10)}
	team.Create()
	return team
}

// helper method creates test player w/ team
func CreateTestPlayer() (player models.Player) {
	team := CreateTestTeam()
	player = models.Player{Name: rand.String(10), Active: true, JerseyNumber: 23, Team: team}
	player.Create()

	return player
}

// helper method creates game w/ away team
func CreateTestGameForHomeTeam(homeTeam models.Team) (game models.Game) {
	away_team := CreateTestTeam()
	game = models.Game{HomeTeam: homeTeam, AwayTeam: away_team}
	game.Create()

	return game
}

// create player associate to test team
func CreateTestPlayerForTeam(team models.Team) {
	player := models.Player{Name: rand.String(10), Active: true, JerseyNumber: 23, Team: team}
	player.Create()
	return
}

// helper creates a game w/ HomeTeam & AwayTeam
func CreateTestGame() (game models.Game) {
	playerTeam1 := CreateTestPlayer()
	playerTeam2 := CreateTestPlayer()

	game = models.Game{HomeTeam: playerTeam1.Team, AwayTeam: playerTeam2.Team}
	game.Create()

	return game
}

// helper creates a game w/ shots
func CreateTestGameWithShots() (game models.Game) {
	playerTeam1 := CreateTestPlayer()
	playerTeam2 := CreateTestPlayer()

	game = models.Game{HomeTeam: playerTeam1.Team, AwayTeam: playerTeam2.Team}
	game.Create()

	shot := models.Shot{Player: playerTeam1, Game: game, Team: playerTeam1.Team, PtValue: 2, Made: true, XAxis: 312, YAxis: 250}
	shot.Create()

	shot2 := models.Shot{Player: playerTeam2, Game: game, Team: playerTeam2.Team, PtValue: 2, Made: true, XAxis: 110, YAxis: 212}
	shot2.Create()

	return game
}

// helper creates test shot
func CreateTestShot() (shot models.Shot) {
	player := CreateTestPlayer()
	shot = models.Shot{Game: CreateTestGameForHomeTeam(player.Team), Team: player.Team, Player: player, PtValue: 2, Made: true, XAxis: 200, YAxis: 300}
	shot.Create()

	return
}
