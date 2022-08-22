package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb://localhost:27017/"

func NewStorage() *MongoStorage {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	db := client.Database("exampleDB")
	coll := db.Collection("exampleColl")

	return &MongoStorage{
		client:     client,
		collection: coll,
	}
}

//go:generate mockery --name=Storage --output ../internal/mocks
type Storage interface {
	Disconnect()
	Create(data CreateData) (CreateResp, error)
	Read() ([]ReadResp, error)
}

type MongoStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (s MongoStorage) Disconnect() {
	if err := s.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
