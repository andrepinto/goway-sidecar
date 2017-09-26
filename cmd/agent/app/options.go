package app

import (
	"gopkg.in/urfave/cli.v1"
)

const(
	Project = "NAVY_PROJECT"
)

type NavyhookClientCmdOptions struct {
	Project 			string

}

func NewNavyhookClientCmdOptions() *NavyhookClientCmdOptions{
	return &NavyhookClientCmdOptions{

	}
}

func (opts *NavyhookClientCmdOptions) AddFlags(app *cli.App){

	flags := []cli.Flag{
		cli.StringFlag{
			Name:        "navy.project",
			Value:       "",
			Usage:       "Navyhook projectId",
			EnvVar:      Project,
			Destination: &opts.Project,
		},

	}

	app.Flags = append(app.Flags, flags...)
}
