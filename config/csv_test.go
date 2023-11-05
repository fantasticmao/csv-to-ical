package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseEventFromFile(t *testing.T) {
	events, err := ParseEventFromFile("testdata/calendar_test.csv")
	assert.Nil(t, err)
	assert.NotNil(t, events)

	validate(t, events)
}

func TestParseEventFromUrl(t *testing.T) {
	events, err := ParseEventFromUrl("https://raw.githubusercontent.com/fantasticmao/csv-to-ical/main/config/testdata/calendar_test.csv")
	assert.Nil(t, err)
	assert.NotNil(t, events)

	validate(t, events)
}

func validate(t *testing.T, events []Event) {
	assert.Equal(t, "小明的毕业纪念日", events[0].name)
	assert.Equal(t, time.June, events[0].month)
	assert.Equal(t, 1, events[0].day)
	assert.Equal(t, 2022, events[0].year)
	assert.Equal(t, Solar, events[0].calendarType)

	assert.Equal(t, "世界人口日", events[1].name)
	assert.Equal(t, time.July, events[1].month)
	assert.Equal(t, 11, events[1].day)
	assert.Equal(t, -1, events[1].year)
	assert.Equal(t, Solar, events[1].calendarType)

	assert.Equal(t, "小明的结婚纪念日", events[2].name)
	assert.Equal(t, time.August, events[2].month)
	assert.Equal(t, 1, events[2].day)
	assert.Equal(t, 2023, events[2].year)
	assert.Equal(t, Lunar, events[2].calendarType)

	assert.Equal(t, "中秋节", events[3].name)
	assert.Equal(t, time.August, events[3].month)
	assert.Equal(t, 15, events[3].day)
	assert.Equal(t, -1, events[3].year)
	assert.Equal(t, Lunar, events[3].calendarType)

	assert.Equal(t, "小明", events[4].name)
	assert.Equal(t, time.September, events[4].month)
	assert.Equal(t, 1, events[4].day)
	assert.Equal(t, 2000, events[4].year)
	assert.Equal(t, BirthdaySolar, events[4].calendarType)

	assert.Equal(t, "小明的朋友", events[5].name)
	assert.Equal(t, time.October, events[5].month)
	assert.Equal(t, 1, events[5].day)
	assert.Equal(t, -1, events[5].year)
	assert.Equal(t, BirthdaySolar, events[5].calendarType)

	assert.Equal(t, "小明的爷爷", events[6].name)
	assert.Equal(t, time.November, events[6].month)
	assert.Equal(t, 1, events[6].day)
	assert.Equal(t, 1955, events[6].year)
	assert.Equal(t, BirthdayLunar, events[6].calendarType)

	assert.Equal(t, "小明的奶奶", events[7].name)
	assert.Equal(t, time.December, events[7].month)
	assert.Equal(t, 1, events[7].day)
	assert.Equal(t, -1, events[7].year)
	assert.Equal(t, BirthdayLunar, events[7].calendarType)
}
