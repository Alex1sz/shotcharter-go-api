package test_helper

import (
	"github.com/alex1sz/shotcharter-go/db"
	"github.com/alex1sz/shotcharter-go/models"
	"github.com/alex1sz/shotcharter-go/test/helpers/rand"
	// /"log"
)

// count based test setup helper used by Create() tests
func SetupBeforeAndAfterCounts(table string) (pre_create_count int, after_create_count int, sql_query string) {
	sql_query = "SELECT COUNT(*) from " + table
	db.Db.Get(&pre_create_count, sql_query)

	return pre_create_count, after_create_count, sql_query
}

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
	awayTeam, homeTeam := CreateTestTeam(), CreateTestTeam()

	CreateTestPlayerForTeam(awayTeam)
	CreateTestPlayerForTeam(homeTeam)

	game = models.Game{HomeTeam: homeTeam, AwayTeam: awayTeam}
	game.Create()

	return game
}
