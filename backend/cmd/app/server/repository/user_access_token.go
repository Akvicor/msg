package repository

import (
	"context"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/model"
)

var UserAccessToken = new(userAccessTokenRepository)

type userAccessTokenRepository struct {
	base[*model.UserAccessToken]
}

/**
查找
*/

// FindAll 获取所有的记录, page为nil时不分页
func (u *userAccessTokenRepository) FindAll(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderAccessToken) (tokens []*model.UserAccessToken, err error) {
	return u.paging(page, u.preload(c, alive, preloader).Order("disabled ASC"))
}

// FindAllByUID 获取所有的记录, page为nil时不分页
func (u *userAccessTokenRepository) FindAllByUID(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderAccessToken, uid int64) (tokens []*model.UserAccessToken, err error) {
	return u.paging(page, u.preload(c, alive, preloader).Where("uid = ?", uid).Order("disabled ASC"))
}

// FindAllByUIDLike 获取所有的记录, page为nil时不分页
func (u *userAccessTokenRepository) FindAllByUIDLike(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderAccessToken, uid int64, like string) (tokens []*model.UserAccessToken, err error) {
	tx := u.preload(c, alive, preloader).Where("uid = ?", uid)
	if len(like) > 0 {
		like = "%" + like + "%"
		tx = tx.Where("name LIKE ? OR token LIKE ?", like, like)
	}
	return u.paging(page, tx.Order("id ASC"))
}

// FindByToken 通过Token获取记录
func (u *userAccessTokenRepository) FindByToken(c context.Context, alive bool, preloader *model.PreloaderAccessToken, tokenStr string) (token *model.UserAccessToken, err error) {
	token = new(model.UserAccessToken)
	err = u.preload(c, alive, preloader).Where("token = ?", tokenStr).First(token).Error
	return token, err
}

/**
创建
*/

// Create 创建记录
func (u *userAccessTokenRepository) Create(c context.Context, token *model.UserAccessToken) error {
	return u.WrapResultErr(u.db(c).Create(token))
}

/**
更新
*/

// UpdateNameByUID 更新名称
func (u *userAccessTokenRepository) UpdateNameByUID(c context.Context, alive bool, uid, id int64, name string) error {
	return u.WrapResultErr(u.alive(c, alive).Where("uid = ? AND id = ?", uid, id).UpdateColumn("name", name))
}

// UpdateLastUsedByToken 更新上次使用时间
func (u *userAccessTokenRepository) UpdateLastUsedByToken(c context.Context, alive bool, tokenStr string, timestamp int64) error {
	return u.WrapResultErr(u.alive(c, alive).Where("token = ?", tokenStr).UpdateColumn("last_used", timestamp))
}

// UpdateDisabledByUID 更新停用时间
func (u *userAccessTokenRepository) UpdateDisabledByUID(c context.Context, alive bool, uid, timestamp int64) error {
	return u.WrapResultErr(u.alive(c, alive).Where("uid = ?", uid).UpdateColumn("disabled", timestamp))
}

/**
删除
*/

// Delete 删除记录
func (u *userAccessTokenRepository) Delete(c context.Context, uid, id int64) error {
	return u.WrapResultErr(u.db(c).Where("uid = ? AND id = ?", uid, id).Delete(&model.UserAccessToken{}))
}
