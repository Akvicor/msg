package period

type Type int64

const (
	TypeSecond          Type = 1  // 每秒
	TypeMinute          Type = 2  // 每分钟
	TypeHour            Type = 3  // 每小时
	TypeDaily           Type = 4  // 每天
	TypeWeekly          Type = 5  // 每周
	TypeMonthly         Type = 6  // 每月
	TypeQuarterly       Type = 7  // 每季度
	TypeYearly          Type = 8  // 每年
	TypeSecondInterval  Type = 9  // 间隔秒
	TypeMinuteInterval  Type = 10 // 间隔分钟
	TypeHourInterval    Type = 11 // 间隔小时
	TypeDayInterval     Type = 12 // 间隔天
	TypeWeekInterval    Type = 13 // 间隔周
	TypeMonthInterval   Type = 14 // 间隔月
	TypeQuarterInterval Type = 15 // 间隔季度
	TypeYearInterval    Type = 16 // 间隔年
)

func (t *Type) Valid() bool {
	switch *t {
	case TypeSecond, TypeMinute, TypeHour, TypeDaily:
		return true
	case TypeWeekly, TypeMonthly, TypeQuarterly, TypeYearly:
		return true
	case TypeSecondInterval, TypeMinuteInterval, TypeHourInterval, TypeDayInterval:
		return true
	case TypeWeekInterval, TypeQuarterInterval, TypeMonthInterval, TypeYearInterval:
		return true
	default:
		return false
	}
}

func (t Type) String() string {
	switch t {
	case TypeSecond:
		return "每秒"
	case TypeMinute:
		return "每分钟"
	case TypeHour:
		return "每小时"
	case TypeDaily:
		return "每天"
	case TypeWeekly:
		return "每周"
	case TypeMonthly:
		return "每月"
	case TypeQuarterly:
		return "每季度"
	case TypeYearly:
		return "每年"
	case TypeSecondInterval:
		return "间隔秒"
	case TypeMinuteInterval:
		return "间隔分"
	case TypeHourInterval:
		return "间隔时"
	case TypeDayInterval:
		return "间隔天"
	case TypeWeekInterval:
		return "间隔周"
	case TypeMonthInterval:
		return "间隔月"
	case TypeQuarterInterval:
		return "间隔季度"
	case TypeYearInterval:
		return "间隔年"
	default:
		return "未知"
	}
}

func (t Type) StringEnglish() string {
	switch t {
	case TypeSecond:
		return "Second"
	case TypeMinute:
		return "Minute"
	case TypeHour:
		return "Hour"
	case TypeDaily:
		return "Daily"
	case TypeWeekly:
		return "Weekly"
	case TypeMonthly:
		return "Monthly"
	case TypeQuarterly:
		return "Quarterly"
	case TypeYearly:
		return "Yearly"
	case TypeSecondInterval:
		return "SecondInterval"
	case TypeMinuteInterval:
		return "MinuteInterval"
	case TypeHourInterval:
		return "HourInterval"
	case TypeDayInterval:
		return "DayInterval"
	case TypeWeekInterval:
		return "WeekInterval"
	case TypeMonthInterval:
		return "MonthInterval"
	case TypeQuarterInterval:
		return "QuarterInterval"
	case TypeYearInterval:
		return "YearInterval"
	default:
		return "Unknown"
	}
}

type TypeDetail struct {
	Type        Type   `json:"type"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name"`
}

var AllType = []*TypeDetail{
	{TypeSecond, TypeSecond.String(), TypeSecond.StringEnglish()},
	{TypeMinute, TypeMinute.String(), TypeMinute.StringEnglish()},
	{TypeHour, TypeHour.String(), TypeHour.StringEnglish()},
	{TypeDaily, TypeDaily.String(), TypeDaily.StringEnglish()},
	{TypeWeekly, TypeWeekly.String(), TypeWeekly.StringEnglish()},
	{TypeMonthly, TypeMonthly.String(), TypeMonthly.StringEnglish()},
	{TypeQuarterly, TypeQuarterly.String(), TypeQuarterly.StringEnglish()},
	{TypeYearly, TypeYearly.String(), TypeYearly.StringEnglish()},
	{TypeSecondInterval, TypeSecondInterval.String(), TypeSecondInterval.StringEnglish()},
	{TypeMinuteInterval, TypeMinuteInterval.String(), TypeMinuteInterval.StringEnglish()},
	{TypeHourInterval, TypeHourInterval.String(), TypeHourInterval.StringEnglish()},
	{TypeDayInterval, TypeDayInterval.String(), TypeDayInterval.StringEnglish()},
	{TypeWeekInterval, TypeWeekInterval.String(), TypeWeekInterval.StringEnglish()},
	{TypeMonthInterval, TypeMonthInterval.String(), TypeMonthInterval.StringEnglish()},
	{TypeQuarterInterval, TypeQuarterInterval.String(), TypeQuarterInterval.StringEnglish()},
	{TypeYearInterval, TypeYearInterval.String(), TypeYearInterval.StringEnglish()},
}
