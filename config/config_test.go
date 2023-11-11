package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseConfig(t *testing.T) {
	config, err := ParseConfig("testdata/config_test.yaml")
	assert.NotNil(t, config)
	assert.Nil(t, err)

	assert.Equal(t, "info", config.logLevel)
	assert.Equal(t, "testdata/calendar_test.csv", config.CsvProviders["user1"].File)
	assert.Equal(t, ZhCn, config.CsvProviders["user1"].Lang)
	assert.Equal(t, "https://raw.githubusercontent.com/fantasticmao/csv-to-ical/main/config/testdata/calendar_test.csv", config.CsvProviders["user2"].Url)
	assert.Equal(t, En, config.CsvProviders["user2"].Lang)
}
