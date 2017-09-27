package app

import (
	"gopkg.in/urfave/cli.v1"
)

const(
	Service = "SERVICE"
	Version = "VERSION"
	ServiceId = "SERVICE_ID"
	Env 	  = "ENV"
)

type NavyhookClientCmdOptions struct {
	Service string
	Version string
	ServiceId string
	Env string
	DelayThreshold int
	BundleCountThreshold int
	Port string
	ElasticIp string
	ElasticIndex string

}

func NewNavyhookClientCmdOptions() *NavyhookClientCmdOptions{
	return &NavyhookClientCmdOptions{

	}
}

func (opts *NavyhookClientCmdOptions) AddFlags(app *cli.App){

	flags := []cli.Flag{
		cli.StringFlag{
			Name:        "svc",
			Value:       "",
			Usage:       "service",
			EnvVar:      Service,
			Destination: &opts.Service,
		},
		cli.StringFlag{
			Name:        "svc.version",
			Value:       "",
			Usage:       "version",
			EnvVar:      Version,
			Destination: &opts.Version,
		},
		cli.StringFlag{
			Name:        "svc.id",
			Value:       "",
			Usage:       "service-id",
			EnvVar:      ServiceId,
			Destination: &opts.ServiceId,
		},
		cli.StringFlag{
			Name:        "env",
			Value:       "",
			Usage:       "env",
			EnvVar:      Env,
			Destination: &opts.Env,
		},
		cli.IntFlag{
			Name:        "delay_threshold",
			Value:       10,
			Usage:       "delay_threshold",
			Destination: &opts.DelayThreshold,
		},
		cli.IntFlag{
			Name:        "bundle_count_threshold",
			Value:       50,
			Usage:       "bundle_count_threshold",
			Destination: &opts.BundleCountThreshold,
		},
		cli.StringFlag{
			Name:        "port",
			Value:       ":5000",
			Usage:       "port",
			Destination: &opts.Port,
		},
		cli.StringFlag{
			Name:        "elastic.endpoint",
			Value:       "http://localhost:9200",
			Usage:       "elastic.endpoint",
			Destination: &opts.ElasticIp,
		},
		cli.StringFlag{
			Name:        "elastic.index",
			Value:       "analytics",
			Usage:       "elastic.index",
			Destination: &opts.ElasticIndex,
		},

	}

	app.Flags = append(app.Flags, flags...)
}
