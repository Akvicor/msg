package telegram

import (
	"encoding/json"
	"fmt"
	"github.com/Akvicor/glog"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"msg/cmd/app/server/bot/status"
	"os"
	"sync"
)

type Bot struct {
	name   string
	api    string
	token  string
	bot    *tgbotapi.BotAPI
	status *status.AtomicSenderStatus

	_debug              bool
	_debugFiltered      *os.File
	_debugFilteredIndex int
	_debugLock          sync.Mutex
}

func NewBot(name, api, token string, debug bool) (*Bot, error) {
	bot := &Bot{
		name:                name,
		api:                 api,
		token:               token,
		bot:                 nil,
		status:              status.NewSenderStatus(status.SenderStatusConnecting),
		_debug:              debug,
		_debugFiltered:      nil,
		_debugFilteredIndex: 0,
		_debugLock:          sync.Mutex{},
	}
	var err error
	bot.bot, err = tgbotapi.NewBotAPIWithAPIEndpoint(token, api)
	if err != nil {
		bot.status.Store(status.SenderStatusDown)
		return nil, fmt.Errorf("new bot api error: %v", err)
	}
	bot.bot.Debug = debug
	if debug {
		glog.SetMask(glog.GetMask() | glog.MaskDEBUG)
		f, err := os.Create(fmt.Sprintf("_debug_%s.txt", name))
		if err != nil {
			return nil, fmt.Errorf("create filtered file error: %v", err)
		}
		bot._debugFiltered = f
		bot.debug(bot.bot.Self, "连接[%s|%s|%d]", bot.bot.Self.FirstName, bot.bot.Self.UserName, bot.bot.Self.ID)
	}
	bot.status.Store(status.SenderStatusOK)
	return bot, nil
}

func (b *Bot) Status() status.SenderStatus {
	return b.status.Load()
}

func (b *Bot) debug(data any, msg string, values ...any) {
	if !b._debug {
		return
	}
	b._debugLock.Lock()
	defer b._debugLock.Unlock()

	b._debugFilteredIndex++

	arg := fmt.Sprintf("[%s][%d] %s", b.name, b._debugFilteredIndex, msg)
	glog.Debug(arg, values...)
	if data != nil {
		ds, _ := json.MarshalIndent(data, "", "  ")
		_, _ = b._debugFiltered.WriteString(fmt.Sprintf("[%d] %s\n", b._debugFilteredIndex, ds))
	}
}
