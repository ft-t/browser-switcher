//go:build windows

package launcher

func (l *Launcher) Launch(_ context.Context, browser *config.Browser) error {
	return exec.Command(
		browser.BinaryPath,
		slices.Concat(
			[]string{l.targetURL},
			browser.LaunchArgs,
		)...,
	).Start()
}
