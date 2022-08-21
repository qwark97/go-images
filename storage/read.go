package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReadResp struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	User        string             `json:"user" bson:"user"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   primitive.DateTime `json:"createdAt" bson:"createdAt"`
}

func (s *Storage) Read() ([]ReadResp, error) {
	var cursor *mongo.Cursor
	var err error
	var response = []ReadResp{}

	ctx := context.TODO()
	if cursor, err = s.collection.Find(ctx, bson.D{}); err != nil {
		return nil, err
	}
	if err := cursor.All(context.TODO(), &response); err != nil {
		return nil, err
	}

	return response, nil
}
