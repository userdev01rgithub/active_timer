package ui

import (
	"github.com/userdev01rgithub/active_timer/internal/log"

	"github.com/andlabs/ui"
)

func StartUI(logger *log.Logger) {
	err := ui.Main(func() {
		mainWindow := ui.NewWindow("Главный экран", 400, 300, false)
		mainWindow.SetMargined(true)

		box := ui.NewVerticalBox()
		box.SetPadded(true)
		mainWindow.SetChild(box)

		label := ui.NewLabel("Добро пожаловать!")
		box.Append(label, false)

		button := ui.NewButton("Начать сессию")
		button.OnClicked(func(*ui.Button) {
			logger.Info("Open window - start session")
			ui.MsgBox(mainWindow, "Сессия", "Начало сессии")
		})
		box.Append(button, false)

		mainWindow.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			logger.Info("Window closing")
			return true
		})

		mainWindow.Show()
	})

	if err != nil {
		logger.Info("Error starting UI: " + err.Error())
		panic(err)
	}
}

func StartSession(logger *log.Logger) {
	err := ui.Main(func() {
		mainWindow := ui.NewWindow("Главный экран", 400, 300, false)
		mainWindow.SetMargined(true)

		box := ui.NewVerticalBox()
		box.SetPadded(true)
		mainWindow.SetChild(box)

		label := ui.NewLabel("Добро пожаловать!")
		box.Append(label, false)

		button := ui.NewButton("Начать сессию")
		button.OnClicked(func(*ui.Button) {
			logger.Info("Open window - start session")
			ui.MsgBox(mainWindow, "Сессия", "Начало сессии")
		})
		box.Append(button, false)

		mainWindow.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			logger.Info("Window closing")
			return true
		})

		mainWindow.Show()
	})

	if err != nil {
		logger.Info("Error starting UI: " + err.Error())
		panic(err)
	}
}
