package utils

import (
	"fmt"
	"msg/cmd/app/server/common/period"
	"strings"
	"time"
)

func FormatSecond(sec int64, deep int) string {
	isFirst := true
	res := strings.Builder{}
	if sec < 0 {
		sec = -sec
		res.WriteString("-")
	}
	if sec == 0 {
		res.WriteString("0s")
		return res.String()
	}

	days := sec / (24 * 3600)
	if days > 0 {
		res.WriteString(fmt.Sprintf("%dd", days))
		isFirst = false
	}
	if !isFirst {
		deep--
		if deep <= 0 {
			return res.String()
		}
	}

	hours := (sec % (24 * 3600)) / 3600
	if hours > 0 {
		res.WriteString(fmt.Sprintf("%dh", hours))
		isFirst = false
	}
	if !isFirst {
		deep--
		if deep <= 0 {
			return res.String()
		}
	}

	minutes := (sec % 3600) / 60
	if minutes > 0 {
		res.WriteString(fmt.Sprintf("%dm", minutes))
		isFirst = false
	}
	if !isFirst {
		deep--
		if deep <= 0 {
			return res.String()
		}
	}

	secs := sec % 60
	if secs > 0 {
		res.WriteString(fmt.Sprintf("%ds", secs))
		isFirst = false
	}
	if !isFirst {
		deep--
		if deep <= 0 {
			return res.String()
		}
	}
	return res.String()
}

func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func YearDays(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

// MonthDays 获取某年某月多少天
func MonthDays(year int, month int) int {
	if year < 1970 {
		year = 1970
	}
	if month < 1 {
		month = 1
	}
	if month > 12 {
		month = 12
	}

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	}
	return 31
}

// MonthOffsetToMonth 识别月份
func MonthOffsetToMonth(month int) int {
	if month > 12 {
		month = 12
	}
	if month < -12 {
		month = -12
	}

	if month > 0 {
		return month
	} else if month < 0 {
		return 13 + month
	} else {
		return 1
	}
}

// MonthOffsetToMonthByQuarter 识别季度内第几月: 0-2
func MonthOffsetToMonthByQuarter(month int) int {
	if month > 2 {
		month = 2
	}
	if month < -3 {
		month = -3
	}

	if month >= 0 {
		return month
	} else {
		return 3 + month
	}
}

// MonthToQuarter 识别月份是第几季度: 0-3
func MonthToQuarter(month int) int {
	if month < 1 {
		month = 1
	}
	if month > 12 {
		month = 12
	}

	return (month - 1) / 3
}

// DayOffsetToDay 识别是某年某月第几天
func DayOffsetToDay(year, month, day int) int {
	dayLimit := MonthDays(year, month)
	if day > dayLimit {
		day = dayLimit
	}
	if day < -dayLimit {
		day = -dayLimit
	}
	if day > 0 {
		return day
	} else if day < 0 {
		return dayLimit + 1 + day
	} else {
		return 1
	}
}

// DayOffsetToWeek 识别是周内第几天(0-6)
func DayOffsetToWeek(day int) int {
	if day > 6 {
		day = 6
	}
	if day < -7 {
		day = -7
	}

	if day >= 0 {
		return day
	} else {
		return 7 + day
	}
}

// HourOffsetToHour 识别小时
func HourOffsetToHour(hour int) int {
	if hour > 23 {
		hour = 23
	}
	if hour < -24 {
		hour = -24
	}

	if hour >= 0 {
		return hour
	} else {
		return 24 + hour
	}
}

// MinuteOffsetToMinute 识别分钟
func MinuteOffsetToMinute(minute int) int {
	if minute > 59 {
		minute = 59
	}
	if minute < -60 {
		minute = -60
	}

	if minute >= 0 {
		return minute
	} else {
		return 60 + minute
	}
}

// SecondOffsetToSecond 识别秒
func SecondOffsetToSecond(second int) int {
	if second > 59 {
		second = 59
	}
	if second < -60 {
		second = -60
	}

	if second >= 0 {
		return second
	} else {
		return 60 + second
	}
}

func WeekdayToDay(weekDay time.Weekday) int {
	if weekDay == time.Sunday {
		return 6
	}
	return int(weekDay) - 1
}

