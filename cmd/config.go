package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"path"
)

type Config struct {
	APIToken string `json:"api_token"`
}

var (
	ErrLoadConfig = errors.New("failed to load config")
)

func loadConfig() (*Config, error) {
	dir, err := configDir()
	if err != nil {
		return nil, err
	}

	name := configFilename(dir)
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	if apiToken != "" {
		cfg.APIToken = apiToken
	}

	return &cfg, nil
}

func configDir() (string, error) {
	hmd, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(hmd, ".todoist"), nil
}

func configFilename(dir string) string {
	return path.Join(dir, "config.json")
}
