package config

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	http.DefaultClient.Transport = &http.Transport{
		Proxy: func(*http.Request) (*url.URL, error) {
			proxy, err := url.Parse("http://127.0.0.1:7890")
			if err != nil {
				return nil, err
			}
			return proxy, nil
		},
	}
}

type Event struct {
	name         string
	month        time.Month
	day          int
	year         int
	calendarType CalendarType
}

func newEvent(name, monthStr, dayStr, yearStr, calType string) (*Event, error) {
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return nil, err
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return nil, err
	}

	var year int
	if strings.TrimSpace(yearStr) == "" {
		year = -1
	} else {
		y, err := strconv.Atoi(yearStr)
		if err != nil {
			return nil, err
		} else {
			year = y
		}
	}

	calendarType := CalendarType(calType)

	return &Event{
		name:         name,
		month:        time.Month(month),
		day:          day,
		year:         year,
		calendarType: calendarType,
	}, nil
}
func ParseEventFromFile(csvFile string) ([]Event, error) {
	data, err := os.ReadFile(csvFile)
	if err != nil {
		return nil, err
	}

	return parseEvent(data)
}

func ParseEventFromUrl(csvUrl string) ([]Event, error) {
	resp, err := http.Get(csvUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return parseEvent(data)
}

func parseEvent(data []byte) ([]Event, error) {
	reader := csv.NewReader(bytes.NewReader(data))

	var nameIdx, monthIdx, dayIdx, yearIdx, calTypeIdx int
	var events []Event
	for i := 0; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if i == 0 {
			for j, v := range record {
				switch v {
				case "name":
					nameIdx = j
				case "month":
					monthIdx = j
				case "day":
					dayIdx = j
				case "year":
					yearIdx = j
				case "calendar_type":
					calTypeIdx = j
				}
			}
		} else {
			event, err := newEvent(record[nameIdx], record[monthIdx], record[dayIdx],
				record[yearIdx], record[calTypeIdx])
			if err != nil {
				return nil, fmt.Errorf("new Event error in line: %v record: %v, cause by %v",
					i, record, err.Error())
			}

			events = append(events, *event)
		}
	}
	return events, nil
}
