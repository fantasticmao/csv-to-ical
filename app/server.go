package app

import (
	"fmt"
	"github.com/fantasticmao/csv-to-ical/config"
	"github.com/fantasticmao/csv-to-ical/csv"
	"github.com/fantasticmao/csv-to-ical/ical"
	"net/http"
	"path"
	"runtime"
	"strconv"
)

func StartServer(addr string) {
	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			fmt.Printf("start HTTP server error: %v\n", err.Error())
		}
	}()
}

func RegisterVersionHandler() {
	http.HandleFunc("/version", func(writer http.ResponseWriter, request *http.Request) {
		writeResponse(writer, fmt.Sprintf("%v %v %v-%v with %v built on commit %v at %v\n",
			config.Name, config.Version, runtime.GOOS, runtime.GOARCH, runtime.Version(),
			config.CommitHash, config.BuildTime))
	})
}

func RegisterRemoteHandler() {
	http.HandleFunc("/remote", func(writer http.ResponseWriter, request *http.Request) {
		url := request.URL.Query().Get("url")
		if url == "" {
			writeResponse400(writer, "query param: url is required")
			return
		}

		var language config.Language
		if param := request.URL.Query().Get("lang"); param == "" {
			language = config.En
		} else {
			lang, err := config.ParseLanguage(param)
			if err != nil {
				writeResponse400(writer, fmt.Sprintf("query param: lang error: %v", err.Error()))
				return
			} else {
				language = lang
			}
		}

		var recurCnt int
		if param := request.URL.Query().Get("recurCnt"); param == "" {
			recurCnt = 5
		} else {
			cnt, err := strconv.Atoi(param)
			if err != nil {
				writeResponse400(writer, fmt.Sprintf("query param: recurCnt error: %v", err.Error()))
				return
			} else {
				if cnt < 0 {
					writeResponse400(writer, "query param: recurCnt cannot be negative")
					return
				} else if cnt > 10 {
					writeResponse400(writer, "query param: recurCnt cannot be grater than 10")
					return
				} else {
					recurCnt = cnt
				}
			}
		}

		events, err := csv.ParseEventFromUrl(url)
		if err != nil {
			writeResponse400(writer, fmt.Sprintf("fetch csv events from url: %v error: %v", url, err.Error()))
			return
		}

		var components []ical.ComponentEvent
		for _, event := range events {
			cmpEvents := csvToIcal(event, language, recurCnt, request.Host)
			components = append(components, cmpEvents...)
		}

		obj := ical.NewObject(components)
		writeResponse(writer, obj.String())
	})
}

func RegisterLocalHandler(configDir, owner string, provider config.CsvProvider) error {
	var events []csv.Event
	if provider.File != "" {
		e, err := csv.ParseEventFromFile(path.Join(configDir, provider.File))
		if err != nil {
			return err
		}
		events = e
	}
	if provider.Url != "" {
		e, err := csv.ParseEventFromUrl(provider.Url)
		if err != nil {
			return err
		}
		events = e
	}

	http.HandleFunc("/local/"+owner, func(writer http.ResponseWriter, request *http.Request) {
		var components []ical.ComponentEvent
		for _, event := range events {
			language, _ := config.ParseLanguage(provider.Language)
			cmpEvents := csvToIcal(event, language, provider.RecurCnt, request.Host)
			components = append(components, cmpEvents...)
		}

		obj := ical.NewObject(components)
		writeResponse(writer, obj.String())
	})
	return nil
}

func csvToIcal(event csv.Event, language config.Language, recurCnt int, host string) []ical.ComponentEvent {
	if event.CalendarType == config.Solar {
		return convertForSolar(event, language, recurCnt, host)
	} else if event.CalendarType == config.Lunar {
		return convertForLunar(event, language, recurCnt, host)
	} else if event.CalendarType == config.BirthdaySolar {
		return convertForBirthdaySolar(event, language, recurCnt, host)
	} else if event.CalendarType == config.BirthdayLunar {
		return convertForBirthdayLunar(event, language, recurCnt, host)
	}
	return nil
}

func writeResponse(writer http.ResponseWriter, response string) {
	writer.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	if _, err := fmt.Fprintln(writer, response); err != nil {
		fmt.Printf("write HTTP response error: %v\n", err.Error())
	}
}

func writeResponse400(writer http.ResponseWriter, response string) {
	writer.WriteHeader(400)
	writer.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	if _, err := fmt.Fprintln(writer, response); err != nil {
		fmt.Printf("write HTTP response error: %v\n", err.Error())
	}
}
