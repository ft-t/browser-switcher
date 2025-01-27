//go:build windows

package main

import (
	"context"
	"os"
	"os/exec"
	"slices"
	"syscall"

	"github.com/rs/zerolog"
	"golang.org/x/sys/windows/registry"

	"github.com/ft-t/browser-switcher/pkg/escaper"
)

func findAppRegistrationPath(ctx context.Context) (string, error) {
	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Software\Clients\StartMenuInternet\Browser Switcher\shell\open\command`,
		registry.QUERY_VALUE,
	)
	if err != nil {
		return "", err
	}

	val, _, err := key.GetStringValue("Proxied")
	if err != nil {
		return "", err
	}

	zerolog.Ctx(ctx).Debug().Msgf("found app registration path: %s", val)

	return val, err
}

func run(ctx context.Context) error {
	appPath, err := findAppRegistrationPath(ctx)
	if err != nil {
		return err
	}

	var escapedArgs []string
	for _, arg := range os.Args[1:] {
		escapedArgs = append(escapedArgs, escaper.Escape(arg))
	}

	args := slices.Concat([]string{
		"/C",
		"start",
		appPath,
	}, escapedArgs)

	zerolog.Ctx(ctx).Debug().Msgf("running command: cmd.exe %v", escapedArgs)

	cmd := exec.Command(
		"cmd.exe",
		args...,
	)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,       // nolint
		CreationFlags: 0x08000000, // nolint
	}

	return cmd.Start()
}
