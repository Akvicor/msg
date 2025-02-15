package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Akvicor/glog"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"msg/cmd/app/server/bot/mail"
	"msg/cmd/app/server/bot/sms"
	"msg/cmd/app/server/bot/status"
	"msg/cmd/app/server/bot/telegram"
	"msg/cmd/app/server/bot/wechat"
	"msg/cmd/app/server/common/types/channel"
	"msg/cmd/app/server/common/types/send"
	"msg/cmd/app/server/model"
	"msg/cmd/config"
	"strconv"
	"strings"
	"sync"
	"time"
)

type maidBot struct {
	SMSSender        *sms.Model            // Maid 短信发送
	MailReceiver     *mail.IMAPModel       // Maid 邮件接收
	MailSender       *mail.SMTPModel       // Maid 邮件发送
	TelegramReceiver *telegram.Bot         // Maid Telegram接收
	TelegramSender   *telegram.Bot         // Maid Telegram发送
	WechatReceiver   *wechat.CallbackModel // Maid 微信企业程序回调
	WechatSender     *wechat.SendModel     // Maid 企业微信程序发送
}

func (b *maidBot) Key() string {
	return SenderKeyMaid
}

func (b *maidBot) Name() string {
	return "Maid"
}

func (b *maidBot) Status(p channel.Type) (sender, receiver status.SenderStatus) {
	sender = status.SenderStatusNotSupported
	receiver = status.SenderStatusNotSupported
	switch p {
	case channel.TypeSMS:
		if config.Global.Bot.Maid.SMS.EnableSender {
			if b.SMSSender == nil {
				sender = status.SenderStatusConnecting
			} else {
				sender = b.SMSSender.Status()
			}
		}
	case channel.TypeMail:
		if config.Global.Bot.Maid.Mail.EnableImap {
			if b.MailReceiver == nil {
				receiver = status.SenderStatusConnecting
			} else {
				receiver = b.MailReceiver.Status()
			}
		}
		if config.Global.Bot.Maid.Mail.EnableSmtp {
			if b.MailSender == nil {
				sender = status.SenderStatusConnecting
			} else {
				sender = b.MailSender.Status()
			}
		}
	case channel.TypeTelegram:
		if config.Global.Bot.Maid.Telegram.EnableReceiver {
			if b.TelegramReceiver == nil {
				receiver = status.SenderStatusConnecting
			} else {
				receiver = b.TelegramReceiver.Status()
			}
		}
		if config.Global.Bot.Maid.Telegram.EnableSender {
			if b.TelegramSender == nil {
				sender = status.SenderStatusConnecting
			} else {
				sender = b.TelegramSender.Status()
			}
		}
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

func (b *maidBot) Send(p channel.Type, target string, msg *model.Send) error {
	switch p {
	case channel.TypeSMS:
		if b.SMSSender == nil {
			if !config.Global.Bot.Maid.SMS.EnableSender {
				return status.SenderErrorNotSupported
			}
			return status.SenderErrorConnecting
		}
		return b.SMSSender.Send(target, msg.Msg)
	case channel.TypeMail:
		if b.MailSender == nil {
			if !config.Global.Bot.Maid.Mail.EnableSmtp {
				return status.SenderErrorNotSupported
			}
			return status.SenderErrorConnecting
		}
		msgType := mail.TypeTextPlain
		switch msg.Type {
		case send.TypeHTML:
			msgType = mail.TypeTextHtml
		}
		return b.MailSender.Send(target, msg.Title, msgType, msg.Msg)
	case channel.TypeTelegram:
		if b.TelegramSender == nil {
			if !config.Global.Bot.Maid.Telegram.EnableSender {
				return status.SenderErrorNotSupported
			}
			return status.SenderErrorConnecting
		}
		chatID, err := strconv.ParseInt(target, 10, 64)
		if err != nil {
			return status.SenderErrorWrongTarget
		}
		msgType := telegram.ModeDefault
		switch msg.Type {
		case send.TypeMarkdown:
			msgType = telegram.ModeMarkdownV2
		case send.TypeHTML:
			msgType = telegram.ModeHTML
		}
		return b.TelegramSender.Send(chatID, msgType, msg.Msg)
	case channel.TypeWechat:
		if b.WechatSender == nil {
			if !config.Global.Bot.Maid.Wechat.EnableSender {
				return status.SenderErrorNotSupported
			}
			return status.SenderErrorConnecting
		}
		msgType := wechat.MessageTypeText
		switch msg.Type {
		case send.TypeMarkdown:
			msgType = wechat.MessageTypeMarkdown
		case send.TypeTextCard:
			msgType = wechat.MessageTypeTextCard
		case send.TypeHTML:
			msgType = wechat.MessageTypeMarkdown
		}
		return b.WechatSender.Send(target, msgType, msg.Title, msg.Msg, "", "")
	default:
		return status.SenderErrorNotSupported
	}
}

func (b *maidBot) init(wg *sync.WaitGroup, ctx context.Context) {
	var err error

	// SMS
	{
		// SMS Sender
		if config.Global.Bot.Maid.SMS.EnableSender {
			glog.Info("Bot Maid SMS OK")
			b.SMSSender = sms.NewModel(config.Global.Bot.Maid.SMS.Api, config.Global.Bot.Maid.SMS.Token)
		}
	}

	// Mail
	{
		// Mail Receiver
		if config.Global.Bot.Maid.Mail.EnableImap {
			b.MailReceiver = mail.NewIMAPModel(config.Global.Bot.Maid.Mail.HostImap, config.Global.Bot.Maid.Mail.PortImap, config.Global.Bot.Maid.Mail.From, config.Global.Bot.Maid.Mail.Username, config.Global.Bot.Maid.Mail.Password)
			wg.Add(2)
			glog.Info("Bot Maid Mail Receiver OK")
			go b.MailReceiver.Listen(wg, ctx)
			go b.mailReceiver(wg, b.MailReceiver.Messages)
		}
		// Mail Sender
		if config.Global.Bot.Maid.Mail.EnableSmtp {
			glog.Info("Bot Maid Mail Sender OK")
			b.MailSender = mail.NewSMTPModel(config.Global.Bot.Maid.Mail.HostSmtp, config.Global.Bot.Maid.Mail.PortSmtp, config.Global.Bot.Maid.Mail.From, config.Global.Bot.Maid.Mail.Username, config.Global.Bot.Maid.Mail.Password)
		}
	}

	// Telegram
	if config.Global.Bot.Maid.Telegram.EnableReceiver || config.Global.Bot.Maid.Telegram.EnableSender {
		var tgbot *telegram.Bot

		wait := time.Duration(0)
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(wait):
				tgbot, err = telegram.NewBot(b.Key(), config.Global.Bot.Maid.Telegram.API, config.Global.Bot.Maid.Telegram.Token, config.Global.Bot.Maid.Telegram.Debug)
				if err != nil {
					glog.Warning("Bot Maid Telegram Connect Failed: %v", err)
					wait = 3 * time.Second
					continue
				}
			}
			break
		}

		// Telegram Receiver
		if config.Global.Bot.Maid.Telegram.EnableReceiver {
			b.TelegramReceiver = tgbot
			wg.Add(1)
			glog.Info("Bot Maid Telegram Receiver OK")
			go b.TelegramReceiver.Listen(wg, ctx, b.telegramReceiver)
		}
		// Telegram Sender
		if config.Global.Bot.Maid.Telegram.EnableSender {
			b.TelegramSender = tgbot
			glog.Info("Bot Maid Telegram Sender OK")
		}
	}

	// Wechat
	{
		// Wechat Receiver
		if config.Global.Bot.Maid.Wechat.EnableReceiver {
			glog.Info("Bot Maid Wechat Receiver OK")
			b.WechatReceiver = wechat.NewCallbackModel(config.Global.Bot.Maid.Wechat.CorpId, config.Global.Bot.Maid.Wechat.Token, config.Global.Bot.Maid.Wechat.AesKey, b.wechatReceiver)
		}
		// Wechat Sender
		if config.Global.Bot.Maid.Wechat.EnableSender {
			glog.Info("Bot Maid Wechat Sender OK")
			b.WechatSender = wechat.NewSendModel(config.Global.Bot.Maid.Wechat.CorpId, config.Global.Bot.Maid.Wechat.Secret, config.Global.Bot.Maid.Wechat.AgentId)
		}
	}
	glog.Info("Bot Maid OK")
}

