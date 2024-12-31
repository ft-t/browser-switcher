package config

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
)

type Config struct {
	Browsers []*Browser `json:"browsers"`
	UI       UI         `json:"ui"`
}

type UI struct {
	Renderer string `json:"renderer"`
}

type Browser struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	LaunchArgs     []string `json:"launch_args"`
	BinaryPath     string   `json:"binary_path"`
	CustomIconPath string   `json:"custom_icon_path"`
	Rules          []string `json:"rules"`
}

func ReadConfig(ctx context.Context) (*Config, error) {
	targetFile := "config.json"

	if homeDir, err := os.UserHomeDir(); err == nil {
		targetFile = filepath.Join(homeDir, "BrowserSwitcher", "config.json")
	} else {
		zerolog.Ctx(ctx).Error().Err(err).Msg("failed to get user home directory")
	}

	zerolog.Ctx(ctx).Info().Msgf("loading config from %s", targetFile)

	fileData, err := os.ReadFile(targetFile)
	if err != nil {
		return nil, err
	}

	var targetConfig Config
	if err = json.Unmarshal(fileData, &targetConfig); err != nil {
		return nil, err
	}

	return &targetConfig, nil
}
