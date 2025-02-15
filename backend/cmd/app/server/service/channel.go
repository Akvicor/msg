package service

import (
	"context"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/channel"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/repository"
)

var Channel = new(channelService)

type channelService struct {
	base
}

// FindAll 获取所有的通知渠道, page为nil时不分页
func (l *channelService) FindAll(page *resp.PageModel, alive bool, preload *model.PreloaderChannel) (channels []*model.Channel, err error) {
	return repository.Channel.FindAll(context.Background(), page, alive, preload)
}

// FindAllByUID 获取所有的通知渠道, page为nil时不分页
func (l *channelService) FindAllByUID(page *resp.PageModel, alive bool, preload *model.PreloaderChannel, uid int64, like string) (channels []*model.Channel, err error) {
	return repository.Channel.FindAllByUID(context.Background(), page, alive, preload, uid, like)
}

// FindByUID 获取通知渠道
func (l *channelService) FindByUID(alive bool, preload *model.PreloaderChannel, uid, id int64) (chl *model.Channel, err error) {
	return repository.Channel.FindByUID(context.Background(), alive, preload, uid, id)
}

// FindByUIDSign 获取通知渠道
func (l *channelService) FindByUIDSign(alive bool, preload *model.PreloaderChannel, uid int64, sign string) (chl *model.Channel, err error) {
	return repository.Channel.FindByUIDSign(context.Background(), alive, preload, uid, sign)
}

// FindByTypeTarget 获取通知渠道
func (l *channelService) FindByTypeTarget(alive bool, preload *model.PreloaderChannel, cType channel.Type, bot, target string) (chl *model.Channel, err error) {
	return repository.Channel.FindByTypeBotTarget(context.Background(), alive, preload, cType, bot, target)
}

// Create 保存通知渠道
func (l *channelService) Create(uid int64, sign, name string, channelType channel.Type, bot, target string) (chl *model.Channel, err error) {
	chl = model.NewChannel(uid, sign, name, channelType, bot, target)
	err = repository.Channel.Create(context.Background(), chl)
	if err != nil {
		return nil, err
	}
	return chl, err
}

// FindAllAlive 获取活跃的通知渠道
func (l *channelService) FindAllAlive(preload *model.PreloaderChannel) (channels []*model.Channel, err error) {
	return repository.Channel.FindAll(context.Background(), nil, true, preload)
}

// FindAllAliveByUID 通过UID获取活跃的通知渠道
func (l *channelService) FindAllAliveByUID(preload *model.PreloaderChannel, uid int64) (channels []*model.Channel, err error) {
	return repository.Channel.FindAllByUID(context.Background(), nil, true, preload, uid, "")
}

// Update 更新通知渠道
func (l *channelService) Update(uid, id int64, sign, name string, channelType channel.Type, bot, target string) (chl *model.Channel, err error) {
	chl = model.NewChannel(uid, sign, name, channelType, bot, target)
	chl.ID = id
	err = repository.Channel.Update(context.Background(), false, chl)
	if err != nil {
		return nil, err
	}
	return chl, err
}

// DeleteByUID 通过UID删除通知渠道
func (l *channelService) DeleteByUID(uid, id int64) (err error) {
	return l.transaction(context.Background(), func(ctx context.Context) error {
		chl, err := repository.Channel.FindByUID(ctx, false, nil, uid, id)
		if err != nil {
			return err
		}
		err = repository.Channel.DeleteByUID(ctx, uid, chl)
		if err != nil {
			return err
		}
		return nil
	})
}

// DeleteBySign 通过Sign删除通知渠道
func (l *channelService) DeleteBySign(uid int64, sign string) (err error) {
	return l.transaction(context.Background(), func(ctx context.Context) error {
		chl, err := repository.Channel.FindByUIDSign(ctx, false, nil, uid, sign)
		if err != nil {
			return err
		}
		err = repository.Channel.DeleteByUID(ctx, uid, chl)
		if err != nil {
			return err
		}
		return nil
	})
}
