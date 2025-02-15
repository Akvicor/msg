package repository

import (
	"context"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/model"
)

var LoginLog = new(loginLogRepository)

type loginLogRepository struct {
	base[*model.LoginLog]
}

/**
查找
*/

// FindAll 获取所有的登录记录, page为nil时不分页
func (l *loginLogRepository) FindAll(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderLoginLog) (logs []*model.LoginLog, err error) {
	return l.paging(page, l.preload(c, alive, preloader))
}

// FindAllByUID 获取登录成功但还未退出登录的指定用户的记录
func (l *loginLogRepository) FindAllByUID(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderLoginLog, uid int64) (logs []*model.LoginLog, err error) {
	return l.paging(page, l.preload(c, alive, preloader).Where("uid = ?", uid))
}

// FindAllByUIDLike 模糊搜索, page为nil时不分页
func (l *loginLogRepository) FindAllByUIDLike(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderLoginLog, uid int64, like string) (logs []*model.LoginLog, err error) {
	tx := l.preload(c, alive, preloader).Where("uid = ?", uid)

	if len(like) > 0 {
		like = "%" + like + "%"
		tx = tx.Where("token LIKE ? OR ip LIKE ?", like, like)
	}

	return l.paging(page, tx)
}

/**
创建
*/

// Create 创建登录记录
func (l *loginLogRepository) Create(c context.Context, log *model.LoginLog) error {
	return l.WrapResultErr(l.db(c).Create(log))
}

/**
更新
*/

// UpdateLogoutByToken 停用用户
func (l *loginLogRepository) UpdateLogoutByToken(c context.Context, alive bool, token string, timestamp int64) error {
	return l.WrapResultErr(l.alive(c, alive).Where("token = ?", token).UpdateColumn("logout_time", timestamp))
}

// UpdateLogoutByUID 停用用户
func (l *loginLogRepository) UpdateLogoutByUID(c context.Context, alive bool, uid, timestamp int64) error {
	return l.WrapResultErr(l.alive(c, alive).Where("uid = ?", uid).UpdateColumn("logout_time", timestamp))
}

// UpdateSuccess 更新成功结果
func (l *loginLogRepository) UpdateSuccess(c context.Context, alive bool, id int64, success bool) error {
	return l.WrapResultErr(l.alive(c, alive).Where("id = ?", id).UpdateColumn("success", success))
}
