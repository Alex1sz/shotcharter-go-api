package db

import (
	"github.com/alex1sz/configor"
	"github.com/alex1sz/shotcharter-go-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var Db *sqlx.DB

var schema = `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS teams (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  name text,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT teams_pkey PRIMARY KEY (id)
);

DROP TRIGGER IF EXISTS update_teams_updated_at ON teams;

CREATE TRIGGER update_teams_updated_at BEFORE UPDATE ON teams FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE IF NOT EXISTS players (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  name text,
  jersey_number integer,
  team_id uuid NOT NULL,
  active boolean NOT NULL DEFAULT true,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT players_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS index_players_on_team_id
  ON players USING btree
  (team_id);

DROP TRIGGER IF EXISTS update_players_updated_at ON players;

CREATE TRIGGER update_players_updated_at BEFORE UPDATE ON players FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE IF NOT EXISTS games (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  start_at timestamp without time zone,
  home_team_id uuid,
  away_team_id uuid,
  home_score integer NOT NULL DEFAULT 0,
  away_score integer NOT NULL DEFAULT 0,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT games_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS index_games_on_home_team_id
  ON games USING btree
  (home_team_id);

CREATE INDEX IF NOT EXISTS index_games_on_away_team_id
  ON games USING btree
  (away_team_id);

DROP TRIGGER IF EXISTS update_games_updated_at ON games;

CREATE TRIGGER update_games_updated_at BEFORE UPDATE ON games FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE IF NOT EXISTS shots (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  player_id uuid,
  game_id uuid,
  team_id uuid,
  pt_value integer NOT NULL DEFAULT 0,
  made boolean NOT NULL DEFAULT false,
  x_axis integer,
  y_axis integer,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT shots_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS index_shots_on_player_id
  ON shots USING btree
  (player_id);

CREATE INDEX IF NOT EXISTS index_shots_on_game_id
  ON shots USING btree
  (game_id);

CREATE INDEX IF NOT EXISTS index_shots_on_team_id
  ON shots USING btree
  (team_id);

DROP TRIGGER IF EXISTS update_shots_updated_at ON shots;

CREATE TRIGGER update_shots_updated_at BEFORE UPDATE ON shots FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
`

func schemaSetup() (err error) {
	Db.MustExec(schema)
	return
}

func init() {
	var appConfig config.Config
	configor.Load(&appConfig, "../config/db_conf.yml")
	log.Println(appConfig)
	//
	Db = sqlx.MustConnect(appConfig.Db.Driver, appConfig.Db.Connection)
	// Db = sqlx.MustConnect("postgres", "dbname=shotcharter_go_development host=localhost sslmode=disable")
	err := schemaSetup()

	if err != nil {
		log.Println(err)
	}
	// sanity check values before deploying production
	Db.SetMaxIdleConns(4)
	Db.SetMaxOpenConns(16)
}
