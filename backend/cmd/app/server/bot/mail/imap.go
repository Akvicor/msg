package mail

import (
	"context"
	"errors"
	"fmt"
	"github.com/Akvicor/glog"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap-idle"
	"github.com/emersion/go-imap/client"
	gmail "github.com/emersion/go-message/mail"
	"msg/cmd/app/server/bot/status"
	"sync"
	"sync/atomic"
	"time"
)

type IMAPModel struct {
	ImapHost string
	ImapPort int
	From     string
	Username string
	Password string

	Messages chan *ImapEmail

	inboxMessagesHandled *atomic.Uint32     // 当前已处理到的邮件ID
	updates              chan client.Update // 新邮件
	done                 chan error         // IMAP IDLE结束
	kill                 chan struct{}      // 主动关闭IDLE
	status               *status.AtomicSenderStatus
}

func NewIMAPModel(imapHost string, imapPort int, from, username, password string) *IMAPModel {
	return &IMAPModel{
		ImapHost:             imapHost,
		ImapPort:             imapPort,
		From:                 from,
		Username:             username,
		Password:             password,
		inboxMessagesHandled: &atomic.Uint32{},
		Messages:             make(chan *ImapEmail, 10),
		updates:              make(chan client.Update, 10),
		done:                 make(chan error, 1),
		kill:                 make(chan struct{}, 1),
		status:               status.NewSenderStatus(status.SenderStatusNotSupported),
	}
}

func (m *IMAPModel) Status() status.SenderStatus {
	return m.status.Load()
}

// Connect 连接IMAP
// done: IDLE结束
// kill: 主动关闭IDLE
func (m *IMAPModel) Connect(box string, updates chan client.Update, done chan error, kill chan struct{}) (*client.Client, uint32, error) {
	c, err := client.DialTLS(fmt.Sprintf("%s:%d", m.ImapHost, m.ImapPort), nil)
	if err != nil {
		return nil, 0, fmt.Errorf("连接IMAP失败: %v", err)
	}

	err = c.Login(m.Username, m.Password)
	if err != nil {
		return nil, 0, fmt.Errorf("登录IMAP失败: %v", err)
	}

	mbox, err := c.Select(box, false)
	if err != nil {
		_ = c.Logout()
		return nil, 0, fmt.Errorf("获取INBOX收件箱失败: %v", err)
	}

	idleClient := idle.NewClient(c)

	c.Updates = updates

	if ok, err := idleClient.SupportIdle(); err != nil || !ok {
		_ = c.Logout()
		return nil, 0, fmt.Errorf("不支持IDLE: %v %v", ok, err)
	}

	go func() {
		done <- idleClient.Idle(kill)
	}()

	return c, mbox.Messages, nil
}

// Listen
// stop: 主动关闭
func (m *IMAPModel) Listen(wg *sync.WaitGroup, ctx context.Context) {
	defer func() {
		glog.Warning("IMAP Listen exist")
		m.status.Store(status.SenderStatusDown)
		wg.Done()
	}()
	m.status.Store(status.SenderStatusConnecting)

	var err error

	closeChan := func() {
		close(m.updates)
		close(m.done)
		close(m.kill)
		close(m.Messages)
	}

	var c *client.Client

	m.inboxMessagesHandled.Store(0)
	var errInit = errors.New("bot init")
	m.done <- errInit
	for {
		select {
		case <-ctx.Done(): // 主动关闭
			m.kill <- struct{}{}
			err = <-m.done
			if c != nil {
				_ = c.Logout()
			}
			closeChan()
			return
		case err = <-m.done:
			isInit := errors.Is(err, errInit)
			if !isInit && c != nil {
				_ = c.Logout()
			}
			var baseSeq uint32
			c, baseSeq, err = m.Connect("INBOX", m.updates, m.done, m.kill)
			if err != nil {
				glog.Warning("Bot Maid Mail Receiver Connect Failed: %v", err)
				time.Sleep(3 * time.Second)
				m.done <- errInit
			}
			if isInit && err == nil {
				glog.Warning("Bot Maid Mail Receiver Connect Successful: init[%v]", isInit)
				m.inboxMessagesHandled.Store(baseSeq)
				m.status.Store(status.SenderStatusOK)
			}
			break
		case update := <-m.updates:
			switch u := update.(type) {
			case *client.MailboxUpdate:
				last := m.inboxMessagesHandled.Load()
				next := u.Mailbox.Messages
				if last != next {
					swapped := m.inboxMessagesHandled.CompareAndSwap(last, next)
					if swapped {
						msg, err := m.FetchEmail(last+1, next)
						if err != nil {
							glog.Warning("获取邮件失败： %v", err)
						} else {
							for _, ms := range msg {
								m.Messages <- ms
							}
						}
					}
				}
			}
		}
	}
}

// FetchEmailEnvelope 获取邮件信封
func (m *IMAPModel) FetchEmailEnvelope(from, to uint32) ([]*ImapEmail, error) {
	c, err := client.DialTLS(fmt.Sprintf("%s:%d", m.ImapHost, m.ImapPort), nil)
	if err != nil {
		return nil, fmt.Errorf("连接错误: %v", err)
	}
	defer func() {
		_ = c.Logout()
	}()

	if err := c.Login(m.Username, m.Password); err != nil {
		return nil, fmt.Errorf("登录错误: %v", err)
	}

	_, err = c.Select("INBOX", false)
	if err != nil {
		return nil, fmt.Errorf("选择邮箱错误: %v", err)
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddRange(from, to)

	items := []imap.FetchItem{imap.FetchEnvelope}

	messages := make(chan *imap.Message, 10)
	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			glog.Warning("获取邮件失败: %v", err)
		}
	}()

	mails := make([]*ImapEmail, 0)
	for msg := range messages {
		if msg.Envelope != nil {
			mail := &ImapEmail{}
			mail.Envelope = parseEmailEnvelope(msg.Envelope)
			mails = append(mails, mail)
		} else {
			glog.Info("No envelope returned for message")
		}
	}
	return mails, nil
}

func (m *IMAPModel) FetchEmail(from, to uint32) ([]*ImapEmail, error) {
	c, err := client.DialTLS(fmt.Sprintf("%s:%d", m.ImapHost, m.ImapPort), nil)
	if err != nil {
		return nil, fmt.Errorf("连接错误: %v", err)
	}
	defer func() {
		_ = c.Logout()
	}()

	if err := c.Login(m.Username, m.Password); err != nil {
		return nil, fmt.Errorf("登录错误: %v", err)
	}

	_, err = c.Select("INBOX", false)
	if err != nil {
		return nil, fmt.Errorf("选择邮箱错误: %v", err)
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddRange(from, to)

	var section imap.BodySectionName
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 10)
	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			glog.Warning("获取邮件失败: %v", err)
		}
	}()

	mails := make([]*ImapEmail, 0)
	for msg := range messages {
		resp := msg.GetBody(&section)
		if resp == nil {
			continue
		}
		mr, err := gmail.CreateReader(resp)
		if err != nil {
			continue
		}
		mails = append(mails, parseEmail(mr))
	}
	return mails, nil
}
