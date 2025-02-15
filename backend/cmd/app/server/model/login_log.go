package model

import "gorm.io/gorm"

// LoginLog 登录日志
type LoginLog struct {
	ID         int64  `gorm:"column:id;primaryKey" json:"id"`
	UID        int64  `gorm:"column:uid;index;not null" json:"uid"`
	Token      string `gorm:"column:token;index;not null" json:"token"`
	IP         string `gorm:"column:ip;index" json:"ip"`                   // 登陆IP
	Agent      string `gorm:"column:agent" json:"agent"`                   // 浏览器Agent
	LoginTime  int64  `gorm:"column:login_time;index" json:"login_time"`   // 登录时间
	LogoutTime int64  `gorm:"column:logout_time;index" json:"logout_time"` // 登出时间
	Remember   bool   `gorm:"column:remember" json:"remember"`             // 是否记住登录
	Success    bool   `gorm:"column:success;index" json:"success"`         // 登录是否成功
	Reason     string `gorm:"column:reason" json:"reason"`                 // 登录结果原因

	User *User `gorm:"foreignKey:UID;references:ID" json:"user"`
}

func (*LoginLog) Alive(tx *gorm.DB) *gorm.DB {
	return tx.Where("success = ? AND logout_time = ?", true, 0)
}

func (*LoginLog) TableName() string {
	return "login_log"
}

func NewLoginLog(uid int64, token, ip, agent string, loginTime int64, remember, success bool, reason string) *LoginLog {
	return &LoginLog{
		ID:         0,
		UID:        uid,
		Token:      token,
		IP:         ip,
		Agent:      agent,
		LoginTime:  loginTime,
		LogoutTime: 0,
		Remember:   remember,
		Success:    success,
		Reason:     reason,
		User:       nil,
	}
}

/**
Preloader
*/

type PreloaderLoginLog struct {
	UserPreload            bool
	UserAccessTokenPreload bool
}

func NewPreloaderLoginLog() *PreloaderLoginLog {
	return &PreloaderLoginLog{
		UserPreload:            false,
		UserAccessTokenPreload: false,
	}
}

func (p *PreloaderLoginLog) Preload(tx *gorm.DB) *gorm.DB {
	if p.UserPreload {
		tx = tx.Preload("User")
	}
	if p.UserAccessTokenPreload {
		tx = tx.Preload("User.AccessToken")
	}
	return tx
}

func (p *PreloaderLoginLog) All() *PreloaderLoginLog {
	p.UserPreload = true
	p.UserAccessTokenPreload = true
	return p
}

func (p *PreloaderLoginLog) User() *PreloaderLoginLog {
	p.UserPreload = true
	return p
}

func (p *PreloaderLoginLog) UserAccessToken() *PreloaderLoginLog {
	p.UserAccessTokenPreload = true
	return p.User()
}
