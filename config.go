package main

type Config struct {
	Browsers []*Browser `json:"browsers"`
}

type Browser struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	LaunchArgs     []string `json:"launch_args"`
	BinaryPath     string   `json:"binary_path"`
	CustomIconPath string   `json:"custom_icon_path"`
	Rules          []string `json:"rules"`
}
