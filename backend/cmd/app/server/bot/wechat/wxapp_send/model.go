package wxapp_send

import (
	"github.com/patrickmn/go-cache"
)

type MessageTypeBaseModel struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Agentid string `json:"agentid"`
}

type MessageTypeBaseMarkdownModel struct {
	Content string `json:"content"`
}

type MessageTypeBaseTextCardModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	BtnTxt      string `json:"btntxt"`
}

type MessageTypeBaseTextModel struct {
	Content string `json:"content"`
}

type MessageTypeMarkdownModel struct {
	MessageTypeBaseModel
	Markdown MessageTypeBaseMarkdownModel `json:"markdown"`
}

type MessageTypeTextCardModel struct {
	MessageTypeBaseModel
	TextCard MessageTypeBaseTextCardModel `json:"textcard"`
}

type MessageTypeTextModel struct {
	MessageTypeBaseModel
	Text MessageTypeBaseTextModel `json:"text"`
}

type GetTokenResultModel struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type SendModel struct {
	cache      *cache.Cache
	CorpId     string
	CorpSecret string
	AgentId    int64
}

type PostJsonResultModel struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
}
