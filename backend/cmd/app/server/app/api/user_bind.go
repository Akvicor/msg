package api

import (
	"github.com/labstack/echo/v4"
	"msg/cmd/app/server/app/dto"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/global/auth"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/service"
)

var UserBind = new(userBindApi)

type userBindApi struct{}

// HomeTipsFind 获取用户主页提示信息
func (u *userBindApi) HomeTipsFind(c echo.Context) error {
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	tips, err := service.UserBindHomeTips.FindByUID(false, nil, user.ID)
	if err != nil {
		tips = model.NewUserBindHomeTips(user.ID, "")
	}
	return resp.SuccessWithData(c, tips)
}

// HomeTipsSave 保存用户主页提示信息
func (u *userBindApi) HomeTipsSave(c echo.Context) (err error) {
	input := new(dto.UserBindHomeTipsSave)
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
	_, err = service.UserBindHomeTips.Save(user.ID, input.Content)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "保存失败: "+err.Error())
	}
	return resp.Success(c)
}
