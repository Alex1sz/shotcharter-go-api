package controllers_test

import (
	"encoding/json"
	"fmt"
	"github.com/alex1sz/shotcharter-go-api/models"
	"github.com/alex1sz/shotcharter-go-api/routers"
	"github.com/alex1sz/shotcharter-go-api/test/helpers/test_helper"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server     *httptest.Server
	reader     io.Reader
	requestURL string
	serverURL  string
)

func init() {
	server = httptest.NewServer(routers.InitRoutes())
	requestURL = fmt.Sprintf("%s/games", server.URL)
	serverURL = server.URL
}

// abstract out request/response error handling for usage in multiple tests
func MakeRequest(httpVerb string, requestURL string, reader io.Reader) (response *http.Response, err error) {
	request, err := http.NewRequest(httpVerb, requestURL, reader)
	if err != nil {
		return
	}
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	return response, err
}

// POST /games
func TestCreateGame(t *testing.T) {
	var game models.Game
	homeTeam := test_helper.CreateTestTeam()
	awayTeam := test_helper.CreateTestTeam()
	game.HomeTeam, game.AwayTeam = homeTeam, awayTeam

	gameJSON, err := json.Marshal(game)
	if err != nil {
		t.Error(err)
	}
	// convert string to reader
	reader = strings.NewReader(string(gameJSON))

	response, err := MakeRequest("POST", requestURL, reader)
	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 200 {
		t.Errorf("Success expected: %d", response.StatusCode)
	}
}

// GET /games/:id
func TestGetGameByID(t *testing.T) {
	game := test_helper.CreateTestGame()
	gameReqJSON, err := json.Marshal(game)

	if err != nil {
		t.Error(err)
	}
	response, err := MakeRequest("GET", fmt.Sprintf("%s/"+game.ID, requestURL), strings.NewReader(string(gameReqJSON)))

	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 200 {
		t.Errorf("Success Expected: %d", response.StatusCode)
	}
}

// POST /players
func TestCreatePlayer(t *testing.T) {
	team := test_helper.CreateTestTeam()
	player := models.Player{Name: "Test player...", Active: true, JerseyNumber: 23, Team: team}

	requestJSON, err := json.Marshal(player)
	if err != nil {
		t.Error(err)
	}
	response, err := MakeRequest("POST", fmt.Sprintf("%s/players", serverURL), strings.NewReader(string(requestJSON)))

	if err != nil {
		t.Error(err)
	}
	var playerResp models.Player
	json.NewDecoder(response.Body).Decode(&playerResp)

	if len(playerResp.ID) < 1 {
		t.Errorf("Expected player.ID in response, got: %s", playerResp.ID)
	}
	if response.StatusCode != 201 {
		t.Errorf("Success Expected: %d", response.StatusCode)
	}
}

// POST /teams
func TestCreateTeam(t *testing.T) {
	team := models.Team{Name: "Walt D's Mighty Ducks"}
	requestJSON, err := json.Marshal(team)

	if err != nil {
		t.Error(err)
	}
	response, err := MakeRequest("POST", fmt.Sprintf("%s/teams", serverURL), strings.NewReader(string(requestJSON)))

	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 201 {
		t.Errorf("Expected 201, got: %d", response.StatusCode)
	}
}

// POST /teams
func TestCreateTeamWithInvalidTeam(t *testing.T) {
	team := models.Game{}
	requestJSON, err := json.Marshal(team)

	if err != nil {
		t.Error(err)
	}
	response, err := MakeRequest("POST", fmt.Sprintf("%s/teams", serverURL), strings.NewReader(string(requestJSON)))

	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 500 {
		t.Errorf("Expected 500, got: %d", response.StatusCode)
	}
}

// PATCH /teams/:id
func TestTeamUpdate(t *testing.T) {
	team := test_helper.CreateTestTeam()
	team.Name = "Donny Trump's Low T"
	requestJSON, err := json.Marshal(team)

	if err != nil {
		t.Error(err)
	}
	resp, err := MakeRequest("PATCH", fmt.Sprintf("%s/teams/"+team.ID, serverURL), strings.NewReader(string(requestJSON)))

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected 200, got: %d", resp.StatusCode)
	}
	var teamResp models.Team
	err = json.NewDecoder(resp.Body).Decode(&teamResp)

	if teamResp.Name != team.Name {
		t.Errorf("Expected player response name: %s, got %s", team.Name, teamResp.Name)
	}
}

