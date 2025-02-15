package example

import (
	"github.com/urfave/cli/v2"
	"msg/cmd/config"
)

var Flags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "config",
		Usage:   "Config",
		Value:   false,
		Aliases: []string{"c"},
	},
	&cli.StringFlag{
		Name:    "path",
		Usage:   "path",
		Value:   "./",
		Aliases: []string{"p"},
	},
}

func Action(ctx *cli.Context) (err error) {
	c := ctx.Bool("config")
	basePath := ctx.String("path")
	if c {
		config.GenerateExample(basePath)
	}
	return nil
}
