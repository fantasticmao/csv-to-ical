package i18n

import (
	"github.com/fantasticmao/csv-to-ical/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_namingTemplate(t *testing.T) {
	name := namingTemplate(common.En, common.BirthdaySolar, true)
	assert.Equal(t, "en:birthday_solar:true", name)

	name = namingTemplate(common.ZhCn, common.BirthdayLunar, false)
	assert.Equal(t, "zh_CN:birthday_lunar:false", name)
}

func TestSummary(t *testing.T) {
	type args struct {
		language common.Language
		calType  common.CalendarType
		name     string
		age      int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"en-solar-age", args{common.En, common.BirthdaySolar, "Tom", 18}, "Tom's 18th Birthday"},
		{"en-solar", args{common.En, common.BirthdaySolar, "Tom", 0}, "Tom's Birthday"},
		{"en-lunar-age", args{common.En, common.BirthdayLunar, "Tom", 18}, "Tom's 18th Chinese Birthday"},
		{"en-lunar", args{common.En, common.BirthdayLunar, "Tom", 0}, "Tom's Chinese Birthday"},
		{"zh_CN-solar-age", args{common.ZhCn, common.BirthdaySolar, "小明", 24}, "小明的 24 岁生日"},
		{"zh_CN-solar", args{common.ZhCn, common.BirthdaySolar, "小明", 0}, "小明的生日"},
		{"zh_CN-lunar-age", args{common.ZhCn, common.BirthdayLunar, "小明", 24}, "小明的 24 岁农历生日"},
		{"zh_CN-lunar", args{common.ZhCn, common.BirthdayLunar, "小明", 0}, "小明的农历生日"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			summary, err := Summary(tt.args.language, tt.args.calType, tt.args.name, tt.args.age)
			assert.Nil(t, err)
			assert.Equalf(t, tt.want, summary, "Summary(%v, %v, %v, %v)", tt.args.language, tt.args.calType, tt.args.name, tt.args.age)
		})
	}
}
