package db

import (
	//"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

var schema = `
CREATE EXTENSION "uuid-ossp";

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE teams (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  name text,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT teams_pkey PRIMARY KEY (id)
);

CREATE TRIGGER update_teams_updated_at BEFORE UPDATE ON teams FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE players (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  name text,
  jersey_number integer,
  team_id uuid NOT NULL,
  active boolean NOT NULL DEFAULT true,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT players_pkey PRIMARY KEY (id)
);

CREATE INDEX index_players_on_team_id
  ON players USING btree
  (team_id);

CREATE TRIGGER update_players_updated_at BEFORE UPDATE ON players FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE teams_players (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  team_id uuid NOT NULL,
  player_id uuid NOT NULL,
  CONSTRAINT teams_players_pkey PRIMARY KEY (id)
);

CREATE INDEX index_teams_players_on_team_id
  ON teams_players USING btree
  (team_id);

CREATE INDEX index_teams_players_on_player_id
  ON teams_players USING btree
  (player_id);

CREATE TABLE games (
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

CREATE INDEX index_games_on_home_team_id
  ON games USING btree
  (home_team_id);

CREATE INDEX index_games_on_away_team_id
  ON games USING btree
  (away_team_id);

CREATE TRIGGER update_games_updated_at BEFORE UPDATE ON games FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE shots (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  player_id uuid,
  game_id uuid,
  pt_value integer NOT NULL DEFAULT 0,
  made boolean NOT NULL DEFAULT false,
  x_axis integer,
  y_axis integer,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT shots_pkey PRIMARY KEY (id)
);

CREATE INDEX index_shots_on_player_id
  ON shots USING btree
  (player_id);

CREATE INDEX index_shots_on_game_id
  ON shots USING btree
  (game_id);

CREATE TRIGGER update_shots_updated_at BEFORE UPDATE ON shots FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
`

func Init() {
	db, err := sqlx.Open("postgres", "dbname=shotcharter_go_development host=localhost sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	// set to reasonable values for production
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(16)

	// err = db.Ping()
}
