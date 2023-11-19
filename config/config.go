package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	BindAddress  string                 `yaml:"bind-address"`
	CsvProviders map[string]CsvProvider `yaml:"csv-providers"`
}

type CsvProvider struct {
	File     string `yaml:"file"`
	Url      string `yaml:"url"`
	Language string `yaml:"language"`
	RecurCnt int    `yaml:"recurCnt"`
}

func (cfg *Config) validate() error {
	for key, val := range cfg.CsvProviders {
		if val.File == "" && val.Url == "" {
			return fmt.Errorf("file and url fields in key: %v cannot both be empty", key)
		}
		if _, err := ParseLanguage(val.Language); err != nil {
			return err
		}
		if val.RecurCnt < 0 {
			return fmt.Errorf("recurCnt in key: %v cannot be negative", key)
		} else if val.RecurCnt > 10 {
			return fmt.Errorf("recurCnt in key: %v cannot be grater than 10", key)
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
		BindAddress: "0.0.0.0:7788",
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
