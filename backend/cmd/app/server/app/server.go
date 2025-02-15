package app

import (
	"fmt"
	"github.com/Akvicor/glog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"msg/cmd/app/server/app/api"
	"msg/cmd/app/server/app/mw"
	"msg/cmd/app/server/bot"
	"msg/cmd/config"
	"net/http"
	"sync"
)

func runServer(wg *sync.WaitGroup, e *echo.Echo) {
	defer wg.Done()

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodDelete},
	}))
	e.Use(mw.Error)
	e.Use(middleware.Gzip())

	publicGroup := e.Group("")
	{
		web := getFS()
		publicGroup.FileFS("/", "index.html", web)
		publicGroup.FileFS("/*", "index.html", web)
		publicGroup.FileFS("/index.html", "/index.html", web)
		publicGroup.FileFS("/favicon.ico", "favicon.ico", web)
		publicGroup.FileFS("/manifest.json", "/manifest.json", web)
		publicGroup.FileFS("/asset-manifest.json", "/asset-manifest.json", web)
		publicGroup.GET("/static/*", WrapHandler(http.FileServer(http.FS(web))))
	}

	apiGroup := e.Group("/api")

	// Set Bot Group
	bot.Bot.SetGroupRoute(apiGroup.Group("/bot"))

	// 管理员
	apiGroupAdmin := apiGroup.Group("")
	apiGroupAdmin.Use(mw.AuthAdmin)
	// 管理员、普通用户
	apiGroupUser := apiGroup.Group("")
	apiGroupUser.Use(mw.AuthUser)
	// 管理员、普通用户、浏览者
	apiGroupViewer := apiGroup.Group("")
	apiGroupViewer.Use(mw.AuthViewer)
	// 已登录用户
	apiGroupAuth := apiGroup.Group("")
	apiGroupAuth.Use(mw.Auth)
	// 公开
	apiGroupPublic := apiGroup.Group("")

	// 调试、性能、系统信息
	{
		apiGroupPublic.GET("/sys/info/version", api.Sys.Version, mw.NewIPLimiter(3, 3))
		apiGroupPublic.GET("/sys/info/version/full", api.Sys.VersionFull, mw.NewIPLimiter(3, 3))
		apiGroupPublic.GET("/sys/info/branding", api.Sys.Branding, mw.NewIPLimiter(3, 3))
		apiGroupAdmin.GET("/sys/info/cache", api.Sys.InfoCache, mw.NewIPLimiter(3, 3))
		apiGroupPublic.GET("/sys/info/health", api.Sys.Health, mw.NewIPLimiter(3, 3))
	}

	// 类型数据
	{
		apiGroupPublic.GET("/type/role/type", api.Type.RoleType, mw.NewIPLimiter(3, 3))
		apiGroupPublic.GET("/type/channel/type", api.Type.ChannelType, mw.NewIPLimiter(3, 3))
		apiGroupPublic.GET("/type/send/type", api.Type.SendType, mw.NewIPLimiter(3, 3))
		apiGroupPublic.GET("/type/bot/senders", api.Type.BotSenders, mw.NewIPLimiter(3, 3))
	}

	// 用户相关
	{
		// 管理员
		{
			apiGroupAdmin.POST("/admin/user/create", api.Admin.CreateUser, mw.NewIPLimiter(1, 1))
			apiGroupAdmin.GET("/admin/user/find", api.Admin.Find, mw.NewIPLimiter(3, 3))
			apiGroupAdmin.POST("/admin/user/update", api.Admin.UpdateUser, mw.NewIPLimiter(3, 3))
			apiGroupAdmin.POST("/admin/user/disable", api.Admin.DisableUser, mw.NewIPLimiter(1, 1))
			apiGroupAdmin.POST("/admin/user/enable", api.Admin.EnableUser, mw.NewIPLimiter(1, 1))
			apiGroupAdmin.GET("/admin/user/access_token/all", api.Admin.AccessTokenInfo, mw.NewIPLimiter(3, 3))
			apiGroupAdmin.GET("/admin/user/login_log/all", api.Admin.LoginLogInfo, mw.NewIPLimiter(3, 3))
		}
		// 公开
		{
			apiGroupPublic.POST("/user/login", api.User.Login, mw.NewIPLimiter(3, 3))
		}
		// 登录用户
		{
			apiGroupAuth.POST("/user/logout", api.User.Logout, mw.NewIPLimiter(1, 1))
			apiGroupAuth.GET("/user/info", api.User.Info, mw.NewIPLimiter(3, 3))
			apiGroupAuth.POST("/user/update", api.User.Update, mw.NewIPLimiter(3, 3))
			apiGroupAuth.GET("/user/access_token/find", api.User.AccessTokenFind, mw.NewIPLimiter(3, 3))
			apiGroupAuth.POST("/user/access_token/create", api.User.AccessTokenCreate, mw.NewIPLimiter(1, 1))
			apiGroupAuth.POST("/user/access_token/update", api.User.AccessTokenUpdate, mw.NewIPLimiter(1, 1))
			apiGroupAuth.POST("/user/access_token/delete", api.User.AccessTokenDelete, mw.NewIPLimiter(1, 1))
			apiGroupAuth.GET("/user/login_log/info", api.User.LoginLogInfo, mw.NewIPLimiter(3, 3))
		}
		// 用户绑定数据
		{
			apiGroupAuth.GET("/user/bind/home/tips/find", api.UserBind.HomeTipsFind, mw.NewIPLimiter(3, 3))
			apiGroupAuth.POST("/user/bind/home/tips/save", api.UserBind.HomeTipsSave, mw.NewIPLimiter(3, 3))
		}
	}

	// 通知渠道
	{
		apiGroupAuth.POST("/channel/create", api.Channel.Create, mw.NewIPLimiter(1, 1))
		apiGroupAuth.GET("/channel/find", api.Channel.Find, mw.NewIPLimiter(3, 3))
		apiGroupAuth.POST("/channel/update", api.Channel.Update, mw.NewIPLimiter(3, 3))
		apiGroupAuth.POST("/channel/delete", api.Channel.Delete, mw.NewIPLimiter(1, 1))
		apiGroupAuth.GET("/channel/test", api.Channel.Test, mw.NewIPLimiter(1, 1))
		apiGroupAuth.GET("/channel/send", api.Channel.Send)
		apiGroupAuth.POST("/channel/send", api.Channel.Send)
	}

	// 发送通知
	{
		apiGroupAuth.GET("/send", api.Send.Send)
		apiGroupAuth.POST("/send", api.Send.Send)
		apiGroupAuth.GET("/send/find", api.Send.Find, mw.NewIPLimiter(3, 3))
		apiGroupAuth.POST("/send/find", api.Send.Find, mw.NewIPLimiter(3, 3))
		apiGroupAuth.GET("/send/cancel", api.Send.Cancel)
		apiGroupAuth.POST("/send/cancel", api.Send.Cancel)
		apiGroupAuth.GET("/send/status", api.Send.Status)
		apiGroupAuth.POST("/send/status", api.Send.Status)
	}

	var err error
	if config.Global.Server.EnableHttps && config.Global.Server.CrtFile != "" && config.Global.Server.KeyFile != "" {
		err = app.Server.StartTLS(fmt.Sprintf("%s:%d", config.Global.Server.HttpIp, config.Global.Server.HttpPort), config.Global.Server.CrtFile, config.Global.Server.KeyFile)
	} else {
		err = app.Server.Start(fmt.Sprintf("%s:%d", config.Global.Server.HttpIp, config.Global.Server.HttpPort))
	}
	glog.Warning("Server Exist with %v", err)
}
