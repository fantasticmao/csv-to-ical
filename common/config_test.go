package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseConfig(t *testing.T) {
	config, err := ParseConfig("testdata/config_test.yaml")
	assert.NotNil(t, config)
	assert.Nil(t, err)

	assert.Equal(t, "127.0.0.1:7788", config.BindAddress)
	assert.Equal(t, "testdata/calendar_test.csv", config.CsvProviders["foo"].File)
	assert.Equal(t, "zh_cn", config.CsvProviders["foo"].Language)
	assert.Equal(t, 3, config.CsvProviders["foo"].RecurCnt)
	assert.Equal(t, "https://raw.githubusercontent.com/fantasticmao/csv-to-ical/main/csv/testdata/calendar_test.csv", config.CsvProviders["bar"].Url)
	assert.Equal(t, "en", config.CsvProviders["bar"].Language)
	assert.Equal(t, 5, config.CsvProviders["bar"].RecurCnt)
}
