package app

import (
	"fmt"
	"github.com/fantasticmao/csv-to-ical/config"
	"github.com/fantasticmao/csv-to-ical/csv"
	"github.com/fantasticmao/csv-to-ical/ical"
	"net/http"
)

func RegisterHandler(owner string, provider config.CsvProvider) error {
	var events []csv.Event
	if provider.File != "" {
		e, err := csv.ParseEventFromFile(provider.File)
		if err != nil {
			return err
		}
		events = e
	}
	if provider.Url != "" {
		e, err := csv.ParseEventFromFile(provider.File)
		if err != nil {
			return err
		}
		events = e
	}

	f := handleEvents(events, provider.Lang)
	http.HandleFunc(owner, f)
	return nil
}

func StartServer(addr string) {
	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			fmt.Printf("start HTTP server error: %v\n", err.Error())
		}
	}()
}

func handleEvents(events []csv.Event, lang config.Language) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var components []ical.ComponentEvent
		for _, event := range events {
			cmpEvents := csvToIcal(event, lang, request.Host)
			components = append(components, cmpEvents...)
		}

		obj := ical.NewObject(components)
		_, err := fmt.Fprintln(writer, obj.String())
		if err != nil {
			fmt.Printf("write HTTP response error: %v\n", err.Error())
		}
	}
}

func csvToIcal(event csv.Event, lang config.Language, host string) []ical.ComponentEvent {
	recurCount := 5
	if event.CalendarType == config.Solar {
		return convertForSolar(event, lang, recurCount, host)
	} else if event.CalendarType == config.Lunar {
		return convertForLunar(event, lang, recurCount, host)
	} else if event.CalendarType == config.BirthdaySolar {
		return convertForBirthdaySolar(event, lang, recurCount, host)
	} else if event.CalendarType == config.BirthdayLunar {
		return convertForBirthdayLunar(event, lang, recurCount, host)
	}
	return nil
}
