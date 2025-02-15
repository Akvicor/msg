package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"msg/cmd/app/server/bot/telegram/dto"
	"msg/cmd/app/server/common/resp"
)

const (
	ModeDefault    = ""
	ModeMarkdown   = tgbotapi.ModeMarkdown
	ModeMarkdownV2 = tgbotapi.ModeMarkdownV2
	ModeHTML       = tgbotapi.ModeHTML
)

func (b *Bot) Send(chatID int64, mode string, text string) error {
	var err error
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = mode
	msg.DisableWebPagePreview = true
	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

// APISend 发送消息
func (b *Bot) APISend(c echo.Context) (err error) {
	input := new(dto.SendModel)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, fmt.Sprintf("错误输入: %v", err))
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, fmt.Sprintf("错误输入: %v", err))
	}
	err = b.Send(input.Chat, input.Mode, input.Msg)
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, fmt.Sprintf("%v", err))
	}
	return resp.Success(c)
}
