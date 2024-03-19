package ical

import (
	"bytes"
	"fmt"
	"github.com/fantasticmao/csv-to-ical/common"
	"text/template"
	"time"
)

var iCalendarObjectTemple *template.Template
var iCalendarComponentTemple *template.Template

func init() {
	var err error
	iCalendarObjectTemple, err = template.New("iCalendarObject").Parse(`BEGIN:VCALENDAR
PRODID:{{ .ProdId }}
VERSION:{{ .Version }}
{{- range .Components }}
{{ . }}
{{- end }}
END:VCALENDAR
`)
	if err != nil {
		panic(err)
	}

	iCalendarComponentTemple, err = template.New("iCalendarComponent").Parse(`BEGIN:VEVENT
DTSTAMP;VALUE=DATE:{{ .DtStamp }}
UID:{{ .Uid }}
DTSTART;VALUE=DATE:{{ .DtStart }}
CLASS:{{ .Class }}
SUMMARY;LANGUAGE={{ .Language }}:{{ .Summary }}
TRANSP:{{ .Transp }}
{{- if .RecurCount }}
RRULE:FREQ=YEARLY;COUNT={{ .RecurCount }}
{{- end }}
END:VEVENT`)
	if err != nil {
		panic(err)
	}
}

type Object struct {
	ProdId     string
	Version    string
	Components []ComponentEvent
}

func (obj Object) String() string {
	output := &bytes.Buffer{}
	if err := iCalendarObjectTemple.Execute(output, obj); err != nil {
		panic(err)
	}
	return output.String()
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
	Class      string
	Language   common.Language
	Summary    string
	Transp     string
	RecurCount int
}

func (cmpEvent ComponentEvent) String() string {
	output := &bytes.Buffer{}
	if err := iCalendarComponentTemple.Execute(output, cmpEvent); err != nil {
		panic(err)
	}
	return output.String()
}

func NewComponentEvent(uid string, language common.Language, summary string, recurCnt int,
	now, start time.Time) ComponentEvent {
	return ComponentEvent{
		DtStamp:    common.FormatDate(now),
		Uid:        uid,
		DtStart:    common.FormatDate(start),
		Class:      "PUBLIC",
		Language:   language,
		Summary:    summary,
		Transp:     "TRANSPARENT",
		RecurCount: recurCnt,
	}
}

// FormatUid generate globally unique identifier for the calendar component,
// for more details see https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.7
func FormatUid(name string, datetime time.Time, calendarType common.CalendarType, host string) string {
	return fmt.Sprintf("%s-%s-%s@%s", name, common.FormatDate(datetime), calendarType, host)
}
