package launcher

import (
	"context"
	"os/exec"
	"slices"

	"github.com/ft-t/browser-switcher/pkg/config"
)

type Launcher struct {
	targetURL string
}

func New(targetURL string) *Launcher {
	return &Launcher{
		targetURL: targetURL,
	}
}

func (l *Launcher) Launch(_ context.Context, browser *config.Browser) error {
	return exec.Command(
		browser.BinaryPath,
		slices.Concat(
			[]string{l.targetURL},
			browser.LaunchArgs,
		)...,
	).Start()
}
