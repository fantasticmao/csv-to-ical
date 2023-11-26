package common

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	BindAddress  string                 `yaml:"bind-address"`
	HttpClient   HttpClient             `yaml:"http-client"`
	CsvProviders map[string]CsvProvider `yaml:"csv-providers"`
}

type HttpClient struct {
	Timeout int    `yaml:"timeout"`
	Proxy   string `yaml:"proxy"`
}

type CsvProvider struct {
	File     string `yaml:"file"`
	Url      string `yaml:"url"`
	Language string `yaml:"language"`
	RecurCnt int    `yaml:"recurCnt"`
}

func (provider *CsvProvider) UnmarshalYAML(value *yaml.Node) error {
	type rawCsvProvider CsvProvider
	raw := rawCsvProvider{
		Language: string(En),
		RecurCnt: 5,
	}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	*provider = CsvProvider(raw)
	return nil
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
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("config file: %v does not exist, falling back to default settings\n", path)
			data = []byte("")
		} else {
			return nil, err
		}
	}

	config := &Config{
		BindAddress: "0.0.0.0:7788",
		HttpClient: HttpClient{
			Timeout: 3_000,
		},
	}

	if err = yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	if err = config.validate(); err != nil {
		return nil, err
	}

	InitHttpClient(config.HttpClient)
	return config, nil
}
