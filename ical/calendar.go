package ical

import (
	"bytes"
	"github.com/fantasticmao/csv-to-ical/date"
	"github.com/fantasticmao/csv-to-ical/i18n"
	"html/template"
	"time"
)

type Object struct {
	ProdId     string
	Version    string
	Components []ComponentEvent
}

func NewObject(prodId string, components []ComponentEvent) Object {
	return Object{
		ProdId:     prodId,
		Version:    "2.0",
		Components: components,
	}
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

type ComponentEvent struct {
	DtStamp    string
	Uid        string
	DtStart    string
	Class      string
	Language   i18n.Language
	Summary    string
	Transp     string
	RecurCount int
}

func NewComponentEvent(uid string, language i18n.Language, summary string, recur int,
	now, start time.Time) ComponentEvent {
	return ComponentEvent{
		DtStamp:    date.FormatTime(now),
		Uid:        uid,
		DtStart:    date.FormatTime(start),
		Class:      "PUBLIC",
		Language:   language,
		Summary:    summary,
		Transp:     "TRANSPARENT",
		RecurCount: recur,
	}
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
