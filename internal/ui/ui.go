package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/userdev01rgithub/active_timer/internal/db"
	"github.com/userdev01rgithub/active_timer/internal/log"
	"github.com/userdev01rgithub/active_timer/internal/session"
)

var sessionLabel *widget.Label

func SetupUI(w fyne.Window, logger *log.Logger, database *db.Database) {
	label := widget.NewLabel("Добро пожаловать!")
	sessionLabel = widget.NewLabel("Время с начала сессии: 0 минут")
	startButton := widget.NewButton("Начать сессию", func() {
		session.StartSession(logger, database)
		go updateSessionTime()
		widget.NewPopUp(widget.NewLabel("Сессия начата"), w.Canvas()).Show()
	})
	stopButton := widget.NewButton("Завершить сессию", func() {
		session.StopSession(logger, database)
	})

	content := container.NewVBox(
		label,
		sessionLabel,
		startButton,
		stopButton,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 300))
}

func updateSessionTime() {
	for session.IsSessionStarted() {
		duration := time.Since(session.GetStartTime()).Minutes()
		sessionLabel.SetText("Время с начала сессии: " + fmt.Sprintf("%.0f", duration) + " минут")
		time.Sleep(1 * time.Minute)
	}
}
