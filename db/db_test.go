package db_test

import (
	"github.com/alex1sz/shotcharter-go/db"
	"testing"
)

func TestDBInit(t *testing.T) {
	if db.Db == nil {
		t.Error("DB init failed!")
	}
}
