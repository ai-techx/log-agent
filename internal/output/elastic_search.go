package output

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/google/logger"
	"log-agent/internal/config"
)

type ElasticSearchClient[T any] struct {
	config config.ElasticSearch
	client *elasticsearch.TypedClient
}

// NewElasticSearchClient creates a new elastic search client
// This client will write the data to elastic search
func NewElasticSearchClient[T any](config config.ElasticSearch) *ElasticSearchClient[T] {
	typedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		CloudID: config.CloudId,
		APIKey:  config.ApiKey,
	})

	// check if the elastic search index exists
	do, err := typedClient.Indices.Get(config.ElasticSearchIndexName).Do(context.TODO())
	if do == nil {
		logger.Info("Creating elastic search index")
		_, err = typedClient.Indices.Create(config.ElasticSearchIndexName).Do(context.TODO())
		if err != nil {
			logger.Fatal(err)
		}
	}

	if err != nil {
		panic(err)
	}

	return &ElasticSearchClient[T]{config: config, client: typedClient}
}

func (e *ElasticSearchClient[T]) Write(data T) error {
	result, err := e.client.Index(e.config.ElasticSearchIndexName).Request(data).Do(context.Background())
	if err != nil {
		return err
	}

	logger.Info(result)
	return nil
}
