package model

import (
	"gorm.io/gorm"
	"msg/cmd/app/server/common/period"
	"msg/cmd/app/server/common/types/send"
	"msg/cmd/app/server/common/utils"
	"time"
)

// Schedule 定时通知
type Schedule struct {
	ID              int64       `gorm:"column:id;primaryKey" json:"id"`
	UID             int64       `gorm:"column:uid;index;not null" json:"uid"`
	Category        string      `gorm:"column:category;not null" json:"category"`
	Type            send.Type   `gorm:"column:type" json:"type"`
	Title           string      `gorm:"column:title" json:"title"`
	Message         string      `gorm:"column:message" json:"message"`
	ChannelID       int64       `gorm:"column:channel_id;index" json:"channel_id"`   // 通知渠道ID
	PeriodType      period.Type `gorm:"column:period_type" json:"period_type"`       // 周期类型
	StartAt         int64       `gorm:"column:start_at" json:"start_at"`             // 第一次的时间
	NextOfPeriod    int64       `gorm:"column:next_of_period" json:"next_of_period"` // 下一次到期时间
	Year            int         `json:"year" gorm:"column:year"`
	Quarter         int         `json:"quarter" gorm:"column:quarter"`
	Month           int         `json:"month" gorm:"column:month"`
	Week            int         `json:"week" gorm:"column:week"`
	Day             int         `json:"day" gorm:"column:day"`
	Hour            int         `json:"hour" gorm:"column:hour"`
	Minute          int         `json:"minute" gorm:"column:minute"`
	Second          int         `json:"second" gorm:"column:second"`
	ExpirationDate  int64       `gorm:"column:expiration_date" json:"expiration_date"`   // 截止日期, -1表示禁用
	ExpirationTimes int64       `gorm:"column:expiration_times" json:"expiration_times"` // 截止次数, -1表示禁用
	Created         int64       `gorm:"column:created" json:"created"`                   // 创建时间
	Sequence        int64       `gorm:"column:sequence;index" json:"sequence"`           // 顺序
	Disabled        int64       `gorm:"column:disabled;index" json:"disabled"`           // 停用时间

	User    *User    `gorm:"foreignKey:UID;references:ID" json:"user"`
	Channel *Channel `gorm:"foreignKey:ChannelID;references:ID" json:"channel"`
}

func (*Schedule) Alive(tx *gorm.DB) *gorm.DB {
	return tx.Where("disabled = ?", 0)
}

func (*Schedule) TableName() string {
	return "schedule"
}

func NewSchedule(uid int64, category string, sendType send.Type, title, message string, channelID int64, periodType period.Type, startAt int64, year, quarter, month, week, day, hour, minute, second int, expirationDate, expirationTimes int64) *Schedule {
	item := &Schedule{
		ID:              0,
		UID:             uid,
		Category:        category,
		Type:            sendType,
		Title:           title,
		Message:         message,
		ChannelID:       channelID,
		PeriodType:      periodType,
		StartAt:         startAt,
		NextOfPeriod:    utils.StartAtOffsetFirst(periodType, startAt, false, year, quarter, month, week, day, hour, minute, second),
		Year:            year,
		Quarter:         quarter,
		Month:           month,
		Week:            week,
		Day:             day,
		Hour:            hour,
		Minute:          minute,
		Second:          second,
		ExpirationDate:  expirationDate,
		ExpirationTimes: expirationTimes,
		Created:         time.Now().Unix(),
		Sequence:        0,
		Disabled:        0,
		User:            nil,
		Channel:         nil,
	}

	return item
}

/**
Preloader
*/

type PreloaderSchedule struct {
	UserPreload    bool
	ChannelPreload bool
}

func NewPreloaderSchedule() *PreloaderSchedule {
	return &PreloaderSchedule{
		UserPreload:    false,
		ChannelPreload: false,
	}
}

func (p *PreloaderSchedule) Preload(tx *gorm.DB) *gorm.DB {
	if p.UserPreload {
		tx = tx.Preload("User")
	}
	if p.ChannelPreload {
		tx = tx.Preload("Channel")
	}
	return tx
}

func (p *PreloaderSchedule) All() *PreloaderSchedule {
	p.UserPreload = true
	p.ChannelPreload = true
	return p
}

func (p *PreloaderSchedule) User() *PreloaderSchedule {
	p.UserPreload = true
	return p
}

func (p *PreloaderSchedule) Channel() *PreloaderSchedule {
	p.ChannelPreload = true
	return p
}

/*
Function
*/

func (p *Schedule) NextPeriod() int64 {
	// 减少通知次数
	if p.ExpirationTimes > 0 {
		p.ExpirationTimes -= 1
	} else if p.ExpirationTimes == 0 {
		return p.NextOfPeriod
	}
	p.NextOfPeriod = utils.StartAtOffsetFirst(p.PeriodType, p.NextOfPeriod, true, p.Year, p.Quarter, p.Month, p.Week, p.Day, p.Hour, p.Minute, p.Second)
	nowT := time.Now().Unix()
	for p.NextOfPeriod < nowT {
		p.NextOfPeriod = utils.StartAtOffsetFirst(p.PeriodType, p.NextOfPeriod, true, p.Year, p.Quarter, p.Month, p.Week, p.Day, p.Hour, p.Minute, p.Second)
	}
	return p.NextOfPeriod
}
