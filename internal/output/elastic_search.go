package output

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log-agent/internal/config"
)

type ElasticSearchClient[T any] struct {
	config config.ElasticSearch
	client *elasticsearch.TypedClient
}

func NewElasticSearchClient[T any](config config.ElasticSearch) *ElasticSearchClient[T] {
	typedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		CloudID: config.CloudId,
		APIKey:  config.ApiKey,
	})

	if err != nil {
		panic(err)
	}

	return &ElasticSearchClient[T]{config: config, client: typedClient}
}

func (e *ElasticSearchClient[T]) Write(data T) error {
	return nil
}
