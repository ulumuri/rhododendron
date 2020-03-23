package database

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

type PostStore struct {
	db *mongo.Database
}

func NewPostStore(db *mongo.Database) *PostStore {
	return &PostStore{
		db: db,
	}
}

const POST_COLL = "posts"

func (s *PostStore) Create(post *Post) (*mongo.InsertOneResult, error) {
	postCollection := s.db.Collection(POST_COLL)
	postResult, err := postCollection.InsertOne(context.Background(), post)
	if err != nil {
		return nil, err
	}

	return postResult, nil
}

func (s *PostStore) Get(id primitive.ObjectID) (*Post, error) {
	post := &Post{}
	filter := bson.M{"_id": id}
	postCollection := s.db.Collection(POST_COLL)
	err := postCollection.FindOne(context.Background(), filter).Decode(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostStore) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	postCollection := s.db.Collection(POST_COLL)
	result, err := postCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *PostStore) ListAll() (*[]Post, error) {
	postCollection := s.db.Collection(POST_COLL)
	cursor, err := postCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var posts []Post
	for cursor.Next(context.Background()) {
		post := Post{}
		err = cursor.Decode(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return &posts, nil
}
