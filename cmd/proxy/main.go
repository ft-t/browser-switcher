package main

import (
	"context"

	logger2 "github.com/ft-t/browser-switcher/pkg/logger"
)

func main() {
	logger := logger2.GetLogger()
	ctx := logger.WithContext(context.Background())

	if err := run(ctx); err != nil {
		logger.Panic().Err(err).Msg("failed to run")
	}
}
