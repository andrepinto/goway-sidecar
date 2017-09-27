package app

import (
	"github.com/andrepinto/goway-sidecar/collector"
	"github.com/andrepinto/goway-sidecar/helpers"
	"github.com/andrepinto/goway-sidecar/outputs/elasticsearch"
)

type NavyhookClientApp struct {
}

func NewNavyhookClientApp() *NavyhookClientApp {
	return &NavyhookClientApp{}
}

func (cli *NavyhookClientApp) Run(options *NavyhookClientCmdOptions) error {

	elasticClient := helpers.CreateElasticSearchConn(options.ElasticIp)
	elasticClient.Conn()

	sContext := map[string]string{
		"service":     options.ServiceId,
		"version":     options.Version,
		"service_id":  options.ServiceId,
		"environment": options.Env,
	}

	elasticsearchOut := repository.NewElasticsearchOutput(elasticClient, options.ElasticIndex)

	collectorAgent := collector.NewCollectorRpcServer(&collector.CollectorRpcServerOptions{
		Port:                 options.Port,
		Output:               elasticsearchOut,
		BundleCountThreshold: options.BundleCountThreshold,
		DelayThreshold:       options.DelayThreshold,
		Context:              sContext,
	})

	err := collectorAgent.Start()

	select {}

	return err
}
