package i18n

import (
	"github.com/fantasticmao/csv-to-ical/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_namingTemplate(t *testing.T) {
	name := namingTemplate(common.En, common.BirthdaySolar, true)
	assert.Equal(t, "en:birthday_solar:summary_with_age", name)

	name = namingTemplate(common.ZhCn, common.BirthdayLunar, false)
	assert.Equal(t, "zh_CN:birthday_lunar:summary", name)
}

func TestSummary(t *testing.T) {
	summary, err := Summary(common.En, common.BirthdaySolar, "Tom", 18)
	assert.Nil(t, err)
	assert.Equal(t, "Tom's 18th Birthday", summary)

	summary, err = Summary(common.ZhCn, common.BirthdayLunar, "小明", 0)
	assert.Nil(t, err)
	assert.Equal(t, "小明的农历生日", summary)

	summary, err = Summary(common.ZhCn, common.BirthdayLunar, "小明", 24)
	assert.Nil(t, err)
	assert.Equal(t, "小明的 24 岁农历生日", summary)
}