// StartAtOffsetFirst 计算出实际的时间戳
//
//	year:
func StartAtOffsetFirst(periodType period.Type, start int64, isNext bool, year, quarter, month, week, day, hour, minute, second int) int64 {
	startT := time.Unix(start, 0)
	switch periodType {
	case period.TypeSecond: // 每秒
		// 忽略所有参数
		nextStart := start
		if isNext {
			nextStart++
		}
		return nextStart
	case period.TypeMinute: // 每分钟
		// 生效： second
		secondT := SecondOffsetToSecond(second) // 计算出是第几秒
		startSec := start % 60
		nextStart := start - startSec + int64(secondT)
		// 保证计算出来的时间是最靠近start且还未到来的时间
		if int64(secondT) < startSec {
			nextStart += 60
		}
		if isNext {
			nextStart += 60
		}
		return nextStart
	case period.TypeHour: // 每小时
		// 生效： minute, second
		secondT := SecondOffsetToSecond(second)
		minuteT := MinuteOffsetToMinute(minute)
		secs := 60*int64(minuteT) + int64(secondT)
		startMSec := start % 3600
		nextStart := start - startMSec + secs
		if secs < startMSec {
			nextStart += 3600
		}
		if isNext {
			nextStart += 3600
		}
		return nextStart
	case period.TypeDaily: // 每天
		// 生效： hour, minute, second
		secondT := SecondOffsetToSecond(second)
		minuteT := MinuteOffsetToMinute(minute)
		hourT := HourOffsetToHour(hour)
		secs := 60*60*int64(hourT) + 60*int64(minuteT) + int64(secondT)
		startHSec := 60*60*int64(startT.Hour()) + 60*int64(startT.Minute()) + int64(startT.Second())
		nextStart := start - startHSec + secs
		if secs < startHSec {
			nextStart += 86400
		}
		if isNext {
			nextStart += 86400
		}
		return nextStart
	case period.TypeWeekly: // 每周
		// 生效： day, hour, minute, second
		secondT := SecondOffsetToSecond(second)
		minuteT := MinuteOffsetToMinute(minute)
		hourT := HourOffsetToHour(hour)
		dayT := DayOffsetToWeek(day)
		secs := 24*60*60*int64(dayT) + 60*60*int64(hourT) + 60*int64(minuteT) + int64(secondT)
		startHSec := 24*60*60*int64(WeekdayToDay(startT.Weekday())) + 60*60*int64(startT.Hour()) + 60*int64(startT.Minute()) + int64(startT.Second())
		nextStart := start - startHSec + secs
		if secs < startHSec {
			nextStart += 604800
		}
		if isNext {
			nextStart += 604800
		}
		return nextStart
	case period.TypeMonthly: // 每月
		// 生效： day, hour, minute, second
		secondT := SecondOffsetToSecond(second)
		minuteT := MinuteOffsetToMinute(minute)
		hourT := HourOffsetToHour(hour)
		monthT := int(startT.Month())
		yearT := startT.Year()
		dayT := DayOffsetToDay(yearT, monthT, day)
		newTime := time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		if newTime.Before(startT) {
			monthT++
			if monthT > 12 {
				monthT = 1
				yearT++
			}
			dayT = DayOffsetToDay(yearT, monthT, day)
			newTime = time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		}
		if isNext {
			monthT++
			if monthT > 12 {
				monthT = 1
				yearT++
			}
			dayT = DayOffsetToDay(yearT, monthT, day)
			newTime = time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		}
		return newTime.Unix()
	case period.TypeQuarterly: // 每季度
		// 生效： month, day, hour, minute, second
		secondT := SecondOffsetToSecond(second)
		minuteT := MinuteOffsetToMinute(minute)
		hourT := HourOffsetToHour(hour)
		yearT := startT.Year()
		quarterT := MonthToQuarter(int(startT.Month()))
		monthT := quarterT*3 + MonthOffsetToMonthByQuarter(month) + 1
		dayT := DayOffsetToDay(yearT, monthT, day)
		newTime := time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		if newTime.Before(startT) {
			monthT += 3
			if monthT > 12 {
				monthT = monthT - 12
				yearT++
			}
			dayT = DayOffsetToDay(yearT, monthT, day)
			newTime = time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		}
		if isNext {
			monthT += 3
			if monthT > 12 {
				monthT = monthT - 12
				yearT++
			}
			dayT = DayOffsetToDay(yearT, monthT, day)
			newTime = time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		}
		return newTime.Unix()
	case period.TypeYearly: // 每年
		// 生效： month, day, hour, minute, second
		secondT := SecondOffsetToSecond(second)
		minuteT := MinuteOffsetToMinute(minute)
		hourT := HourOffsetToHour(hour)
		monthT := MonthOffsetToMonth(month)
		yearT := startT.Year()
		dayT := DayOffsetToDay(yearT, monthT, day)
		newTime := time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		if newTime.Before(startT) {
			yearT++
			dayT = DayOffsetToDay(yearT, monthT, day)
			newTime = time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		}
		if isNext {
			yearT++
			dayT = DayOffsetToDay(yearT, monthT, day)
			newTime = time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		}
		return newTime.Unix()
	case period.TypeSecondInterval: // 间隔秒
		// 生效: second
		nextStart := start
		if isNext {
			nextStart += int64(second)
		}
		return nextStart
	case period.TypeMinuteInterval: // 间隔分钟
		// 生效： minute, second
		secondT := SecondOffsetToSecond(second) // 计算出是第几秒
		startSec := start % 60
		nextStart := start - startSec + int64(secondT)
		// 保证计算出来的时间是最靠近start且还未到来的时间
		if int64(secondT) < startSec {
			nextStart += int64(minute) * 60
		}
		if isNext {
			nextStart += int64(minute) * 60
		}
		return nextStart
	case period.TypeHourInterval: // 间隔小时
		// 生效： hour, minute, second
		secondT := SecondOffsetToSecond(second)
		minuteT := MinuteOffsetToMinute(minute)
		secs := 60*int64(minuteT) + int64(secondT)
		startMSec := start % 3600
		nextStart := start - startMSec + secs
		if secs < startMSec {
			nextStart += int64(hour) * 3600
		}
		if isNext {
			nextStart += int64(hour) * 3600
		}
		return nextStart
	case period.TypeDayInterval: // 间隔天
		// 生效： day, hour, minute, second
		secondT := SecondOffsetToSecond(second)
		minuteT := MinuteOffsetToMinute(minute)
		hourT := HourOffsetToHour(hour)
		secs := 60*60*int64(hourT) + 60*int64(minuteT) + int64(secondT)
		startHSec := 60*60*int64(startT.Hour()) + 60*int64(startT.Minute()) + int64(startT.Second())
		nextStart := start - startHSec + secs
		if secs < startHSec {
			nextStart += int64(day) * 86400
		}
		if isNext {
			nextStart += int64(day) * 86400
		}
		return nextStart
	case period.TypeWeekInterval: // 间隔周
		// 生效： week, day, hour, minute, second
		secondT := SecondOffsetToSecond(second)
		minuteT := MinuteOffsetToMinute(minute)
		hourT := HourOffsetToHour(hour)
		dayT := DayOffsetToWeek(day)
		secs := 24*60*60*int64(dayT) + 60*60*int64(hourT) + 60*int64(minuteT) + int64(secondT)
		startHSec := 24*60*60*int64(WeekdayToDay(startT.Weekday())) + 60*60*int64(startT.Hour()) + 60*int64(startT.Minute()) + int64(startT.Second())
		nextStart := start - startHSec + secs
		if secs < startHSec {
			nextStart += int64(week) * 604800
		}
		if isNext {
			nextStart += int64(week) * 604800
		}
		return nextStart
	case period.TypeMonthInterval: // 间隔月
		// 生效： month, day, hour, minute, second
		secondT := SecondOffsetToSecond(second)
		minuteT := MinuteOffsetToMinute(minute)
		hourT := HourOffsetToHour(hour)
		monthT := int(startT.Month())
		yearT := startT.Year()
		dayT := DayOffsetToDay(yearT, monthT, day)
		newTime := time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		if newTime.Before(startT) {
			monthT += month
			for monthT > 12 {
				monthT -= 12
				yearT++
			}
			dayT = DayOffsetToDay(yearT, monthT, day)
			newTime = time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		}
		if isNext {
			monthT += month
			for monthT > 12 {
				monthT -= 12
				yearT++
			}
			dayT = DayOffsetToDay(yearT, monthT, day)
			newTime = time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		}
		return newTime.Unix()
	case period.TypeQuarterInterval: // 间隔季度
		// 生效： quarter, month, day, hour, minute, second
		secondT := SecondOffsetToSecond(second)
		minuteT := MinuteOffsetToMinute(minute)
		hourT := HourOffsetToHour(hour)
		yearT := startT.Year()
		quarterT := MonthToQuarter(int(startT.Month()))
		monthT := quarterT*3 + MonthOffsetToMonthByQuarter(month) + 1
		dayT := DayOffsetToDay(yearT, monthT, day)
		newTime := time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		if newTime.Before(startT) {
			monthT += quarter * 3
			for monthT > 12 {
				monthT -= 12
				yearT++
			}
			dayT = DayOffsetToDay(yearT, monthT, day)
			newTime = time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		}
		if isNext {
			monthT += quarter * 3
			for monthT > 12 {
				monthT -= 12
				yearT++
			}
			dayT = DayOffsetToDay(yearT, monthT, day)
			newTime = time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		}
		return newTime.Unix()
	case period.TypeYearInterval: // 间隔年
		// 生效： year, month, day, hour, minute, second
		secondT := SecondOffsetToSecond(second)
		minuteT := MinuteOffsetToMinute(minute)
		hourT := HourOffsetToHour(hour)
		monthT := MonthOffsetToMonth(month)
		yearT := startT.Year()
		dayT := DayOffsetToDay(yearT, monthT, day)
		newTime := time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		if newTime.Before(startT) {
			yearT += year
			dayT = DayOffsetToDay(yearT, monthT, day)
			newTime = time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		}
		if isNext {
			yearT += year
			dayT = DayOffsetToDay(yearT, monthT, day)
			newTime = time.Date(yearT, time.Month(monthT), dayT, hourT, minuteT, secondT, 0, startT.Location())
		}
		return newTime.Unix()
	}

	return 0
}
