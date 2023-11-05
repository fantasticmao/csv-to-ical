package config

type CalendarType string

const (
	Solar         CalendarType = "solar"
	Lunar         CalendarType = "lunar"
	BirthdaySolar CalendarType = "birthday_solar"
	BirthdayLunar CalendarType = "birthday_lunar"
)

type Language string

const (
	En   Language = "en"
	ZhCn Language = "zh_CN"
)
