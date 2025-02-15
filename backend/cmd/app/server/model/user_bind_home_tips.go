package model

import (
	"gorm.io/gorm"
)

// UserBindHomeTips 用户主页提示信息
type UserBindHomeTips struct {
	ID       int64  `gorm:"column:id;primaryKey" json:"id"`
	UID      int64  `gorm:"column:uid;uniqueIndex;" json:"uid"`
	Content  string `gorm:"column:content" json:"content"`
	LastUsed int64  `gorm:"column:last_used;index" json:"last_used"`

	User *User `gorm:"foreignKey:UID;references:ID" json:"user"`
}

func (*UserBindHomeTips) Alive(tx *gorm.DB) *gorm.DB {
	return tx
}

func (*UserBindHomeTips) TableName() string {
	return "user_bind_home_tips"
}

func NewUserBindHomeTips(uid int64, content string) *UserBindHomeTips {
	return &UserBindHomeTips{
		ID:       0,
		UID:      uid,
		Content:  content,
		LastUsed: 0,
		User:     nil,
	}
}

/**
Preloader
*/

type PreloaderUserBindHomeTips struct {
	UserPreload               bool
	UserMasterCurrencyPreload bool
	UserAccessTokenPreload    bool
}

func NewPreloaderUserBindHomeTips() *PreloaderUserBindHomeTips {
	return &PreloaderUserBindHomeTips{
		UserPreload:               false,
		UserMasterCurrencyPreload: false,
		UserAccessTokenPreload:    false,
	}
}

func (p *PreloaderUserBindHomeTips) Preload(tx *gorm.DB) *gorm.DB {
	if p.UserPreload {
		tx = tx.Preload("User")
		if p.UserMasterCurrencyPreload {
			tx = tx.Preload("User.MasterCurrency")
		}
		if p.UserAccessTokenPreload {
			tx = tx.Preload("User.AccessToken")
		}
	}
	return tx
}

func (p *PreloaderUserBindHomeTips) All() *PreloaderUserBindHomeTips {
	p.UserPreload = true
	p.UserMasterCurrencyPreload = true
	p.UserAccessTokenPreload = true
	return p
}

func (p *PreloaderUserBindHomeTips) User() *PreloaderUserBindHomeTips {
	p.UserPreload = true
	return p
}

func (p *PreloaderUserBindHomeTips) UserMasterCurrency() *PreloaderUserBindHomeTips {
	p.UserMasterCurrencyPreload = true
	return p.User()
}

func (p *PreloaderUserBindHomeTips) UserAccessToken() *PreloaderUserBindHomeTips {
	p.UserAccessTokenPreload = true
	return p.User()
}
