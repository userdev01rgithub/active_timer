package main

import (
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/userdev01rgithub/active_timer/internal/app"
	"github.com/userdev01rgithub/active_timer/internal/db"
	"github.com/userdev01rgithub/active_timer/internal/log"
)

func main() {
	logger := log.NewLogger()
	defer logger.Close()

	logger.Info("Application starting")

	database := db.NewDatabase("app.db", logger)
	defer database.Close()

	app.Run(logger)

	logger.Info("Application stopping")
}
