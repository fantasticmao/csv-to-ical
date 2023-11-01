package ical

import (
	"github.com/fantasticmao/csv-to-ical/date"
	"github.com/fantasticmao/csv-to-ical/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObject_String(t *testing.T) {
	obj := NewObject("github.com/fantasticmao/csv-to-ical", []ComponentEvent{})
	str := obj.String()
	assert.Equal(t, `BEGIN:VCALENDAR
PRODID:github.com/fantasticmao/csv-to-ical
VERSION:2.0
END:VCALENDAR
`, str)
}

func TestComponentEvent_String(t *testing.T) {
	datetime, err := date.ParseTime("20231031")
	assert.Nil(t, err)

	cmpEvent := NewComponentEvent("Tom-20231031-birthday_solar@localhost", i18n.En,
		"Tom’s 18th Birthday", 5, datetime, datetime)
	str := cmpEvent.String()
	assert.Equal(t, `BEGIN:VEVENT
DTSTAMP;VALUE=DATE:20231031
UID:Tom-20231031-birthday_solar@localhost
DTSTART;VALUE=DATE:20231031
CLASS:PUBLIC
SUMMARY;LANGUAGE=en:Tom’s 18th Birthday
TRANSP:TRANSPARENT
RRULE:FREQ=YEARLY;COUNT=5
END:VEVENT`, str)

	cmpEvent = NewComponentEvent("小明-20231031-birthday_lunar@localhost", i18n.ZhCn,
		"小明的18岁农历生日", 0, datetime, datetime)
	str = cmpEvent.String()
	assert.Equal(t, `BEGIN:VEVENT
DTSTAMP;VALUE=DATE:20231031
UID:小明-20231031-birthday_lunar@localhost
DTSTART;VALUE=DATE:20231031
CLASS:PUBLIC
SUMMARY;LANGUAGE=zh_CN:小明的18岁农历生日
TRANSP:TRANSPARENT
END:VEVENT`, str)
}
