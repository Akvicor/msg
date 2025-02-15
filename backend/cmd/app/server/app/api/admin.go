package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"msg/cmd/app/server/app/dto"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/global/auth"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/service"
)

var Admin = new(adminApi)

type adminApi struct{}

// CreateUser 创建用户
func (a *adminApi) CreateUser(c echo.Context) (err error) {
	input := new(dto.UserCreate)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	_, err = service.User.Create(input.Username, input.Password, input.Nickname, input.Avatar, input.Mail, input.Phone, input.Role)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "创建失败: "+err.Error())
	}
	return resp.Success(c)
}

// Find 获取用户的信息 (全部, 单个, 模糊搜索)
func (a *adminApi) Find(c echo.Context) (err error) {
	input := new(dto.UserFind)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	if input.ID == 0 {
		var users = make([]*model.User, 0)
		if len(input.Search) != 0 {
			users, err = service.User.FindAllLike(&input.PageModel, false, model.NewPreloaderUser().All(), input.Search)
		} else {
			users, err = service.User.FindAll(&input.PageModel, false, model.NewPreloaderUser().All())
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.FailWithMsg(c, resp.Failed, "系统错误")
		}
		return resp.SuccessWithPageData(c, &input.PageModel, users)
	} else {
		var user *model.User
		user, err = service.User.FindByID(false, model.NewPreloaderUser().All(), input.ID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.FailWithMsg(c, resp.Failed, "系统错误")
		}
		return resp.SuccessWithData(c, user)
	}
}

// UpdateUser 更新用户
func (a *adminApi) UpdateUser(c echo.Context) (err error) {
	input := new(dto.UserUpdate)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	err = service.User.Update(false, input.ID, input.Username, input.Password, input.Nickname, input.Avatar, input.Mail, input.Phone, input.Role)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "更新失败: "+err.Error())
	}
	return resp.Success(c)
}

// DisableUser 停用用户
func (a *adminApi) DisableUser(c echo.Context) (err error) {
	input := new(dto.UserDisableEnable)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入")
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	self := auth.GetUser(c)
	if self == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	if input.ID == self.ID || input.Username == self.Username {
		return resp.FailWithMsg(c, resp.BadRequest, "不能停用自己")
	}
	var user *model.User
	if input.ID != 0 {
		user, err = service.User.FindByID(true, nil, input.ID)
		if err != nil {
			return resp.FailWithMsg(c, resp.BadRequest, "未找到指定ID的活跃用户")
		}
	} else if len(input.Username) > 0 {
		user, err = service.User.FindByUsername(true, nil, input.Username)
		if err != nil {
			return resp.FailWithMsg(c, resp.BadRequest, "未找到指定用户名的活跃用户")
		}
	} else {
		return resp.FailWithMsg(c, resp.BadRequest, "请提供ID或用户名")
	}
	err = service.User.DisableById(user.ID)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "停用失败: "+err.Error())
	}

	return resp.Success(c)
}

func (a *adminApi) EnableUser(c echo.Context) (err error) {
	input := new(dto.UserDisableEnable)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入")
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	var user *model.User
	if input.ID != 0 {
		user, err = service.User.FindByID(false, nil, input.ID)
		if err != nil {
			return resp.FailWithMsg(c, resp.BadRequest, "未找到指定ID的活跃用户")
		}
	} else if len(input.Username) > 0 {
		user, err = service.User.FindByUsername(false, nil, input.Username)
		if err != nil {
			return resp.FailWithMsg(c, resp.BadRequest, "未找到指定用户名的活跃用户")
		}
	} else {
		return resp.FailWithMsg(c, resp.BadRequest, "请提供ID或用户名")
	}
	if user.Disabled == 0 {
		return resp.FailWithMsg(c, resp.BadRequest, "用户已经启用")
	}
	err = service.User.EnableById(user.ID)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "启用失败: "+err.Error())
	}

	return resp.Success(c)
}

// AccessTokenInfo AccessToken信息
func (a *adminApi) AccessTokenInfo(c echo.Context) error {
	tokens, err := service.UserAccessToken.FindAll(nil, false, model.NewPreloaderAccessToken().All())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return resp.FailWithMsg(c, resp.Failed, "系统错误")
	}
	return resp.SuccessWithData(c, tokens)
}

// LoginLogInfo 登录信息
func (a *adminApi) LoginLogInfo(c echo.Context) error {
	logs, err := service.LoginLog.FindAll(nil, false, model.NewPreloaderLoginLog().All())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return resp.FailWithMsg(c, resp.Failed, "系统错误")
	}
	return resp.SuccessWithData(c, logs)
}
