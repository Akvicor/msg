package dro

type SysHealth struct {
	Status   string   `json:"status"`
	BotError []string `json:"bot_error,omitempty"`
}
