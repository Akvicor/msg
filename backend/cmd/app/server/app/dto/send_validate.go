package dto

import (
	"errors"
	"msg/cmd/app/server/common/types/send"
	"strings"
	"time"
)

func (b *SendFind) Validate() error {
	b.Search = strings.TrimSpace(b.Search)
	return nil
}

func (b *Send) Validate() error {
	b.Sign = strings.TrimSpace(b.Sign)
	if b.ID == 0 && len(b.Sign) == 0 {
		return errors.New("请输入ID或Sign")
	}
	if !b.Type.Valid() {
		b.Type = send.TypeText
	}
	if b.At <= 0 {
		b.At = time.Now().Unix()
	}
	b.Title = strings.TrimSpace(b.Title)
	b.Msg = strings.TrimSpace(b.Msg)
	if len(b.Title) == 0 && len(b.Msg) == 0 {
		return errors.New("请输入通知内容")
	}
	return nil
}

func (b *SendCancel) Validate() error {
	return nil
}

func (b *SendStatus) Validate() error {
	return nil
}
