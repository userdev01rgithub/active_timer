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

	// Создаем таблицу для сессий, если ее нет
	createTableSQL := `CREATE TABLE IF NOT EXISTS sessions (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		logger.Info("Failed to create table: " + err.Error())
	}

	return &Database{DB: db, logger: logger}
}

func (d *Database) Close() {
	if err := d.DB.Close(); err != nil {
		d.logger.Info("Failed to close database: " + err.Error())
	}
}

func (d *Database) AddSession() error {
	_, err := d.DB.Exec("INSERT INTO sessions (timestamp) VALUES (CURRENT_TIMESTAMP)")
	if err != nil {
		d.logger.Info("Failed to insert session: " + err.Error())
	}
	return err
}
