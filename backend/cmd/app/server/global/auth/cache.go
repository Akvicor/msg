package auth

import (
	"errors"
	"github.com/Akvicor/glog"
	"gorm.io/gorm"
	"msg/cmd/app/server/common/token"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/service"
)

// LoadFromDBToCache 从登录记录中加载token
func LoadFromDBToCache() error {
	logs, err := service.LoginLog.FindAllAlive(model.NewPreloaderLoginLog().All())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	for _, log := range logs {
		if log.User == nil {
			glog.Error("登录记录[%d]未成功加载用户, UID[%d]", log.ID, log.UID)
			_ = service.LoginLog.UpdateSuccess(true, log.ID, false)
			continue
		}
		user := log.User
		log.User = nil
		_ = NewLoginAuthorization(log.Token, log.Remember, user, log)
		glog.Debug("加载用户登录Token, Nickname:%s, Token:%s", user.Nickname, log.Token)
	}
	return nil
}

// OnTokenEvicted 当Token过期或被删除时，标记用户为登出状态
func OnTokenEvicted(tokenStr string, value any) {
	authorization, ok := value.(*Authorization)
	if ok && authorization.Type == token.TypeLoginToken {
		glog.Debug("删除用户LoginToken: %s", tokenStr)
		_ = service.LoginLog.UpdateLogout(true, tokenStr)
	} else {
		glog.Debug("删除用户AccessToken: %s", tokenStr)
	}
}
