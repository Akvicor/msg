package service

import (
	"context"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/repository"
	"time"
)

var LoginLog = new(loginLogService)

type loginLogService struct {
	base
}

// FindAll 获取所有的登录记录, page为nil时不分页
func (l *loginLogService) FindAll(page *resp.PageModel, alive bool, preload *model.PreloaderLoginLog) (logs []*model.LoginLog, err error) {
	return repository.LoginLog.FindAll(context.Background(), page, alive, preload)
}

// FindAllByUID 获取所有的登录记录, page为nil时不分页
func (l *loginLogService) FindAllByUID(page *resp.PageModel, alive bool, preload *model.PreloaderLoginLog, uid int64) (logs []*model.LoginLog, err error) {
	return repository.LoginLog.FindAllByUID(context.Background(), page, alive, preload, uid)
}

// FindAllByUIDLike 获取所有的登录记录, page为nil时不分页
func (l *loginLogService) FindAllByUIDLike(page *resp.PageModel, alive bool, preload *model.PreloaderLoginLog, uid int64, like string) (logs []*model.LoginLog, err error) {
	return repository.LoginLog.FindAllByUIDLike(context.Background(), page, alive, preload, uid, like)
}

// Save 保存登录记录
func (l *loginLogService) Save(uid int64, token, ip, agent string, remember, success bool, reason string) (log *model.LoginLog, err error) {
	if uid < 0 {
		uid = 0
	}
	log = model.NewLoginLog(uid, token, ip, agent, time.Now().Unix(), remember, success, reason)
	err = repository.LoginLog.Create(context.Background(), log)
	if err != nil {
		return nil, err
	}
	return log, err
}

// FindAllAlive 获取活跃的登录记录
func (l *loginLogService) FindAllAlive(preload *model.PreloaderLoginLog) (logs []*model.LoginLog, err error) {
	return repository.LoginLog.FindAll(context.Background(), nil, true, preload)
}

// FindAllAliveByUID 通过UID获取活跃的登录记录
func (l *loginLogService) FindAllAliveByUID(preload *model.PreloaderLoginLog, uid int64) (logs []*model.LoginLog, err error) {
	return repository.LoginLog.FindAllByUID(context.Background(), nil, true, preload, uid)
}

// UpdateSuccess 更新成功结果
func (l *loginLogService) UpdateSuccess(alive bool, id int64, success bool) error {
	return repository.LoginLog.UpdateSuccess(context.Background(), alive, id, success)
}

// UpdateLogout 更新退出登录时间
func (l *loginLogService) UpdateLogout(alive bool, token string) error {
	return repository.LoginLog.UpdateLogoutByToken(context.Background(), alive, token, time.Now().Unix())
}
