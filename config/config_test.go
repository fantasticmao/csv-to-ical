package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseConfig(t *testing.T) {
	config, err := ParseConfig("testdata/config_test.yaml")
	assert.NotNil(t, config)
	assert.Nil(t, err)

	assert.Equal(t, "info", config.LogLevel)
	assert.Equal(t, "0.0.0.0:7788", config.BindAddress)
	assert.Equal(t, "testdata/calendar_test.csv", config.CsvProviders["foo"].File)
	assert.Equal(t, ZhCn, config.CsvProviders["foo"].Language)
	assert.Equal(t, "https://raw.githubusercontent.com/fantasticmao/csv-to-ical/main/csv/testdata/calendar_test.csv", config.CsvProviders["bar"].Url)
	assert.Equal(t, En, config.CsvProviders["bar"].Language)
}
