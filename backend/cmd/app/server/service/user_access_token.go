package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/repository"
	"time"
)

var UserAccessToken = new(userAccessTokenService)

type userAccessTokenService struct {
	base
}

// FindAll 获取所有访问密钥
func (u *userAccessTokenService) FindAll(page *resp.PageModel, alive bool, preload *model.PreloaderAccessToken) (tokens []*model.UserAccessToken, err error) {
	return repository.UserAccessToken.FindAll(context.Background(), page, alive, preload)
}

// FindAllByUID 通过UID获取所有访问密钥
func (u *userAccessTokenService) FindAllByUID(page *resp.PageModel, alive bool, preload *model.PreloaderAccessToken, uid int64) (tokens []*model.UserAccessToken, err error) {
	return repository.UserAccessToken.FindAllByUID(context.Background(), page, alive, preload, uid)
}

// FindAllByUIDLike 通过UID获取所有访问密钥
func (u *userAccessTokenService) FindAllByUIDLike(page *resp.PageModel, alive bool, preload *model.PreloaderAccessToken, uid int64, like string) (tokens []*model.UserAccessToken, err error) {
	return repository.UserAccessToken.FindAllByUIDLike(context.Background(), page, alive, preload, uid, like)
}

// Create 创建用户访问密钥
func (u *userAccessTokenService) Create(uid int64, name, token string) (accessToken *model.UserAccessToken, err error) {
	accessToken = model.NewUserAccessToken(uid, name, token)
	err = u.transaction(context.Background(), func(ctx context.Context) error {
		// 判断用户是否存在且未被停用
		exist, e := repository.User.ExistById(ctx, true, uid)
		if !exist || errors.Is(e, gorm.ErrRecordNotFound) {
			return errors.New("用户已被停用")
		}
		// 创建访问密钥
		e = repository.UserAccessToken.Create(ctx, accessToken)
		if e != nil {
			return e
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

// UpdateName 更新名称
func (u *userAccessTokenService) UpdateName(alive bool, uid, id int64, name string) error {
	return repository.UserAccessToken.UpdateNameByUID(context.Background(), alive, uid, id, name)
}

// FindByToken 通过访问密钥获取信息
func (u *userAccessTokenService) FindByToken(alive bool, preload *model.PreloaderAccessToken, token string) (accessToken *model.UserAccessToken, err error) {
	return repository.UserAccessToken.FindByToken(context.Background(), alive, preload, token)
}

// UpdateLastUsed 更新上次使用时间
func (u *userAccessTokenService) UpdateLastUsed(alive bool, token string) error {
	return repository.UserAccessToken.UpdateLastUsedByToken(context.Background(), alive, token, time.Now().Unix())
}

// Delete 删除
func (u *userAccessTokenService) Delete(uid, id int64) error {
	return repository.UserAccessToken.Delete(context.Background(), uid, id)
}
