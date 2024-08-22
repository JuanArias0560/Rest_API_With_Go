package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createdTables(DB)
}

func createdTables(db *sql.DB) {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 name TEXT NOT NULL,
	 description TEXT NOT NULL,
	 location TEXT NOT NULL,
	 dateTime DATETIME NOT NULL,
	 user_id INTEGER	
	)	
	`

	_, err := db.Exec(createEventsTable)

	if err != nil {
		panic("Could not created events table.")
	}
}
