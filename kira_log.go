package kira

import (
	"io"
	"os"

	"github.com/go-kira/config"
	"github.com/go-kira/log"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func setupLogger(config *config.Config) *log.Logger {
	logger := log.New(
		setupWriter(config),
		setupFormatter(),
	)
	logger.SetLevel(log.LevelStrings[config.GetString("log.level", "info")])

	return logger
}

func setupWriter(config *config.Config) io.Writer {
	switch config.GetString("log.log") {
	case "stderr":
		return os.Stderr
	case "stdin":
		return os.Stdin
	case "stdout":
		return os.Stdout
	case "file":
		return &lumberjack.Logger{
			Filename:   config.GetString("log.file", "./storage/logs/kira.log"),
			MaxSize:    config.GetInt("log.file_max_size", 100),
			MaxBackups: config.GetInt("log.file_max_backups", 3),
			MaxAge:     config.GetInt("log.file_max_age", 28),
			Compress:   config.GetBool("log.file_max_compress", false),
		}
	}

	return os.Stderr
}

// setupFormatter to setup the logger formatter.
func setupFormatter() log.Formatter {
	// TODO
	// - Add color formatter
	return log.NewDefaultFormatter()
}
