package dto

import (
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/send"
)

type SendFind struct {
	resp.PageModel
	Id
	Search     string      `json:"search" form:"search" query:"search"`
	Sent       []int64     `json:"sent" form:"sent" query:"sent"`
	ChannelIds []int64     `json:"channel_ids" form:"channel_ids" query:"channel_ids"`
	Types      []send.Type `json:"types" form:"types" query:"types"`
}

type Send struct {
	Id
	Sign  string    `json:"sign" form:"sign" query:"sign"` // 唯一标记
	Type  send.Type `json:"type" form:"type" query:"type"` // 消息类型
	At    int64     `json:"at" form:"at" query:"at"`       // SendAt
	Title string    `json:"title" form:"title" query:"title"`
	Msg   string    `json:"msg" form:"msg" query:"msg"`
}

type SendCancel struct {
	Id
}

type SendStatus struct {
	Id
}
