package server

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"bbs-test-webapp/util"
)

var DB *sql.DB

func InitDB() {
	dbURL := util.Getenv("DATABASE_URL", "")
	_db, dbErr := sql.Open("postgres", dbURL)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	DB = _db;
}