func (b *maidBot) mailReceiver(wg *sync.WaitGroup, messages chan *mail.ImapEmail) {
	defer func() {
		glog.Warning("Mail Receiver exist")
		wg.Done()
	}()

	for msg := range messages {
		if msg.Envelope != nil {
			glog.Info("New Message: [%s] ID [%s] From [%v] To [%v]", msg.Envelope.Subject, msg.Envelope.MessageID, msg.Envelope.From, msg.Envelope.To)
		}
		if msg.Content != nil {
			glog.Info("    Content: Type [%s] Params [%v]", msg.Content.Type, msg.Content.Params)
			for _, c := range msg.Content.Contents {
				glog.Info("    Content: %s", c)
			}
		}
	}
}

func (b *maidBot) telegramReceiver(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	defer func() {
		glog.Warning("Telegram Receiver exist")
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case update := <-updates:
			if update.Message != nil { // Message new incoming message of any kind — text, photo, sticker, etc.
				msg := update.Message
				if len(msg.NewChatMembers) > 0 {
					for _, user := range msg.NewChatMembers {
						glog.Info("[%s] 对话[%s|%s|%d|%d] 用户[%s]邀请成员[%s]进入聊天", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.MessageID, msg.From.FirstName, user.FirstName)
					}
				} else if msg.GroupChatCreated {
					glog.Info("[%s] 用户[%s]创建群组[%s|%s|%d|%d]", time.Unix(int64(msg.Date), 0).String(), msg.From.FirstName, msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.MessageID)
				} else if msg.LeftChatMember != nil {
					glog.Info("[%s] 对话[%s|%s|%d|%d] 用户[%s]删除成员[%s]", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.MessageID, msg.From.FirstName, msg.LeftChatMember.FirstName)
				} else if msg.MigrateToChatID != 0 {
					glog.Info("[%s] 对话[%s|%s|%d|%d] 用户[%s]变更对话ID[%d]->[%d]", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.MessageID, msg.From.FirstName, msg.Chat.ID, msg.MigrateToChatID)
				} else if msg.IsCommand() {
					if msg.Command() == "/get_id" {
						err := b.TelegramReceiver.Send(msg.Chat.ID, telegram.ModeDefault, fmt.Sprintf("%d", msg.Chat.ID))
						if err != nil {
							glog.Warning("response chat id failed: %v", err)
						}
					}
					glog.Info("[%s] 对话[%s|%s|%d|%d] 用户[%s]发送命令[%s][%s]", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.MessageID, msg.From.FirstName, msg.Command(), msg.CommandArguments())
				} else if len(msg.Text) != 0 {
					entities := strings.Builder{}
					if len(msg.Entities) != 0 {
						entities.WriteString("标签[")
						for k, e := range msg.Entities {
							if k != 0 {
								entities.WriteString(",")
							}
							entities.WriteString(e.Type)
						}
						entities.WriteString("]")
					}
					glog.Info("[%s] 对话[%s|%s|%d|%d] 用户[%s]发送%s信息[%s]", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.MessageID, msg.From.FirstName, entities.String(), msg.Text)
				} else {
					d, _ := json.MarshalIndent(msg, "", "  ")
					glog.Info("Message: %s", string(d))
				}
			}
			if update.EditedMessage != nil { // EditedMessage new version of a message that is known to the bot and was edited
				d, _ := json.MarshalIndent(update.EditedMessage, "", "  ")
				glog.Info("EditedMessage: %s", string(d))
			}
			if update.ChannelPost != nil { // ChannelPost new version of a message that is known to the bot and was edited
				msg := update.ChannelPost
				if len(msg.Text) != 0 {
					if msg.Text == "/get_id" {
						err := b.TelegramReceiver.Send(msg.Chat.ID, telegram.ModeDefault, fmt.Sprintf("%d", msg.Chat.ID))
						if err != nil {
							glog.Warning("response chat id failed: %v", err)
						}
					}
					glog.Info("[%s] 频道[%s|%s|%d|%s] 发布信息[%s]", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.Chat.UserName, msg.Text)
				} else {
					d, _ := json.MarshalIndent(update.ChannelPost, "", "  ")
					glog.Info("ChannelPost: %s", string(d))
				}
			}
			if update.EditedChannelPost != nil { // EditedChannelPost new incoming channel post of any kind — text, photo,  sticker, etc.
				d, _ := json.MarshalIndent(update.EditedChannelPost, "", "  ")
				glog.Info("EditedChannelPost: %s", string(d))
			}
			if update.InlineQuery != nil { // InlineQuery new incoming inline query
				d, _ := json.MarshalIndent(update.InlineQuery, "", "  ")
				glog.Info("InlineQuery: %s", string(d))
			}
			if update.ChosenInlineResult != nil { // ChosenInlineResult is the result of an inline query that was chosen by a user and sent to their chat partner.
				d, _ := json.MarshalIndent(update.ChosenInlineResult, "", "  ")
				glog.Info("ChosenInlineResult: %s", string(d))
			}
			if update.CallbackQuery != nil { // CallbackQuery new incoming callback query
				d, _ := json.MarshalIndent(update.CallbackQuery, "", "  ")
				glog.Info("CallbackQuery: %s", string(d))
			}
			if update.ShippingQuery != nil { // ShippingQuery new incoming shipping query. Only for invoices with flexible price
				d, _ := json.MarshalIndent(update.ShippingQuery, "", "  ")
				glog.Info("ShippingQuery: %s", string(d))
			}
			if update.PreCheckoutQuery != nil { // PreCheckoutQuery new incoming pre-checkout query. Contains full information about checkout
				d, _ := json.MarshalIndent(update.PreCheckoutQuery, "", "  ")
				glog.Info("PreCheckoutQuery: %s", string(d))
			}
			if update.Poll != nil { // Pool new poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
				d, _ := json.MarshalIndent(update.Poll, "", "  ")
				glog.Info("Poll: %s", string(d))
			}
			if update.PollAnswer != nil { // PollAnswer user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
				d, _ := json.MarshalIndent(update.PollAnswer, "", "  ")
				glog.Info("PollAnswer: %s", string(d))
			}
			if update.MyChatMember != nil { // MyChatMember is the bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
				msg := update.MyChatMember
				glog.Info("[%s] 对话[%s|%s|%d] 用户[%s]变更了用户状态[%s|%s]->[%s|%s]", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.From.FirstName, msg.OldChatMember.User.FirstName, msg.OldChatMember.Status, msg.NewChatMember.User.FirstName, msg.NewChatMember.Status)
			}
			if update.ChatMember != nil { // ChatMember is a chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
				d, _ := json.MarshalIndent(update.ChatMember, "", "  ")
				glog.Info("ChatMember: %s", string(d))
			}
			if update.ChatJoinRequest != nil { //  ChatJoinRequest is a request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
				d, _ := json.MarshalIndent(update.ChatJoinRequest, "", "  ")
				glog.Info("ChatJoinRequest: %s", string(d))
			}
		}
	}
}

func (b *maidBot) wechatReceiver(toUserName, fromUsername, createTime, msgType, content, msgId string, agentId int64) (rToUserName, rFromUsername string, rCreateTime int64, rMsgType, rContent string) {
	content = strings.TrimSpace(content)

	glog.Info("Maid Received From [%s] To [%s] Created [%s] agent [%d] msgid [%s] type [%s] content [%s]", fromUsername, toUserName, createTime, agentId, msgId, msgType, content)

	rToUserName = fromUsername
	rFromUsername = b.WechatReceiver.CorpID
	rCreateTime = time.Now().Unix()
	rMsgType = "text"
	rContent = fmt.Sprintf("Received [%s]", content)
	return
}