// PATCH /teams/:id
func TestUpdateTeamRespondsWith500(t *testing.T) {
	team := models.Team{ID: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380b11", Name: "Daffy Duck"}
	requestJSON, err := json.Marshal(team)

	if err != nil {
		t.Error("Error marshaling json")
	}
	response, err := MakeRequest("PATCH", fmt.Sprintf("%s/teams/"+team.ID, serverURL), strings.NewReader(string(requestJSON)))

	if response.StatusCode != 404 {
		t.Errorf("Expected 404, got: %d", response.StatusCode)
	}
}

// GET /teams/:id
func TestGetTeamByID(t *testing.T) {
	team := test_helper.CreateTestTeam()
	requestJSON, err := json.Marshal(team)

	if err != nil {
		t.Error(err)
	}
	response, err := MakeRequest("GET", fmt.Sprintf("%s/teams/"+team.ID, serverURL), strings.NewReader(string(requestJSON)))

	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 200 {
		t.Errorf("Success Expected: %d", response.StatusCode)
	}
}

// POST /shots
func TestCreateShot(t *testing.T) {
	player := test_helper.CreateTestPlayer()
	game := models.Game{HomeTeam: player.Team, AwayTeam: test_helper.CreateTestTeam()}
	game.Create()

	shot := models.Shot{
		Player:  models.Player{ID: player.ID},
		Game:    game,
		Team:    models.Team{ID: player.Team.ID},
		PtValue: 3,
		XAxis:   312,
		YAxis:   250}
	requestJSON, err := json.Marshal(shot)

	if err != nil {
		t.Errorf("TestCreateShot returns err: %s", err.Error())
	}
	response, err := MakeRequest("POST", fmt.Sprintf("%s/shots", serverURL), strings.NewReader(string(requestJSON)))

	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 201 {
		t.Errorf("Expected 201 got: %d", response.StatusCode)
	}
}

// GET /games/:id game w/shots
func TestGetGameByIDForGameWithShots(t *testing.T) {
	game := test_helper.CreateTestGameWithShots()
	gameReqJSON, err := json.Marshal(game)

	response, err := MakeRequest("GET", fmt.Sprintf("%s/"+game.ID, requestURL), strings.NewReader(string(gameReqJSON)))

	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 200 {
		t.Errorf("Success Expected: %d", response.StatusCode)
	}
	var gameResp models.Game
	err = json.NewDecoder(response.Body).Decode(&gameResp)

	if len(gameResp.HomeShots) != 1 && len(gameResp.AwayShots) != 1 {
		t.Error("JSON response does not contain game's shots")
	}
}

// PATCH /shots/:id
func TestUpdateShotWithBadID(t *testing.T) {
	shot := models.Shot{
		ID:      "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380b11",
		PtValue: 3,
		Made:    true,
		XAxis:   200,
		YAxis:   300,
		Team: models.Team{
			ID: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380b11",
		},
		Player: models.Player{
			ID: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380b11",
		},
	}
	requestJSON, err := json.Marshal(shot)

	if err != nil {
		t.Error("Error marshaling json")
	}
	response, err := MakeRequest("PATCH", fmt.Sprintf("%s/shots/"+shot.ID, serverURL), strings.NewReader(string(requestJSON)))

	if response.StatusCode != 404 {
		t.Errorf("Expected 404, got: %d", response.StatusCode)
	}
}

// PATCH /shots/:id
// valid shot
func TestShotUpdate(t *testing.T) {
	shot := test_helper.CreateTestShot()
	shot.XAxis = 100
	requestJSON, err := json.Marshal(shot)

	response, err := MakeRequest("PATCH", fmt.Sprintf("%s/shots/"+shot.ID, serverURL), strings.NewReader(string(requestJSON)))

	if err != nil {
		t.Errorf("Expected updated shot, got err: %s", err.Error())
	}
	if response.StatusCode != 200 {
		t.Errorf("Expected 200, got: %d", response.StatusCode)
	}
}

// PATCH /players/:id
func TestPlayerUpdate(t *testing.T) {
	// valid player
	player := test_helper.CreateTestPlayer()
	player.Name = "50 Cent"
	requestJSON, err := json.Marshal(player)

	resp, err := MakeRequest("PATCH", fmt.Sprintf("%s/players/"+player.ID, serverURL), strings.NewReader(string(requestJSON)))

	if err != nil {
		t.Errorf("Expected player to update, got err %s", err.Error())
	}
	if resp.StatusCode != 201 {
		t.Errorf("Expected 201, got %d", resp.StatusCode)
	}
	var playerResp models.Player
	err = json.NewDecoder(resp.Body).Decode(&playerResp)

	if playerResp.Name != player.Name {
		t.Errorf("Expected player response name: %s, got %s", playerResp.Name, player.Name)
	}
}
