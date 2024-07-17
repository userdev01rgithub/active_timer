package session

import (
	"time"

	"github.com/shirou/gopsutil/process"
	"github.com/userdev01rgithub/active_timer/internal/db"
	"github.com/userdev01rgithub/active_timer/internal/log"
)

var startTime time.Time
var sessionID int64
var sessionStarted bool

func StartSession(logger *log.Logger, database *db.Database) {
	if sessionStarted {
		return
	}
	startTime = time.Now()
	sessionStarted = true

	logger.Info("Start session ...")

	insertSessionSQL := `INSERT INTO sessions (duration) VALUES (?)`
	res, err := database.DB.Exec(insertSessionSQL, 0)
	if err != nil {
		logger.Info("Failed to insert session: " + err.Error())
		return
	}

	sessionID, err = res.LastInsertId()
	if err != nil {
		logger.Info("Failed to get session ID: " + err.Error())
		return
	}

	go updateActiveWindow(logger, database)
}

func StopSession(logger *log.Logger, database *db.Database) {
	if !sessionStarted {
		return
	}
	duration := time.Since(startTime).Seconds()

	updateSessionSQL := `UPDATE sessions SET duration = ? WHERE id = ?`
	_, err := database.DB.Exec(updateSessionSQL, duration, sessionID)
	if err != nil {
		logger.Info("Failed to update session: " + err.Error())
	}
	logger.Info("Stop session ...")

	sessionStarted = false
}

func updateActiveWindow(logger *log.Logger, database *db.Database) {
	for sessionStarted {
		activeWindows := getActiveWindowTitles(logger)
		insertActiveWindowSQL := `INSERT INTO active_windows (session_id, window_title, timestamp) VALUES (?, ?, ?)`
		_, err := database.DB.Exec(insertActiveWindowSQL, sessionID, activeWindows, time.Now())
		if err != nil {
			logger.Info("Failed to insert active window: " + err.Error())
		}
		time.Sleep(10 * time.Second)
	}
}

func getActiveWindowTitles(logger *log.Logger) string {
	processList, err := process.Processes()
	if err != nil {
		logger.Info("Failed to get process list: " + err.Error())
		return ""
	}

	activeWindows := ""
	for _, proc := range processList {
		name, err := proc.Name()
		if err == nil {
			activeWindows += name + "; "
		}
	}
	return activeWindows
}
