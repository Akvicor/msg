package utils

import (
	"fmt"
	"msg/cmd/app/server/common/period"
	"strings"
	"testing"
	"time"
)

func TestFormatSecond(t *testing.T) {
	var sec int64
	var deep int
	var want string
	{
		sec = -1
		deep = 0
		want = "-1s"
		got := FormatSecond(sec, deep)
		if got != want {
			t.Errorf("FormatSecond(%d, %d) = %s; want %s", sec, deep, got, want)
		}
	}
	{
		sec = 0
		deep = 0
		want = "0s"
		got := FormatSecond(sec, deep)
		if got != want {
			t.Errorf("FormatSecond(%d, %d) = %s; want %s", sec, deep, got, want)
		}
	}
	{
		sec = 60
		deep = 0
		want = "1m"
		got := FormatSecond(sec, deep)
		if got != want {
			t.Errorf("FormatSecond(%d, %d) = %s; want %s", sec, deep, got, want)
		}
	}
	{
		sec = 61
		deep = 0
		want = "1m"
		got := FormatSecond(sec, deep)
		if got != want {
			t.Errorf("FormatSecond(%d, %d) = %s; want %s", sec, deep, got, want)
		}
	}
	{
		sec = 61
		deep = 1
		want = "1m"
		got := FormatSecond(sec, deep)
		if got != want {
			t.Errorf("FormatSecond(%d, %d) = %s; want %s", sec, deep, got, want)
		}
	}
	{
		sec = 61
		deep = 2
		want = "1m1s"
		got := FormatSecond(sec, deep)
		if got != want {
			t.Errorf("FormatSecond(%d, %d) = %s; want %s", sec, deep, got, want)
		}
	}
	{
		sec = 3600
		deep = 2
		want = "1h"
		got := FormatSecond(sec, deep)
		if got != want {
			t.Errorf("FormatSecond(%d, %d) = %s; want %s", sec, deep, got, want)
		}
	}
	{
		sec = 3601
		deep = 2
		want = "1h"
		got := FormatSecond(sec, deep)
		if got != want {
			t.Errorf("FormatSecond(%d, %d) = %s; want %s", sec, deep, got, want)
		}
	}
	{
		sec = 3661
		deep = 2
		want = "1h1m"
		got := FormatSecond(sec, deep)
		if got != want {
			t.Errorf("FormatSecond(%d, %d) = %s; want %s", sec, deep, got, want)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	var input int
	var want bool
	{
		input = 2000
		want = true
		got := IsLeapYear(input)
		if got != want {
			t.Errorf("IsLeapYear(%d) = %t, want %t", input, got, want)
		}
	}
	{
		input = 2007
		want = false
		got := IsLeapYear(input)
		if got != want {
			t.Errorf("IsLeapYear(%d) = %t, want %t", input, got, want)
		}
	}
	{
		input = 2008
		want = true
		got := IsLeapYear(input)
		if got != want {
			t.Errorf("IsLeapYear(%d) = %t, want %t", input, got, want)
		}
	}
	{
		input = 2200
		want = false
		got := IsLeapYear(input)
		if got != want {
			t.Errorf("IsLeapYear(%d) = %t, want %t", input, got, want)
		}
	}

}

func TestMonthOffsetToMonth(t *testing.T) {
	var input int
	var want int
	var month int
	{
		input = 0
		want = 1
		month = MonthOffsetToMonth(input)
		if month != want {
			t.Errorf("MonthOffsetToMonth(%d) = %d, want %d", input, month, want)
		}
	}
	{
		input = -1
		want = 12
		month = MonthOffsetToMonth(input)
		if month != want {
			t.Errorf("MonthOffsetToMonth(%d) = %d, want %d", input, month, want)
		}
	}
	{
		input = -2
		want = 11
		month = MonthOffsetToMonth(input)
		if month != want {
			t.Errorf("MonthOffsetToMonth(%d) = %d, want %d", input, month, want)
		}
	}
	{
		input = -12
		want = 1
		month = MonthOffsetToMonth(input)
		if month != want {
			t.Errorf("MonthOffsetToMonth(%d) = %d, want %d", input, month, want)
		}
	}
	{
		input = -13
		want = 1
		month = MonthOffsetToMonth(input)
		if month != want {
			t.Errorf("MonthOffsetToMonth(%d) = %d, want %d", input, month, want)
		}
	}
}

func TestDayOffsetToWeek(t *testing.T) {
	var input int
	var want int
	var week int
	{
		input = 0
		want = 0
		week = DayOffsetToWeek(input)
		if week != want {
			t.Errorf("DayOffsetToWeek(%d) = %d, want %d", input, week, want)
		}
	}
	{
		input = 1
		want = 1
		week = DayOffsetToWeek(input)
		if week != want {
			t.Errorf("DayOffsetToWeek(%d) = %d, want %d", input, week, want)
		}
	}
	{
		input = -1
		want = 6
		week = DayOffsetToWeek(input)
		if week != want {
			t.Errorf("DayOffsetToWeek(%d) = %d, want %d", input, week, want)
		}
	}
	{
		input = 8
		want = 6
		week = DayOffsetToWeek(input)
		if week != want {
			t.Errorf("DayOffsetToWeek(%d) = %d, want %d", input, week, want)
		}
	}
	{
		input = -8
		want = 0
		week = DayOffsetToWeek(input)
		if week != want {
			t.Errorf("DayOffsetToWeek(%d) = %d, want %d", input, week, want)
		}
	}
}

//func TestDayOffsetToDayByYear(t *testing.T) {
//	var inputYear int
//	var inputDay int
//	var want int
//	var got int
//	{
//		inputYear = 2024
//		inputDay = 1
//		want = 1
//		got = DayOffsetToDayByYear(inputYear, inputDay)
//		if got != want {
//			t.Errorf("DayOffsetToDay(%d,%d) = %d, want %d", inputYear, inputDay, got, want)
//		}
//	}
//	{
//		inputYear = 2025
//		inputDay = 1
//		want = 1
//		got = DayOffsetToDayByYear(inputYear, inputDay)
//		if got != want {
//			t.Errorf("DayOffsetToDay(%d,%d) = %d, want %d", inputYear, inputDay, got, want)
//		}
//	}
//	{
//		inputYear = 2024
//		inputDay = -1
//		want = 366
//		got = DayOffsetToDayByYear(inputYear, inputDay)
//		if got != want {
//			t.Errorf("DayOffsetToDay(%d,%d) = %d, want %d", inputYear, inputDay, got, want)
//		}
//	}
//	{
//		inputYear = 2025
//		inputDay = -1
//		want = 365
//		got = DayOffsetToDayByYear(inputYear, inputDay)
//		if got != want {
//			t.Errorf("DayOffsetToDay(%d,%d) = %d, want %d", inputYear, inputDay, got, want)
//		}
//	}
//	{
//		inputYear = 2024
//		inputDay = 367
//		want = 366
//		got = DayOffsetToDayByYear(inputYear, inputDay)
//		if got != want {
//			t.Errorf("DayOffsetToDay(%d,%d) = %d, want %d", inputYear, inputDay, got, want)
//		}
//	}
//	{
//		inputYear = 2025
//		inputDay = 367
//		want = 365
//		got = DayOffsetToDayByYear(inputYear, inputDay)
//		if got != want {
//			t.Errorf("DayOffsetToDay(%d,%d) = %d, want %d", inputYear, inputDay, got, want)
//		}
//	}
//}

func TestHourOffsetToHour(t *testing.T) {
	var inputHour int
	var want int
	var got int
	{
		inputHour = 1
		want = 1
		got = HourOffsetToHour(inputHour)
		if got != want {
			t.Errorf("HourOffsetToHour(%d) = %d, want %d", inputHour, got, want)
		}
	}
	{
		inputHour = -1
		want = 23
		got = HourOffsetToHour(inputHour)
		if got != want {
			t.Errorf("HourOffsetToHour(%d) = %d, want %d", inputHour, got, want)
		}
	}
	{
		inputHour = 0
		want = 0
		got = HourOffsetToHour(inputHour)
		if got != want {
			t.Errorf("HourOffsetToHour(%d) = %d, want %d", inputHour, got, want)
		}
	}
	{
		inputHour = 24
		want = 0
		got = HourOffsetToHour(inputHour)
		if got != want {
			t.Errorf("HourOffsetToHour(%d) = %d, want %d", inputHour, got, want)
		}
	}
	{
		inputHour = 25
		want = 0
		got = HourOffsetToHour(inputHour)
		if got != want {
			t.Errorf("HourOffsetToHour(%d) = %d, want %d", inputHour, got, want)
		}
	}
	{
		inputHour = -24
		want = 0
		got = HourOffsetToHour(inputHour)
		if got != want {
			t.Errorf("HourOffsetToHour(%d) = %d, want %d", inputHour, got, want)
		}
	}
	{
		inputHour = -25
		want = 0
		got = HourOffsetToHour(inputHour)
		if got != want {
			t.Errorf("HourOffsetToHour(%d) = %d, want %d", inputHour, got, want)
		}
	}
}

func TestMinuteOffsetToMinute(t *testing.T) {
	var input int
	var want int
	var got int
	{
		input = 1
		want = 1
		got = MinuteOffsetToMinute(input)
		if got != want {
			t.Errorf("MinuteOffsetToMinute(%d) = %d, want %d", input, got, want)
		}
	}
	{
		input = -1
		want = 59
		got = MinuteOffsetToMinute(input)
		if got != want {
			t.Errorf("MinuteOffsetToMinute(%d) = %d, want %d", input, got, want)
		}
	}
	{
		input = 0
		want = 0
		got = MinuteOffsetToMinute(input)
		if got != want {
			t.Errorf("MinuteOffsetToMinute(%d) = %d, want %d", input, got, want)
		}
	}
	{
		input = 60
		want = 0
		got = MinuteOffsetToMinute(input)
		if got != want {
			t.Errorf("MinuteOffsetToMinute(%d) = %d, want %d", input, got, want)
		}
	}
	{
		input = 61
		want = 0
		got = MinuteOffsetToMinute(input)
		if got != want {
			t.Errorf("MinuteOffsetToMinute(%d) = %d, want %d", input, got, want)
		}
	}
	{
		input = -60
		want = 0
		got = MinuteOffsetToMinute(input)
		if got != want {
			t.Errorf("MinuteOffsetToMinute(%d) = %d, want %d", input, got, want)
		}
	}
	{
		input = -61
		want = 0
		got = MinuteOffsetToMinute(input)
		if got != want {
			t.Errorf("MinuteOffsetToMinute(%d) = %d, want %d", input, got, want)
		}
	}
}

func TestSecondOffsetToSecond(t *testing.T) {
	var input int
	var want int
	var got int
	{
		input = 0
		want = 0
		got = SecondOffsetToSecond(input)
		if got != want {
			t.Errorf("SecondOffsetToSecond(%d) = %d, want %d", input, got, want)
		}
	}
	{
		input = 1
		want = 1
		got = SecondOffsetToSecond(input)
		if got != want {
			t.Errorf("SecondOffsetToSecond(%d) = %d, want %d", input, got, want)
		}
	}
	{
		input = -1
		want = 59
		got = SecondOffsetToSecond(input)
		if got != want {
			t.Errorf("SecondOffsetToSecond(%d) = %d, want %d", input, got, want)
		}
	}
	{
		input = 60
		want = 0
		got = SecondOffsetToSecond(input)
		if got != want {
			t.Errorf("SecondOffsetToSecond(%d) = %d, want %d", input, got, want)
		}
	}
	{
		input = -60
		want = 0
		got = SecondOffsetToSecond(input)
		if got != want {
			t.Errorf("SecondOffsetToSecond(%d) = %d, want %d", input, got, want)
		}
	}
}

func TestStartAtOffsetFirst(t *testing.T) {
	type TestCaseModel struct {
		title        string
		inputPeriod  period.Type
		inputStart   int64
		inputYear    int
		inputQuarter int
		inputMonth   int
		inputWeek    int
		inputDay     int
		inputHour    int
		inputMinute  int
		inputSecond  int
		want         int64
		wantNext     int64
	}

	const (
		TestTime1 = 1709177260 // 2024-02-29 11:27:40
		TestTime2 = 1740758399 // 2025-02-28 23:59:59
		TestTime3 = 1738252800 // 2025-01-31 00:00:00
	)
	var (
		TestTime1T = time.Unix(TestTime1, 0)
		TestTime2T = time.Unix(TestTime2, 0)
		TestTime3T = time.Unix(TestTime3, 0)
	)

	var TestCase = []*TestCaseModel{
		{title: "每秒1", inputPeriod: period.TypeSecond, want: TestTime1, wantNext: TestTime1 + 1, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: 0},
		{title: "每秒2", inputPeriod: period.TypeSecond, want: TestTime2, wantNext: TestTime2 + 1, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: 0},
		{title: "每秒3", inputPeriod: period.TypeSecond, want: TestTime3, wantNext: TestTime3 + 1, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: 0},
		{title: "每分钟1", inputPeriod: period.TypeMinute, want: TestTime1, wantNext: 1709177320, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: TestTime1T.Second()},
		{title: "每分钟2", inputPeriod: period.TypeMinute, want: TestTime2, wantNext: 1740758459, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: TestTime2T.Second()},
		{title: "每分钟3", inputPeriod: period.TypeMinute, want: TestTime3, wantNext: 1738252860, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: TestTime3T.Second()},
		{title: "每分钟：第一秒1", inputPeriod: period.TypeMinute, want: 1709177280, wantNext: 1709177340, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: 0},
		{title: "每分钟：第一秒2", inputPeriod: period.TypeMinute, want: 1740758400, wantNext: 1740758460, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: 0},
		{title: "每分钟：第一秒3", inputPeriod: period.TypeMinute, want: TestTime3, wantNext: 1738252860, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: 0},
		{title: "每分钟：倒数第一秒1", inputPeriod: period.TypeMinute, want: 1709177279, wantNext: 1709177339, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: -1},
		{title: "每分钟：倒数第一秒2", inputPeriod: period.TypeMinute, want: TestTime2, wantNext: 1740758459, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: -1},
		{title: "每分钟：倒数第一秒2", inputPeriod: period.TypeMinute, want: 1738252859, wantNext: 1738252919, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: -1},
		{title: "每小时1", inputPeriod: period.TypeHour, want: TestTime1, wantNext: 1709180860, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: TestTime1T.Minute(), inputSecond: TestTime1T.Second()},
		{title: "每小时2", inputPeriod: period.TypeHour, want: TestTime2, wantNext: 1740761999, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: TestTime2T.Minute(), inputSecond: TestTime2T.Second()},
		{title: "每小时3", inputPeriod: period.TypeHour, want: TestTime3, wantNext: 1738256400, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: TestTime3T.Minute(), inputSecond: TestTime3T.Second()},
		{title: "每小时：第一秒1", inputPeriod: period.TypeHour, want: 1709180820, wantNext: 1709184420, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "每小时：第一秒2", inputPeriod: period.TypeHour, want: 1740761940, wantNext: 1740765540, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "每小时：第一秒3", inputPeriod: period.TypeHour, want: TestTime3, wantNext: 1738256400, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "每小时：倒数第一秒1", inputPeriod: period.TypeHour, want: 1709177279, wantNext: 1709180879, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "每小时：倒数第一秒2", inputPeriod: period.TypeHour, want: TestTime2, wantNext: 1740761999, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "每小时：倒数第一秒2", inputPeriod: period.TypeHour, want: 1738252859, wantNext: 1738256459, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "每天1", inputPeriod: period.TypeDaily, want: TestTime1, wantNext: 1709263660, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: TestTime1T.Second()},
		{title: "每天2", inputPeriod: period.TypeDaily, want: TestTime2, wantNext: 1740844799, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: TestTime2T.Second()},
		{title: "每天3", inputPeriod: period.TypeDaily, want: TestTime3, wantNext: 1738339200, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: TestTime3T.Second()},
		{title: "每天：第一秒1", inputPeriod: period.TypeDaily, want: 1709263620, wantNext: 1709350020, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "每天：第一秒2", inputPeriod: period.TypeDaily, want: 1740844740, wantNext: 1740931140, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "每天：第一秒3", inputPeriod: period.TypeDaily, want: TestTime3, wantNext: 1738339200, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "每天：倒数第一秒1", inputPeriod: period.TypeDaily, want: 1709177279, wantNext: 1709263679, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "每天：倒数第一秒2", inputPeriod: period.TypeDaily, want: TestTime2, wantNext: 1740844799, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "每天：倒数第一秒2", inputPeriod: period.TypeDaily, want: 1738252859, wantNext: 1738339259, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "每周1", inputPeriod: period.TypeWeekly, want: TestTime1, wantNext: 1709782060, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: WeekdayToDay(TestTime1T.Weekday()), inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: TestTime1T.Second()},
		{title: "每周2", inputPeriod: period.TypeWeekly, want: TestTime2, wantNext: 1741363199, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: WeekdayToDay(TestTime2T.Weekday()), inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: TestTime2T.Second()},
		{title: "每周3", inputPeriod: period.TypeWeekly, want: TestTime3, wantNext: 1738857600, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: WeekdayToDay(TestTime3T.Weekday()), inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: TestTime3T.Second()},
		{title: "每周：第一天1", inputPeriod: period.TypeWeekly, want: 1709522820, wantNext: 1710127620, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "每周：第一天2", inputPeriod: period.TypeWeekly, want: 1741017540, wantNext: 1741622340, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "每周：第一天3", inputPeriod: period.TypeWeekly, want: 1738512000, wantNext: 1739116800, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "每周：倒数第一天1", inputPeriod: period.TypeWeekly, want: 1709436479, wantNext: 1710041279, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: -1, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "每周：倒数第一天2", inputPeriod: period.TypeWeekly, want: 1740931199, wantNext: 1741535999, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: -1, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "每周：倒数第一天2", inputPeriod: period.TypeWeekly, want: 1738425659, wantNext: 1739030459, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: -1, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "每月1", inputPeriod: period.TypeMonthly, want: TestTime1, wantNext: 1711682860, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: TestTime1T.Day(), inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: TestTime1T.Second()},
		{title: "每月2", inputPeriod: period.TypeMonthly, want: TestTime2, wantNext: 1743177599, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: TestTime2T.Day(), inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: TestTime2T.Second()},
		{title: "每月3", inputPeriod: period.TypeMonthly, want: TestTime3, wantNext: 1740672000, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: TestTime3T.Day(), inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: TestTime3T.Second()},
		{title: "每月：第一天1", inputPeriod: period.TypeMonthly, want: 1709263620, wantNext: 1711942020, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "每月：第一天2", inputPeriod: period.TypeMonthly, want: 1740844740, wantNext: 1743523140, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "每月：第一天3", inputPeriod: period.TypeMonthly, want: 1738339200, wantNext: 1740758400, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "每月：倒数第一天1", inputPeriod: period.TypeMonthly, want: 1709177279, wantNext: 1711855679, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: -1, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "每月：倒数第一天2", inputPeriod: period.TypeMonthly, want: 1740758399, wantNext: 1743436799, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: -1, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "每月：倒数第一天2", inputPeriod: period.TypeMonthly, want: 1738252859, wantNext: 1740672059, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: -1, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "每季度1", inputPeriod: period.TypeQuarterly, want: TestTime1, wantNext: 1716953260, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: (int(TestTime1T.Month()) - 1) % 3, inputWeek: 0, inputDay: TestTime1T.Day(), inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: TestTime1T.Second()},
		{title: "每季度2", inputPeriod: period.TypeQuarterly, want: TestTime2, wantNext: 1748447999, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: (int(TestTime2T.Month()) - 1) % 3, inputWeek: 0, inputDay: TestTime2T.Day(), inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: TestTime2T.Second()},
		{title: "每季度3", inputPeriod: period.TypeQuarterly, want: TestTime3, wantNext: 1745942400, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: (int(TestTime3T.Month()) - 1) % 3, inputWeek: 0, inputDay: TestTime3T.Day(), inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: TestTime3T.Second()},
		{title: "每季度：第一天1", inputPeriod: period.TypeQuarterly, want: 1714534020, wantNext: 1722482820, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: (int(TestTime1T.Month()) - 1) % 3, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "每季度：第一天2", inputPeriod: period.TypeQuarterly, want: 1746115140, wantNext: 1754063940, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: (int(TestTime2T.Month()) - 1) % 3, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "每季度：第一天3", inputPeriod: period.TypeQuarterly, want: 1743436800, wantNext: 1751299200, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: (int(TestTime3T.Month()) - 1) % 3, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "每季度：倒数第一天1", inputPeriod: period.TypeQuarterly, want: 1709177279, wantNext: 1717126079, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: (int(TestTime1T.Month()) - 1) % 3, inputWeek: 0, inputDay: -1, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "每季度：倒数第一天2", inputPeriod: period.TypeQuarterly, want: 1740758399, wantNext: 1748707199, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: (int(TestTime2T.Month()) - 1) % 3, inputWeek: 0, inputDay: -1, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "每季度：倒数第一天2", inputPeriod: period.TypeQuarterly, want: 1738252859, wantNext: 1745942459, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: (int(TestTime3T.Month()) - 1) % 3, inputWeek: 0, inputDay: -1, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "每季度：第一月1", inputPeriod: period.TypeQuarterly, want: 1711942020, wantNext: 1719804420, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "每季度：第一月2", inputPeriod: period.TypeQuarterly, want: 1743523140, wantNext: 1751385540, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "每季度：第一月3", inputPeriod: period.TypeQuarterly, want: 1743436800, wantNext: 1751299200, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "每季度：倒数第一月1", inputPeriod: period.TypeQuarterly, want: 1711855679, wantNext: 1719718079, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: -1, inputWeek: 0, inputDay: -1, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "每季度：倒数第一月2", inputPeriod: period.TypeQuarterly, want: 1743436799, wantNext: 1751299199, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: -1, inputWeek: 0, inputDay: -1, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "每季度：倒数第一月2", inputPeriod: period.TypeQuarterly, want: 1743350459, wantNext: 1751212859, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: -1, inputWeek: 0, inputDay: -1, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "每年1", inputPeriod: period.TypeYearly, want: TestTime1, wantNext: 1740713260, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: int(TestTime1T.Month()), inputWeek: 0, inputDay: TestTime1T.Day(), inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: TestTime1T.Second()},
		{title: "每年2", inputPeriod: period.TypeYearly, want: TestTime2, wantNext: 1772294399, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: int(TestTime2T.Month()), inputWeek: 0, inputDay: TestTime2T.Day(), inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: TestTime2T.Second()},
		{title: "每年3", inputPeriod: period.TypeYearly, want: TestTime3, wantNext: 1769788800, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: int(TestTime3T.Month()), inputWeek: 0, inputDay: TestTime3T.Day(), inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: TestTime3T.Second()},
		{title: "每年：第一天1", inputPeriod: period.TypeYearly, want: 1738380420, wantNext: 1769916420, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: int(TestTime1T.Month()), inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "每年：第一天2", inputPeriod: period.TypeYearly, want: 1769961540, wantNext: 1801497540, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: int(TestTime2T.Month()), inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "每年：第一天3", inputPeriod: period.TypeYearly, want: 1767196800, wantNext: 1798732800, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: int(TestTime3T.Month()), inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "每年：倒数第一天1", inputPeriod: period.TypeYearly, want: 1709177279, wantNext: 1740713279, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: int(TestTime1T.Month()), inputWeek: 0, inputDay: -1, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "每年：倒数第一天2", inputPeriod: period.TypeYearly, want: 1740758399, wantNext: 1772294399, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: int(TestTime2T.Month()), inputWeek: 0, inputDay: -1, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "每年：倒数第一天2", inputPeriod: period.TypeYearly, want: 1738252859, wantNext: 1769788859, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: int(TestTime3T.Month()), inputWeek: 0, inputDay: -1, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "每年：第一月1", inputPeriod: period.TypeYearly, want: 1735702020, wantNext: 1767238020, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 1, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "每年：第一月2", inputPeriod: period.TypeYearly, want: 1767283140, wantNext: 1798819140, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 1, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "每年：第一月3", inputPeriod: period.TypeYearly, want: 1767196800, wantNext: 1798732800, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 1, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "每年：倒数第一月1", inputPeriod: period.TypeYearly, want: 1735615679, wantNext: 1767151679, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: -1, inputWeek: 0, inputDay: -1, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "每年：倒数第一月2", inputPeriod: period.TypeYearly, want: 1767196799, wantNext: 1798732799, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: -1, inputWeek: 0, inputDay: -1, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "每年：倒数第一月2", inputPeriod: period.TypeYearly, want: 1767110459, wantNext: 1798646459, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: -1, inputWeek: 0, inputDay: -1, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "间隔2秒1", inputPeriod: period.TypeSecondInterval, want: TestTime1, wantNext: 1709177262, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: 2},
		{title: "间隔2秒2", inputPeriod: period.TypeSecondInterval, want: TestTime2, wantNext: 1740758401, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: 2},
		{title: "间隔2秒3", inputPeriod: period.TypeSecondInterval, want: TestTime3, wantNext: 1738252802, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 0, inputSecond: 2},
		{title: "间隔2分钟1", inputPeriod: period.TypeMinuteInterval, want: TestTime1, wantNext: 1709177380, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 2, inputSecond: TestTime1T.Second()},
		{title: "间隔2分钟2", inputPeriod: period.TypeMinuteInterval, want: TestTime2, wantNext: 1740758519, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 2, inputSecond: TestTime2T.Second()},
		{title: "间隔2分钟3", inputPeriod: period.TypeMinuteInterval, want: TestTime3, wantNext: 1738252920, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 2, inputSecond: TestTime3T.Second()},
		{title: "间隔2分钟：第一秒1", inputPeriod: period.TypeMinuteInterval, want: 1709177340, wantNext: 1709177460, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 2, inputSecond: 0},
		{title: "间隔2分钟：第一秒2", inputPeriod: period.TypeMinuteInterval, want: 1740758460, wantNext: 1740758580, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 2, inputSecond: 0},
		{title: "间隔2分钟：第一秒3", inputPeriod: period.TypeMinuteInterval, want: TestTime3, wantNext: 1738252920, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 2, inputSecond: 0},
		{title: "间隔2分钟：倒数第一秒1", inputPeriod: period.TypeMinuteInterval, want: 1709177279, wantNext: 1709177399, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 2, inputSecond: -1},
		{title: "间隔2分钟：倒数第一秒2", inputPeriod: period.TypeMinuteInterval, want: TestTime2, wantNext: 1740758519, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 2, inputSecond: -1},
		{title: "间隔2分钟：倒数第一秒2", inputPeriod: period.TypeMinuteInterval, want: 1738252859, wantNext: 1738252979, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 0, inputMinute: 2, inputSecond: -1},
		{title: "间隔2小时1", inputPeriod: period.TypeHourInterval, want: TestTime1, wantNext: 1709184460, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 2, inputMinute: TestTime1T.Minute(), inputSecond: TestTime1T.Second()},
		{title: "间隔2小时2", inputPeriod: period.TypeHourInterval, want: TestTime2, wantNext: 1740765599, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 2, inputMinute: TestTime2T.Minute(), inputSecond: TestTime2T.Second()},
		{title: "间隔2小时3", inputPeriod: period.TypeHourInterval, want: TestTime3, wantNext: 1738260000, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 2, inputMinute: TestTime3T.Minute(), inputSecond: TestTime3T.Second()},
		{title: "间隔2小时：第一秒1", inputPeriod: period.TypeHourInterval, want: 1709184420, wantNext: 1709191620, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 2, inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "间隔2小时：第一秒2", inputPeriod: period.TypeHourInterval, want: 1740765540, wantNext: 1740772740, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 2, inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "间隔2小时：第一秒3", inputPeriod: period.TypeHourInterval, want: TestTime3, wantNext: 1738260000, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 2, inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "间隔2小时：倒数第一秒1", inputPeriod: period.TypeHourInterval, want: 1709177279, wantNext: 1709184479, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 2, inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "间隔2小时：倒数第一秒2", inputPeriod: period.TypeHourInterval, want: TestTime2, wantNext: 1740765599, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 2, inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "间隔2小时：倒数第一秒2", inputPeriod: period.TypeHourInterval, want: 1738252859, wantNext: 1738260059, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: 2, inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "间隔2天1", inputPeriod: period.TypeDayInterval, want: TestTime1, wantNext: 1709350060, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 2, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: TestTime1T.Second()},
		{title: "间隔2天2", inputPeriod: period.TypeDayInterval, want: TestTime2, wantNext: 1740931199, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 2, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: TestTime2T.Second()},
		{title: "间隔2天3", inputPeriod: period.TypeDayInterval, want: TestTime3, wantNext: 1738425600, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 2, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: TestTime3T.Second()},
		{title: "间隔2天：第一秒1", inputPeriod: period.TypeDayInterval, want: 1709350020, wantNext: 1709522820, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 2, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "间隔2天：第一秒2", inputPeriod: period.TypeDayInterval, want: 1740931140, wantNext: 1741103940, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 2, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "间隔2天：第一秒3", inputPeriod: period.TypeDayInterval, want: TestTime3, wantNext: 1738425600, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 2, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "间隔2天：倒数第一秒1", inputPeriod: period.TypeDayInterval, want: 1709177279, wantNext: 1709350079, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 2, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "间隔2天：倒数第一秒2", inputPeriod: period.TypeDayInterval, want: TestTime2, wantNext: 1740931199, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 2, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "间隔2天：倒数第一秒2", inputPeriod: period.TypeDayInterval, want: 1738252859, wantNext: 1738425659, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 0, inputDay: 2, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "间隔2周1", inputPeriod: period.TypeWeekInterval, want: TestTime1, wantNext: 1710386860, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 2, inputDay: WeekdayToDay(TestTime1T.Weekday()), inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: TestTime1T.Second()},
		{title: "间隔2周2", inputPeriod: period.TypeWeekInterval, want: TestTime2, wantNext: 1741967999, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 2, inputDay: WeekdayToDay(TestTime2T.Weekday()), inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: TestTime2T.Second()},
		{title: "间隔2周3", inputPeriod: period.TypeWeekInterval, want: TestTime3, wantNext: 1739462400, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 2, inputDay: WeekdayToDay(TestTime3T.Weekday()), inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: TestTime3T.Second()},
		{title: "间隔2周：第一天1", inputPeriod: period.TypeWeekInterval, want: 1710127620, wantNext: 1711337220, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 2, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "间隔2周：第一天2", inputPeriod: period.TypeWeekInterval, want: 1741622340, wantNext: 1742831940, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 2, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "间隔2周：第一天3", inputPeriod: period.TypeWeekInterval, want: 1739116800, wantNext: 1740326400, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 2, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "间隔2周：倒数第一天1", inputPeriod: period.TypeWeekInterval, want: 1709436479, wantNext: 1710646079, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 2, inputDay: -1, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "间隔2周：倒数第一天2", inputPeriod: period.TypeWeekInterval, want: 1740931199, wantNext: 1742140799, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 2, inputDay: -1, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "间隔2周：倒数第一天2", inputPeriod: period.TypeWeekInterval, want: 1738425659, wantNext: 1739635259, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 0, inputWeek: 2, inputDay: -1, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "间隔2月1", inputPeriod: period.TypeMonthInterval, want: TestTime1, wantNext: 1714361260, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 2, inputWeek: 0, inputDay: TestTime1T.Day(), inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: TestTime1T.Second()},
		{title: "间隔2月2", inputPeriod: period.TypeMonthInterval, want: TestTime2, wantNext: 1745855999, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 2, inputWeek: 0, inputDay: TestTime2T.Day(), inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: TestTime2T.Second()},
		{title: "间隔2月3", inputPeriod: period.TypeMonthInterval, want: TestTime3, wantNext: 1743350400, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 2, inputWeek: 0, inputDay: TestTime3T.Day(), inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: TestTime3T.Second()},
		{title: "间隔2月：第一天1", inputPeriod: period.TypeMonthInterval, want: 1711942020, wantNext: 1717212420, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 2, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "间隔2月：第一天2", inputPeriod: period.TypeMonthInterval, want: 1743523140, wantNext: 1748793540, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 2, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "间隔2月：第一天3", inputPeriod: period.TypeMonthInterval, want: 1740758400, wantNext: 1746028800, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 2, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "间隔2月：倒数第一天1", inputPeriod: period.TypeMonthInterval, want: 1709177279, wantNext: 1714447679, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 2, inputWeek: 0, inputDay: -1, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "间隔2月：倒数第一天2", inputPeriod: period.TypeMonthInterval, want: 1740758399, wantNext: 1746028799, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 2, inputWeek: 0, inputDay: -1, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "间隔2月：倒数第一天2", inputPeriod: period.TypeMonthInterval, want: 1738252859, wantNext: 1743350459, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 2, inputWeek: 0, inputDay: -1, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "间隔11月：第一天1", inputPeriod: period.TypeMonthInterval, want: 1735702020, wantNext: 1764559620, inputStart: TestTime1, inputYear: 0, inputQuarter: 0, inputMonth: 11, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "间隔11月：第一天2", inputPeriod: period.TypeMonthInterval, want: 1767283140, wantNext: 1796140740, inputStart: TestTime2, inputYear: 0, inputQuarter: 0, inputMonth: 11, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "间隔11月：第一天3", inputPeriod: period.TypeMonthInterval, want: 1764518400, wantNext: 1793462400, inputStart: TestTime3, inputYear: 0, inputQuarter: 0, inputMonth: 11, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "间隔2季度1", inputPeriod: period.TypeQuarterInterval, want: TestTime1, wantNext: 1724902060, inputStart: TestTime1, inputYear: 0, inputQuarter: 2, inputMonth: (int(TestTime1T.Month()) - 1) % 3, inputWeek: 0, inputDay: TestTime1T.Day(), inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: TestTime1T.Second()},
		{title: "间隔2季度2", inputPeriod: period.TypeQuarterInterval, want: TestTime2, wantNext: 1756396799, inputStart: TestTime2, inputYear: 0, inputQuarter: 2, inputMonth: (int(TestTime2T.Month()) - 1) % 3, inputWeek: 0, inputDay: TestTime2T.Day(), inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: TestTime2T.Second()},
		{title: "间隔2季度3", inputPeriod: period.TypeQuarterInterval, want: TestTime3, wantNext: 1753891200, inputStart: TestTime3, inputYear: 0, inputQuarter: 2, inputMonth: (int(TestTime3T.Month()) - 1) % 3, inputWeek: 0, inputDay: TestTime3T.Day(), inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: TestTime3T.Second()},
		{title: "间隔2季度：第一天1", inputPeriod: period.TypeQuarterInterval, want: 1722482820, wantNext: 1738380420, inputStart: TestTime1, inputYear: 0, inputQuarter: 2, inputMonth: (int(TestTime1T.Month()) - 1) % 3, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "间隔2季度：第一天2", inputPeriod: period.TypeQuarterInterval, want: 1754063940, wantNext: 1769961540, inputStart: TestTime2, inputYear: 0, inputQuarter: 2, inputMonth: (int(TestTime2T.Month()) - 1) % 3, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "间隔2季度：第一天3", inputPeriod: period.TypeQuarterInterval, want: 1751299200, wantNext: 1767196800, inputStart: TestTime3, inputYear: 0, inputQuarter: 2, inputMonth: (int(TestTime3T.Month()) - 1) % 3, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "间隔2季度：倒数第一天1", inputPeriod: period.TypeQuarterInterval, want: 1709177279, wantNext: 1725074879, inputStart: TestTime1, inputYear: 0, inputQuarter: 2, inputMonth: (int(TestTime1T.Month()) - 1) % 3, inputWeek: 0, inputDay: -1, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "间隔2季度：倒数第一天2", inputPeriod: period.TypeQuarterInterval, want: 1740758399, wantNext: 1756655999, inputStart: TestTime2, inputYear: 0, inputQuarter: 2, inputMonth: (int(TestTime2T.Month()) - 1) % 3, inputWeek: 0, inputDay: -1, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "间隔2季度：倒数第一天2", inputPeriod: period.TypeQuarterInterval, want: 1738252859, wantNext: 1753891259, inputStart: TestTime3, inputYear: 0, inputQuarter: 2, inputMonth: (int(TestTime3T.Month()) - 1) % 3, inputWeek: 0, inputDay: -1, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "间隔2季度：第一月1", inputPeriod: period.TypeQuarterInterval, want: 1719804420, wantNext: 1735702020, inputStart: TestTime1, inputYear: 0, inputQuarter: 2, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "间隔2季度：第一月2", inputPeriod: period.TypeQuarterInterval, want: 1751385540, wantNext: 1767283140, inputStart: TestTime2, inputYear: 0, inputQuarter: 2, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "间隔2季度：第一月3", inputPeriod: period.TypeQuarterInterval, want: 1751299200, wantNext: 1767196800, inputStart: TestTime3, inputYear: 0, inputQuarter: 2, inputMonth: 0, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "间隔2季度：倒数第一月1", inputPeriod: period.TypeQuarterInterval, want: 1711855679, wantNext: 1727666879, inputStart: TestTime1, inputYear: 0, inputQuarter: 2, inputMonth: -1, inputWeek: 0, inputDay: -1, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "间隔2季度：倒数第一月2", inputPeriod: period.TypeQuarterInterval, want: 1743436799, wantNext: 1759247999, inputStart: TestTime2, inputYear: 0, inputQuarter: 2, inputMonth: -1, inputWeek: 0, inputDay: -1, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "间隔2季度：倒数第一月2", inputPeriod: period.TypeQuarterInterval, want: 1743350459, wantNext: 1759161659, inputStart: TestTime3, inputYear: 0, inputQuarter: 2, inputMonth: -1, inputWeek: 0, inputDay: -1, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "间隔2年1", inputPeriod: period.TypeYearInterval, want: TestTime1, wantNext: 1772249260, inputStart: TestTime1, inputYear: 2, inputQuarter: 0, inputMonth: int(TestTime1T.Month()), inputWeek: 0, inputDay: TestTime1T.Day(), inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: TestTime1T.Second()},
		{title: "间隔2年2", inputPeriod: period.TypeYearInterval, want: TestTime2, wantNext: 1803830399, inputStart: TestTime2, inputYear: 2, inputQuarter: 0, inputMonth: int(TestTime2T.Month()), inputWeek: 0, inputDay: TestTime2T.Day(), inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: TestTime2T.Second()},
		{title: "间隔2年3", inputPeriod: period.TypeYearInterval, want: TestTime3, wantNext: 1801324800, inputStart: TestTime3, inputYear: 2, inputQuarter: 0, inputMonth: int(TestTime3T.Month()), inputWeek: 0, inputDay: TestTime3T.Day(), inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: TestTime3T.Second()},
		{title: "间隔2年：第一天1", inputPeriod: period.TypeYearInterval, want: 1769916420, wantNext: 1832988420, inputStart: TestTime1, inputYear: 2, inputQuarter: 0, inputMonth: int(TestTime1T.Month()), inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "间隔2年：第一天2", inputPeriod: period.TypeYearInterval, want: 1801497540, wantNext: 1864655940, inputStart: TestTime2, inputYear: 2, inputQuarter: 0, inputMonth: int(TestTime2T.Month()), inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "间隔2年：第一天3", inputPeriod: period.TypeYearInterval, want: 1798732800, wantNext: 1861891200, inputStart: TestTime3, inputYear: 2, inputQuarter: 0, inputMonth: int(TestTime3T.Month()), inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "间隔2年：倒数第一天1", inputPeriod: period.TypeYearInterval, want: 1709177279, wantNext: 1772249279, inputStart: TestTime1, inputYear: 2, inputQuarter: 0, inputMonth: int(TestTime1T.Month()), inputWeek: 0, inputDay: -1, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "间隔2年：倒数第一天2", inputPeriod: period.TypeYearInterval, want: 1740758399, wantNext: 1803830399, inputStart: TestTime2, inputYear: 2, inputQuarter: 0, inputMonth: int(TestTime2T.Month()), inputWeek: 0, inputDay: -1, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "间隔2年：倒数第一天2", inputPeriod: period.TypeYearInterval, want: 1738252859, wantNext: 1801324859, inputStart: TestTime3, inputYear: 2, inputQuarter: 0, inputMonth: int(TestTime3T.Month()), inputWeek: 0, inputDay: -1, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
		{title: "间隔2年：第一月1", inputPeriod: period.TypeYearInterval, want: 1767238020, wantNext: 1830310020, inputStart: TestTime1, inputYear: 2, inputQuarter: 0, inputMonth: 1, inputWeek: 0, inputDay: 0, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: 0},
		{title: "间隔2年：第一月2", inputPeriod: period.TypeYearInterval, want: 1798819140, wantNext: 1861977540, inputStart: TestTime2, inputYear: 2, inputQuarter: 0, inputMonth: 1, inputWeek: 0, inputDay: 0, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: 0},
		{title: "间隔2年：第一月3", inputPeriod: period.TypeYearInterval, want: 1798732800, wantNext: 1861891200, inputStart: TestTime3, inputYear: 2, inputQuarter: 0, inputMonth: 1, inputWeek: 0, inputDay: 0, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: 0},
		{title: "间隔2年：倒数第一月1", inputPeriod: period.TypeYearInterval, want: 1735615679, wantNext: 1798687679, inputStart: TestTime1, inputYear: 2, inputQuarter: 0, inputMonth: -1, inputWeek: 0, inputDay: -1, inputHour: TestTime1T.Hour(), inputMinute: TestTime1T.Minute(), inputSecond: -1},
		{title: "间隔2年：倒数第一月2", inputPeriod: period.TypeYearInterval, want: 1767196799, wantNext: 1830268799, inputStart: TestTime2, inputYear: 2, inputQuarter: 0, inputMonth: -1, inputWeek: 0, inputDay: -1, inputHour: TestTime2T.Hour(), inputMinute: TestTime2T.Minute(), inputSecond: -1},
		{title: "间隔2年：倒数第一月2", inputPeriod: period.TypeYearInterval, want: 1767110459, wantNext: 1830182459, inputStart: TestTime3, inputYear: 2, inputQuarter: 0, inputMonth: -1, inputWeek: 0, inputDay: -1, inputHour: TestTime3T.Hour(), inputMinute: TestTime3T.Minute(), inputSecond: -1},
	}

	for k := range TestCase {
		c := TestCase[k]
		ok := true
		got := StartAtOffsetFirst(c.inputPeriod, c.inputStart, false, c.inputYear, c.inputQuarter, c.inputMonth, c.inputWeek, c.inputDay, c.inputHour, c.inputMinute, c.inputSecond)
		result := strings.Builder{}
		result.WriteString(fmt.Sprintf("\n%s %s P[%s] Year[%d] Quarter[%d] Month[%d] Week[%d] Day[%d] Hour[%d] Minute[%d] Second[%d]\n%s", time.Unix(c.inputStart, 0).Format("2006-01-02 15:04:05"), c.title, c.inputPeriod.String(), c.inputYear, c.inputQuarter, c.inputMonth, c.inputWeek, c.inputDay, c.inputHour, c.inputMinute, c.inputSecond, time.Unix(got, 0).Format("2006-01-02 15:04:05")))
		if got != c.want {
			ok = false
			result.WriteString(fmt.Sprintf("\nWrong: StartAtOffsetFirst(%s,%d,false,%d,%d,%d,%d,%d,%d,%d,%d) = %d, want %d", c.inputPeriod.String(), c.inputStart, c.inputYear, c.inputQuarter, c.inputMonth, c.inputWeek, c.inputDay, c.inputHour, c.inputMinute, c.inputSecond, got, c.want))
		}
		gotNext := StartAtOffsetFirst(c.inputPeriod, c.inputStart, true, c.inputYear, c.inputQuarter, c.inputMonth, c.inputWeek, c.inputDay, c.inputHour, c.inputMinute, c.inputSecond)
		result.WriteString(fmt.Sprintf("\n%s Next", time.Unix(gotNext, 0).Format("2006-01-02 15:04:05")))
		if gotNext != c.wantNext {
			ok = false
			result.WriteString(fmt.Sprintf("\nWrong: StartAtOffsetFirst(%s,%d,true,%d,%d,%d,%d,%d,%d,%d,%d) = %d, want-next %d", c.inputPeriod.String(), c.inputStart, c.inputYear, c.inputQuarter, c.inputMonth, c.inputWeek, c.inputDay, c.inputHour, c.inputMinute, c.inputSecond, gotNext, c.wantNext))
		}
		if ok {
			t.Log(result.String())
		} else {
			t.Errorf(result.String())
		}
	}

}
