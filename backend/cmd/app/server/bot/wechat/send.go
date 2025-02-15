package wechat

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"msg/cmd/app/server/bot/status"
	"msg/cmd/app/server/bot/wechat/dto"
	"msg/cmd/app/server/bot/wechat/wxapp_send"
	"msg/cmd/app/server/common/resp"
)

const (
	MessageTypeText     = "text"
	MessageTypeTextCard = "textcard"
	MessageTypeMarkdown = "markdown"
)

type SendModel struct {
	wxappSend *wxapp_send.SendModel
}

func NewSendModel(corpID, corpSecret string, agentId int64) *SendModel {
	return &SendModel{
		wxappSend: wxapp_send.NewSendModel(corpID, corpSecret, agentId),
	}
}

func (s *SendModel) Status() status.SenderStatus {
	return status.SenderStatusOK
}

// Send 发送消息
func (s *SendModel) Send(to string, messageType, title, message, url, btn string) (err error) {
	return s.wxappSend.Send(to, messageType, title, message, url, btn)
}

// APISend 发送消息
func (s *SendModel) APISend(c echo.Context) (err error) {
	input := new(dto.SendModel)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, fmt.Sprintf("错误输入: %v", err))
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, fmt.Sprintf("错误输入: %v", err))
	}
	err = s.wxappSend.Send(input.Touser, input.Type, input.Title, input.Msg, input.Url, input.Btn)
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, fmt.Sprintf("%v", err))
	}
	return resp.Success(c)
}
