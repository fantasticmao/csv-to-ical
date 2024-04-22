package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/fantasticmao/csv-to-ical/common"
	"io"
	"os"
	"strconv"
	"strings"
)

type Event struct {
	Name         string
	Month        int
	Day          int
	Year         int
	CalendarType common.CalendarType
}

func ParseEventFromFile(csvFile string) ([]Event, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}

	return parseEvent(file)
}

func ParseEventFromUrl(csvUrl string) ([]Event, error) {
	resp, err := common.HttpGet(csvUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return parseEvent(resp.Body)
}

func parseEvent(reader io.Reader) ([]Event, error) {
	csvReader := csv.NewReader(reader)

	var nameIdx, monthIdx, dayIdx, yearIdx, calTypeIdx int
	var events []Event
	for i := 0; ; i++ {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if i == 0 {
			for idx, v := range record {
				switch strings.ToLower(v) {
				case "name":
					nameIdx = idx
				case "month":
					monthIdx = idx
				case "day":
					dayIdx = idx
				case "year":
					yearIdx = idx
				case "calendar_type", "calendarType":
					calTypeIdx = idx
				}
			}
		} else {
			event, err := newEvent(record[nameIdx], record[monthIdx], record[dayIdx],
				record[yearIdx], record[calTypeIdx])
			if err != nil {
				return nil, fmt.Errorf("new csv event error at line: '%v' in record: '%v', cause by %v",
					i, record, err.Error())
			}

			events = append(events, *event)
		}
	}
	return events, nil
}

func newEvent(name, monthStr, dayStr, yearStr, calTypeStr string) (*Event, error) {
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

	calType, err := common.ParseCalendarType(calTypeStr)
	if err != nil {
		return nil, err
	}

	return &Event{
		Name:         name,
		Month:        month,
		Day:          day,
		Year:         year,
		CalendarType: calType,
	}, nil
}
