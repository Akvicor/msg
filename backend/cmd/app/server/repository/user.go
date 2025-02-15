package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/role"
	"msg/cmd/app/server/model"
)

var User = new(userRepository)

type userRepository struct {
	base[*model.User]
}

/**
查找
*/

// FindAll 获取所有的用户信息, page为nil时不分页
func (u *userRepository) FindAll(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderUser) (users []*model.User, err error) {
	return u.paging(page, u.preload(c, alive, preloader).Order("disabled ASC"))
}

// FindAllByUID 获取所有的用户信息, page为nil时不分页
func (u *userRepository) FindAllByUID(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderUser, uid int64) (users []*model.User, err error) {
	return u.paging(page, u.preload(c, alive, preloader).Where("uid = ?", uid).Order("disabled ASC"))
}

// FindAllLike 模糊搜索, page为nil时不分页
func (u *userRepository) FindAllLike(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderUser, like string) (users []*model.User, err error) {
	tx := u.preload(c, alive, preloader)

	if len(like) > 0 {
		like = "%" + like + "%"
		tx = tx.Where("username LIKE ? OR nickname LIKE ? OR mail LIKE ? OR phone LIKE ?", like, like, like, like)
	}

	return u.paging(page, tx.Order("disabled ASC"))
}

// FindAllLikeUsername 模糊搜索用户名
func (u *userRepository) FindAllLikeUsername(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderUser, username string) (users []*model.User, err error) {
	tx := u.preload(c, alive, preloader)

	if len(username) > 0 {
		tx = tx.Where("username LIKE ?", "%"+username+"%")
	}

	return u.paging(page, tx.Order("disabled ASC"))
}

// FindAllLikeNickname 模糊搜索用户昵称
func (u *userRepository) FindAllLikeNickname(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderUser, nickname string) (users []*model.User, err error) {
	tx := u.preload(c, alive, preloader)

	if len(nickname) > 0 {
		tx = tx.Where("nickname LIKE ?", "%"+nickname+"%")
	}

	return u.paging(page, tx.Order("disabled ASC"))
}

// FindById 通过ID获取用户
func (u *userRepository) FindById(c context.Context, alive bool, preloader *model.PreloaderUser, id int64) (user *model.User, err error) {
	user = new(model.User)
	err = u.WrapResultErr(u.preload(c, alive, preloader).Where("id = ?", id).First(user))
	return
}

// FindByUsername 通过Username获取用户
func (u *userRepository) FindByUsername(c context.Context, alive bool, preloader *model.PreloaderUser, username string) (user *model.User, err error) {
	user = new(model.User)
	err = u.WrapResultErr(u.preload(c, alive, preloader).Where("username = ?", username).First(user))
	return
}

// ExistByUsername 通过Username判断用户是否存在
func (u *userRepository) ExistByUsername(c context.Context, alive bool, username string) (exist bool, err error) {
	var count uint64
	err = u.alive(c, alive).Where("username = ?", username).Select("count(*)").Find(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistById 通过ID判断用户是否存在
func (u *userRepository) ExistById(c context.Context, alive bool, id int64) (exist bool, err error) {
	var count uint64
	err = u.alive(c, alive).Where("id = ?", id).Select("count(*)").Find(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Count 统计用户数量
func (u *userRepository) Count(c context.Context, alive bool) (total int64, err error) {
	err = u.alive(c, alive).Count(&total).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, err
		}
		return 0, nil
	}
	return total, nil
}

/**
创建
*/

// Create 创建用户
func (u *userRepository) Create(c context.Context, user *model.User) error {
	return u.WrapResultErr(u.db(c).Create(user))
}

/**
更新
*/

// Update 更新用户
func (u *userRepository) Update(c context.Context, alive bool, user *model.User) error {
	if user.Nickname == "" {
		user.Nickname = user.Username
	}
	tx := u.alive(c, alive).Select("*")
	omits := []string{"disabled"}
	if user.Password == "" {
		omits = append(omits, "password")
	}
	return u.WrapResultErr(tx.Omit(omits...).Where("id = ?", user.ID).Updates(user))
}

// UpdateUsernameByID 通过ID更新用户名
func (u *userRepository) UpdateUsernameByID(c context.Context, alive bool, id int64, username string) error {
	return u.WrapResultErr(u.alive(c, alive).Where("id = ?", id).UpdateColumn("username", username))
}

// UpdatePasswordByID 通过ID更新用户密码
func (u *userRepository) UpdatePasswordByID(c context.Context, alive bool, id int64, password string) error {
	return u.WrapResultErr(u.alive(c, alive).Where("id = ?", id).UpdateColumn("password", password))
}

// UpdateNicknameByID 通过ID更新用户昵称
func (u *userRepository) UpdateNicknameByID(c context.Context, alive bool, id int64, nickname string) error {
	return u.WrapResultErr(u.alive(c, alive).Where("id = ?", id).UpdateColumn("nickname", nickname))
}

// UpdateMailByID 通过ID更新用户邮箱
func (u *userRepository) UpdateMailByID(c context.Context, alive bool, id int64, mail string) error {
	return u.WrapResultErr(u.alive(c, alive).Where("id = ?", id).UpdateColumn("mail", mail))
}

// UpdatePhoneByID 通过ID更新用户手机
func (u *userRepository) UpdatePhoneByID(c context.Context, alive bool, id int64, phone string) error {
	return u.WrapResultErr(u.alive(c, alive).Where("id = ?", id).UpdateColumn("phone", phone))
}

// UpdateRoleByID 通过ID更新用户角色
func (u *userRepository) UpdateRoleByID(c context.Context, alive bool, id int64, rol role.Type) error {
	return u.WrapResultErr(u.alive(c, alive).Where("id = ?", id).UpdateColumn("role", rol))
}

// UpdateDisabledByID 通过ID更新用户是否停用
func (u *userRepository) UpdateDisabledByID(c context.Context, alive bool, id, timestamp int64) error {
	return u.WrapResultErr(u.alive(c, alive).Where("id = ?", id).UpdateColumn("disabled", timestamp))
}

// UpdateDisabledByUsername 通过Username更新用户是否停用
func (u *userRepository) UpdateDisabledByUsername(c context.Context, alive bool, username string, timestamp int64) error {
	return u.WrapResultErr(u.alive(c, alive).Where("username = ?", username).UpdateColumn("disabled", timestamp))
}
