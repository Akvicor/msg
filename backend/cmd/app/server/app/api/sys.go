package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"msg/cmd/app/server/app/dro"
	"msg/cmd/app/server/bot"
	"msg/cmd/app/server/bot/status"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/types/channel"
	"msg/cmd/app/server/global/sys"
	"msg/cmd/config"
	"msg/cmd/def"
)

var Sys = new(sysApi)

type sysApi struct{}

func (a *sysApi) Version(c echo.Context) (err error) {
	return resp.SuccessWithData(c, def.Version)
}

func (a *sysApi) VersionFull(c echo.Context) (err error) {
	return resp.SuccessWithData(c, def.AppVersion())
}

func (a *sysApi) Branding(c echo.Context) error {
	return resp.SuccessWithData(c, resp.Map{
		"name":      config.Global.AppName,
		"copyright": "MSG",
	})
}

func (a *sysApi) InfoCache(c echo.Context) (err error) {
	tokenManager := sys.TokenManagerItems()
	loginFailedManager := sys.LoginFailedManagerItems()

	return resp.SuccessWithData(c, resp.Map{
		"token_manager_count":        len(tokenManager),
		"token_manager":              tokenManager,
		"login_failed_manager_count": len(loginFailedManager),
		"login_failed_manager":       loginFailedManager,
	})
}

func (a *sysApi) Health(c echo.Context) error {
	health := &dro.SysHealth{
		Status:   "healthy",
		BotError: make([]string, 0),
	}
	bots := bot.Bot.List()
	for _, b := range bots {
		for _, cType := range channel.AllType {
			senderStatus, receiverStatus := b.Status(cType.Type)
			if senderStatus != status.SenderStatusNotSupported && senderStatus != status.SenderStatusOK {
				health.Status = "unhealthy"
				health.BotError = append(health.BotError, fmt.Sprintf("%s %s sender %s", b.Name(), cType.Name, senderStatus.ToString()))
			}
			if receiverStatus != status.SenderStatusNotSupported && receiverStatus != status.SenderStatusOK {
				health.Status = "unhealthy"
				health.BotError = append(health.BotError, fmt.Sprintf("%s %s receiver %s", b.Name(), cType.Name, receiverStatus.ToString()))
			}
		}
	}
	if health.Status == "healthy" {
		return resp.Healthy(c, health)
	} else {
		return resp.Unhealthy(c, health)
	}
}
