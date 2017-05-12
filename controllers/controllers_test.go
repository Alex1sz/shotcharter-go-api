package controllers_test

import (
	"encoding/json"
	"fmt"
	"github.com/alex1sz/shotcharter-go/models"
	"github.com/alex1sz/shotcharter-go/routers"
	//"github.com/alex1sz/shotcharter-go/test/helpers/rand"
	"github.com/alex1sz/shotcharter-go/test/helpers/test_helper"
	"io"
	// "log"
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
	if response.StatusCode != 200 {
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

	if response.StatusCode != 200 {
		t.Errorf("Success Expected: %d", response.StatusCode)
	}
}
