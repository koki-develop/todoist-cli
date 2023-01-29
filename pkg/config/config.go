package config

import (
	"encoding/json"
	"os"
	"path"

	"github.com/koki-develop/todoist-cli/pkg/renderer"
)

type Config struct {
	APIToken *string          `json:"api_token,omitempty"`
	Format   *renderer.Format `json:"format,omitempty"`
}

func format(f renderer.Format) *renderer.Format { return &f }

func Load(def *Config) (*Config, error) {
	dir, err := Dir()
	if err != nil {
		return nil, err
	}

	name := Filename(dir)
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg := &Config{}
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	if def != nil {
		if def.APIToken != nil {
			cfg.APIToken = def.APIToken
		}
		if def.Format != nil {
			cfg.Format = def.Format
		}
	}
	if cfg.Format == nil {
		cfg.Format = format(renderer.FormatTable)
	}

	return cfg, nil
}

func Dir() (string, error) {
	hmd, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(hmd, ".todoist-cli"), nil
}

func Filename(dir string) string {
	return path.Join(dir, "config.json")
}
