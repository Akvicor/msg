package app

import (
	"github.com/urfave/cli/v2"
	"msg/cmd/app/example"
	"msg/cmd/app/migrate"
	"msg/cmd/app/server"
)

var commands = []*cli.Command{
	{
		Name:                   "server",
		Usage:                  "HTTP Server",
		UseShortOptionHandling: true,
		Action:                 server.Action,
		Flags:                  server.Flags,
	},
	{
		Name:                   "migrate",
		Usage:                  "Migrate Database",
		UseShortOptionHandling: true,
		Action:                 migrate.Action,
		Flags:                  migrate.Flags,
	},
	{
		Name:                   "example",
		Usage:                  "Generate Example",
		UseShortOptionHandling: true,
		Action:                 example.Action,
		Flags:                  example.Flags,
	},
}
