package models

type Game struct {
	ID         string
	HomeTeamID string
	AwayTeamID string
	HomeScore  uint8
	AwayScore  uint8
}
