package ui

import (
	"context"

	"github.com/cockroachdb/errors"

	"github.com/ft-t/browser-switcher/pkg/config"
	"github.com/ft-t/browser-switcher/pkg/ui/internal/bubbles"
)

type UI interface {
	ShowManualSelect(
		ctx context.Context,
	) error
}

func NewUi(
	cfg *config.Config,
	launcher Launcher,
) (UI, error) {
	switch cfg.UI.Renderer {
	case "bubbles", "":
		return bubbles.NewBubbles(cfg, launcher), nil
	default:
		return nil, errors.Newf("unknown renderer")
	}
}
