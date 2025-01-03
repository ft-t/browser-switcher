//go:build linux

package main

import (
	"context"
	"os"
	"os/exec"
	"syscall"

	"github.com/rs/zerolog"

	"github.com/ft-t/browser-switcher/pkg/escaper"
)

func run(ctx context.Context) error {
	// Escape arguments
	var escapedArgs []string
	for _, arg := range os.Args[1:] {
		escapedArgs = append(escapedArgs, escaper.Escape(arg))
	}

	args := append([]string{"-e", "/usr/local/bin/browser-switcher-proxied"}, escapedArgs...) // todo: remove hardcoded path

	zerolog.Ctx(ctx).Debug().Msgf("running command: terminal %v", args)

	cmd := exec.Command(
		"x-terminal-emulator",
		args...,
	)

	cmd.SysProcAttr = &syscall.SysProcAttr{}

	return cmd.Start()
}
