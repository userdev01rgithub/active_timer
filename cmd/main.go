package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/userdev01rgithub/active_timer/internal/app"
	"github.com/userdev01rgithub/active_timer/internal/db"
	"github.com/userdev01rgithub/active_timer/internal/log"
	"github.com/userdev01rgithub/active_timer/internal/session"
)

func main() {
	logger := log.NewLogger()
	defer logger.Close()

	logger.Info("Application starting")

	database := db.NewDatabase("app.db", logger)
	defer database.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		session.StopSession(logger, database)
		logger.Info("Application stopping")
		os.Exit(0)
	}()

	app.Run(logger, database)

	session.StopSession(logger, database)
	logger.Info("Application stopping")
}
