package models

import (
	"database/sql"
	"github.com/alex1sz/shotcharter-go-api/db"
)

func RowExists(query string, args ...interface{}) (exists bool, err error) {
	err = db.Db.QueryRowx(query, args...).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, err
}
