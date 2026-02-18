package infra

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB
}

func (db *DB) Connect2() (*gorm.DB, error) {
	conn, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_DSN")), )

	if err == nil {
		db.conn = conn
	}

	return conn, err
}

var Database = DB{}
