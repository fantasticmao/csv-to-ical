package app

import (
	"github.com/fantasticmao/csv-to-ical/common"
	"github.com/fantasticmao/csv-to-ical/csv"
	"github.com/fantasticmao/csv-to-ical/i18n"
	"github.com/fantasticmao/csv-to-ical/ical"
	"time"
)

func convertForSolar(event csv.Event, language common.Language, recurCnt, backCnt int, host string) []ical.ComponentEvent {
	now := time.Now()
	startDate := common.NewDate(now.Year()-backCnt, event.Month, event.Day)

	summary := event.Name
	uid := ical.FormatUid(event.Name, startDate, event.CalendarType, host)
	cmpEvent := ical.NewComponentEvent(uid, language, summary, recurCnt+backCnt, now, startDate)
	return []ical.ComponentEvent{cmpEvent}
}

func convertForLunar(event csv.Event, language common.Language, recurCnt, backCnt int, host string) []ical.ComponentEvent {
	now := time.Now()
	lunarYear, _, _, _ := common.SolarToLunar(now)
	var cmpEvents []ical.ComponentEvent
	for i := -backCnt; i < recurCnt; i++ {
		startDate := common.LunarToSolar(lunarYear, event.Month, event.Day, i)

		summary := event.Name
		uid := ical.FormatUid(event.Name, startDate, event.CalendarType, host)
		cmpEvent := ical.NewComponentEvent(uid, language, summary, 0, now, startDate)
		cmpEvents = append(cmpEvents, cmpEvent)
	}
	return cmpEvents
}

func convertForBirthdaySolar(event csv.Event, language common.Language, recurCnt, backCnt int, host string) []ical.ComponentEvent {
	now := time.Now()
	var cmpEvents []ical.ComponentEvent
	for i := -backCnt; i < recurCnt; i++ {
		startDate := common.NewDate(now.Year()+i, event.Month, event.Day)

		var summary string
		var err error
		if event.Year > 0 {
			age := common.CalcAge(event.Year, event.Month, event.Day, startDate)
			if age < 0 {
				continue
			}
			summary, err = i18n.Summary(language, event.CalendarType, event.Name, age)
		} else {
			summary, err = i18n.Summary(language, event.CalendarType, event.Name, -1)
		}
		if err != nil {
			// FIXME 兼容错误处理
			summary = err.Error()
		}

		uid := ical.FormatUid(event.Name, startDate, event.CalendarType, host)
		cmpEvent := ical.NewComponentEvent(uid, language, summary, 0, now, startDate)
		cmpEvents = append(cmpEvents, cmpEvent)
	}
	return cmpEvents
}

func convertForBirthdayLunar(event csv.Event, language common.Language, recurCnt, backCnt int, host string) []ical.ComponentEvent {
	now := time.Now()
	lunarYear, _, _, _ := common.SolarToLunar(now)
	var cmpEvents []ical.ComponentEvent
	for i := -backCnt; i < recurCnt; i++ {
		startDate := common.LunarToSolar(lunarYear, event.Month, event.Day, i)

		var summary string
		var err error
		if event.Year > 0 {
			age := common.CalcLunarAge(event.Year, startDate)
			if age < 1 {
				continue
			}
			summary, err = i18n.Summary(language, event.CalendarType, event.Name, age)
		} else {
			summary, err = i18n.Summary(language, event.CalendarType, event.Name, -1)
		}
		if err != nil {
			// FIXME 兼容错误处理
			summary = err.Error()
		}

		uid := ical.FormatUid(event.Name, startDate, event.CalendarType, host)
		cmpEvent := ical.NewComponentEvent(uid, language, summary, 0, now, startDate)
		cmpEvents = append(cmpEvents, cmpEvent)
	}
	return cmpEvents
}
