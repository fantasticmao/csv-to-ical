package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseEventFromFile(t *testing.T) {
	events := ParseEventFromFile("testdata/calendar_test.csv")
	assert.NotNil(t, events)

	assert.Equal(t, "小明", events[0].name)
	assert.Equal(t, time.May, events[0].month)
	assert.Equal(t, 1, events[0].day)
	assert.Equal(t, 2000, events[0].year)
	assert.Equal(t, BirthdaySolar, events[0].tag)

	assert.Equal(t, "小红", events[1].name)
	assert.Equal(t, time.October, events[0].month)
	assert.Equal(t, 1, events[0].day)
	assert.Equal(t, 2002, events[0].year)
	assert.Equal(t, BirthdayLunar, events[1].tag)
}
