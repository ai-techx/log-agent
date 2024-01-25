package output

import (
	"context"
	"github.com/google/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log-agent/internal/config"
)

type MongoDB[T any] struct {
	config *config.MongoDB
	client *mongo.Client
}

func NewMongoDBClient[T any](conf config.MongoDB) *MongoDB[T] {
	if len(conf.DatabaseName) == 0 || len(conf.CollectionName) == 0 || len(conf.Endpoint) == 0 {
		logger.Fatal("MongoDB configuration is not valid, missing database name, collection name or endpoint")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.Endpoint))
	if err != nil {
		logger.Fatal(err)
	}

	return &MongoDB[T]{config: &conf, client: client}
}

func (m *MongoDB[T]) Write(data T) error {
	_, err := m.client.Database(m.config.DatabaseName).Collection(m.config.CollectionName).InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}
