package storage

import (
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateData struct {
	User        string             `json:"user" bson:"user"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   primitive.DateTime `bson:"createdAt"`
}

type CreateResp struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
}

func (s *MongoStorage) Create(data CreateData) (CreateResp, error) {
	ctx := context.TODO()
	data.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	if _, err := s.collection.InsertOne(ctx, data); err != nil {
		return CreateResp{}, err
	}
	return CreateResp{
		Status: http.StatusNoContent,
		Msg:    "ok",
	}, nil
}
