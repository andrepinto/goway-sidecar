package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"gopkg.in/olivere/elastic.v3"
)

const (
	pageLimit = 20
)

type ElasticSearch struct {
	uri    string
	client *elastic.Client
}

func CreateElasticSearchConn(uri string) *ElasticSearch {
	return &ElasticSearch{uri, nil}
}

func (es *ElasticSearch) Conn() error {

	fmt.Println(es.uri, 10*time.Second)

	client, err := elastic.NewClient(
		elastic.SetURL(es.uri),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetMaxRetries(5))

	if err != nil {
		// Handle error
		return err
	}

	es.client = client
	return nil
}

func (es *ElasticSearch) Find(index string, query elastic.Query, table string, params ...int) ([]interface{}, int64, error) {
	var objects []interface{}

	skipCount := 0

	if len(params) >= 1 {
		if params[0] > 1 {
			skipCount = (params[0] - 1) * pageLimit
		}
	}

	searchResult, err := es.client.Search().
		Index(index).
		Type(table).
		Query(query).
		From(skipCount).Size(pageLimit).
		Pretty(true).
		Do()
	if err != nil {
		return nil, 0, err
	}

	log.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	if searchResult.Hits != nil {
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			var model interface{}

			err := json.Unmarshal(*hit.Source, &model)
			if err != nil {
				return nil, 0, err
			}

			objects = append(objects, model)

		}
	}

	return objects, searchResult.Hits.TotalHits, nil
}

func (es *ElasticSearch) FindById(index string, id string, table string, typeOf interface{}) (interface{}, int64, error) {

	skipCount := 0

	query := elastic.NewTermQuery("_id", id)

	searchResult, err := es.client.Search().
		Index(index).
		Type(table).
		Query(query).
		From(skipCount).Size(pageLimit).
		Pretty(true).
		Do()
	if err != nil {
		return nil, 0, err
	}

	log.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	var model interface{}

	for _, item := range searchResult.Each(reflect.TypeOf(typeOf)) {
		return item, searchResult.Hits.TotalHits, nil
	}

	return model, searchResult.Hits.TotalHits, nil
}

func (es *ElasticSearch) Insert(index string, table string, id string, model interface{}) error {
	_, err := es.client.Index().
		Index(index).
		Type(table).
		Id(id).
		BodyJson(model).
		Do()

	if err != nil {
		return err
	}

	// Flush data
	_, err = es.client.Flush().Index(index).Do()

	if err != nil {
		return err
	}

	return nil
}

func (es *ElasticSearch) InsertByID(index string, table, id string, model interface{}) error {
	_, err := es.client.Index().
		Index(index).
		Type(table).
		Id(id).
		BodyJson(model).
		Do()

	if err != nil {
		return err
	}

	// Flush data
	_, err = es.client.Flush().Index(index).Do()

	if err != nil {
		return err
	}

	return nil
}

func (es *ElasticSearch) Delete(index string, table string, query elastic.Query) error {
	res, err := es.client.DeleteByQuery().Index(index).Type(table).Query(query).Do()
	if err != nil {
		return err
	}

	if res == nil {
		return errors.New("response is nil")
	}

	_, found := res.Indices[index]
	if !found {
		log.Printf("expected Found = true; got: %v", found)
	}

	_, err = es.client.Flush().Index(index).Do()
	if err != nil {
		return err
	}
	return nil
}

func (es *ElasticSearch) DeleteID(index string, table string, id string) error {
	res, err := es.client.Delete().Index(index).Type(table).Id(id).Do()
	if err != nil {
		return err
	}

	if res.Found != true {
		return errors.New("document not found")
	}

	_, err = es.client.Flush().Index(index).Do()
	if err != nil {
		return err
	}
	return nil
}

func (es *ElasticSearch) Update(index string, table string, id string, data map[string]interface{}) error {
	_, err := es.client.Update().Index(index).Type(table).Id(id).Doc(data).Do()

	if err != nil {
		return err
	}

	return nil
}

//bulk methods

func (es *ElasticSearch) NewBulk() *elastic.BulkService {
	return es.client.Bulk()
}

func (es *ElasticSearch) AddToBulk(index string, bulk *elastic.BulkService, table string, model interface{}, id string) {
	bulk.Add(elastic.NewBulkIndexRequest().Index(index).Type(table).Doc(model).Id(id))
}

func (es *ElasticSearch) SendBulk(bulk *elastic.BulkService) {
	bulk.Do()
}
