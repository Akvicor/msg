package app

import (
	"github.com/urfave/cli/v2"
	"msg/cmd/def"
)

var App *cli.App

func init() {
	App = &cli.App{
		Name:                   def.AppName,
		Usage:                  def.AppUsage,
		Version:                def.AppVersion(),
		Description:            def.AppDescription,
		Authors:                []*cli.Author{{Name: "Akvicor", Email: "akvicor@akvicor.com"}},
		UseShortOptionHandling: true,
		DefaultCommand:         "help",
		Flags:                  flags,
		Commands:               commands,
		Action:                 action,
	}
}
