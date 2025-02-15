package app

import (
	"github.com/Akvicor/glog"
	"github.com/go-co-op/gocron/v2"
	"msg/cmd/app/server/app/cron"
	"sync"
)

func runCron(wg *sync.WaitGroup, scheduler gocron.Scheduler) {
	defer wg.Done()

	var gErr []error
	var err error

	// 每日通知
	err = cron.Debug.Daily(scheduler)
	if err != nil {
		gErr = append(gErr, err)
		glog.Warning("Cron Debug [daily] failed: %v", err)
	}

	if len(gErr) == 0 {
		glog.Info("Cron OK")
	}
	scheduler.Start()
	return
}
