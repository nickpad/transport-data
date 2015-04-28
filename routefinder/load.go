package routefinder

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type RouteDatabase struct {
	db *sql.DB
}

func NewRouteDatabase() *RouteDatabase {
	db, err := sql.Open("sqlite3", "data/nsw.db")

	if err != nil {
		log.Fatal(err)
	}

	return &RouteDatabase{db}
}
