package model

import (
	"gorm.io/gorm"
	"msg/cmd/app/server/common/types/send"
)

// Send 通知渠道
type Send struct {
	ID         int64  `gorm:"column:id;primaryKey" json:"id"`
	ReadyAt    int64  `gorm:"column:ready_at;index" json:"ready_at"` // 创建时间
	SendAt     int64  `gorm:"column:send_at;index" json:"send_at"`   // 发送时间
	SentAt     int64  `gorm:"column:sent_at;index" json:"sent_at"`   // 发送成功时间
	UID        int64  `gorm:"column:uid;index" json:"uid"`
	ChannelID  int64  `gorm:"column:channel_id;index" json:"channel_id"`             // 通知渠道ID
	ScheduleID int64  `gorm:"column:schedule_id;index;default:0" json:"schedule_id"` // 绑定的定时通知
	IP         string `gorm:"column:ip;index" json:"ip"`                             // 调用者IP

	Type  send.Type `gorm:"column:type;index" json:"type"`
	Title string    `gorm:"column:title" json:"title"`
	Msg   string    `gorm:"column:msg" json:"msg"`

	ErrMsg string `gorm:"column:err_msg" json:"err_msg"` // 错误信息

	User     *User     `gorm:"foreignKey:UID;references:ID" json:"user"`
	Channel  *Channel  `gorm:"foreignKey:ChannelID;references:ID" json:"channel"`
	Schedule *Schedule `gorm:"foreignKey:ScheduleID;references:ID" json:"schedule"`
}

func (*Send) Alive(tx *gorm.DB) *gorm.DB {
	return tx
}

func (*Send) TableName() string {
	return "send"
}

func NewSend(readyAt, sendAt, sentAt int64, uid, channelID, scheduleID int64, ip string, cType send.Type, title, msg string) *Send {
	return &Send{
		ID:         0,
		ReadyAt:    readyAt,
		SendAt:     sendAt,
		SentAt:     sentAt,
		UID:        uid,
		ChannelID:  channelID,
		ScheduleID: scheduleID,
		IP:         ip,
		Type:       cType,
		Title:      title,
		Msg:        msg,
		ErrMsg:     "",
		User:       nil,
		Channel:    nil,
	}
}

func NewInternalSend(cType send.Type, title, msg string) *Send {
	return &Send{
		ID:         0,
		ReadyAt:    0,
		SendAt:     0,
		SentAt:     0,
		UID:        0,
		ChannelID:  0,
		ScheduleID: 0,
		IP:         "",
		Type:       cType,
		Title:      title,
		Msg:        msg,
		ErrMsg:     "",
		User:       nil,
		Channel:    nil,
	}
}

/**
Preloader
*/

type PreloaderSend struct {
	UserPreload     bool
	ChannelPreload  bool
	SchedulePreload bool
}

func NewPreloaderSend() *PreloaderSend {
	return &PreloaderSend{
		UserPreload:     false,
		ChannelPreload:  false,
		SchedulePreload: false,
	}
}

func (p *PreloaderSend) Preload(tx *gorm.DB) *gorm.DB {
	if p.UserPreload {
		tx = tx.Preload("User")
	}
	if p.ChannelPreload {
		tx = tx.Preload("Channel")
	}
	if p.SchedulePreload {
		tx = tx.Preload("Schedule")
	}
	return tx
}

func (p *PreloaderSend) All() *PreloaderSend {
	p.UserPreload = true
	p.ChannelPreload = true
	p.SchedulePreload = true
	return p
}

func (p *PreloaderSend) User() *PreloaderSend {
	p.UserPreload = true
	return p
}

func (p *PreloaderSend) Channel() *PreloaderSend {
	p.ChannelPreload = true
	return p
}

func (p *PreloaderSend) Schedule() *PreloaderSend {
	p.SchedulePreload = true
	return p
}
