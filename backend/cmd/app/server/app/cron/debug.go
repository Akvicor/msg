package cron

import (
	"github.com/Akvicor/glog"
	"github.com/go-co-op/gocron/v2"
	"msg/cmd/config"
)

var Debug = new(debugCron)

type debugCron struct{}

// Daily 每日通知
func (s *debugCron) Daily(scheduler gocron.Scheduler) (err error) {
	// 执行时间
	definition := gocron.DailyJob(
		1,
		gocron.NewAtTimes(
			gocron.NewAtTime(config.Global.Cron.Debug.Hour, config.Global.Cron.Debug.Minute, config.Global.Cron.Debug.Second),
		),
	)
	// 任务
	task := gocron.NewTask(
		func() {
			glog.Debug("daily ")

		},
	)
	// 创建cron
	_, err = scheduler.NewJob(definition, task)
	return err
}
