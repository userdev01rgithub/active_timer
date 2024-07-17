package session

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

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
	logger.Info("Stop session ...")

	updateSessionSQL := `UPDATE sessions SET duration = ? WHERE id = ?`
	_, err := database.DB.Exec(updateSessionSQL, duration, sessionID)
	if err != nil {
		logger.Info("Failed to update session: " + err.Error())
	}

	sessionStarted = false
}

func updateActiveWindow(logger *log.Logger, database *db.Database) {
	for sessionStarted {
		activeWindow, err := getActiveWindowTitle()
		if err != nil {
			logger.Info("Failed to get active window title: " + err.Error())
		}

		insertActiveWindowSQL := `INSERT INTO active_windows (session_id, window_title, timestamp) VALUES (?, ?, ?)`
		_, err = database.DB.Exec(insertActiveWindowSQL, sessionID, activeWindow, time.Now())
		if err != nil {
			logger.Info("Failed to insert active window: " + err.Error())
		}
		time.Sleep(10 * time.Second)
	}
}

func getActiveWindowTitle() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return getActiveWindowWindows()
	case "darwin":
		return getActiveWindowMac()
	case "linux":
		return getActiveWindowLinux()
	default:
		return "", fmt.Errorf("unsupported platform")
	}
}

func getActiveWindowWindows() (string, error) {

	return "TODO: ", nil
}

func getActiveWindowMac() (string, error) {
	return "", nil
}

func getActiveWindowLinux() (string, error) {
	cmd := exec.Command("xdotool", "getactivewindow", "getwindowname")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get active window title: %v", err)
	}

	title := strings.TrimSpace(string(output))
	if title == "" {
		return "", fmt.Errorf("no active window title found")
	}

	return title, nil
}

func GetStartTime() time.Time {
	return startTime
}

func IsSessionStarted() bool {
	return sessionStarted
}
