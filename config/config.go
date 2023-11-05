package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	logLevel     string                 `yaml:"log-level"`
	CsvProviders map[string]CsvProvider `yaml:"csv-providers"`
}

type CsvProvider struct {
	File string   `yaml:"file"`
	Url  string   `yaml:"url"`
	Lang Language `yaml:"lang"`
}

func ParseConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config := &Config{
		logLevel: "info",
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}

	return *config
}
