package api

import (
	"github.com/labstack/echo/v4"
	"msg/cmd/app/server/app/dro"
	"msg/cmd/app/server/bot"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/channel"
	"msg/cmd/app/server/common/types/role"
	"msg/cmd/app/server/common/types/send"
)

var Type = new(typeApi)

type typeApi struct{}

func (a *typeApi) RoleType(c echo.Context) (err error) {
	return resp.SuccessWithData(c, role.AllType)
}

func (a *typeApi) ChannelType(c echo.Context) (err error) {
	return resp.SuccessWithData(c, channel.AllType)
}

func (a *typeApi) SendType(c echo.Context) (err error) {
	return resp.SuccessWithData(c, send.AllType)
}

func (a *typeApi) BotSenders(c echo.Context) (err error) {
	bots := bot.Bot.List()
	senders := make([]*dro.BotSender, 0, len(bots))
	for _, b := range bots {
		sender := &dro.BotSender{
			Key:    b.Key(),
			Name:   b.Name(),
			Status: make([]*dro.BotSenderStatus, 0, len(channel.AllType)),
		}
		for _, cType := range channel.AllType {
			senderStatus, receiverStatus := b.Status(cType.Type)
			sender.Status = append(sender.Status, &dro.BotSenderStatus{
				ChannelType:        string(cType.Type),
				ChannelName:        cType.Name,
				ChannelEnglishName: cType.EnglishName,
				SenderStatus:       senderStatus,
				SenderStatusStr:    senderStatus.ToString(),
				ReceiverStatus:     receiverStatus,
				ReceiverStatusStr:  receiverStatus.ToString(),
			})
		}
		senders = append(senders, sender)
	}
	return resp.SuccessWithData(c, senders)
}
