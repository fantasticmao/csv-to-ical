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

type stuff struct {
	Summary        string `yaml:"summary"`
	SummaryWithAge string `yaml:"summary_with_age"`
}

type stuffs map[string]stuff

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

		language, err := parseLanguage(fileName)
		if err != nil {
			panic(err)
		}

		stuffs, err := parseStuffs(fileName)
		if err != nil {
			panic(err)
		}

		for calTypeStr, stuff := range *stuffs {
			calType, err := common.ParseCalendarType(calTypeStr)
			if err != nil {
				panic(err)
			}

			name := namingTemplate(language, calType, false)
			template.Must(t.New(name).Parse(stuff.Summary))

			name = namingTemplate(language, calType, true)
			template.Must(t.New(name).Parse(stuff.SummaryWithAge))
		}
	}
}

func parseLanguage(fileName string) (common.Language, error) {
	lang := fileName[:len(fileName)-len(".yaml")]
	return common.ParseLanguage(lang)
}

func parseStuffs(fileName string) (*stuffs, error) {
	if data, err := i18nFs.ReadFile(fileName); err != nil {
		return nil, err
	} else {
		stuffs := &stuffs{}
		if err = yaml.Unmarshal(data, stuffs); err != nil {
			return nil, err
		} else {
			return stuffs, err
		}
	}
}

func namingTemplate(language common.Language, calType common.CalendarType, withAge bool) string {
	return fmt.Sprintf("%v:%v:%v", language, calType, withAge)
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
