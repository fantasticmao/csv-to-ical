package ical

import (
	"bytes"
	"fmt"
	"github.com/fantasticmao/csv-to-ical/common"
	"text/template"
	"time"
)

var t = template.New("iCalendar")
var nameObject = "object"
var nameComponent = "component"

func init() {
	template.Must(t.New(nameObject).Parse(`BEGIN:VCALENDAR
PRODID:{{ .ProdId }}
VERSION:{{ .Version }}
{{- range .Components }}
{{ .Transform }}
{{- end }}
END:VCALENDAR
`))

	template.Must(t.New(nameComponent).Parse(`BEGIN:VEVENT
DTSTAMP;VALUE=DATE:{{ .DtStamp }}
UID:{{ .Uid }}
DTSTART:{{ .DtStart }}
DTEND:{{ .DtEnd }}
CLASS:{{ .Class }}
SUMMARY;LANGUAGE={{ .Language }}:{{ .Summary }}
TRANSP:{{ .Transp }}
{{- if .RecurCount }}
RRULE:FREQ=YEARLY;COUNT={{ .RecurCount }}
{{- end }}
END:VEVENT`))
}

type Object struct {
	ProdId     string
	Version    string
	Components []ComponentEvent
}

func (obj Object) Transform() (string, error) {
	output := &bytes.Buffer{}
	if err := t.ExecuteTemplate(output, nameObject, obj); err != nil {
		return "", err
	} else {
		return output.String(), nil
	}
}

func NewObject(components []ComponentEvent) Object {
	return Object{
		ProdId:     common.FullName,
		Version:    "2.0",
		Components: components,
	}
}

type ComponentEvent struct {
	DtStamp    string
	Uid        string
	DtStart    string
	DtEnd      string
	Class      string
	Language   common.Language
	Summary    string
	Transp     string
	RecurCount int
}

func (cmpEvent ComponentEvent) Transform() (string, error) {
	output := &bytes.Buffer{}
	if err := t.ExecuteTemplate(output, nameComponent, cmpEvent); err != nil {
		return "", err
	} else {
		return output.String(), nil
	}
}

func NewComponentEvent(uid string, language common.Language, summary string, recurCnt int,
	now, start time.Time) ComponentEvent {
	start = common.ResetTime(start, 8, 0, 0)
	end := common.ResetTime(start, 23, 59, 59)
	return ComponentEvent{
		DtStamp:    common.FormatDate(now),
		Uid:        uid,
		DtStart:    common.FormatDatetime(start),
		DtEnd:      common.FormatDatetime(end),
		Class:      "PUBLIC",
		Language:   language,
		Summary:    summary,
		Transp:     "TRANSPARENT",
		RecurCount: recurCnt,
	}
}

// FormatUid generate globally unique identifier for the calendar component,
// for more details see https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.7
func FormatUid(name string, date time.Time, calendarType common.CalendarType, host string) string {
	return fmt.Sprintf("%s-%s-%s@%s", name, common.FormatDate(date), calendarType, host)
}
