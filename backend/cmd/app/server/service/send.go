package service

import (
	"context"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/send"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/repository"
	"time"
)

var Send = new(sendService)

type sendService struct {
	base
}

// FindAll 获取所有的通知渠道, page为nil时不分页
func (l *sendService) FindAll(page *resp.PageModel, alive bool, preload *model.PreloaderSend) (sends []*model.Send, err error) {
	return repository.Send.FindAll(context.Background(), page, alive, preload)
}

// FindAllByUID 获取所有的通知渠道, page为nil时不分页
func (l *sendService) FindAllByUID(page *resp.PageModel, alive bool, preload *model.PreloaderSend, uid int64, sent []int64, search string, channelIds []int64, types []send.Type) (sends []*model.Send, err error) {
	return repository.Send.FindAllByUID(context.Background(), page, alive, preload, uid, sent, search, channelIds, types)
}

// FindByUID 获取通知渠道
func (l *sendService) FindByUID(alive bool, preload *model.PreloaderSend, uid, id int64) (send *model.Send, err error) {
	return repository.Send.FindByUID(context.Background(), alive, preload, uid, id)
}

// Create 保存通知渠道
func (l *sendService) Create(readyAt, sendAt int64, uid, channelID int64, ip string, cType send.Type, title, msg string) (send *model.Send, err error) {
	send = model.NewSend(readyAt, sendAt, 0, uid, channelID, ip, cType, title, msg)
	err = repository.Send.Create(context.Background(), send)
	if err != nil {
		return nil, err
	}
	return send, err
}

// UpdateSentByID 通过ID更新送达时间
func (l *sendService) UpdateSentByID(id, sentAt int64, errMsg string) (err error) {
	return repository.Send.UpdateSentByID(context.Background(), id, sentAt, errMsg)
}

// UpdateSentByIDFinished 成功送达，通过ID更新送达时间为现在
func (l *sendService) UpdateSentByIDFinished(id int64) (err error) {
	return repository.Send.UpdateSentByID(context.Background(), id, time.Now().Unix(), "")
}

// UpdateSentByIDCancel 主动取消，通过ID更新送达时间为-现在
func (l *sendService) UpdateSentByIDCancel(id int64) (err error) {
	return repository.Send.UpdateSentByID(context.Background(), id, -time.Now().Unix(), string(send.StatusCancel))
}

// UpdateSentByIDFailed 发送错误，通过ID更新送达时间为-现在
func (l *sendService) UpdateSentByIDFailed(id int64, errMsg string) (err error) {
	return repository.Send.UpdateSentByID(context.Background(), id, -time.Now().Unix(), errMsg)
}
