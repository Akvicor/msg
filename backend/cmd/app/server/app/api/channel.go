package api

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"msg/cmd/app/server/app/dto"
	"msg/cmd/app/server/bot"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/send"
	"msg/cmd/app/server/global/auth"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/service"
	"time"
)

var Channel = new(channelApi)

type channelApi struct{}

// Create 创建
func (a *channelApi) Create(c echo.Context) (err error) {
	input := new(dto.ChannelCreate)
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
	_, err = service.Channel.Create(user.ID, input.Sign, input.Name, input.Type, input.Bot, input.Target)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "创建失败: "+err.Error())
	}
	return resp.Success(c)
}

// Find 获取通知渠道的信息 (全部, 单个, 模糊搜索)
func (a *channelApi) Find(c echo.Context) (err error) {
	input := new(dto.ChannelFind)
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
	if input.ID == 0 {
		var channels = make([]*model.Channel, 0)
		channels, err = service.Channel.FindAllByUID(&input.PageModel, true, model.NewPreloaderChannel(), user.ID, input.Search)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.FailWithMsg(c, resp.Failed, "系统错误")
		}
		return resp.SuccessWithPageData(c, &input.PageModel, channels)
	} else {
		var channel *model.Channel
		channel, err = service.Channel.FindByUID(true, model.NewPreloaderChannel(), user.ID, input.ID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.FailWithMsg(c, resp.Failed, "系统错误")
		}
		return resp.SuccessWithData(c, channel)
	}
}

// Update 更新
func (a *channelApi) Update(c echo.Context) (err error) {
	input := new(dto.ChannelUpdate)
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
	_, err = service.Channel.Update(user.ID, input.ID, input.Sign, input.Name, input.Type, input.Bot, input.Target)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "更新失败: "+err.Error())
	}
	return resp.Success(c)
}

// Delete 删除
func (a *channelApi) Delete(c echo.Context) (err error) {
	input := new(dto.ChannelDelete)
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
	if input.ID != 0 {
		err = service.Channel.DeleteByUID(user.ID, input.ID)
	} else if input.Sign != "" {
		err = service.Channel.DeleteBySign(user.ID, input.Sign)
	}
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "删除失败: "+err.Error())
	}
	return resp.Success(c)
}

// Test 测试
func (a *channelApi) Test(c echo.Context) (err error) {
	input := new(dto.ChannelTest)
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
	var userChannel *model.Channel
	if input.ID != 0 {
		userChannel, err = service.Channel.FindByUID(false, nil, user.ID, input.ID)
	} else if input.Sign != "" {
		userChannel, err = service.Channel.FindByUIDSign(false, nil, user.ID, input.Sign)
	}
	if err != nil || userChannel == nil {
		return resp.FailWithMsg(c, resp.Failed, fmt.Sprintf("找不到渠道: %v", err))
	}
	err = bot.Bot.Send(userChannel.Bot, userChannel.Type, userChannel.Target, model.NewInternalSend(send.TypeText, "Test", fmt.Sprintf("Test on %s", time.Now().Format(time.DateTime))))
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, fmt.Sprintf("发送失败: %v", err))
	}
	return resp.Success(c)
}

// Send 发送
func (a *channelApi) Send(c echo.Context) (err error) {
	input := new(dto.ChannelSend)
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
	var userChannel *model.Channel
	if input.ID != 0 {
		userChannel, err = service.Channel.FindByUID(true, nil, user.ID, input.ID)
	} else if input.Sign != "" {
		userChannel, err = service.Channel.FindByUIDSign(true, nil, user.ID, input.Sign)
	}
	if err != nil || userChannel == nil {
		return resp.FailWithMsg(c, resp.Failed, fmt.Sprintf("找不到渠道: %v", err))
	}
	_, err = bot.Sender.Send(user.ID, userChannel.ID, 0, input.At, c.RealIP(), input.Type, input.Title, input.Msg)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, fmt.Sprintf("发送失败: %v", err))
	}
	return resp.Success(c)
}
