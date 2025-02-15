package repository

import (
	"context"
	"msg/cmd/app/server/model"
)

var UserBindHomeTips = new(userBindHomeTipsRepository)

type userBindHomeTipsRepository struct {
	base[*model.UserBindHomeTips]
}

/**
查找
*/

// FindByUID 获取用户的记录
func (u *userBindHomeTipsRepository) FindByUID(c context.Context, alive bool, preloader *model.PreloaderAccessToken, uid int64) (tips *model.UserBindHomeTips, err error) {
	tips = new(model.UserBindHomeTips)
	err = u.WrapResultErr(u.preload(c, alive, preloader).Where("uid = ?", uid).First(tips))
	return tips, err
}

/**
保存
*/

// Save 保存记录
func (u *userBindHomeTipsRepository) Save(c context.Context, tips *model.UserBindHomeTips) error {
	return u.WrapResultErr(u.db(c).Save(tips))
}
