package dto

import (
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/channel"
	"msg/cmd/app/server/common/types/send"
)

type ChannelCreate struct {
	Sign   string       `json:"sign" form:"sign" query:"sign"`       // 唯一标记
	Name   string       `json:"name" form:"name" query:"name"`       // 名字
	Type   channel.Type `json:"type" form:"type" query:"type"`       // 渠道类型
	Bot    string       `json:"bot" form:"bot" query:"bot"`          // 发送者Bot
	Target string       `json:"target" form:"target" query:"target"` // 目标
}

type ChannelFind struct {
	resp.PageModel
	Id
	Search string `json:"search" form:"search" query:"search"`
}

type ChannelUpdate struct {
	Id
	ChannelCreate
}

type ChannelDelete struct {
	Id
	Sign string `json:"sign" form:"sign" query:"sign"`
}

type ChannelTest struct {
	Id
	Sign string `json:"sign" form:"sign" query:"sign"`
}

type ChannelSend struct {
	Id
	Sign  string    `json:"sign" form:"sign" query:"sign"` // 唯一标记
	Type  send.Type `json:"type" form:"type" query:"type"` // 消息类型
	At    int64     `json:"at" form:"at" query:"at"`       // SendAt
	Title string    `json:"title" form:"title" query:"title"`
	Msg   string    `json:"msg" form:"msg" query:"msg"`
}
