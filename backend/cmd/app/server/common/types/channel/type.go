package channel

type Type string

const (
	TypeSMS      Type = "sms"
	TypeMail     Type = "mail"
	TypeTelegram Type = "telegram"
	TypeWechat   Type = "wechat"
)

func (t Type) Valid() bool {
	switch t {
	case TypeSMS, TypeMail, TypeTelegram, TypeWechat:
		return true
	default:
		return false
	}
}

func (t Type) ToString() string {
	switch t {
	case TypeSMS:
		return "短信"
	case TypeMail:
		return "邮件"
	case TypeTelegram:
		return "Telegram"
	case TypeWechat:
		return "微信"
	default:
		return "未知"
	}
}

func (t Type) ToStringEnglish() string {
	switch t {
	case TypeSMS:
		return "SMS"
	case TypeMail:
		return "Mail"
	case TypeTelegram:
		return "Telegram"
	case TypeWechat:
		return "Wechat"
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
	{TypeSMS, TypeSMS.ToString(), TypeSMS.ToStringEnglish()},
	{TypeMail, TypeMail.ToString(), TypeMail.ToStringEnglish()},
	{TypeTelegram, TypeTelegram.ToString(), TypeTelegram.ToStringEnglish()},
	{TypeWechat, TypeWechat.ToString(), TypeWechat.ToStringEnglish()},
}
