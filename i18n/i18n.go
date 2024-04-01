package i18n

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/fantasticmao/csv-to-ical/common"
	"gopkg.in/yaml.v3"
	"text/template"
)

var t = template.New("i18n")

func namingTemplate(language common.Language, calType common.CalendarType, addAge bool) string {
	if addAge {
		return fmt.Sprintf("%v:%v:summary_with_age", language, calType)
	} else {
		return fmt.Sprintf("%v:%v:summary", language, calType)
	}
}

func Summary(language common.Language, calType common.CalendarType, name string, age int) (string, error) {
	data := map[string]interface{}{
		"Name": name,
		"Age":  age,
	}

	output := &bytes.Buffer{}
	templateName := namingTemplate(language, calType, age > 0)
	if err := t.ExecuteTemplate(output, templateName, data); err != nil {
		return "", err
	}
	return output.String(), nil
}

//go:embed *.yaml
var i18nFs embed.FS

func init() {
	entries, err := i18nFs.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			panic(err)
		}
		if info.IsDir() {
			continue
		}

		fileName := info.Name()
		lang := fileName[:len(fileName)-len(".yaml")]
		language, err := common.ParseLanguage(lang)
		if err != nil {
			panic(err)
		}

		data, err := i18nFs.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		summaryMap := make(map[string]map[string]string)
		if err = yaml.Unmarshal(data, summaryMap); err != nil {
			panic(err)
		}

		for calTypeStr, subMap := range summaryMap {
			calType, err := common.ParseCalendarType(calTypeStr)
			if err != nil {
				panic(err)
			}

			for key, val := range subMap {
				var name string
				if key == "summary" {
					name = namingTemplate(language, calType, false)
				} else if key == "summary_with_age" {
					name = namingTemplate(language, calType, true)
				}
				template.Must(t.New(name).Parse(val))
			}
		}
	}
}
