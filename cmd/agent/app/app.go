package app

import (
	"github.com/andrepinto/goway-sidecar/collector"
	"github.com/andrepinto/goway-sidecar/helpers"
	"github.com/andrepinto/goway-sidecar/outputs/elasticsearch"
)

type NavyhookClientApp struct {

}

func NewNavyhookClientApp() *NavyhookClientApp{
	return &NavyhookClientApp{}
}


func (cli *NavyhookClientApp)Run(options *NavyhookClientCmdOptions) error{

	elasticClient := helpers.CreateElasticSearchConn("http://localhost")
	elasticClient.Conn()

	elasticsearchOut := repository.NewElasticsearchOutput(elasticClient, "analytics")

	collectorAgent := collector.NewCollectorRpcServer(&collector.CollectorRpcServerOptions{
		Port:":5000",
		Output: elasticsearchOut,
	})

	err := collectorAgent.Start()


	select {}

	return err
}