package repository

import (
	"context"
	"github.com/Akvicor/util"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/channel"
	"msg/cmd/app/server/model"
	"time"
)

var Channel = new(channelRepository)

type channelRepository struct {
	base[*model.Channel]
}

/**
查找
*/

// FindAll 获取所有的通知渠道, page为nil时不分页
func (l *channelRepository) FindAll(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderChannel) (channels []*model.Channel, err error) {
	return l.paging(page, l.preload(c, alive, preloader))
}

// FindAllByUID 获取指定用户的记录
func (l *channelRepository) FindAllByUID(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderChannel, uid int64, like string) (channels []*model.Channel, err error) {
	tx := l.preload(c, alive, preloader).Where("uid = ?", uid)

	if len(like) > 0 {
		like = "%" + like + "%"
		tx = tx.Where("name LIKE ? OR target LIKE ?", like, like)
	}

	return l.paging(page, tx)
}

// FindByUID 通过ID获取
func (l *channelRepository) FindByUID(c context.Context, alive bool, preloader *model.PreloaderChannel, uid, id int64) (channel *model.Channel, err error) {
	channel = new(model.Channel)
	err = l.WrapResultErr(l.preload(c, alive, preloader).Where("uid = ? AND id = ?", uid, id).First(channel))
	return
}

// FindByUIDSign 通过Sign获取
func (l *channelRepository) FindByUIDSign(c context.Context, alive bool, preloader *model.PreloaderChannel, uid int64, sign string) (channel *model.Channel, err error) {
	channel = new(model.Channel)
	err = l.WrapResultErr(l.preload(c, alive, preloader).Where("uid = ? AND sign = ?", uid, sign).First(channel))
	return
}

// FindByTypeBotTarget 通过渠道类型,bot和发送对象获取
func (l *channelRepository) FindByTypeBotTarget(c context.Context, alive bool, preloader *model.PreloaderChannel, cType channel.Type, bot, target string) (channel *model.Channel, err error) {
	channel = new(model.Channel)
	err = l.WrapResultErr(l.preload(c, alive, preloader).Where("type = ? AND bot = ? AND target = ?", cType, bot, target).First(channel))
	return
}

/**
创建
*/

// Create 创建通知渠道
func (l *channelRepository) Create(c context.Context, channel *model.Channel) error {
	return l.WrapResultErr(l.db(c).Create(channel))
}

/**
更新
*/

// Update 更新用户通知渠道
func (l *channelRepository) Update(c context.Context, alive bool, chl *model.Channel) error {
	tx := l.alive(c, alive).Select("*")
	omits := []string{"UID", "Disabled", "User"}
	return l.WrapResultErr(tx.Omit(omits...).Where("uid = ? AND id = ?", chl.UID, chl.ID).Updates(chl))
}

/**
删除
*/

// DeleteByUID 删除
func (l *channelRepository) DeleteByUID(c context.Context, uid int64, chl *model.Channel) error {
	tx := l.alive(c, true).Select("*")
	omits := []string{"UID", "User"}
	chl.Sign = util.RandomStringWithTimestamp(32) + "_" + chl.Sign
	chl.Disabled = time.Now().Unix()
	return l.WrapResultErr(tx.Omit(omits...).Where("uid = ? AND id = ?", uid, chl.ID).Updates(chl))
}
