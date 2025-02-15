package model

import "gorm.io/gorm"

// UserAccessToken 用户访问密钥
type UserAccessToken struct {
	ID       int64  `gorm:"column:id;primaryKey" json:"id"`
	UID      int64  `gorm:"column:uid;uniqueIndex:uid_name;index;not null" json:"uid"`
	Name     string `gorm:"column:name;uniqueIndex:uid_name;index;not null" json:"name"`
	Token    string `gorm:"column:token;uniqueIndex;not null" json:"token"` // 通行证,用户可以访问的API
	LastUsed int64  `gorm:"column:last_used;index" json:"last_used"`
	Disabled int64  `gorm:"column:disabled;index" json:"disabled"`

	User *User `gorm:"foreignKey:UID;references:ID" json:"user"`
}

func (*UserAccessToken) Alive(tx *gorm.DB) *gorm.DB {
	return tx.Where("disabled = ?", 0)
}

func (*UserAccessToken) TableName() string {
	return "user_access_token"
}

func NewUserAccessToken(uid int64, name, token string) *UserAccessToken {
	return &UserAccessToken{
		ID:       0,
		UID:      uid,
		Name:     name,
		Token:    token,
		LastUsed: 0,
		Disabled: 0,
		User:     nil,
	}
}

/**
Preloader
*/

type PreloaderAccessToken struct {
	UserPreload bool
}

func NewPreloaderAccessToken() *PreloaderAccessToken {
	return &PreloaderAccessToken{
		UserPreload: false,
	}
}

func (p *PreloaderAccessToken) Preload(tx *gorm.DB) *gorm.DB {
	if p.UserPreload {
		tx = tx.Preload("User")
	}
	return tx
}

func (p *PreloaderAccessToken) All() *PreloaderAccessToken {
	p.UserPreload = true
	return p
}

func (p *PreloaderAccessToken) User() *PreloaderAccessToken {
	p.UserPreload = true
	return p
}
