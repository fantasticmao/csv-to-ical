package app

import (
	"github.com/fantasticmao/csv-to-ical/common"
	"github.com/fantasticmao/csv-to-ical/csv"
	"github.com/fantasticmao/csv-to-ical/ical"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"runtime"
	"strconv"
)

func HomeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, `If you see this message, the %v is successfully installed and working.
For more informations see https://github.com/fantasticmao/csv-to-ical
`, common.Name)
	}
}

func VersionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "%v %v %v-%v with %v built on commit %v at %v\n",
			common.Name, common.Version, runtime.GOOS, runtime.GOARCH, runtime.Version(),
			common.CommitHash, common.BuildTime)
	}
}

func RemoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Query("url")
		if url == "" {
			c.String(http.StatusBadRequest, "query param: url is required")
			return
		}

		languageParam := c.DefaultQuery("lang", string(common.En))
		language, err := common.ParseLanguage(languageParam)
		if err != nil {
			c.String(http.StatusBadRequest, "query param: lang error: %v", err.Error())
			return
		}

		recurCntParam := c.DefaultQuery("recurCnt", "5")
		recurCnt, err := strconv.Atoi(recurCntParam)
		if err != nil {
			c.String(http.StatusBadRequest, "query param: recurCnt error: %v", err.Error())
			return
		} else {
			if recurCnt < 0 {
				c.String(http.StatusBadRequest, "query param: recurCnt cannot be negative")
				return
			} else if recurCnt > 10 {
				c.String(http.StatusBadRequest, "query param: recurCnt cannot be grater than 10")
				return
			}
		}

		events, err := csv.ParseEventFromUrl(url)
		if err != nil {
			c.String(http.StatusBadRequest, "fetch csv events from url: '%v' error: %v", url, err.Error())
			return
		}

		var components []ical.ComponentEvent
		for _, event := range events {
			cmpEvents := csvToIcal(event, language, recurCnt, c.Request.Host)
			components = append(components, cmpEvents...)
		}

		obj := ical.NewObject(components)
		if response, err := obj.Transform(); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.String(http.StatusOK, response)
		}
	}
}

func LocalHandler(configDir string, provider common.CsvProvider) (gin.HandlerFunc, error) {
	var events []csv.Event
	if provider.File != "" {
		e, err := csv.ParseEventFromFile(path.Join(configDir, provider.File))
		if err != nil {
			return nil, err
		}
		events = e
	}
	if provider.Url != "" {
		e, err := csv.ParseEventFromUrl(provider.Url)
		if err != nil {
			return nil, err
		}
		events = e
	}

	return func(c *gin.Context) {
		var components []ical.ComponentEvent
		for _, event := range events {
			language, _ := common.ParseLanguage(provider.Language)
			cmpEvents := csvToIcal(event, language, provider.RecurCnt, c.Request.Host)
			components = append(components, cmpEvents...)
		}

		obj := ical.NewObject(components)
		if response, err := obj.Transform(); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.String(http.StatusOK, response)
		}
	}, nil
}

func csvToIcal(event csv.Event, language common.Language, recurCnt int, host string) []ical.ComponentEvent {
	switch event.CalendarType {
	case common.Solar:
		return convertForSolar(event, language, recurCnt, host)
	case common.Lunar:
		return convertForLunar(event, language, recurCnt, host)
	case common.BirthdaySolar:
		return convertForBirthdaySolar(event, language, recurCnt, host)
	case common.BirthdayLunar:
		return convertForBirthdayLunar(event, language, recurCnt, host)
	default:
		return nil
	}
}
