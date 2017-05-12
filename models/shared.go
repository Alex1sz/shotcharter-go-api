package models

import (
	"database/sql"
	"fmt"
	"github.com/alex1sz/shotcharter-go/db"
	// "log"
)

func RowExists(query string, args ...interface{}) (exists bool, err error) {
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err = db.Db.QueryRow(query, args...).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, err
}
