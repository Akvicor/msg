package send

type Type string

const (
	TypeText     Type = "text"
	TypeTextCard Type = "textcard" // 微信卡片信息
	TypeMarkdown Type = "markdown"
	TypeHTML     Type = "html"
)

func (t Type) Valid() bool {
	switch t {
	case TypeText, TypeTextCard, TypeMarkdown, TypeHTML:
		return true
	default:
		return false
	}
}

func (t Type) ToString() string {
	switch t {
	case TypeText:
		return "Text"
	case TypeTextCard:
		return "TextCard"
	case TypeMarkdown:
		return "Markdown"
	case TypeHTML:
		return "HTML"
	default:
		return "未知"
	}
}

func (t Type) ToStringEnglish() string {
	switch t {
	case TypeText:
		return "Text"
	case TypeTextCard:
		return "TextCard"
	case TypeMarkdown:
		return "Markdown"
	case TypeHTML:
		return "HTML"
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
	{TypeText, TypeText.ToString(), TypeText.ToStringEnglish()},
	{TypeTextCard, TypeTextCard.ToString(), TypeTextCard.ToStringEnglish()},
	{TypeMarkdown, TypeMarkdown.ToString(), TypeMarkdown.ToStringEnglish()},
	{TypeHTML, TypeHTML.ToString(), TypeHTML.ToStringEnglish()},
}
