package model

import (
	"gorm.io/gorm"
	"msg/cmd/app/server/common/types/channel"
)

// Channel 通知渠道
type Channel struct {
	ID       int64        `gorm:"column:id;primaryKey" json:"id"`
	UID      int64        `gorm:"column:uid;index" json:"uid"`
	Sign     string       `gorm:"column:sign;uniqueIndex" json:"sign"` // 唯一标记，可以根据sign向对应Channel发送通知
	Name     string       `gorm:"column:name" json:"name"`
	Type     channel.Type `gorm:"column:type;index" json:"type"` // 通知渠道类型
	Bot      string       `gorm:"column:bot" json:"bot"`         // 发送的Bot
	Target   string       `gorm:"column:target" json:"target"`   // 发送目标(phone,email,telegramChatID,wechatUser)
	Disabled int64        `gorm:"column:disabled" json:"disabled"`

	User *User `gorm:"foreignKey:UID;references:ID" json:"user"`
}

func (*Channel) Alive(tx *gorm.DB) *gorm.DB {
	return tx.Where("disabled = ?", 0)
}

func (*Channel) TableName() string {
	return "channel"
}

func NewChannel(uid int64, sign, name string, channelType channel.Type, bot, target string) *Channel {
	return &Channel{
		ID:       0,
		UID:      uid,
		Sign:     sign,
		Name:     name,
		Type:     channelType,
		Bot:      bot,
		Target:   target,
		Disabled: 0,
		User:     nil,
	}
}

/**
Preloader
*/

type PreloaderChannel struct {
	UserPreload bool
}

func NewPreloaderChannel() *PreloaderChannel {
	return &PreloaderChannel{
		UserPreload: false,
	}
}

func (p *PreloaderChannel) Preload(tx *gorm.DB) *gorm.DB {
	if p.UserPreload {
		tx = tx.Preload("User")
	}
	return tx
}

func (p *PreloaderChannel) All() *PreloaderChannel {
	p.UserPreload = true
	return p
}

func (p *PreloaderChannel) User() *PreloaderChannel {
	p.UserPreload = true
	return p
}
