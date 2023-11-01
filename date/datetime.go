package date

import (
	"github.com/isee15/Lunar-Solar-Calendar-Converter/Go/lunarsolar"
	"time"
)

// SolarToLunar 公历转农历
func SolarToLunar(sYear, sMonth, sDay int) (lYear, lMonth, lDay int) {
	solar := lunarsolar.Solar{
		SolarYear:  sYear,
		SolarMonth: sMonth,
		SolarDay:   sDay,
	}
	lunar := lunarsolar.SolarToLunar(solar)
	return lunar.LunarYear, lunar.LunarMonth, lunar.LunarDay
}

// CalcAge 计算周岁
func CalcAge(year, month, day int, now time.Time) int {
	age := now.Year() - year
	if nowMonth := int(now.Month()); nowMonth > month {
		age++
	} else if nowMonth == month {
		if now.Day() >= day {
			age++
		}
	}
	return age
}

// CalcLunarAge 计算虚岁
func CalcLunarAge(year, month, day int, now time.Time) int {
	return now.Year() - year + 1
}

func ParseTime(value string) (time.Time, error) {
	return time.Parse("20060102", value)
}

func FormatTime(time time.Time) string {
	return time.Format("20060102")
}
