package mongo

import (
	"context"

	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var DB *MongoClient

type MongoClient struct {
	*mongo.Client
}

// Example code for reference
func newMongoClient(opts ...*options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.Connect(opts...)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client, nil
}

func Init() {
	mongoClient, err := newMongoClient()
	if err != nil {
		panic(err)
	}
	_ = mongox.NewClient(mongoClient, &mongox.Config{})
	// database := client.NewDatabase("db-test")
}
