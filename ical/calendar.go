package ical

import (
	"bytes"
	"fmt"
	"github.com/fantasticmao/csv-to-ical/config"
	"github.com/fantasticmao/csv-to-ical/date"
	"html/template"
	"time"
)

type Object struct {
	ProdId     string
	Version    string
	Components []ComponentEvent
}

func (obj Object) String() string {
	temp, err := template.New("iCalendarObject").Parse(`BEGIN:VCALENDAR
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

	output := &bytes.Buffer{}
	err = temp.Execute(output, obj)
	if err != nil {
		panic(err)
	}
	return output.String()
}

func NewObject(components []ComponentEvent) Object {
	return Object{
		ProdId:     config.FullName,
		Version:    "2.0",
		Components: components,
	}
}

type ComponentEvent struct {
	DtStamp    string
	Uid        string
	DtStart    string
	Class      string
	Language   config.Language
	Summary    string
	Transp     string
	RecurCount int
}

func (cmpEvent ComponentEvent) String() string {
	temp, err := template.New("iCalendarComponents").Parse(`BEGIN:VEVENT
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

	output := &bytes.Buffer{}
	err = temp.Execute(output, cmpEvent)
	if err != nil {
		panic(err)
	}
	return output.String()
}

func NewComponentEvent(uid string, language config.Language, summary string, recur int,
	now, start time.Time) ComponentEvent {
	return ComponentEvent{
		DtStamp:    date.FormatDate(now),
		Uid:        uid,
		DtStart:    date.FormatDate(start),
		Class:      "PUBLIC",
		Language:   language,
		Summary:    summary,
		Transp:     "TRANSPARENT",
		RecurCount: recur,
	}
}

func FormatUid(name string, datetime time.Time, calendarType config.CalendarType, host string) string {
	return fmt.Sprintf("%s-%s-%s@%s", name, date.FormatDate(datetime), calendarType, host)
}
