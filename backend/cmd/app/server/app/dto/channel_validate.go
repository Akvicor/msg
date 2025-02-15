package dto

import (
	"errors"
	"msg/cmd/app/server/common/types/send"
	"strings"
	"time"
)

func (b *ChannelCreate) Validate() error {
	b.Sign = strings.TrimSpace(b.Sign)
	if len(b.Sign) == 0 {
		return errors.New("请输入唯一标记")
	}
	b.Name = strings.TrimSpace(b.Name)
	if len(b.Name) == 0 {
		b.Name = b.Sign
	}
	if !b.Type.Valid() {
		return errors.New("错误渠道类型")
	}
	b.Target = strings.TrimSpace(b.Target)
	if len(b.Target) == 0 {
		return errors.New("请输入渠道目标")
	}
	return nil
}

func (b *ChannelFind) Validate() error {
	b.Search = strings.TrimSpace(b.Search)
	return nil
}

func (b *ChannelUpdate) Validate() error {
	if b.ID <= 0 {
		return errors.New("请输入ID")
	}
	b.Sign = strings.TrimSpace(b.Sign)
	if len(b.Sign) == 0 {
		return errors.New("请输入唯一标记")
	}
	b.Name = strings.TrimSpace(b.Name)
	if len(b.Name) == 0 {
		b.Name = b.Sign
	}
	if !b.Type.Valid() {
		return errors.New("错误渠道类型")
	}
	b.Target = strings.TrimSpace(b.Target)
	if len(b.Target) == 0 {
		return errors.New("请输入渠道目标")
	}
	return nil
}

func (b *ChannelDelete) Validate() error {
	b.Sign = strings.TrimSpace(b.Sign)
	if len(b.Sign) == 0 && b.ID == 0 {
		return errors.New("请输入渠道ID或唯一标记")
	}
	return nil
}

func (b *ChannelTest) Validate() error {
	b.Sign = strings.TrimSpace(b.Sign)
	if len(b.Sign) == 0 && b.ID == 0 {
		return errors.New("请输入渠道ID或唯一标记")
	}
	return nil
}

func (b *ChannelSend) Validate() error {
	b.Sign = strings.TrimSpace(b.Sign)
	if len(b.Sign) == 0 && b.ID == 0 {
		return errors.New("请输入渠道ID或唯一标记")
	}
	if !b.Type.Valid() {
		b.Type = send.TypeText
	}
	if b.At <= 0 {
		b.At = time.Now().Unix()
	}
	b.Title = strings.TrimSpace(b.Title)
	b.Msg = strings.TrimSpace(b.Msg)
	if b.Title == "" && b.Msg == "" {
		return errors.New("请输入标题或内容")
	}
	return nil
}
