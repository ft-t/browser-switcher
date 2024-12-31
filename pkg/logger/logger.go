package logger

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func GetLogger() zerolog.Logger {
	currentDir := filepath.Dir(os.Args[0])

	logFile := &lumberjack.Logger{
		Filename:   filepath.Join(currentDir, "logs", "log.log"),
		MaxSize:    30,
		MaxBackups: 3,
		MaxAge:     10,
		Compress:   false,
	}
	lg := zerolog.New(zerolog.MultiLevelWriter(os.Stdout, logFile)).With().Timestamp().Logger()
	log.Logger = lg

	return lg
}
