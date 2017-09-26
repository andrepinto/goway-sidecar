package repository

import (
	"github.com/andrepinto/goway-sidecar/helpers"
	"github.com/twinj/uuid"
	"github.com/andrepinto/goway-sidecar/outputs"
	"fmt"
)

const (
	HTTP_LOGGER_INDEX_TYPE = "http-logger"
)

type ElasticSearchRepo struct {
	Client *helpers.ElasticSearch
	Index string
}

func NewElasticsearchOutput(client *helpers.ElasticSearch, index string) *ElasticSearchRepo{
	return &ElasticSearchRepo{
		Client: client,
		Index: index,
	}
}

func(repo *ElasticSearchRepo) Send(data []*outputs.HttpLoggerClient) error{
	bulk := repo.Client.NewBulk()
	for _, v := range data{
		repo.Client.AddToBulk(repo.Index, bulk, HTTP_LOGGER_INDEX_TYPE, v, uuid.NewV4().String())
	}

	fmt.Println("*******************")

	fmt.Println(data[0].Timestamp)

	repo.Client.SendBulk(bulk)

	return nil
}
