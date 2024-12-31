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

	args := slices.Concat([]string{
		"/C",
		"start",
		appPath,
	}, os.Args[1:])

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
