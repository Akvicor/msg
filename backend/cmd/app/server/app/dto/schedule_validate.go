package dto

import (
	"errors"
	"msg/cmd/app/server/common/types/send"
	"strings"
)

func (b *ScheduleCreate) Validate() error {
	b.Category = strings.TrimSpace(b.Category)
	if len(b.Category) == 0 {
		return errors.New("请输入分类")
	}
	if !b.Type.Valid() {
		b.Type = send.TypeText
	}
	b.Title = strings.TrimSpace(b.Title)
	b.Message = strings.TrimSpace(b.Message)
	if !b.PeriodType.Valid() {
		return errors.New("错误渠道类型")
	}
	return nil
}

func (b *ScheduleFind) Validate() error {
	b.Search = strings.TrimSpace(b.Search)
	return nil
}

func (b *ScheduleUpdate) Validate() error {
	if b.ID <= 0 {
		return errors.New("请输入ID")
	}
	b.Category = strings.TrimSpace(b.Category)
	if len(b.Category) == 0 {
		return errors.New("请输入分类")
	}
	if !b.Type.Valid() {
		b.Type = send.TypeText
	}
	b.Title = strings.TrimSpace(b.Title)
	b.Message = strings.TrimSpace(b.Message)
	if !b.PeriodType.Valid() {
		return errors.New("错误渠道类型")
	}
	return nil
}

func (b *ScheduleUpdateNext) Validate() error {
	if b.ID <= 0 {
		return errors.New("请输入ID")
	}
	return nil
}

func (b *ScheduleUpdateSequence) Validate() error {
	if b.ID <= 0 {
		return errors.New("请输入ID")
	}
	return nil
}

func (b *ScheduleDisable) Validate() error {
	if b.ID <= 0 {
		return errors.New("请输入ID")
	}
	return nil
}

func (b *ScheduleDelete) Validate() error {
	if b.ID <= 0 {
		return errors.New("请输入ID")
	}
	return nil
}
