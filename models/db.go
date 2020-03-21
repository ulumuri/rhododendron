package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Requester interface {
	Insert(post *Post) (*mongo.InsertOneResult, error)
	FindByID(id string) (*Post, error)
	DeleteByID(id string) (*mongo.DeleteResult, error)
	DefineEndPoints()
}

type Conn struct {
	db  *mongo.Database
	ctx context.Context
}

func Connect(client *mongo.Client, ctx context.Context, dbName string) *Conn {
	return &Conn{
		client.Database(dbName),
		ctx,
	}
}
