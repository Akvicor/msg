package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"msg/cmd/app/server/common/period"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/send"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/repository"
	"time"
)

var Schedule = new(scheduleService)

type scheduleService struct {
	base
}

// FindAll 获取所有
func (u *scheduleService) FindAll(page *resp.PageModel, alive bool, preload *model.PreloaderSchedule) (tokens []*model.Schedule, err error) {
	return repository.Schedule.FindAll(context.Background(), page, alive, preload)
}

// FindAllRecentBySecond 获取所有
func (u *scheduleService) FindAllRecentBySecond(preload *model.PreloaderSchedule, seconds int64) (tokens []*model.Schedule, err error) {
	return repository.Schedule.FindAllRecentBySecond(context.Background(), preload, seconds)
}

// FindAllByUID 获取用户全部
func (u *scheduleService) FindAllByUID(page *resp.PageModel, alive bool, preload *model.PreloaderSchedule, uid int64) (tokens []*model.Schedule, err error) {
	return repository.Schedule.FindAllByUID(context.Background(), page, alive, preload, uid)
}

// FindAllByUIDLike 搜索
func (u *scheduleService) FindAllByUIDLike(page *resp.PageModel, alive bool, preload *model.PreloaderSchedule, uid int64, like string) (tokens []*model.Schedule, err error) {
	return repository.Schedule.FindAllByUIDLike(context.Background(), page, alive, preload, uid, like)
}

// FindByID 获取
func (u *scheduleService) FindByID(alive bool, preload *model.PreloaderSchedule, uid, id int64) (*model.Schedule, error) {
	return repository.Schedule.FindByID(context.Background(), alive, preload, uid, id)
}

// Create 创建
func (u *scheduleService) Create(uid int64, category string, sendType send.Type, title, message string, channelID int64, periodType period.Type, startAt int64, year int, quarter int, month int, week int, day int, hour int, minute int, second int, expirationDate int64, expirationTimes int64) (schedule *model.Schedule, err error) {
	schedule = model.NewSchedule(uid, category, sendType, title, message, channelID, periodType, startAt, year, quarter, month, week, day, hour, minute, second, expirationDate, expirationTimes)
	err = u.transaction(context.Background(), func(ctx context.Context) error {
		// 判断用户是否存在且未被停用
		exist, e := repository.User.ExistById(ctx, true, uid)
		if !exist || errors.Is(e, gorm.ErrRecordNotFound) {
			return errors.New("用户已被停用")
		}
		// 创建访问密钥
		e = repository.Schedule.Create(ctx, schedule)
		if e != nil {
			return e
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

// Update 更新
func (u *scheduleService) Update(uid, id int64, category string, sendType send.Type, title, message string, channelID int64, periodType period.Type, startAt int64, year int, quarter int, month int, week int, day int, hour int, minute int, second int, expirationDate int64, expirationTimes int64) error {
	schedule := model.NewSchedule(uid, category, sendType, title, message, channelID, periodType, startAt, year, quarter, month, week, day, hour, minute, second, expirationDate, expirationTimes)
	schedule.ID = id
	return repository.Schedule.Update(context.Background(), false, schedule)
}

// UpdateNext 更新下次提醒时间
func (u *scheduleService) UpdateNext(uid, id int64) error {
	err := u.transaction(context.Background(), func(ctx context.Context) error {
		// 判断用户是否存在且未被停用
		exist, e := repository.User.ExistById(ctx, true, uid)
		if !exist || errors.Is(e, gorm.ErrRecordNotFound) {
			return errors.New("用户已被停用")
		}
		// 获取
		schedule, e := repository.Schedule.FindByID(ctx, false, nil, uid, id)
		if e != nil {
			return e
		}
		// 创建访问密钥
		e = repository.Schedule.UpdateNext(ctx, false, uid, id, schedule.NextPeriod())
		if e != nil {
			return e
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateSequenceByUID 更新序号
func (u *scheduleService) UpdateSequenceByUID(uid, id, targetSequence int64) error {
	err := u.transaction(context.Background(), func(ctx context.Context) error {
		origin, err := repository.Schedule.FindByID(ctx, false, nil, uid, id)
		if err != nil {
			return err
		}
		maxSequence, err := repository.Schedule.GetMaxSequenceByUID(ctx, false, uid)
		if err != nil {
			return nil
		}
		if targetSequence < 1 {
			targetSequence = 1
		}
		if targetSequence > maxSequence {
			targetSequence = maxSequence
		}
		err = repository.Schedule.UpdatesSequenceByUID(ctx, false, uid, origin.Sequence, targetSequence)
		if err != nil {
			return err
		}
		err = repository.Schedule.UpdateSequenceByUID(ctx, false, uid, id, targetSequence)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateNextBy 更新下次提醒时间
func (u *scheduleService) UpdateNextBy(uid, id, nextOfPeriod int64) error {
	return repository.Schedule.UpdateNext(context.Background(), false, uid, id, nextOfPeriod)
}

// UpdateDisabledByUID 停用/启用
func (u *scheduleService) UpdateDisabledByUID(uid, id int64, disable bool) error {
	if disable {
		return repository.Schedule.UpdateDisabledByUID(context.Background(), false, uid, id, time.Now().Unix())
	} else {
		return repository.Schedule.UpdateDisabledByUID(context.Background(), false, uid, id, 0)
	}
}

// Delete 删除
func (u *scheduleService) Delete(uid, id int64) error {
	return repository.Schedule.Delete(context.Background(), uid, id)
}
