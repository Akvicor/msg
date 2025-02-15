package repository

import (
	"context"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/send"
	"msg/cmd/app/server/model"
	"time"
)

var Send = new(sendRepository)

type sendRepository struct {
	base[*model.Send]
}

/**
查找
*/

// FindAll 获取所有的通知渠道, page为nil时不分页
func (l *sendRepository) FindAll(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderSend) (sends []*model.Send, err error) {
	return l.paging(page, l.preload(c, alive, preloader))
}

// FindAllByUID 获取指定用户的记录
func (l *sendRepository) FindAllByUID(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderSend, uid int64, sent []int64, search string, channelIds []int64, types []send.Type) (sends []*model.Send, err error) {
	tx := l.preload(c, alive, preloader).Where("uid = ?", uid)
	if len(sent) > 0 {
		cond := ""
		args := make([]any, 0, len(sent))
		for _, si := range sent {
			if cond != "" {
				cond += " OR "
			}
			if si > 0 {
				cond += "sent_at > ?"
				args = append(args, si)
			} else if si < 0 {
				cond += "sent_at < ?"
				args = append(args, si)
			} else {
				cond += "sent_at = ?"
				args = append(args, si)
			}
		}
		tx = tx.Where(cond, args...)
	}
	if len(types) > 0 {
		tx = tx.Where("type IN (?)", types)
	}
	if len(search) > 0 {
		search = "%" + search + "%"
		tx = tx.Where("ip LIKE ? OR title LIKE ? OR msg LIKE ?", search, search, search)
	}
	if len(channelIds) > 0 {
		tx = tx.Where("channel_id IN (?)", channelIds)
	}

	return l.paging(page, tx.Order("CASE WHEN sent_at = 0 THEN 0 ELSE 1 END ASC, CASE WHEN sent_at = 0 THEN send_at ELSE ABS(sent_at)*-1 END ASC"))
}

// FindByUID 通过ID获取
func (l *sendRepository) FindByUID(c context.Context, alive bool, preloader *model.PreloaderSend, uid, id int64) (send *model.Send, err error) {
	send = new(model.Send)
	err = l.WrapResultErr(l.preload(c, alive, preloader).Where("uid = ? AND id = ?", uid, id).First(send))
	return
}

// FindAllPending 获取所有等待发送的消息
func (l *sendRepository) FindAllPending(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderSend) (sends []*model.Send, err error) {
	return l.paging(page, l.preload(c, alive, preloader).Where("sent_at = ? AND send_at > ?", 0, time.Now().AddDate(0, 0, -1).Unix()))
}

/**
创建
*/

// Create 创建通知渠道
func (l *sendRepository) Create(c context.Context, send *model.Send) error {
	return l.WrapResultErr(l.db(c).Create(send))
}

/**
更新
*/

// UpdateSentByID 通过ID更新送达时间
func (l *sendRepository) UpdateSentByID(c context.Context, id, sentAt int64, err string) error {
	return l.WrapResultErr(l.dbm(c).Where("id = ?", id).UpdateColumns(map[string]any{
		"sent_at": sentAt,
		"err_msg": err,
	}))
}
