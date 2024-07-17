package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"github.com/userdev01rgithub/active_timer/internal/db"
	"github.com/userdev01rgithub/active_timer/internal/log"
	"github.com/userdev01rgithub/active_timer/internal/session"
	"github.com/userdev01rgithub/active_timer/internal/ui"
)

func Run(logger *log.Logger, database *db.Database) {
	a := app.New()
	w := a.NewWindow("SysTray")

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}),
			fyne.NewMenuItem("Start Session", func() {
				session.StartSession(logger, database)
			}),
			fyne.NewMenuItem("Stop Session", func() {
				session.StopSession(logger, database)
			}),
			fyne.NewMenuItem("Quit", func() {
				session.StopSession(logger, database)
				a.Quit()
			}),
		)
		desk.SetSystemTrayMenu(m)
		desk.SetSystemTrayIcon(theme.FyneLogo())
	}

	ui.SetupUI(w, logger, database)
	w.SetCloseIntercept(func() {
		w.Hide()
	})

	w.ShowAndRun()
}
