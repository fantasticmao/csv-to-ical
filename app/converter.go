package app

import (
	"fmt"
	"github.com/fantasticmao/csv-to-ical/config"
	"github.com/fantasticmao/csv-to-ical/csv"
	"github.com/fantasticmao/csv-to-ical/date"
	"github.com/fantasticmao/csv-to-ical/ical"
	"time"
)

func convertForSolar(event csv.Event, lang config.Language, recurCount int, host string) []ical.ComponentEvent {
	now := time.Now()
	// FIXME 是否需要回溯过往年份？
	startTime := date.NewDate(now.Year(), event.Month, event.Day)

	summary := event.Name
	uid := ical.FormatUid(event.Name, startTime, event.CalendarType, host)
	cmpEvent := ical.NewComponentEvent(uid, lang, summary, recurCount, now, startTime)
	return []ical.ComponentEvent{cmpEvent}
}

func convertForLunar(event csv.Event, lang config.Language, recurCount int, host string) []ical.ComponentEvent {
	now := time.Now()
	var cmpEvents []ical.ComponentEvent
	for i := 0; i < recurCount; i++ {
		// FIXME 是否需要回溯过往年份？
		startTime := date.NewDate(now.Year(), event.Month, event.Day)
		startTime = date.LunarToSolar(startTime, i)

		summary := event.Name
		uid := ical.FormatUid(event.Name, startTime, event.CalendarType, host)
		cmpEvent := ical.NewComponentEvent(uid, lang, summary, 0, now, startTime)
		cmpEvents = append(cmpEvents, cmpEvent)
	}
	return cmpEvents
}

func convertForBirthdaySolar(event csv.Event, lang config.Language, recurCount int, host string) []ical.ComponentEvent {
	now := time.Now()
	var cmpEvents []ical.ComponentEvent
	for i := 0; i < recurCount; i++ {
		// FIXME 是否需要回溯过往年份？
		startTime := date.NewDate(now.Year()+i, event.Month, event.Day)

		var summary string
		if event.Year > 0 {
			age := date.CalcAge(event.Year, event.Month, event.Day, startTime)
			summary = fmt.Sprintf(config.SummaryMap[lang][event.CalendarType][true], event.Name, age)
		} else {
			summary = fmt.Sprintf(config.SummaryMap[lang][event.CalendarType][false], event.Name)
		}
		uid := ical.FormatUid(event.Name, startTime, event.CalendarType, host)
		cmpEvent := ical.NewComponentEvent(uid, lang, summary, 0, now, startTime)
		cmpEvents = append(cmpEvents, cmpEvent)
	}
	return cmpEvents
}

func convertForBirthdayLunar(event csv.Event, lang config.Language, recurCount int, host string) []ical.ComponentEvent {
	now := time.Now()
	var cmpEvents []ical.ComponentEvent
	for i := 0; i < recurCount; i++ {
		// FIXME 是否需要回溯过往年份？
		startTime := date.NewDate(now.Year(), event.Month, event.Day)
		startTime = date.LunarToSolar(startTime, i)

		var summary string
		if event.Year > 0 {
			age := date.CalcLunarAge(event.Year, startTime)
			summary = fmt.Sprintf(config.SummaryMap[lang][event.CalendarType][true], event.Name, age)
		} else {
			summary = fmt.Sprintf(config.SummaryMap[lang][event.CalendarType][false], event.Name)
		}
		uid := ical.FormatUid(event.Name, startTime, event.CalendarType, host)
		cmpEvent := ical.NewComponentEvent(uid, lang, summary, 0, now, startTime)
		cmpEvents = append(cmpEvents, cmpEvent)
	}
	return cmpEvents
}
