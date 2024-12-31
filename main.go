package main

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"

	config2 "github.com/ft-t/browser-switcher/pkg/config"
	"github.com/ft-t/browser-switcher/pkg/launcher"
	"github.com/ft-t/browser-switcher/pkg/selector"
	"github.com/ft-t/browser-switcher/pkg/ui"
)

func main() {
	logFile := &lumberjack.Logger{
		Filename:   "logs/log.log",
		MaxSize:    30,
		MaxBackups: 3,
		MaxAge:     10,
		Compress:   false,
	}
	lg := zerolog.New(zerolog.MultiLevelWriter(os.Stdout, logFile)).With().Timestamp().Logger()
	log.Logger = lg

	ctx := lg.WithContext(context.Background())

	if len(os.Args) < 2 {
		lg.Panic().Msg("No arguments provided")
		return
	}

	targetURL := os.Args[1]

	ctx = lg.With().Str("targetURL", targetURL).Logger().WithContext(ctx)

	browserConfig, err := config2.ReadConfig(ctx)

	if err != nil {
		lg.Err(err).Msg("Failed to read config")
		return
	}

	browserSelector := selector.New(browserConfig)

	targetBrowser := browserSelector.SelectBrowser(ctx, targetURL)

	browserLauncher := launcher.New(targetURL)
	uiRenderer, err := ui.NewUi(browserConfig, browserLauncher)
	if err != nil {
		lg.Panic().Err(err).Msg("Failed to create UI")
		return
	}

	if targetBrowser == nil {
		if err = uiRenderer.ShowManualSelect(ctx); err != nil {
			lg.Panic().Err(err).Msg("Failed to show manual browser selection")
			return
		}
	}
}
