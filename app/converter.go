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
	summary := event.Name
	// FIXME 是否需要回溯过往年份？
	startTime := date.NewDate(time.Now().Year(), event.Month, event.Day)
	uid := ical.FormatUid(event.Name, startTime, event.CalendarType, host)
	cmpEvent := ical.NewComponentEvent(uid, lang, summary, recurCount, time.Now(), startTime)
	return []ical.ComponentEvent{cmpEvent}
}

func convertForLunar(event csv.Event, lang config.Language, recurCount int, host string) []ical.ComponentEvent {
	var cmpEvents []ical.ComponentEvent
	for i := 0; i < recurCount; i++ {
		summary := event.Name
		// FIXME 是否需要回溯过往年份？
		startTime := date.NewDate(time.Now().Year(), event.Month, event.Day)
		startTime = date.AddLunarYears(startTime, i)
		uid := ical.FormatUid(event.Name, startTime, event.CalendarType, host)
		cmpEvent := ical.NewComponentEvent(uid, lang, summary, 0, time.Now(), startTime)
		cmpEvents = append(cmpEvents, cmpEvent)
	}
	return cmpEvents
}

func convertForBirthdaySolar(event csv.Event, lang config.Language, recurCount int, host string) []ical.ComponentEvent {
	var cmpEvents []ical.ComponentEvent
	for i := 0; i < recurCount; i++ {
		var summary string
		if event.Year > 0 {
			age := date.CalcAge(event.Year, event.Month, event.Day, time.Now())
			summary = fmt.Sprintf(config.SummaryMap[lang][event.CalendarType][true], event.Name, age)
		} else {
			summary = fmt.Sprintf(config.SummaryMap[lang][event.CalendarType][false], event.Name)
		}
		// FIXME 是否需要回溯过往年份？
		startTime := date.NewDate(time.Now().Year()+i, event.Month, event.Day)
		uid := ical.FormatUid(event.Name, startTime, event.CalendarType, host)
		cmpEvent := ical.NewComponentEvent(uid, lang, summary, 0, time.Now(), startTime)
		cmpEvents = append(cmpEvents, cmpEvent)
	}
	return cmpEvents
}

func convertForBirthdayLunar(event csv.Event, lang config.Language, recurCount int, host string) []ical.ComponentEvent {
	var cmpEvents []ical.ComponentEvent
	for i := 0; i < recurCount; i++ {
		var summary string
		if event.Year > 0 {
			age := date.CalcLunarAge(event.Year, time.Now())
			summary = fmt.Sprintf(config.SummaryMap[lang][event.CalendarType][true], event.Name, age)
		} else {
			summary = fmt.Sprintf(config.SummaryMap[lang][event.CalendarType][false], event.Name)
		}
		// FIXME 是否需要回溯过往年份？
		startTime := date.NewDate(time.Now().Year(), event.Month, event.Day)
		startTime = date.AddLunarYears(startTime, i)
		uid := ical.FormatUid(event.Name, startTime, event.CalendarType, host)
		cmpEvent := ical.NewComponentEvent(uid, lang, summary, 0, time.Now(), startTime)
		cmpEvents = append(cmpEvents, cmpEvent)
	}
	return cmpEvents
}
