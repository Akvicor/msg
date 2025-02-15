package migrate

import (
	"github.com/Akvicor/glog"
	"github.com/urfave/cli/v2"
	"msg/cmd/app/server/common/db"
	"msg/cmd/config"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "config",
		Usage:   "config file path",
		Value:   "./data/config.toml",
		Aliases: []string{"c"},
	},
	&cli.StringFlag{
		Name:    "username",
		Usage:   "Admin Username",
		Value:   "admin",
		Aliases: []string{"u"},
	},
	&cli.StringFlag{
		Name:    "password",
		Usage:   "Admin Password",
		Value:   "password",
		Aliases: []string{"p"},
	},
	&cli.StringFlag{
		Name:    "nickname",
		Usage:   "Admin Nickname",
		Value:   "Akvicor",
		Aliases: []string{"n"},
	},
	&cli.StringFlag{
		Name:    "avatar",
		Usage:   "Admin Avatar",
		Value:   "",
		Aliases: []string{"a"},
	},
	&cli.StringFlag{
		Name:    "mail",
		Usage:   "Admin Mail",
		Value:   "",
		Aliases: []string{"m"},
	},
	&cli.StringFlag{
		Name:    "phone",
		Usage:   "Admin Phone",
		Value:   "",
		Aliases: []string{"o"},
	},
}

func Action(ctx *cli.Context) (err error) {
	// 加载配置文件
	config.Load(ctx.String("config"))
	// 加载数据库
	_ = db.Create()

	d := db.Get()

	if err = d.AutoMigrate(list...); err != nil {
		glog.Fatal("初始化数据库表结构异常: %v", err)
		return err
	}

	err = initUser(ctx)
	if err != nil {
		return err
	}

	return nil
}
