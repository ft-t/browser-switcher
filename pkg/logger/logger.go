package logger

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func GetLogger() zerolog.Logger {
	currentDir := filepath.Join(filepath.Dir(os.Args[0]), "logs")

	if runtime.GOOS != "windows" {
		currentDir = "/tmp/"
	}

	logFile := &lumberjack.Logger{
		Filename:   filepath.Join(currentDir, "browser-switcher.log"),
		MaxSize:    30,
		MaxBackups: 3,
		MaxAge:     10,
		Compress:   false,
	}
	lg := zerolog.New(zerolog.MultiLevelWriter(os.Stdout, logFile)).With().Timestamp().Logger()
	log.Logger = lg

	return lg
}
