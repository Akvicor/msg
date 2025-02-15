package mail

import (
	"github.com/emersion/go-imap"
	gmail "github.com/emersion/go-message/mail"
	"io"
	"strings"
	"time"
)

// Email

type ImapEmail struct {
	Envelope *ImapEmailEnvelope `json:"envelope"` // 信封
	Content  *ImapEmailContent  `json:"content"`  // 内容
}

type ImapEmailContent struct {
	Type        string            `json:"type"`        // 内容类型
	Params      map[string]string `json:"params"`      // 内容参数
	Contents    []string          `json:"contents"`    // 内容
	Attachments []string          `json:"attachments"` // 附件名
}

// Envelope

type ImapEmailEnvelope struct {
	Date      time.Time           `json:"date"`        // 邮件日期
	Subject   string              `json:"subject"`     // 邮件标题
	MessageID string              `json:"message_id"`  // 邮件ID
	InReplyTo string              `json:"in_reply_to"` // 回复于邮件ID
	From      []*ImapEmailAddress `json:"from"`        // 编写邮件的人地址
	Sender    []*ImapEmailAddress `json:"sender"`      // 投递邮件的人的地址
	ReplyTo   []*ImapEmailAddress `json:"reply_to"`    // 回复的收件人
	To        []*ImapEmailAddress `json:"to"`          // 发送给
	Cc        []*ImapEmailAddress `json:"cc"`          // 收到邮件副本的收件人地址
	Bcc       []*ImapEmailAddress `json:"bcc"`         // 密送的收件人地址，收件人不知道其他收信人
}

// Address

type ImapEmailAddress struct {
	PersonalName string `json:"personal_name"`  // 个人
	AtDomainList string `json:"at_domain_list"` // The SMTP at-domain-list (source route)
	MailboxName  string `json:"mailbox_name"`   // 邮局名字
	HostName     string `json:"host_name"`      // 邮局域名
}

func newImapEmailAddress(personalName, atDomainList, mailboxName, hostname string) *ImapEmailAddress {
	return &ImapEmailAddress{
		PersonalName: personalName,
		AtDomainList: atDomainList,
		MailboxName:  mailboxName,
		HostName:     hostname,
	}
}

func parseEmailEnvelope(envelope *imap.Envelope) *ImapEmailEnvelope {
	i := &ImapEmailEnvelope{}
	i.Date = envelope.Date
	i.Subject = envelope.Subject
	i.MessageID = envelope.MessageId
	i.InReplyTo = envelope.InReplyTo
	i.From = make([]*ImapEmailAddress, 0, len(envelope.From))
	for _, from := range envelope.From {
		i.From = append(i.From, newImapEmailAddress(from.PersonalName, from.AtDomainList, from.MailboxName, from.HostName))
	}
	i.Sender = make([]*ImapEmailAddress, 0, len(envelope.Sender))
	for _, sender := range envelope.Sender {
		i.Sender = append(i.Sender, newImapEmailAddress(sender.PersonalName, sender.AtDomainList, sender.MailboxName, sender.HostName))
	}
	i.ReplyTo = make([]*ImapEmailAddress, 0, len(envelope.ReplyTo))
	for _, replyTo := range envelope.ReplyTo {
		i.ReplyTo = append(i.ReplyTo, newImapEmailAddress(replyTo.PersonalName, replyTo.AtDomainList, replyTo.MailboxName, replyTo.HostName))
	}
	i.To = make([]*ImapEmailAddress, 0, len(envelope.To))
	for _, to := range envelope.To {
		i.From = append(i.From, newImapEmailAddress(to.PersonalName, to.AtDomainList, to.MailboxName, to.HostName))
	}
	i.Cc = make([]*ImapEmailAddress, 0, len(envelope.Cc))
	for _, cc := range envelope.Cc {
		i.Cc = append(i.Cc, newImapEmailAddress(cc.PersonalName, cc.AtDomainList, cc.MailboxName, cc.HostName))
	}
	i.Bcc = make([]*ImapEmailAddress, 0, len(envelope.Bcc))
	for _, bcc := range envelope.Bcc {
		i.Bcc = append(i.Bcc, newImapEmailAddress(bcc.PersonalName, bcc.AtDomainList, bcc.MailboxName, bcc.HostName))
	}
	return i
}

