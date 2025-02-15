package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"msg/cmd/app/server/app/dto"
	"msg/cmd/app/server/common/cache"
	"msg/cmd/app/server/common/passwd"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/token"
	"msg/cmd/app/server/global/auth"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/service"
)

var User = new(userApi)

type userApi struct{}

// Login 登录
func (u *userApi) Login(c echo.Context) (err error) {
	input := new(dto.UserLogin)
	if err := c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入")
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	// 存储登录失败次数信息
	loginFailCount := cache.GetFailedCountByIpUsername(c.RealIP(), input.Username)
	if loginFailCount >= 5 {
		return resp.FailWithMsg(c, resp.TooManyRequests, "登录失败次数过多，请等待5分钟后再试")
	}
	// 获取用户信息
	user, err := service.User.Login(nil, input.Username)
	if err != nil {
		loginFailCount++
		cache.SetFailedCountByIpUsername(c.RealIP(), input.Username, loginFailCount)
		return resp.FailWithMsg(c, resp.NotFound, "账号已停用或您输入的账号或密码不正确")
	}
	// 判断密码是否匹配
	err = passwd.MatchString(user.Password, input.Password)
	if err != nil {
		loginFailCount++
		cache.SetFailedCountByIpUsername(c.RealIP(), input.Username, loginFailCount)
		_, _ = service.LoginLog.Save(user.ID, "", c.RealIP(), c.Request().UserAgent(), input.Remember, false, "密码不正确")
		return resp.FailWithMsg(c, resp.NotFound, "账号已停用或您输入的账号或密码不正确")
	}
	// 生成登录Token
	tokenStr := token.NewLoginToken(user.ID)
	log, err := service.LoginLog.Save(user.ID, tokenStr, c.RealIP(), c.Request().UserAgent(), input.Remember, true, "")
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "登录失败")
	}
	auth.NewLoginAuthorization(tokenStr, input.Remember, user, log)

	return resp.SuccessWithData(c, resp.Map{
		"token": tokenStr,
		"user":  user,
	})
}

// Logout 登出
func (u *userApi) Logout(c echo.Context) error {
	authorization := auth.GetAuthorization(c)
	authorization.Delete()
	return resp.Success(c)
}

// AccessTokenCreate 创建访问密钥
func (u *userApi) AccessTokenCreate(c echo.Context) (err error) {
	input := new(dto.UserCreateAccessToken)
	if err := c.Bind(input); err != nil || len(input.Name) <= 0 {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入")
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	authorization := auth.GetAuthorization(c)
	if authorization == nil {
		return resp.FailWithMsg(c, resp.BadRequest, "请登录")
	}
	user := authorization.User
	if user == nil {
		return resp.FailWithMsg(c, resp.BadRequest, "请登录")
	}
	defer authorization.Refresh()

	tokenStr := token.NewAccessToken(user.ID)
	accessToken, err := service.UserAccessToken.Create(user.ID, input.Name, tokenStr)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "创建失败")
	}
	return resp.SuccessWithData(c, accessToken.Token)
}

// AccessTokenUpdate 更新访问密钥
func (u *userApi) AccessTokenUpdate(c echo.Context) (err error) {
	input := new(dto.UserUpdateAccessToken)
	if err := c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入")
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.BadRequest, "请登录")
	}

	err = service.UserAccessToken.UpdateName(false, user.ID, input.ID, input.Name)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "更新失败")
	}
	return resp.Success(c)
}

// AccessTokenDelete 删除访问密钥
func (u *userApi) AccessTokenDelete(c echo.Context) (err error) {
	input := new(dto.UserDeleteAccessToken)
	if err := c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入")
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.BadRequest, "请登录")
	}

	err = service.UserAccessToken.Delete(user.ID, input.ID)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "删除失败")
	}
	return resp.Success(c)
}

// AccessTokenFind 用户AccessToken信息
func (u *userApi) AccessTokenFind(c echo.Context) (err error) {
	input := new(dto.UserFindAccessToken)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	tokens, err := service.UserAccessToken.FindAllByUIDLike(&input.PageModel, false, nil, user.ID, input.Search)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return resp.FailWithMsg(c, resp.Failed, "系统错误")
	}
	return resp.SuccessWithData(c, tokens)
}

// LoginLogInfo 用户登录信息
func (u *userApi) LoginLogInfo(c echo.Context) (err error) {
	input := new(dto.UserLoginLogFind)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	var logs = make([]*model.LoginLog, 0)
	if input.Search == "" {
		logs, err = service.LoginLog.FindAllByUID(&input.PageModel, false, nil, user.ID)
	} else {
		logs, err = service.LoginLog.FindAllByUIDLike(&input.PageModel, false, nil, user.ID, input.Search)
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return resp.FailWithMsg(c, resp.Failed, "系统错误")
	}
	return resp.SuccessWithData(c, logs)
}

// Info 用户信息
func (u *userApi) Info(c echo.Context) error {
	authorization := auth.GetAuthorization(c)
	if authorization == nil || authorization.User == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	user, err := service.User.FindByID(false, nil, authorization.User.ID)
	if err != nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	authorization.User = user
	return resp.SuccessWithData(c, user)
}

// Update 用户更新信息
func (u *userApi) Update(c echo.Context) (err error) {
	input := new(dto.UserInfoUpdate)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	err = service.User.Update(true, user.ID, input.Username, input.Password, input.Nickname, input.Avatar, input.Mail, input.Phone, user.Role)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "更新失败: "+err.Error())
	}
	return resp.Success(c)
}
