package model

import (
	"gorm.io/gorm"
	"msg/cmd/app/server/common/types/role"
)

// User 用户信息
type User struct {
	ID       int64     `gorm:"column:id;primaryKey" json:"id"`
	Username string    `gorm:"column:username;uniqueIndex;not null" json:"username"`
	Password string    `gorm:"column:password;not null" json:"-"`
	Nickname string    `gorm:"column:nickname;index" json:"nickname"`
	Avatar   string    `gorm:"column:avatar;index" json:"avatar"`
	Mail     string    `gorm:"column:mail;index" json:"mail"`
	Phone    string    `gorm:"column:phone;index" json:"phone"`
	Role     role.Type `gorm:"column:role;not null" json:"role"` // 用户身份（管理员, 普通用户, 浏览者
	Disabled int64     `gorm:"column:disabled" json:"disabled"`

	AccessToken []*UserAccessToken `gorm:"foreignKey:UID;references:ID" json:"access_token"` // 访问密钥,用户可以通过密钥访问API
}

func (*User) Alive(tx *gorm.DB) *gorm.DB {
	return tx.Where("disabled = ?", 0)
}

func (*User) TableName() string {
	return "user"
}

func NewUser(username, password, nickname, avatar, mail, phone string, rol role.Type) *User {
	return &User{
		ID:          0,
		Username:    username,
		Password:    password,
		Nickname:    nickname,
		Avatar:      avatar,
		Mail:        mail,
		Phone:       phone,
		Role:        rol,
		Disabled:    0,
		AccessToken: nil,
	}
}

/**
Preloader
*/

type PreloaderUser struct {
	AccessTokenPreload bool
}

func NewPreloaderUser() *PreloaderUser {
	return &PreloaderUser{
		AccessTokenPreload: false,
	}
}

func (p *PreloaderUser) Preload(tx *gorm.DB) *gorm.DB {
	if p.AccessTokenPreload {
		tx = tx.Preload("AccessToken")
	}
	return tx
}

func (p *PreloaderUser) All() *PreloaderUser {
	p.AccessTokenPreload = true
	return p
}

func (p *PreloaderUser) AccessToken() *PreloaderUser {
	p.AccessTokenPreload = true
	return p
}
