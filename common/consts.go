package common

import (
	"fmt"
	"strings"
)

const (
	Name     = "csv-to-ical"
	FullName = "github.com/fantasticmao/" + Name
)

var (
	Version    = "unknown version"
	BuildTime  = "unknown time"
	CommitHash = "unknown commit"
)

type Language string

const (
	En   Language = "en"
	ZhCn Language = "zh_CN"
)

func ParseLanguage(language string) (Language, error) {
	switch strings.ToLower(language) {
	case "en":
		return En, nil
	case "zh-cn", "zh_cn":
		return ZhCn, nil
	default:
		return "", fmt.Errorf("unsupported language: '%v'", language)
	}
}

type CalendarType string

const (
	Solar         CalendarType = "solar"
	Lunar         CalendarType = "lunar"
	BirthdaySolar CalendarType = "birthday_solar"
	BirthdayLunar CalendarType = "birthday_lunar"
)

func ParseCalendarType(calendarType string) (CalendarType, error) {
	switch strings.ToLower(calendarType) {
	case "solar":
		return Solar, nil
	case "lunar":
		return Lunar, nil
	case "birthday_solar":
		return BirthdaySolar, nil
	case "birthday_lunar":
		return BirthdayLunar, nil
	default:
		return "", fmt.Errorf("unsupported calendar type: '%v'", calendarType)
	}
}
