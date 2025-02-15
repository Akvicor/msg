package auth

import (
	"github.com/Akvicor/glog"
	"msg/cmd/app/server/common/cache"
	"msg/cmd/app/server/common/token"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/service"
)

// Authorization 用户访问信息
type Authorization struct {
	Token    string
	Type     token.Type
	Remember bool
	User     *model.User
	Log      *model.LoginLog
	Session  *Session
}

// UpdateUser 更新用户信息
func (a *Authorization) UpdateUser(user *model.User) {
	a.User = user
}

// Refresh 刷新用户信息
func (a *Authorization) Refresh() {
	user, err := service.User.FindByID(true, model.NewPreloaderUser().All(), a.User.ID)
	if err != nil {
		a.Delete()
		return
	}
	a.UpdateUser(user)
	// 更新AccessToken对应的数据
	for _, accessToken := range user.AccessToken {
		authorization := GetCacheAuthorizationByToken(accessToken.Token)
		if authorization != nil {
			authorization.UpdateUser(user)
		}
	}
}

// UpdateExpiration 更新令牌过期时间
func (a *Authorization) UpdateExpiration() {
	if a.Type == token.TypeLoginToken {
		glog.Debug("UpdateExpiration LoginToken: [%s][%s]", a.User.Username, a.Token)
		if a.Remember {
			cache.TokenManager.Set(a.Token, a, cache.ExpirationRemember)
		} else {
			cache.TokenManager.Set(a.Token, a, cache.ExpirationNotRemember)
		}
	} else if a.Type == token.TypeAccessToken {
		glog.Debug("UpdateExpiration AccessToken: [%s][%s]", a.User.Username, a.Token)
		cache.TokenManager.Set(a.Token, a, cache.ExpirationAccessToken)
		_ = service.UserAccessToken.UpdateLastUsed(true, a.Token)
	} else {
		glog.Debug("UpdateExpiration: [%s][%s]", a.User.Username, a.Token)
		cache.TokenManager.Set(a.Token, a, 0)
	}
}

// Delete 删除令牌
func (a *Authorization) Delete() {
	cache.TokenManager.Delete(a.Token)
}

// NewLoginAuthorization 创建LoginToken的Authorization并写入Cache
func NewLoginAuthorization(tokenStr string, remember bool, user *model.User, log *model.LoginLog) *Authorization {
	authorization := &Authorization{
		Token:    tokenStr,
		Type:     token.TypeLoginToken,
		Remember: remember,
		User:     user,
		Log:      log,
		Session:  newSession(),
	}
	authorization.UpdateExpiration()
	return authorization
}

// NewAccessAuthorization 创建AccessToken的Authorization并写入Cache
func NewAccessAuthorization(tokenStr string, user *model.User) *Authorization {
	authorization := &Authorization{
		Token:    tokenStr,
		Type:     token.TypeAccessToken,
		Remember: true,
		User:     user,
		Log:      nil,
		Session:  nil,
	}
	authorization.UpdateExpiration()
	return authorization
}
