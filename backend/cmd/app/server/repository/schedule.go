package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/model"
	"time"
)

var Schedule = new(scheduleRepository)

type scheduleRepository struct {
	base[*model.Schedule]
}

/**
查找
*/

// FindAll 获取所有的记录, page为nil时不分页
func (u *scheduleRepository) FindAll(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderSchedule) (schedules []*model.Schedule, err error) {
	return u.paging(page, u.preload(c, alive, preloader).Order("disabled ASC").Order("sequence ASC"))
}

// FindAllRecentBySecond 获取所有的记录
func (u *scheduleRepository) FindAllRecentBySecond(c context.Context, preloader *model.PreloaderSchedule, seconds int64) (schedules []*model.Schedule, err error) {
	nowT := time.Now().Unix()
	return u.paging(nil, u.preload(c, true, preloader).Where("expiration_date > ? AND expiration_times != ? AND next_of_period <= ?", nowT, 0, nowT+seconds))
}

// FindAllByUID 获取所有的记录, page为nil时不分页
func (u *scheduleRepository) FindAllByUID(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderSchedule, uid int64) (schedules []*model.Schedule, err error) {
	return u.paging(page, u.preload(c, alive, preloader).Where("uid = ?", uid).Order("disabled ASC").Order("sequence ASC"))
}

// FindAllByChannel 获取所有的记录, page为nil时不分页
func (u *scheduleRepository) FindAllByChannel(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderSchedule, uid, channelId int64) (schedules []*model.Schedule, err error) {
	return u.paging(page, u.preload(c, alive, preloader).Where("uid = ?", uid).Where("channel_id = ?", channelId).Order("disabled ASC").Order("sequence ASC"))
}

// FindAllByUIDLike 获取所有的记录, page为nil时不分页
func (u *scheduleRepository) FindAllByUIDLike(c context.Context, page *resp.PageModel, alive bool, preloader *model.PreloaderSchedule, uid int64, like string) (schedules []*model.Schedule, err error) {
	tx := u.preload(c, alive, preloader).Where("uid = ?", uid)
	if len(like) > 0 {
		like = "%" + like + "%"
		tx = tx.Where("category LIKE ? OR title LIKE ? OR message LIKE ?", like, like)
	}
	return u.paging(page, tx.Order("disabled ASC").Order("sequence ASC"))
}

// FindByID 通过Token获取记录
func (u *scheduleRepository) FindByID(c context.Context, alive bool, preloader *model.PreloaderSchedule, uid, id int64) (schedule *model.Schedule, err error) {
	schedule = new(model.Schedule)
	err = u.preload(c, alive, preloader).Where("uid = ?", uid).Where("id = ?", id).First(schedule).Error
	return schedule, err
}

// GetMaxSequenceByUID 获取最大序号
func (u *scheduleRepository) GetMaxSequenceByUID(c context.Context, alive bool, uid int64) (seq int64, err error) {
	schedule := new(model.Schedule)
	err = u.WrapResultErr(u.alive(c, alive).Where("uid = ?", uid).Order("sequence DESC").First(schedule))
	if err != nil {
		return 0, err
	}
	return schedule.Sequence, nil
}

/**
创建
*/

// Create 创建记录
func (u *scheduleRepository) Create(c context.Context, schedule *model.Schedule) error {
	return u.WrapResultErr(u.db(c).Create(schedule))
}

/**
更新
*/

// Update 更新
func (u *scheduleRepository) Update(c context.Context, alive bool, schedule *model.Schedule) error {
	tx := u.alive(c, alive).Select("*")
	omits := []string{"ID", "UID", "Created", "Sequence", "Disabled", "User", "Channel"}
	return u.WrapResultErr(tx.Omit(omits...).Where("uid = ? AND id = ?", schedule.UID, schedule.ID).Updates(schedule))
}

// UpdateNext 更新名称
func (u *scheduleRepository) UpdateNext(c context.Context, alive bool, uid, id, nextOfPeriod int64) error {
	return u.WrapResultErr(u.alive(c, alive).Where("uid = ? AND id = ?", uid, id).UpdateColumn("next_of_period", nextOfPeriod))
}

// UpdatesSequenceByUID 更新范围序号，需配合 UpdateSequenceByUID 使用
func (u *scheduleRepository) UpdatesSequenceByUID(c context.Context, alive bool, uid int64, origin, target int64) (err error) {
	if origin == target {
		return nil
	} else if origin < target {
		// update (origin, target] -= 1, origin = target
		err = u.WrapResultErr(u.alive(c, alive).Where("uid = ? AND sequence > ? AND sequence <= ?", uid, origin, target).UpdateColumn("sequence", gorm.Expr("sequence - ?", 1)))
	} else if target < origin {
		// update [target, origin) += 1, origin = target
		err = u.WrapResultErr(u.alive(c, alive).Where("uid = ? AND sequence >= ? AND sequence < ?", uid, target, origin).UpdateColumn("sequence", gorm.Expr("sequence + ?", 1)))
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return nil
}

// UpdateSequenceByUID 更新序号，需配合 UpdatesSequenceByUID 使用
func (u *scheduleRepository) UpdateSequenceByUID(c context.Context, alive bool, uid, id, targetSequence int64) error {
	return u.WrapResultErr(u.alive(c, alive).Where("uid = ? AND id = ?", uid, id).UpdateColumn("sequence", targetSequence))
}

// UpdateDisabledByUID 更新停用时间
func (u *scheduleRepository) UpdateDisabledByUID(c context.Context, alive bool, uid, id, timestamp int64) error {
	tx := u.alive(c, alive).Where("uid = ?", uid)
	if id != 0 {
		tx = tx.Where("id = ?", id)
	}
	return u.WrapResultErr(tx.UpdateColumn("disabled", timestamp))
}

/**
删除
*/

// Delete 删除记录
func (u *scheduleRepository) Delete(c context.Context, uid, id int64) error {
	return u.WrapResultErr(u.db(c).Where("uid = ? AND id = ?", uid, id).Delete(&model.Schedule{}))
}
