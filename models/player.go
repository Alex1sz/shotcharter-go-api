package models

type Player struct {
	ID           string
	Name         string
	Active       bool
	JerseyNumber uint8
	// TeamID    string
	Team  *Team
	Shots []Shot
}
