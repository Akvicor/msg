package app

import (
	"context"
	"github.com/Akvicor/glog"
	"github.com/go-co-op/gocron/v2"
	"github.com/labstack/echo/v4"
	"msg/cmd/app/server/bot"
	"msg/cmd/app/server/global/auth"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var app *App

type App struct {
	Cron   gocron.Scheduler
	Server *echo.Echo
}

func newApp() *App {
	return &App{}
}

func init() {
	app = newApp()
}

func Run() error {
	var err error

	if err = auth.LoadFromDBToCache(); err != nil {
		glog.Fatal("cannot load token from login log. %v", err)
	}

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	// Cron
	app.Cron, err = gocron.NewScheduler()
	if err != nil {
		glog.Fatal("cannot create cron. %v", err)
	}
	wg.Add(1)
	go runCron(wg, app.Cron)

	// Bot
	wg.Add(1)
	go bot.Bot.Run(wg, ctx)

	// Sender
	wg.Add(1)
	go bot.Sender.Run(wg, ctx)

	// Server
	app.Server = echo.New()
	app.Server.HideBanner = true
	wg.Add(1)
	go runServer(wg, app.Server)

	// Shutdown Signal
	shutdown := func() {
		cancel()
		_ = app.Server.Close()
		_ = app.Cron.Shutdown()
		wg.Wait()
	}
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			shutdown()
			return nil
		case syscall.SIGUSR1:
			glog.Info("Sig User 1")
		case syscall.SIGUSR2:
			glog.Info("Sig User 2")
		default:
		}
	}
	shutdown()
	return nil
}
