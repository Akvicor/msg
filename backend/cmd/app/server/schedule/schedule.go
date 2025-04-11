package schedule

import (
	"context"
	"errors"
	"github.com/Akvicor/glog"
	"gorm.io/gorm"
	"msg/cmd/app/server/bot"
	"msg/cmd/app/server/common/smap"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/service"
	"sync"
	"time"
)

type Schedule struct {
	refreshSig chan struct{}
	cancelsWg  *sync.WaitGroup
	cancels    *smap.SMap[int64, context.CancelFunc]
}

func NewSchedule() *Schedule {
	return &Schedule{
		cancelsWg: &sync.WaitGroup{},
		cancels:   smap.NewSMap[int64, context.CancelFunc](),
	}
}

func (s *Schedule) Refresh() {
	s.refreshSig <- struct{}{}
}

func (s *Schedule) Run(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	glog.Info("Schedule OK")
	wait := 5 * time.Second
	recent := int64(10)
	for {
		select {
		case <-ctx.Done():
			glog.Info("Schedule existing")
			s.cancels.Range(func(key int64, cancel context.CancelFunc) bool {
				cancel()
				return true
			})
			s.cancelsWg.Wait()
			glog.Info("Schedule existed")
			return
		case <-time.After(wait): // 等待下一个信息的发送时间
		case <-s.refreshSig: // 主动刷新
		}
		schedules, err := service.Schedule.FindAllRecentBySecond(nil, recent)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			glog.Warning("find schedule error: %v", err)
			continue
		}
		for _, item := range schedules {
			itemCtx, itemCancel := context.WithCancel(context.Background())
			if s.cancels.Set(item.ID, itemCancel) {
				continue
			}
			s.cancelsWg.Add(1)
			go s.send(itemCtx, item)
		}
	}
}

func (s *Schedule) send(ctx context.Context, schedule *model.Schedule) {
	defer func() {
		s.cancels.Delete(schedule.ID)
		s.cancelsWg.Done()
	}()
	timeout := max(schedule.NextOfPeriod-time.Now().Unix(), 0)
	select {
	case <-ctx.Done():
		glog.Warning("send schedule cancel")
		return
	case <-time.After(time.Duration(timeout) * time.Second):
	}
	target, err := service.Schedule.FindByID(true, nil, schedule.UID, schedule.ID)
	if err != nil {
		glog.Warning("find schedule error: %v", err)
		return
	}
	if target.NextOfPeriod != schedule.NextOfPeriod {
		// 已经发送过
		return
	}
	_, err = bot.Sender.Send(target.UID, target.ChannelID, target.ID, target.NextOfPeriod, "", target.Type, target.Title, target.Message)
	if err != nil {
		glog.Warning("send schedule error: %v", err)
		return
	}
	_ = service.Schedule.UpdateNextBy(target.UID, target.ID, target.NextPeriod())
}
