package app

import (
	"github.com/userdev01rgithub/active_timer/internal/log"

	"github.com/userdev01rgithub/active_timer/internal/ui"
)

func Run(logger *log.Logger) {
	ui.StartUI(logger)
}
