package service

import (
	"context"
	"errors"
	"fmt"
	"msg/cmd/app/server/common/cache"
	"msg/cmd/app/server/common/passwd"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/role"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/repository"
	"time"
)

var User = new(userService)

type userService struct {
	base
}

// FindAll 获取所有的用户信息, page为nil时不分页
func (u *userService) FindAll(page *resp.PageModel, alive bool, preload *model.PreloaderUser) (users []*model.User, err error) {
	return repository.User.FindAll(context.Background(), page, alive, preload)
}

// FindAllLike 搜索符合条件的所有的用户信息, page为nil时不分页
func (u *userService) FindAllLike(page *resp.PageModel, alive bool, preload *model.PreloaderUser, like string) (users []*model.User, err error) {
	return repository.User.FindAllLike(context.Background(), page, alive, preload, like)
}

// FindByID 获取指定ID用户信息
func (u *userService) FindByID(alive bool, preload *model.PreloaderUser, id int64) (user *model.User, err error) {
	return repository.User.FindById(context.Background(), alive, preload, id)
}

// FindByUsername 通过Username获取用户
func (u *userService) FindByUsername(alive bool, preload *model.PreloaderUser, username string) (user *model.User, err error) {
	return repository.User.FindByUsername(context.Background(), alive, preload, username)
}

// Login 通过用户名获取登录用户
func (u *userService) Login(preload *model.PreloaderUser, username string) (user *model.User, err error) {
	user, err = repository.User.FindByUsername(context.Background(), true, preload, username)
	if err != nil {
		return nil, err
	}
	return
}

// Create 创建用户
func (u *userService) Create(username, password, nickname, avatar, email, phone string, rol role.Type) (user *model.User, err error) {
	if len(username) == 0 || len(password) == 0 {
		return nil, errors.New("empty username or password")
	}
	password, err = passwd.Filter(password)
	if err != nil {
		return nil, err
	}
	var pass string
	pass, err = passwd.EncodeString(password)
	if err != nil {
		return nil, err
	}
	user = model.NewUser(username, pass, nickname, avatar, email, phone, rol)

	err = u.transaction(context.Background(), func(ctx context.Context) error {
		exist, err := repository.User.ExistByUsername(ctx, false, user.Username)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("username %s is already used", user.Username)
		}

		if err := repository.User.Create(ctx, user); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update 更新用户
func (u *userService) Update(alive bool, id int64, username, password, nickname, avatar, mail, phone string, rol role.Type) (err error) {
	if len(username) == 0 {
		return errors.New("empty username")
	}
	var pass = ""
	if len(password) > 0 {
		password, err = passwd.Filter(password)
		if err != nil {
			return err
		}
		pass, err = passwd.EncodeString(password)
		if err != nil {
			return err
		}
	}
	user := model.NewUser(username, pass, nickname, avatar, mail, phone, rol)
	user.ID = id
	return repository.User.Update(context.Background(), alive, user)
}

// DisableById 通过ID停用用户
func (u *userService) DisableById(id int64) (err error) {
	timestamp := time.Now().Unix()

	// 停用用户
	err = repository.User.UpdateDisabledByID(context.Background(), true, id, timestamp)
	if err != nil {
		return err
	}

	// 获取所有活跃访问密钥
	accessTokens, _ := repository.UserAccessToken.FindAllByUID(context.Background(), nil, true, nil, id)
	// 获取所有活跃登录记录
	logs, _ := repository.LoginLog.FindAllByUID(context.Background(), nil, true, nil, id)

	// 停用用户所有访问密钥
	_ = repository.UserAccessToken.UpdateDisabledByUID(context.Background(), true, id, timestamp)
	// 登出所有用户登录记录
	_ = repository.LoginLog.UpdateLogoutByUID(context.Background(), true, id, timestamp)

	// 处理所有活跃Token
	tokens := make([]string, 0, len(accessTokens)+len(logs))
	for _, access := range accessTokens {
		tokens = append(tokens, access.Token)
	}
	for _, log := range logs {
		tokens = append(tokens, log.Token)
	}
	cache.DeleteTokens(tokens)
	return nil
}

// EnableById 通过ID启用用户
func (u *userService) EnableById(id int64) (err error) {
	return repository.User.UpdateDisabledByID(context.Background(), false, id, 0)
}
