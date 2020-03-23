package api

import (
	"github.com/ulumuri/rhododendron/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type API struct {
	Post *PostResource
}

func NewAPI(db *mongo.Database) *API {
	postStore := database.NewPostStore(db)
	post := NewPostResource(postStore)

	return &API{
		Post: post,
	}
}
