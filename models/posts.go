package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	ID        primitive.ObjectID    `bson:"_id,omitempty" json:"id,omitempty"`
	Timestamp primitive.Timestamp   `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Author    string                `bson:"author,omitempty" json:"author,omitempty"`
	Message   string                `bson:"message,omitempty" json:"message,omitempty"`
	Tags      []string              `bson:"tags,omitempty" json:"tags,omitempty"`
	Shared    []primitive.DBPointer `bson:"shared,omitempty" json:"shared,omitempty"`
}

const POST_COLL = "posts"

func (c *Conn) Insert(post *Post) (*mongo.InsertOneResult, error) {
	postCollection := c.Collection(POST_COLL)
	postResult, err := postCollection.InsertOne(context.Background(), post)
	if err != nil {
		return nil, err
	}

	return postResult, nil
}

func (c *Conn) FindByID(id primitive.ObjectID) (*Post, error) {
	post := &Post{}
	filter := bson.M{"_id": id}
	postCollection := c.Collection(POST_COLL)
	err := postCollection.FindOne(context.Background(), filter).Decode(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (c *Conn) DeleteByID(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	postCollection := c.Collection(POST_COLL)
	result, err := postCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
