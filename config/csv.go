package config

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/url"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Event struct {
	name  string     `yaml:"name"`
	month time.Month `yaml:"month"`
	day   int        `yaml:"day"`
	year  int        `yaml:"year"`
	tag   Tag        `yaml:"tag"`
}

func newEvent(name, month, day, year, calendarType string) Event {
	m, err := strconv.Atoi(month)
	if err != nil {
		panic(err)
	}

	d, err := strconv.Atoi(day)
	if err != nil {
		panic(err)
	}

	var y int
	if strings.TrimSpace(year) == "" {
		y = -1
	} else {
		y, err = strconv.Atoi(year)
		if err != nil {
			panic(err)
		}
	}

	event := Event{
		name:  name,
		month: time.Month(m),
		day:   d,
		year:  y,
		tag:   Tag(calendarType),
	}
	return event
}

var legalSchemes = []string{"file", "http", "https"}

func ParseEventFromUrl(csvUrl string) []Event {
	u, err := url.Parse(csvUrl)
	if err != nil {
		panic(err)
	}

	existed := slices.Contains(legalSchemes, u.Scheme)
	if !existed {
		panic(fmt.Errorf("illegal URL sheme: %v", u.Scheme))
	}

	return nil
}

func ParseEventFromFile(csvFile string) []Event {
	data, err := os.ReadFile(csvFile)
	if err != nil {
		panic(err)
	}

	var nameIndex, monthIndex, dayIndex, yearIndex, calendarTypeIndex int
	reader := csv.NewReader(bytes.NewReader(data))
	events := make([]Event, 32)
	for i := 0; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if i == 0 {
			for i, v := range record {
				switch v {
				case "name":
					nameIndex = i
				case "month":
					monthIndex = i
				case "day":
					dayIndex = i
				case "year":
					yearIndex = i
				case "calendar_type", "calendarType":
					calendarTypeIndex = i
				}
			}
		} else {
			event := newEvent(record[nameIndex], record[monthIndex], record[dayIndex],
				record[yearIndex], record[calendarTypeIndex])
			events = append(events, event)
		}
	}
	return events
}
