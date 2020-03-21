package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Requester interface {
	Insert(post *Post) (*mongo.InsertOneResult, error)
	FindByID(id primitive.ObjectID) (*Post, error)
	DeleteByID(id primitive.ObjectID) (*mongo.DeleteResult, error)
}

type Conn struct {
	*mongo.Database
}

func Connect(client *mongo.Client, dbName string) *Conn {
	return &Conn{client.Database(dbName)}
}
