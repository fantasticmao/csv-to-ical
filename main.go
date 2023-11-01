package main

import (
	"fmt"
	"github.com/fantasticmao/csv-to-ical/date"
	"github.com/fantasticmao/csv-to-ical/i18n"
	"github.com/fantasticmao/csv-to-ical/ical"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/fantasticmao/birthdays.ics", func(writer http.ResponseWriter, request *http.Request) {
		datetime, err := date.ParseTime("20231031")
		if err != nil {
			panic(err)
		}

		events := []ical.ComponentEvent{
			ical.NewComponentEvent("Tom-20231031-birthday_solar@localhost", i18n.En,
				"Tom’s 18th Birthday", 5, time.Now(), datetime),
			ical.NewComponentEvent("小明-20231031-birthday_lunar@localhost", i18n.ZhCn,
				"小明的18岁农历生日", 0, time.Now(), datetime),
		}
		obj := ical.NewObject("github.com/fantasticmao/csv-to-ical", events)

		_, err = fmt.Fprintln(writer, obj.String())
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
