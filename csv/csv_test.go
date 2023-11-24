package csv

import (
	"github.com/fantasticmao/csv-to-ical/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseEventFromFile(t *testing.T) {
	events, err := ParseEventFromFile("testdata/calendar_test.csv")
	assert.Nil(t, err)
	assert.NotNil(t, events)

	validate(t, events)
}

func TestParseEventFromUrl(t *testing.T) {
	events, err := ParseEventFromUrl("https://raw.githubusercontent.com/fantasticmao/csv-to-ical/main/csv/testdata/calendar_test.csv")
	assert.Nil(t, err)
	assert.NotNil(t, events)

	validate(t, events)
}

func validate(t *testing.T, events []Event) {
	assert.Equal(t, "小明的毕业纪念日", events[0].Name)
	assert.Equal(t, 6, events[0].Month)
	assert.Equal(t, 1, events[0].Day)
	assert.Equal(t, 2022, events[0].Year)
	assert.Equal(t, common.Solar, events[0].CalendarType)

	assert.Equal(t, "世界人口日", events[1].Name)
	assert.Equal(t, 7, events[1].Month)
	assert.Equal(t, 11, events[1].Day)
	assert.Equal(t, -1, events[1].Year)
	assert.Equal(t, common.Solar, events[1].CalendarType)

	assert.Equal(t, "小明的结婚纪念日", events[2].Name)
	assert.Equal(t, 8, events[2].Month)
	assert.Equal(t, 1, events[2].Day)
	assert.Equal(t, 2023, events[2].Year)
	assert.Equal(t, common.Lunar, events[2].CalendarType)

	assert.Equal(t, "中秋节", events[3].Name)
	assert.Equal(t, 8, events[3].Month)
	assert.Equal(t, 15, events[3].Day)
	assert.Equal(t, -1, events[3].Year)
	assert.Equal(t, common.Lunar, events[3].CalendarType)

	assert.Equal(t, "小明", events[4].Name)
	assert.Equal(t, 9, events[4].Month)
	assert.Equal(t, 1, events[4].Day)
	assert.Equal(t, 2000, events[4].Year)
	assert.Equal(t, common.BirthdaySolar, events[4].CalendarType)

	assert.Equal(t, "小明朋友", events[5].Name)
	assert.Equal(t, 10, events[5].Month)
	assert.Equal(t, 1, events[5].Day)
	assert.Equal(t, -1, events[5].Year)
	assert.Equal(t, common.BirthdaySolar, events[5].CalendarType)

	assert.Equal(t, "小明爷爷", events[6].Name)
	assert.Equal(t, 11, events[6].Month)
	assert.Equal(t, 1, events[6].Day)
	assert.Equal(t, 1955, events[6].Year)
	assert.Equal(t, common.BirthdayLunar, events[6].CalendarType)

	assert.Equal(t, "小明奶奶", events[7].Name)
	assert.Equal(t, 12, events[7].Month)
	assert.Equal(t, 1, events[7].Day)
	assert.Equal(t, -1, events[7].Year)
	assert.Equal(t, common.BirthdayLunar, events[7].CalendarType)
}
