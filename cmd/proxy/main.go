package main

import (
	"context"
	"os"

	"github.com/rs/zerolog"

	logger2 "github.com/ft-t/browser-switcher/pkg/logger"
)

func main() {
	logger := logger2.GetLogger()
	ctx := logger.WithContext(context.Background())

	zerolog.Ctx(ctx).Debug().Msgf("starting proxy with arguments: %v", os.Args)

	if err := run(ctx); err != nil {
		logger.Panic().Err(err).Msg("failed to run")
	}
}
