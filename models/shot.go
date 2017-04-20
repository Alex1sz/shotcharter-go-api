package models

type Shot struct {
	ID string
	// PlayerID string
	Player  *Player
	GameID  string
	PtValue uint8
	Made    bool
	XAxis   int
	YAxis   int
}
