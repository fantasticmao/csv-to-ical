package app

import (
	"fmt"
	"github.com/fantasticmao/csv-to-ical/common"
	"github.com/fantasticmao/csv-to-ical/csv"
	"github.com/fantasticmao/csv-to-ical/date"
	"github.com/fantasticmao/csv-to-ical/ical"
	"time"
)

func convertForSolar(event csv.Event, language common.Language, recurCnt int, host string) []ical.ComponentEvent {
	now := time.Now()
	// FIXME 是否需要回溯过往年份？
	startTime := date.NewDate(now.Year(), event.Month, event.Day)

	summary := event.Name
	uid := ical.FormatUid(event.Name, startTime, event.CalendarType, host)
	cmpEvent := ical.NewComponentEvent(uid, language, summary, recurCnt, now, startTime)
	return []ical.ComponentEvent{cmpEvent}
}

func convertForLunar(event csv.Event, language common.Language, recurCnt int, host string) []ical.ComponentEvent {
	now := time.Now()
	var cmpEvents []ical.ComponentEvent
	for i := 0; i < recurCnt; i++ {
		// FIXME 是否需要回溯过往年份？
		startTime := date.LunarToSolar(now.Year(), event.Month, event.Day, i)

		summary := event.Name
		uid := ical.FormatUid(event.Name, startTime, event.CalendarType, host)
		cmpEvent := ical.NewComponentEvent(uid, language, summary, 0, now, startTime)
		cmpEvents = append(cmpEvents, cmpEvent)
	}
	return cmpEvents
}

func convertForBirthdaySolar(event csv.Event, language common.Language, recurCnt int, host string) []ical.ComponentEvent {
	now := time.Now()
	var cmpEvents []ical.ComponentEvent
	for i := 0; i < recurCnt; i++ {
		// FIXME 是否需要回溯过往年份？
		startTime := date.NewDate(now.Year()+i, event.Month, event.Day)

		var summary string
		if event.Year > 0 {
			age := date.CalcAge(event.Year, event.Month, event.Day, startTime)
			summary = fmt.Sprintf(common.SummaryMap[language][event.CalendarType][true], event.Name, age)
		} else {
			summary = fmt.Sprintf(common.SummaryMap[language][event.CalendarType][false], event.Name)
		}
		uid := ical.FormatUid(event.Name, startTime, event.CalendarType, host)
		cmpEvent := ical.NewComponentEvent(uid, language, summary, 0, now, startTime)
		cmpEvents = append(cmpEvents, cmpEvent)
	}
	return cmpEvents
}

func convertForBirthdayLunar(event csv.Event, language common.Language, recurCnt int, host string) []ical.ComponentEvent {
	now := time.Now()
	var cmpEvents []ical.ComponentEvent
	for i := 0; i < recurCnt; i++ {
		// FIXME 是否需要回溯过往年份？
		startTime := date.LunarToSolar(now.Year(), event.Month, event.Day, i)

		var summary string
		if event.Year > 0 {
			age := date.CalcLunarAge(event.Year, startTime)
			summary = fmt.Sprintf(common.SummaryMap[language][event.CalendarType][true], event.Name, age)
		} else {
			summary = fmt.Sprintf(common.SummaryMap[language][event.CalendarType][false], event.Name)
		}
		uid := ical.FormatUid(event.Name, startTime, event.CalendarType, host)
		cmpEvent := ical.NewComponentEvent(uid, language, summary, 0, now, startTime)
		cmpEvents = append(cmpEvents, cmpEvent)
	}
	return cmpEvents
}
