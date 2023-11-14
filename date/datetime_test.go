package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddLunarYears(t *testing.T) {
	type args struct {
		date  time.Time
		years int
	}
	tests := []struct {
		name     string
		args     args
		wantDate time.Time
	}{
		{"2023-10-31+0", args{NewDate(2023, 10, 31), 0}, NewDate(2023, 10, 31)},
		{"2023-10-31+1", args{NewDate(2023, 10, 31), 1}, NewDate(2024, 10, 19)},
		{"2023-10-31+2", args{NewDate(2023, 10, 31), 2}, NewDate(2025, 11, 6)},
		{"2023-10-31+3", args{NewDate(2023, 10, 31), 3}, NewDate(2026, 10, 26)},
		{"2023-10-31+4", args{NewDate(2023, 10, 31), 4}, NewDate(2027, 10, 16)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDate := AddLunarYears(tt.args.date, tt.args.years)
			assert.Equalf(t, tt.wantDate, gotDate, "AddLunarYears(%v, %v)", tt.args.date, tt.args.years)
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
		year int
		now  time.Time
	}

	// 2023-10-31
	now := time.UnixMilli(1698681600000)
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2022-10-31", args{2022, now}, 2},
		{"2022-11-1", args{2022, now}, 2},
		{"2023-10-31", args{2023, now}, 1},
		{"2077-01-01", args{2077, now}, -53},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CalcLunarAge(tt.args.year, tt.args.now),
				"CalcLunarAge(%v, %v)", tt.args.year, tt.args.now)
		})
	}
}

func TestNewDate(t *testing.T) {
	datetime := NewDate(2023, 10, 31)

	assert.Equal(t, 2023, datetime.Year())
	assert.Equal(t, time.October, datetime.Month())
	assert.Equal(t, 31, datetime.Day())
}

func TestFormatTime(t *testing.T) {
	datetime, err := time.Parse("2006-01-02", "2023-10-31")
	assert.Nil(t, err)

	timeStr := FormatDate(datetime)
	assert.Equal(t, "20231031", timeStr)
}
