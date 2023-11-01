package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSolarToLunar(t *testing.T) {
	type args struct {
		sYear  int
		sMonth int
		sDay   int
	}
	tests := []struct {
		name       string
		args       args
		wantLYear  int
		wantLMonth int
		wantLDay   int
	}{
		{"2023-10-31", args{2023, 10, 31}, 2023, 9, 17},
		{"2077-01-01", args{2077, 1, 1}, 2076, 12, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLYear, gotLMonth, gotLDay := SolarToLunar(tt.args.sYear, tt.args.sMonth, tt.args.sDay)
			assert.Equalf(t, tt.wantLYear, gotLYear, "SolarToLunar(%v, %v, %v)", tt.args.sYear, tt.args.sMonth, tt.args.sDay)
			assert.Equalf(t, tt.wantLMonth, gotLMonth, "SolarToLunar(%v, %v, %v)", tt.args.sYear, tt.args.sMonth, tt.args.sDay)
			assert.Equalf(t, tt.wantLDay, gotLDay, "SolarToLunar(%v, %v, %v)", tt.args.sYear, tt.args.sMonth, tt.args.sDay)
		})
	}
}

func TestCalcAge(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
		now   time.Time
	}

	// 2023-10-31
	now := time.UnixMilli(1698681600000)
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2022-10-31", args{2022, 10, 30, now}, 2},
		{"2022-11-1", args{2022, 11, 1, now}, 1},
		{"2023-10-31", args{2023, 10, 30, now}, 1},
		{"2077-01-01", args{2077, 1, 1, now}, -53},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CalcAge(tt.args.year, tt.args.month, tt.args.day, tt.args.now),
				"CalcAge(%v, %v, %v, %v)", tt.args.year, tt.args.month, tt.args.day, tt.args.now)
		})
	}
}

func TestCalcLunarAge(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
		now   time.Time
	}

	// 2023-10-31
	now := time.UnixMilli(1698681600000)
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2022-10-31", args{2022, 10, 30, now}, 2},
		{"2022-11-1", args{2022, 11, 1, now}, 2},
		{"2023-10-31", args{2023, 10, 31, now}, 1},
		{"2077-01-01", args{2077, 1, 1, now}, -53},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CalcLunarAge(tt.args.year, tt.args.month, tt.args.day, tt.args.now),
				"CalcLunarAge(%v, %v, %v, %v)", tt.args.year, tt.args.month, tt.args.day, tt.args.now)
		})
	}
}

func TestParseTime(t *testing.T) {
	datetime, err := ParseTime("20231031")
	assert.Nil(t, err)

	assert.Equal(t, 2023, datetime.Year())
	assert.Equal(t, time.October, datetime.Month())
	assert.Equal(t, 31, datetime.Day())
}

func TestFormatTime(t *testing.T) {
	datetime, err := time.Parse("2006-01-02", "2023-10-31")
	assert.Nil(t, err)

	timeStr := FormatTime(datetime)
	assert.Equal(t, "20231031", timeStr)
}
