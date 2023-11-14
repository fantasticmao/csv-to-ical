package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	LogLevel     string                 `yaml:"log-level"`
	BindAddress  string                 `yaml:"bind-address"`
	CsvProviders map[string]CsvProvider `yaml:"csv-providers"`
}

type CsvProvider struct {
	File string   `yaml:"file"`
	Url  string   `yaml:"url"`
	Lang Language `yaml:"lang"`
}

func (cfg *Config) validate() error {
	for k, v := range cfg.CsvProviders {
		if v.File == "" && v.Url == "" {
			return fmt.Errorf("file and url fields in the key: %v cannot be empty at the same time", k)
		}
	}
	return nil
}

func ParseConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{
		LogLevel:    "info",
		BindAddress: "127.0.0.1:7788",
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	err = config.validate()
	if err != nil {
		return nil, err
	}
	return config, nil
}
