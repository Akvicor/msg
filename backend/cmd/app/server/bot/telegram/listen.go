package telegram

import (
	"context"
	"encoding/json"
	"github.com/Akvicor/glog"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"msg/cmd/app/server/bot/status"
	"strings"
	"sync"
	"time"
)

func (b *Bot) Listen(wg *sync.WaitGroup, ctx context.Context, handle func(context.Context, tgbotapi.UpdatesChannel)) {
	defer func() {
		glog.Warning("Telegram Listen exist")
		b.status.Store(status.SenderStatusDown)
		wg.Done()
	}()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	if handle != nil {
		handle(ctx, updates)
		return
	}

	for {
		select {
		case <-ctx.Done():
			b.debug(nil, "tel stop")
			return
		case update := <-updates:
			if update.Message != nil { // Message new incoming message of any kind — text, photo, sticker, etc.
				b.debug(nil, "Message")
				msg := update.Message
				if len(msg.NewChatMembers) > 0 {
					for _, user := range msg.NewChatMembers {
						b.debug(msg, "[%s] 对话[%s|%s|%d|%d] 用户[%s]邀请成员[%s]进入聊天", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.MessageID, msg.From.FirstName, user.FirstName)
					}
				} else if msg.GroupChatCreated {
					b.debug(msg, "[%s] 用户[%s]创建群组[%s|%s|%d|%d]", time.Unix(int64(msg.Date), 0).String(), msg.From.FirstName, msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.MessageID)
				} else if msg.LeftChatMember != nil {
					b.debug(msg, "[%s] 对话[%s|%s|%d|%d] 用户[%s]删除成员[%s]", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.MessageID, msg.From.FirstName, msg.LeftChatMember.FirstName)
				} else if msg.MigrateToChatID != 0 {
					b.debug(msg, "[%s] 对话[%s|%s|%d|%d] 用户[%s]变更对话ID[%d]->[%d]", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.MessageID, msg.From.FirstName, msg.Chat.ID, msg.MigrateToChatID)
				} else if msg.IsCommand() {
					b.debug(msg, "[%s] 对话[%s|%s|%d|%d] 用户[%s]发送命令[%s][%s]", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.MessageID, msg.From.FirstName, msg.Command(), msg.CommandArguments())
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
					b.debug(msg, "[%s] 对话[%s|%s|%d|%d] 用户[%s]发送%s信息[%s]", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.MessageID, msg.From.FirstName, entities.String(), msg.Text)
				} else {
					d, _ := json.MarshalIndent(msg, "", "  ")
					b.debug(nil, "Message: %s", string(d))
				}
			}
			if update.EditedMessage != nil { // EditedMessage new version of a message that is known to the bot and was edited
				b.debug(nil, "EditedMessage")
				d, _ := json.MarshalIndent(update.EditedMessage, "", "  ")
				b.debug(nil, "EditedMessage: %s", string(d))
			}
			if update.ChannelPost != nil { // ChannelPost new version of a message that is known to the bot and was edited
				b.debug(nil, "ChannelPost")
				msg := update.ChannelPost
				if len(msg.Text) != 0 {
					b.debug(msg, "[%s] 频道[%s|%s|%d|%s] 发布信息[%s]", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.Chat.UserName, msg.Text)
				} else {
					d, _ := json.MarshalIndent(update.ChannelPost, "", "  ")
					b.debug(nil, "ChannelPost: %s", string(d))
				}
			}
			if update.EditedChannelPost != nil { // EditedChannelPost new incoming channel post of any kind — text, photo,  sticker, etc.
				b.debug(nil, "EditedChannelPost")
				d, _ := json.MarshalIndent(update.EditedChannelPost, "", "  ")
				b.debug(nil, "EditedChannelPost: %s", string(d))
			}
			if update.InlineQuery != nil { // InlineQuery new incoming inline query
				b.debug(nil, "InlineQuery")
				d, _ := json.MarshalIndent(update.InlineQuery, "", "  ")
				b.debug(nil, "InlineQuery: %s", string(d))
			}
			if update.ChosenInlineResult != nil { // ChosenInlineResult is the result of an inline query that was chosen by a user and sent to their chat partner.
				b.debug(nil, "ChosenInlineResult")
				d, _ := json.MarshalIndent(update.ChosenInlineResult, "", "  ")
				b.debug(nil, "ChosenInlineResult: %s", string(d))
			}
			if update.CallbackQuery != nil { // CallbackQuery new incoming callback query
				b.debug(nil, "CallbackQuery")
				d, _ := json.MarshalIndent(update.CallbackQuery, "", "  ")
				b.debug(nil, "CallbackQuery: %s", string(d))
			}
			if update.ShippingQuery != nil { // ShippingQuery new incoming shipping query. Only for invoices with flexible price
				b.debug(nil, "ShippingQuery")
				d, _ := json.MarshalIndent(update.ShippingQuery, "", "  ")
				b.debug(nil, "ShippingQuery: %s", string(d))
			}
			if update.PreCheckoutQuery != nil { // PreCheckoutQuery new incoming pre-checkout query. Contains full information about checkout
				b.debug(nil, "PreCheckoutQuery")
				d, _ := json.MarshalIndent(update.PreCheckoutQuery, "", "  ")
				b.debug(nil, "PreCheckoutQuery: %s", string(d))
			}
			if update.Poll != nil { // Pool new poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
				b.debug(nil, "Poll")
				d, _ := json.MarshalIndent(update.Poll, "", "  ")
				b.debug(nil, "Poll: %s", string(d))
			}
			if update.PollAnswer != nil { // PollAnswer user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
				b.debug(nil, "PollAnswer")
				d, _ := json.MarshalIndent(update.PollAnswer, "", "  ")
				b.debug(nil, "PollAnswer: %s", string(d))
			}
			if update.MyChatMember != nil { // MyChatMember is the bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
				b.debug(nil, "MyChatMember")
				msg := update.MyChatMember
				b.debug(msg, "[%s] 对话[%s|%s|%d] 用户[%s]变更了用户状态[%s|%s]->[%s|%s]", time.Unix(int64(msg.Date), 0).String(), msg.Chat.Type, msg.Chat.Title, msg.Chat.ID, msg.From.FirstName, msg.OldChatMember.User.FirstName, msg.OldChatMember.Status, msg.NewChatMember.User.FirstName, msg.NewChatMember.Status)
			}
			if update.ChatMember != nil { // ChatMember is a chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
				b.debug(nil, "ChatMember")
				d, _ := json.MarshalIndent(update.ChatMember, "", "  ")
				b.debug(nil, "ChatMember: %s", string(d))
			}
			if update.ChatJoinRequest != nil { //  ChatJoinRequest is a request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
				b.debug(nil, "ChatJoinRequest")
				d, _ := json.MarshalIndent(update.ChatJoinRequest, "", "  ")
				b.debug(nil, "ChatJoinRequest: %s", string(d))
			}
		}
	}
}
