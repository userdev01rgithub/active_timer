package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/userdev01rgithub/active_timer/internal/log"
)

type Database struct {
	DB     *sql.DB
	logger *log.Logger
}

func NewDatabase(dataSourceName string, logger *log.Logger) *Database {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		logger.Info("Failed to open database: " + err.Error())
	}

	createSessionsTableSQL := `CREATE TABLE IF NOT EXISTS sessions (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
		"duration" INTEGER
	);`

	_, err = db.Exec(createSessionsTableSQL)
	if err != nil {
		logger.Info("Failed to create sessions table: " + err.Error())
	}

	createActiveWindowsTableSQL := `CREATE TABLE IF NOT EXISTS active_windows (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"session_id" INTEGER,
		"window_title" TEXT,
		"timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createActiveWindowsTableSQL)
	if err != nil {
		logger.Info("Failed to create active_windows table: " + err.Error())
	}

	return &Database{DB: db, logger: logger}
}

func (database *Database) Close() {
	database.DB.Close()
}
