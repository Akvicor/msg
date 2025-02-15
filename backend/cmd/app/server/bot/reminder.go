package bot

import (
	"context"
	"fmt"
	"github.com/Akvicor/glog"
	"msg/cmd/app/server/bot/status"
	"msg/cmd/app/server/bot/wechat"
	"msg/cmd/app/server/common/types/channel"
	"msg/cmd/app/server/common/types/send"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/service"
	"msg/cmd/config"
	"strconv"
	"strings"
	"sync"
	"time"
)

type reminderBot struct {
	WechatReceiver *wechat.CallbackModel // Reminder 微信企业程序回调
	WechatSender   *wechat.SendModel     // Reminder 企业微信程序发送
}

func (b *reminderBot) Key() string {
	return SenderKeyReminder
}

func (b *reminderBot) Name() string {
	return "Reminder"
}

func (b *reminderBot) Status(p channel.Type) (sender, receiver status.SenderStatus) {
	sender = status.SenderStatusNotSupported
	receiver = status.SenderStatusNotSupported
	switch p {
	case channel.TypeWechat:
		if config.Global.Bot.Maid.Wechat.EnableReceiver {
			if b.WechatReceiver == nil {
				receiver = status.SenderStatusConnecting
			} else {
				receiver = b.WechatReceiver.Status()
			}
		}
		if config.Global.Bot.Maid.Wechat.EnableSender {
			if b.WechatSender == nil {
				sender = status.SenderStatusConnecting
			} else {
				sender = b.WechatSender.Status()
			}
		}
	default:
	}
	return
}

func (b *reminderBot) Send(p channel.Type, target string, msg *model.Send) error {
	switch p {
	case channel.TypeWechat:
		if b.WechatSender == nil {
			if !config.Global.Bot.Reminder.Wechat.EnableSender {
				return status.SenderErrorNotSupported
			}
			return status.SenderErrorConnecting
		}
		msgType := wechat.MessageTypeText
		switch msg.Type {
		case send.TypeTextCard:
			msgType = wechat.MessageTypeTextCard
		case send.TypeMarkdown:
			msgType = wechat.MessageTypeMarkdown
		case send.TypeHTML:
			msgType = wechat.MessageTypeMarkdown
		}
		return b.WechatSender.Send(target, msgType, msg.Title, msg.Msg, "", "")
	default:
		return status.SenderErrorNotSupported
	}
}

func (b *reminderBot) init(*sync.WaitGroup, context.Context) {
	// Wechat
	{
		// Wechat Receiver
		if config.Global.Bot.Reminder.Wechat.EnableReceiver {
			glog.Info("Bot Reminder Wechat Receiver OK")
			b.WechatReceiver = wechat.NewCallbackModel(config.Global.Bot.Reminder.Wechat.CorpId, config.Global.Bot.Reminder.Wechat.Token, config.Global.Bot.Reminder.Wechat.AesKey, b.wechatReceiver)
		}
		// Wechat Sender
		if config.Global.Bot.Reminder.Wechat.EnableSender {
			glog.Info("Bot Reminder Wechat Sender OK")
			b.WechatSender = wechat.NewSendModel(config.Global.Bot.Reminder.Wechat.CorpId, config.Global.Bot.Reminder.Wechat.Secret, config.Global.Bot.Reminder.Wechat.AgentId)
		}
	}
	glog.Info("Bot Reminder OK")
}

func (b *reminderBot) wechatReceiver(toUserName, fromUsername, createTime, msgType, content, msgId string, agentId int64) (rToUserName, rFromUsername string, rCreateTime int64, rMsgType, rContent string) {
	content = strings.TrimSpace(content)

	rToUserName = fromUsername
	rFromUsername = b.WechatReceiver.CorpID
	rCreateTime = time.Now().Unix()
	rMsgType = "text"

	glog.Info("Reminder Received From [%s] To [%s] Created [%s] agent [%d] msgid [%s] type [%s] content [%s]", fromUsername, toUserName, createTime, agentId, msgId, msgType, content)

	fromChannel, err := service.Channel.FindByTypeTarget(true, nil, channel.TypeWechat, b.Key(), fromUsername)
	if err != nil {
		rContent = "您并未绑定此通知渠道"
		return
	}

	request := strings.Split(content, ":|:")
	for k := range request {
		request[k] = strings.TrimSpace(request[k])
	}

	rContent = ""
	func() {
		if len(request) <= 0 {
			return
		}
		switch strings.ToUpper(request[0]) {
		case "TODO":
			rContent = b.wechatReceiverTodo(fromChannel, request[1:])
		}
	}()
	if rContent == "" {
		rBuilder := strings.Builder{}
		rBuilder.WriteString("Tips")
		rBuilder.WriteString(fmt.Sprintf("\nTODO :|: %s 00:00:00 :|: 内容", time.Now().Format("2006-01-02")))
		rBuilder.WriteString(fmt.Sprintf("\nTODO :|: CANCEL :|: 序号"))
		rContent = rBuilder.String()
	}
	return
}

func (b *reminderBot) wechatReceiverTodo(fromChannel *model.Channel, args []string) string {
	if len(args) < 1 {
		return ""
	}

	if strings.ToUpper(args[0]) == "CANCEL" && len(args) > 1 {
		idStr := args[1]
		sendId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return fmt.Sprintf("[%s] is not valid id", idStr)
		}
		err = Sender.Cancel(fromChannel.UID, sendId)
		if err != nil {
			return fmt.Sprintf("cancel [%d] failed with %v", sendId, err.Error())
		}
		return fmt.Sprintf("Cancel [%d] Successful", sendId)
	}

	if len(args) < 2 {
		return ""
	}
	timeStr := args[0]
	content := args[1]

	targetTime, err := time.ParseInLocation(time.DateTime, timeStr, time.Local)
	if err != nil {
		return fmt.Sprintf("[%s] is not valid time", timeStr)
	}

	sendId, err := Sender.Send(fromChannel.UID, fromChannel.ID, targetTime.Unix(), "", send.TypeTextCard, content, timeStr)

	return fmt.Sprintf("Todo [%d] send at [%s]", sendId, targetTime.Format(time.DateTime))
}
