package config

type Tag string

const (
	Solar         Tag = "solar"
	Lunar         Tag = "lunar"
	BirthdaySolar Tag = "birthday_solar"
	BirthdayLunar Tag = "birthday_lunar"
)

type Language string

const (
	En   Language = "en"
	ZhCn Language = "zh_CN"
)
