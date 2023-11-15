package date

import (
	"github.com/isee15/Lunar-Solar-Calendar-Converter/Go/lunarsolar"
	"time"
)

// LunarToSolar 农历转公历，并新增年份
func LunarToSolar(year, month, day, addYears int) time.Time {
	lunar := lunarsolar.Lunar{
		IsLeap:     false,
		LunarYear:  year + addYears,
		LunarMonth: month,
		LunarDay:   day,
	}
	solar := lunarsolar.LunarToSolar(lunar)
	return NewDate(solar.SolarYear, solar.SolarMonth, solar.SolarDay)
}

// CalcAge 计算周岁：出生是零岁，每过一个生日就长一岁
func CalcAge(year, month, day int, now time.Time) int {
	age := now.Year() - year - 1
	if nowMonth := int(now.Month()); nowMonth > month {
		age++
	} else if nowMonth == month {
		if now.Day() >= day {
			age++
		}
	}
	return age
}

// CalcLunarAge 计算虚岁：出生是一岁，每过一个春节就长一岁
func CalcLunarAge(year int, now time.Time) int {
	nowLunar := lunarsolar.SolarToLunar(lunarsolar.Solar{
		SolarYear:  now.Year(),
		SolarMonth: int(now.Month()),
		SolarDay:   now.Day(),
	})
	return nowLunar.LunarYear - year + 1
}

func NewDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func FormatDate(time time.Time) string {
	return time.Format("20060102")
}
