package app

import (
	"github.com/andrepinto/goway-sidecar/collector"
	"github.com/andrepinto/goway-sidecar/helpers"
	"github.com/andrepinto/goway-sidecar/outputs/elasticsearch"
	log "github.com/sirupsen/logrus"
)

type NavyhookClientApp struct {
}

func NewNavyhookClientApp() *NavyhookClientApp {
	return &NavyhookClientApp{}
}

func (cli *NavyhookClientApp) Run(options *NavyhookClientCmdOptions) error {

	log.SetLevel(log.DebugLevel)

	elasticClient := helpers.CreateElasticSearchConn(options.ElasticIp)
	err := elasticClient.Conn()

	log.Debug(err)

	sContext := map[string]string{
		"service":     options.Service,
		"version":     options.Version,
		"service_id":  options.ServiceId,
		"environment": options.Env,
	}

	log.Debug(sContext)

	elasticsearchOut := repository.NewElasticsearchOutput(elasticClient, options.ElasticIndex)

	collectorAgent := collector.NewCollectorRpcServer(&collector.CollectorRpcServerOptions{
		Port:                 options.Port,
		Output:               elasticsearchOut,
		BundleCountThreshold: options.BundleCountThreshold,
		DelayThreshold:       options.DelayThreshold,
		Context:              sContext,
	})

	err = collectorAgent.Start()

	log.Info("Server started")

	select {}

	return err
}