func parseEmail(mr *gmail.Reader) *ImapEmail {
	email := &ImapEmail{
		Envelope: &ImapEmailEnvelope{},
		Content: &ImapEmailContent{
			Type:        "",
			Params:      nil,
			Contents:    make([]string, 0),
			Attachments: make([]string, 0),
		},
	}
	header := mr.Header
	if date, err := header.Date(); err == nil {
		email.Envelope.Date = date
	}
	if subject, err := header.Subject(); err == nil {
		email.Envelope.Subject = subject
	}
	if messageId, err := header.MessageID(); err == nil {
		email.Envelope.MessageID = messageId
	}
	if contentType, params, err := header.ContentType(); err == nil {
		email.Content.Type = contentType
		email.Content.Params = params
	}
	email.Envelope.InReplyTo = header.Get("In-Reply-To")
	if addrs, err := header.AddressList("From"); err == nil {
		email.Envelope.From = make([]*ImapEmailAddress, 0, len(addrs))
		for _, addr := range addrs {
			name := strings.Split(addr.Address, "@")
			if len(name) >= 2 {
				email.Envelope.From = append(email.Envelope.From, newImapEmailAddress(addr.Name, "", name[0], name[1]))
			}
		}
	}
	if addrs, err := header.AddressList("Sender"); err == nil {
		email.Envelope.Sender = make([]*ImapEmailAddress, 0, len(addrs))
		for _, addr := range addrs {
			name := strings.Split(addr.Address, "@")
			if len(name) >= 2 {
				email.Envelope.Sender = append(email.Envelope.Sender, newImapEmailAddress(addr.Name, "", name[0], name[1]))
			}
		}
	}
	if addrs, err := header.AddressList("Reply-To"); err == nil {
		email.Envelope.ReplyTo = make([]*ImapEmailAddress, 0, len(addrs))
		for _, addr := range addrs {
			name := strings.Split(addr.Address, "@")
			if len(name) >= 2 {
				email.Envelope.ReplyTo = append(email.Envelope.ReplyTo, newImapEmailAddress(addr.Name, "", name[0], name[1]))
			}
		}
	}
	if addrs, err := header.AddressList("To"); err == nil {
		email.Envelope.To = make([]*ImapEmailAddress, 0, len(addrs))
		for _, addr := range addrs {
			name := strings.Split(addr.Address, "@")
			if len(name) >= 2 {
				email.Envelope.To = append(email.Envelope.To, newImapEmailAddress(addr.Name, "", name[0], name[1]))
			}
		}
	}
	if addrs, err := header.AddressList("Cc"); err == nil {
		email.Envelope.Cc = make([]*ImapEmailAddress, 0, len(addrs))
		for _, addr := range addrs {
			name := strings.Split(addr.Address, "@")
			if len(name) >= 2 {
				email.Envelope.Cc = append(email.Envelope.Cc, newImapEmailAddress(addr.Name, "", name[0], name[1]))
			}
		}
	}
	if addrs, err := header.AddressList("Bcc"); err == nil {
		email.Envelope.Bcc = make([]*ImapEmailAddress, 0, len(addrs))
		for _, addr := range addrs {
			name := strings.Split(addr.Address, "@")
			if len(name) >= 2 {
				email.Envelope.Bcc = append(email.Envelope.Bcc, newImapEmailAddress(addr.Name, "", name[0], name[1]))
			}
		}
	}

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			continue
		}

		switch h := p.Header.(type) {
		case *gmail.InlineHeader:
			b, _ := io.ReadAll(p.Body)
			email.Content.Contents = append(email.Content.Contents, string(b))
		case *gmail.AttachmentHeader:
			filename, _ := h.Filename()
			email.Content.Attachments = append(email.Content.Attachments, filename)
		}
	}
	return email
}
