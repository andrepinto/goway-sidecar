package app

import (
	"sort"
	"gopkg.in/urfave/cli.v1"
)

func NewCliApp() *cli.App {

	app := cli.NewApp()

	app.Name = "navyhook-client"

	opts := NewNavyhookClientCmdOptions()
	opts.AddFlags(app)

	app.Action = func(c *cli.Context) error {
		//worker := opts.NewWorker()
		proc := NewNavyhookClientApp()
		return proc.Run(opts)
	}

	// sort flags by name
	sort.Sort(cli.FlagsByName(app.Flags))

	return app
}
