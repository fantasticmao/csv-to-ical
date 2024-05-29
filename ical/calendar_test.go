package ical

import (
	"github.com/fantasticmao/csv-to-ical/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObject_String(t *testing.T) {
	obj := NewObject([]ComponentEvent{})
	str, err := obj.Transform()
	assert.Nil(t, err)
	assert.Equal(t, `BEGIN:VCALENDAR
PRODID:github.com/fantasticmao/csv-to-ical
VERSION:2.0
END:VCALENDAR
`, str)
}

func TestComponentEvent_String(t *testing.T) {
	datetime := common.NewDate(2023, 10, 31)

	cmpEvent := NewComponentEvent("Tom-20231031-birthday_solar@localhost", common.En,
		"Tom’s 18th Birthday", 5, datetime, datetime)
	str, err := cmpEvent.Transform()
	assert.Nil(t, err)
	assert.Equal(t, `BEGIN:VEVENT
DTSTAMP;VALUE=DATE:20231031
UID:Tom-20231031-birthday_solar@localhost
DTSTART;VALUE=DATE:20231031
CLASS:PUBLIC
SUMMARY;LANGUAGE=en:Tom’s 18th Birthday
TRANSP:TRANSPARENT
RRULE:FREQ=YEARLY;COUNT=5
END:VEVENT`, str)

	cmpEvent = NewComponentEvent("小明-20231031-birthday_lunar@localhost", common.ZhCn,
		"小明的18岁农历生日", 0, datetime, datetime)
	str, err = cmpEvent.Transform()
	assert.Nil(t, err)
	assert.Equal(t, `BEGIN:VEVENT
DTSTAMP;VALUE=DATE:20231031
UID:小明-20231031-birthday_lunar@localhost
DTSTART;VALUE=DATE:20231031
CLASS:PUBLIC
SUMMARY;LANGUAGE=zh_CN:小明的18岁农历生日
TRANSP:TRANSPARENT
END:VEVENT`, str)
}
