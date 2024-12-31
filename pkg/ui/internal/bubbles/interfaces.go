package bubbles

import (
	"context"

	"github.com/ft-t/browser-switcher/pkg/config"
)

type Launcher interface {
	Launch(_ context.Context, browser *config.Browser) error
}
