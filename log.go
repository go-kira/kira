package kira

import (
	"io"
	"os"

	"github.com/go-kira/kog"
	"github.com/go-kira/kon"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func setupLogger(config *kon.Kon) *kog.Logger {
	logger := kog.New(setupWriter(config), setupFormatter())
	logger.SetLevel(kog.LevelStrings[config.GetString("LOG_LEVEL")])

	return logger
}

func setupWriter(config *kon.Kon) io.Writer {
	switch config.GetString("LOG") {
	case "stderr":
		return os.Stderr
	case "stdin":
		return os.Stdin
	case "stdout":
		return os.Stdout
	case "file":
		return logToFile(config)
	}

	return os.Stderr
}

// setupFormatter to setup the logger formatter.
func setupFormatter() kog.Formatter {
	// TODO
	// - Add color formatter
	return kog.NewDefaultFormatter()
}

// LoggerToFile - make evrey log in log file
// append log to this destination file: storage/logs/year/month/day/logs.log
func logToFile(config *kon.Kon) io.Writer {
	// TODO
	// Rotate file log
	// set a max size of log file
	// when the file rish the limit, create new one.

	return &lumberjack.Logger{
		Filename:   config.GetString("LOG_FILE"),
		MaxSize:    config.GetInt("LOG_FILE_MAX_SIZE"),
		MaxBackups: config.GetInt("LOG_FILE_MAX_BACKUPS"),
		MaxAge:     config.GetInt("LOG_FILE_MAX_AGE"),
		Compress:   config.GetBool("LOG_FILE_MAX_COMPRESS"),
	}
}
