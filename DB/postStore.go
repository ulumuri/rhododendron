package DB

import (
	"context"

	"github.com/ulumuri/rhododendron/status"
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
	//ok, _ := s.db.ListCollectionNames(context.TODO(), bson.M{"name": postCollectionHandle})
	//if len(ok) == 0 {
	//	return nil, status.NewCollectionNotFound(postCollectionHandle)
	//}

	postCollection := s.db.Collection(postCollectionHandle)
	postResult, err := postCollection.InsertOne(context.TODO(), post)
	if err != nil {
		return nil, status.NewUnknown("TODO", err)
	}

	return postResult, nil
}

func (s *PostStore) Get(id string) (*Post, error) {
	//ok, _ := s.db.ListCollectionNames(context.TODO(), bson.M{"name": postCollectionHandle})
	//if len(ok) == 0 {
	//	return nil, status.NewCollectionNotFound(postCollectionHandle)
	//}
	postCollection := s.db.Collection(postCollectionHandle)

	post := &Post{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.NewInvalidID("", err)
	}

	//opts := options.FindOne().SetSort(bson.D{{"author", getAuthor()}})
	filter := bson.D{{"_id", objectID}}
	err = postCollection.FindOne(context.TODO(), filter).Decode(post)
	if err != nil {
		return nil, status.NewIDNotFound("", err)
	}

	return post, nil
}

func (s *PostStore) Delete(id string) (*Post, error) {
	ok, _ := s.db.ListCollectionNames(context.TODO(), bson.M{"name": postCollectionHandle})
	if len(ok) == 0 {
		return nil, status.NewCollectionNotFound(postCollectionHandle)
	}

	postCollection := s.db.Collection(postCollectionHandle)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.NewInvalidID("", err)
	}

	post := &Post{}
	//opts := options.FindOne().SetSort(bson.D{{"author", getAuthor()}})
	filter := bson.D{{"_id", objectID}}
	err = postCollection.FindOneAndDelete(context.TODO(), filter).Decode(post)
	if err != nil {
		return nil, status.NewIDNotFound("", err)
	}

	return post, nil
}

func (s *PostStore) ListAll() (*[]Post, error) {
	ok, _ := s.db.ListCollectionNames(context.TODO(), bson.M{"name": postCollectionHandle})
	if len(ok) == 0 {
		return nil, status.NewCollectionNotFound(postCollectionHandle)
	}

	postCollection := s.db.Collection(postCollectionHandle)
	cursor, err := postCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, status.NewUnknown("", err)
	}
	defer cursor.Close(context.TODO())

	var posts []Post
	for cursor.Next(context.TODO()) {
		post := Post{}
		err = cursor.Decode(&post)
		if err != nil {
			return nil, status.NewUnknown("", err)
		}
		posts = append(posts, post)
	}

	return &posts, nil
}
