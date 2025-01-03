//go:build linux

package launcher

import (
	"context"
	"os/exec"
	"slices"

	"github.com/ft-t/browser-switcher/pkg/config"
)

func (l *Launcher) Launch(ctx context.Context, browser *config.Browser) error {
	return exec.Command(
		"setsid",
		slices.Concat(
			[]string{
				browser.BinaryPath,
				l.targetURL,
			},
			browser.LaunchArgs,
		)...,
	).Start()
}
