package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/repository"
	"time"
)

var UserBindHomeTips = new(userBindHomeTipsService)

type userBindHomeTipsService struct {
	base
}

// FindByUID 获取用户的记录
func (u *userBindHomeTipsService) FindByUID(alive bool, preload *model.PreloaderAccessToken, uid int64) (tips *model.UserBindHomeTips, err error) {
	return repository.UserBindHomeTips.FindByUID(context.Background(), alive, preload, uid)
}

// Save 保存记录
func (u *userBindHomeTipsService) Save(uid int64, content string) (tips *model.UserBindHomeTips, err error) {
	err = u.transaction(context.Background(), func(ctx context.Context) error {
		// 判断用户是否存在且未被停用
		exist, e := repository.User.ExistById(ctx, true, uid)
		if !exist || errors.Is(e, gorm.ErrRecordNotFound) {
			return errors.New("用户已被停用")
		}
		// 查询记录
		tips, e = repository.UserBindHomeTips.FindByUID(ctx, true, nil, uid)
		if e != nil {
			tips = model.NewUserBindHomeTips(uid, content)
			tips.LastUsed = time.Now().Unix()
		} else {
			tips.Content = content
			tips.LastUsed = time.Now().Unix()
		}
		// 保存
		e = repository.UserBindHomeTips.Save(ctx, tips)
		if e != nil {
			return e
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tips, nil
}
