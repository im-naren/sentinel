package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
)

type AppConfig struct {
	Name                  string `koanf:"name"`
	PortfolioDataPath     string `koanf:"portfolioDataPath"`
	PortfolioDataFilePrefix string `koanf:"portfolioDataFilePrefix"`
	Version               string `koanf:"version"`
}

type Config struct {
	App AppConfig `koanf:"app"`
}

func LoadConfig(filePath string) (*Config, error) {
	var config Config
	k := koanf.New(".")

	if err := k.Load(file.Provider(filePath), toml.Parser()); err != nil {
		return nil, err
	}

	if err := k.Unmarshal("", &config); err != nil {
		return nil, err
	}

	return &config, nil
}
