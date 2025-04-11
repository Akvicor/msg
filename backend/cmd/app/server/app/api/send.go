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
)

var Send = new(sendApi)

type sendApi struct{}

// Find 获取发送历史 (全部, 单个, 模糊搜索)
func (a *sendApi) Find(c echo.Context) (err error) {
	input := new(dto.SendFind)
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
		var sends = make([]*model.Send, 0)
		sends, err = service.Send.FindAllByUID(&input.PageModel, false, model.NewPreloaderSend().Channel(), user.ID, input.Sent, input.Search, input.ChannelIds, input.Types)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.FailWithMsg(c, resp.Failed, "系统错误")
		}
		return resp.SuccessWithPageData(c, &input.PageModel, sends)
	} else {
		var se *model.Send
		se, err = service.Send.FindByUID(false, model.NewPreloaderSend().Channel(), user.ID, input.ID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.FailWithMsg(c, resp.Failed, "系统错误")
		}
		return resp.SuccessWithData(c, se)
	}
}

// Send 发送信息
func (a *sendApi) Send(c echo.Context) (err error) {
	input := new(dto.Send)
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

	var channel *model.Channel
	if input.ID != 0 {
		channel, err = service.Channel.FindByUID(true, nil, user.ID, input.ID)
	} else {
		channel, err = service.Channel.FindByUIDSign(true, nil, user.ID, input.Sign)
	}
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, fmt.Sprintf("find channel failed: %v", err))
	}
	id, err := bot.Sender.Send(user.ID, channel.ID, 0, input.At, c.RealIP(), input.Type, input.Title, input.Msg)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, fmt.Sprintf("find channel failed: %v", err))
	}
	return resp.SuccessWithData(c, id)
}

// Cancel 取消发送信息
func (a *sendApi) Cancel(c echo.Context) (err error) {
	input := new(dto.SendCancel)
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

	err = bot.Sender.Cancel(user.ID, input.ID)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, fmt.Sprintf("cancel failed: %v", err))
	}

	return resp.Success(c)
}

// Status 信息状态
func (a *sendApi) Status(c echo.Context) (err error) {
	input := new(dto.SendStatus)
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

	msg, err := service.Send.FindByUID(false, nil, user.ID, input.ID)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, fmt.Sprintf("status failed: %v", err))
	}

	if msg.SentAt > 0 {
		return resp.SuccessWithData(c, send.NewStatusOK(msg.SendAt, msg.SentAt))
	} else if msg.SentAt == 0 {
		return resp.SuccessWithData(c, send.NewStatusWait(msg.SendAt, msg.SentAt))
	} else if msg.SentAt < 0 {
		if msg.ErrMsg == string(send.StatusCancel) {
			return resp.SuccessWithData(c, send.NewStatusCancel(msg.SendAt, msg.SentAt))
		} else {
			return resp.SuccessWithData(c, send.NewStatusFailed(msg.SendAt, msg.SentAt, msg.ErrMsg))
		}
	}

	return resp.Success(c)
}
