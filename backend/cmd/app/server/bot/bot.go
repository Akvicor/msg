package bot

import (
	"context"
	"errors"
	"github.com/Akvicor/glog"
	"github.com/labstack/echo/v4"
	"msg/cmd/app/server/app/mw"
	"msg/cmd/app/server/bot/status"
	"msg/cmd/app/server/common/types/channel"
	"msg/cmd/app/server/model"
	"sync"
)

const (
	SenderKeyMaid     = "maid"
	SenderKeyRosemary = "rosemary"
	SenderKeyReminder = "reminder"
)

var (
	ErrorNotFound    = errors.New("bot not found")
	ErrorMsgNotFound = errors.New("msg not found")
	ErrorMsgInvalid  = errors.New("msg invalid")
)

type Model struct {
	ready    *sync.WaitGroup
	status   map[string]SenderCommon
	list     []SenderCommon
	maid     *maidBot
	reminder *reminderBot
}

var Bot = newBot()

func newBot() *Model {
	bot := &Model{
		ready:    &sync.WaitGroup{},
		status:   make(map[string]SenderCommon),
		maid:     &maidBot{},
		reminder: &reminderBot{},
	}
	bot.ready.Add(1)
	bot.addSender(bot.maid)
	bot.addSender(bot.reminder)
	return bot
}

func (b *Model) addSender(sender SenderCommon) {
	b.status[sender.Key()] = sender
	b.list = append(b.list, sender)
}

func (b *Model) List() []SenderCommon {
	return b.list
}

func (b *Model) Run(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	bwg := &sync.WaitGroup{}
	b.maid.init(bwg, ctx)
	b.reminder.init(bwg, ctx)

	b.ready.Done()
	glog.Info("Bot OK")
	bwg.Wait()
	select {
	case <-ctx.Done():
	}
	glog.Info("Bot exist")
}

func (b *Model) Wait() {
	b.ready.Wait()
}

func (b *Model) Send(sender string, p channel.Type, target string, msg *model.Send) error {
	bot, ok := b.status[sender]
	if !ok {
		return ErrorNotFound
	}
	stat, _ := bot.Status(p)
	if stat != status.SenderStatusOK {
		return stat.ToError()
	}
	if msg.Title == "" && msg.Msg == "" {
		return ErrorMsgInvalid
	}
	switch p {
	case channel.TypeSMS:
		if msg.Msg == "" {
			msg.Msg = msg.Title
			msg.Title = ""
		}
	case channel.TypeMail:
		if msg.Title == "" {
			tMsg := []rune(msg.Msg)
			msg.Title = string(tMsg[:min(len(tMsg), 10)])
		}
	case channel.TypeTelegram:
		if msg.Msg == "" {
			msg.Msg = msg.Title
			msg.Title = ""
		}
	case channel.TypeWechat:
		if msg.Msg == "" {
			msg.Msg = msg.Title
			msg.Title = ""
		}
	}
	return bot.Send(p, target, msg)
}

func (b *Model) SetGroupRoute(api *echo.Group) {
	b.ready.Wait()

	// 管理员
	apiGroupAdmin := api.Group("")
	apiGroupAdmin.Use(mw.AuthAdmin)
	// 管理员、普通用户
	apiGroupUser := api.Group("")
	apiGroupUser.Use(mw.AuthUser)
	// 管理员、普通用户、浏览者
	apiGroupViewer := api.Group("")
	apiGroupViewer.Use(mw.AuthViewer)
	// 已登录用户
	apiGroupAuth := api.Group("")
	apiGroupAuth.Use(mw.Auth)
	// 公开
	apiGroupPublic := api.Group("")

	// maid
	{
		// SMS
		if b.maid.SMSSender != nil {
			apiGroupAuth.GET("/maid/sms/send", b.maid.SMSSender.APISend)
			apiGroupAuth.POST("/maid/sms/send", b.maid.SMSSender.APISend)
		}
		// Mail
		if b.maid.MailSender != nil {
			apiGroupAuth.GET("/maid/mail/send", b.maid.MailSender.APISend)
			apiGroupAuth.POST("/maid/mail/send", b.maid.MailSender.APISend)
		}
		// Telegram
		if b.maid.TelegramSender != nil {
			apiGroupAuth.GET("/maid/telegram/send", b.maid.TelegramSender.APISend)
			apiGroupAuth.POST("/maid/telegram/send", b.maid.TelegramSender.APISend)
		}
		// Wechat
		if b.maid.WechatReceiver != nil {
			apiGroupPublic.GET("/maid/wechat/callback", b.maid.WechatReceiver.Verify)
			apiGroupPublic.POST("/maid/wechat/callback", b.maid.WechatReceiver.Callback)
		}
		if b.maid.WechatSender != nil {
			apiGroupAuth.GET("/maid/wechat/send", b.maid.WechatSender.APISend)
			apiGroupAuth.POST("/maid/wechat/send", b.maid.WechatSender.APISend)
		}
	}
	// reminder
	{
		// Wechat
		if b.maid.WechatReceiver != nil {
			apiGroupPublic.GET("/reminder/wechat/callback", b.reminder.WechatReceiver.Verify)
			apiGroupPublic.POST("/reminder/wechat/callback", b.reminder.WechatReceiver.Callback)
		}
		if b.maid.WechatSender != nil {
			apiGroupAuth.GET("/reminder/wechat/send", b.reminder.WechatSender.APISend)
			apiGroupAuth.POST("/reminder/wechat/send", b.reminder.WechatSender.APISend)
		}
	}

}
