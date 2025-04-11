package dto

import (
	"msg/cmd/app/server/common/period"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/send"
)

type ScheduleCreate struct {
	Category        string      `json:"category" form:"category" query:"category"`          // 分类
	Type            send.Type   `json:"type" form:"type" query:"type"`                      // 消息类型
	Title           string      `json:"title" form:"title" query:"title"`                   // 标题
	Message         string      `json:"message" form:"message" query:"message"`             // 信息
	ChannelID       int64       `json:"channel_id" form:"channel_id" query:"channel_id"`    // 通知渠道ID
	PeriodType      period.Type `json:"period_type" form:"period_type" query:"period_type"` // 循环类型
	StartAt         int64       `json:"start_at" form:"start_at" query:"start_at"`          // 开始时间
	Year            int         `json:"year" form:"year" query:"year"`
	Quarter         int         `json:"quarter" form:"quarter" query:"quarter"`
	Month           int         `json:"month" form:"month" query:"month"`
	Week            int         `json:"week" form:"week" query:"week"`
	Day             int         `json:"day" form:"day" query:"day"`
	Hour            int         `json:"hour" form:"hour" query:"hour"`
	Minute          int         `json:"minute" form:"minute" query:"minute"`
	Second          int         `json:"second" form:"second" query:"second"`
	ExpirationDate  int64       `json:"expiration_date" form:"expiration_date" query:"expiration_date"`    // 过期时间
	ExpirationTimes int64       `json:"expiration_times" form:"expiration_times" query:"expiration_times"` // 过期次数
}

type ScheduleFind struct {
	resp.PageModel
	Id
	All    bool   `json:"all" form:"all" query:"all"`
	Search string `json:"search" form:"search" query:"search"`
}

type ScheduleUpdate struct {
	Id
	ScheduleCreate
}

type ScheduleUpdateNext struct {
	Id
}

type ScheduleUpdateSequence struct {
	Id
	Target int64 `json:"target" form:"target" query:"target"` // 目标序号
}

type ScheduleDisable struct {
	Id
}

type ScheduleDelete struct {
	Id
}
