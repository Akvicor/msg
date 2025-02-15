package server

import (
	"github.com/urfave/cli/v2"
	"msg/cmd/app/server/app"
	"msg/cmd/app/server/common/cache"
	"msg/cmd/app/server/common/db"
	"msg/cmd/app/server/global/auth"
	"msg/cmd/config"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "config",
		Usage:   "config file path",
		Value:   "./data/config.toml",
		Aliases: []string{"c"},
	},
}

func Action(ctx *cli.Context) (err error) {
	// 加载配置文件
	config.Load(ctx.String("config"))
	// 加载数据库
	db.Load()
	// 配置缓存
	cache.SetTokenManagerOnEvicted(auth.OnTokenEvicted)
	// 运行服务
	return app.Run()
}
