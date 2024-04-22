package app

import (
	"fmt"
	"github.com/fantasticmao/csv-to-ical/common"
	"github.com/fantasticmao/csv-to-ical/csv"
	"github.com/fantasticmao/csv-to-ical/ical"
	"github.com/fantasticmao/csv-to-ical/log"
	"net/http"
	"path"
	"runtime"
	"strconv"
)

func StartServer(addr string) {
	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			log.Error("start HTTP server error: %v", err.Error())
		}
	}()
}

func RegisterDefaultHandler() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writeResponse(writer, fmt.Sprintf(`If you see this message, the %v is successfully installed and working.
For more informations see https://github.com/fantasticmao/csv-to-ical
`, common.Name))
	})
	http.HandleFunc("/version", func(writer http.ResponseWriter, request *http.Request) {
		writeResponse(writer, fmt.Sprintf("%v %v %v-%v with %v built on commit %v at %v\n",
			common.Name, common.Version, runtime.GOOS, runtime.GOARCH, runtime.Version(),
			common.CommitHash, common.BuildTime))
	})
}

func RegisterRemoteHandler() {
	http.HandleFunc("/remote", func(writer http.ResponseWriter, request *http.Request) {
		url := request.URL.Query().Get("url")
		if url == "" {
			writeResponse400(writer, "query param: url is required")
			return
		}

		var language common.Language
		if param := request.URL.Query().Get("lang"); param == "" {
			language = common.En
		} else {
			lang, err := common.ParseLanguage(param)
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
			writeResponse400(writer, fmt.Sprintf("fetch csv events from url: '%v' error: %v", url, err.Error()))
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

func RegisterLocalHandler(configDir, owner string, provider common.CsvProvider) error {
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
			language, _ := common.ParseLanguage(provider.Language)
			cmpEvents := csvToIcal(event, language, provider.RecurCnt, request.Host)
			components = append(components, cmpEvents...)
		}

		obj := ical.NewObject(components)
		writeResponse(writer, obj.String())
	})
	return nil
}

func csvToIcal(event csv.Event, language common.Language, recurCnt int, host string) []ical.ComponentEvent {
	switch event.CalendarType {
	case common.Solar:
		return convertForSolar(event, language, recurCnt, host)
	case common.Lunar:
		return convertForLunar(event, language, recurCnt, host)
	case common.BirthdaySolar:
		return convertForBirthdaySolar(event, language, recurCnt, host)
	case common.BirthdayLunar:
		return convertForBirthdayLunar(event, language, recurCnt, host)
	default:
		return nil
	}
}

func writeResponse(writer http.ResponseWriter, response string) {
	writer.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	if _, err := fmt.Fprintln(writer, response); err != nil {
		log.Error("write HTTP response error: %v", err.Error())
	}
}

func writeResponse400(writer http.ResponseWriter, response string) {
	writer.WriteHeader(http.StatusBadRequest)
	writer.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	if _, err := fmt.Fprintln(writer, response); err != nil {
		log.Error("write HTTP response error: %v", err.Error())
	}
}
