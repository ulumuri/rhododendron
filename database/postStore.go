package database

import (
	"context"

	"github.com/ulumuri/rhododendron/errors"
	"github.com/ulumuri/rhododendron/util/runtime"
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

const postCollectionHandle string = "posts"

func (s *PostStore) Create(post *Post) (*mongo.InsertOneResult, error) {
	ok, _ := s.db.ListCollectionNames(context.TODO(), bson.M{"name": postCollectionHandle})
	if len(ok) == 0 {
		dbErr := errors.NewCollectionNotFound(runtime.Trace(), "")
		apiErr := errors.NewInternalError(dbErr, "")
		return nil, apiErr
	}
	postCollection := s.db.Collection(postCollectionHandle)
	postResult, err := postCollection.InsertOne(context.TODO(), post)
	if err != nil {
		dbErr := errors.NewUnknownCause(err, runtime.Trace())
		apiErr := errors.NewUnknown(dbErr, "")
		return nil, apiErr
	}

	return postResult, nil
}

func (s *PostStore) Get(id string) (*Post, error) {
	ok, _ := s.db.ListCollectionNames(context.TODO(), bson.M{"name": postCollectionHandle})
	if len(ok) == 0 {
		dbErr := errors.NewCollectionNotFound(runtime.Trace(), "")
		apiErr := errors.NewInternalError(dbErr, "")
		return nil, apiErr
	}
	postCollection := s.db.Collection(postCollectionHandle)

	post := &Post{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		dbErr := errors.NewInvalidID("")
		apiErr := errors.NewInvalidData(dbErr, "")
		return nil, apiErr
	}
	filter := bson.M{"_id": objectID}

	err = postCollection.FindOne(context.TODO(), filter).Decode(post)
	if err != nil {
		dbErr := errors.NewIDNotFound("")
		apiErr := errors.NewInvalidData(dbErr, "")
		return nil, apiErr
	}

	return post, nil
}

func (s *PostStore) Delete(id string) (*mongo.DeleteResult, error) {
	ok, _ := s.db.ListCollectionNames(context.TODO(), bson.M{"name": postCollectionHandle})
	if len(ok) == 0 {
		dbErr := errors.NewCollectionNotFound(runtime.Trace(), "")
		apiErr := errors.NewInternalError(dbErr, "")
		return nil, apiErr
	}
	postCollection := s.db.Collection(postCollectionHandle)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		dbErr := errors.NewInvalidID("")
		apiErr := errors.NewInvalidData(dbErr, "")
		return nil, apiErr
	}
	filter := bson.M{"_id": objectID}

	result, err := postCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		dbErr := errors.NewIDNotFound("")
		apiErr := errors.NewInvalidData(dbErr, "")
		return nil, apiErr
	}

	return result, nil
}

func (s *PostStore) ListAll() (*[]Post, error) {
	ok, _ := s.db.ListCollectionNames(context.TODO(), bson.M{"name": postCollectionHandle})
	if len(ok) == 0 {
		dbErr := errors.NewCollectionNotFound(runtime.Trace(), "")
		apiErr := errors.NewInternalError(dbErr, "")
		return nil, apiErr
	}
	postCollection := s.db.Collection(postCollectionHandle)

	cursor, err := postCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		dbErr := errors.NewUnknownCause(err, runtime.Trace())
		apiErr := errors.NewUnknown(dbErr, "")
		return nil, apiErr
	}
	defer cursor.Close(context.TODO())

	var posts []Post
	for cursor.Next(context.TODO()) {
		post := Post{}
		err = cursor.Decode(&post)
		if err != nil {
			dbErr := errors.NewUnknownCause(err, runtime.Trace())
			apiErr := errors.NewUnknown(dbErr, "")
			return nil, apiErr
		}
		posts = append(posts, post)
	}

	return &posts, nil
}
