package config

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

type CalendarType string

const (
	Solar         CalendarType = "solar"
	Lunar         CalendarType = "lunar"
	BirthdaySolar CalendarType = "birthday_solar"
	BirthdayLunar CalendarType = "birthday_lunar"
)

var SummaryMap = make(map[Language]map[CalendarType]map[bool]string)

func init() {
	SummaryMap[En] = make(map[CalendarType]map[bool]string)

	SummaryMap[En][BirthdaySolar] = make(map[bool]string)
	SummaryMap[En][BirthdaySolar][true] = "%s's %dth Birthday"
	SummaryMap[En][BirthdaySolar][false] = "%s's Birthday"

	SummaryMap[En][BirthdayLunar] = make(map[bool]string)
	SummaryMap[En][BirthdayLunar][true] = "%s's %dth Chinese Birthday"
	SummaryMap[En][BirthdayLunar][false] = "%s's Chinese Birthday"

	SummaryMap[ZhCn] = make(map[CalendarType]map[bool]string)

	SummaryMap[ZhCn][BirthdaySolar] = make(map[bool]string)
	SummaryMap[ZhCn][BirthdaySolar][true] = "%s 的 %d 岁生日"
	SummaryMap[ZhCn][BirthdaySolar][false] = "%s 的生日"

	SummaryMap[ZhCn][BirthdayLunar] = make(map[bool]string)
	SummaryMap[ZhCn][BirthdayLunar][true] = "%s 的 %d 岁农历生日"
	SummaryMap[ZhCn][BirthdayLunar][false] = "%s 的农历生日"
}
