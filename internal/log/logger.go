package log

import (
	"log"
	"os"
)

type Logger struct {
	file   *os.File
	logger *log.Logger
}

func NewLogger() *Logger {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger := log.New(file, "", log.LstdFlags)
	return &Logger{file: file, logger: logger}
}

func (l *Logger) Close() {
	if err := l.file.Close(); err != nil {
		log.Fatalf("Failed to close log file: %v", err)
	}
}

func (l *Logger) Info(msg string) {
	l.logger.Println("INFO: " + msg)
}
