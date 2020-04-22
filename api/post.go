package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ulumuri/rhododendron/database"
	"github.com/ulumuri/rhododendron/errors"
	"github.com/ulumuri/rhododendron/util/api"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostStore interface {
	Create(*database.Post) (*mongo.InsertOneResult, error)
	Get(string) (*database.Post, error)
	Delete(string) (*database.Post, error)
	ListAll() (*[]database.Post, error)
}

type PostResource struct {
	Store PostStore
}

func NewPostResource(store PostStore) *PostResource {
	return &PostResource{
		Store: store,
	}
}

func (rs *PostResource) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	post := &database.Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		api.RespondWithJsonStatus(w, errors.NewBadRequest("TODO", err))
		return
	}

	if len(post.Message) > errors.PostMaxSize {
		api.RespondWithJsonStatus(w, errors.NewPostMaxSizeExceeded())
		return
	}

	_, err = rs.Store.Create(post)
	if err != nil {
		api.RespondWithJsonStatus(w, err.(*errors.StatusError))
		return
	}

	api.RespondWithJson(w, post)
}

func (rs *PostResource) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	post, err := rs.Store.Get(ps.ByName("id"))
	if err != nil {
		api.RespondWithJsonStatus(w, err.(*errors.StatusError))
		return
	}

	api.RespondWithJson(w, post)
}

func (rs *PostResource) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := rs.Store.Delete(ps.ByName("id"))
	if err != nil {
		api.RespondWithJsonStatus(w, err.(*errors.StatusError))
		return
	}

	api.RespondWithJsonStatus(w, errors.NewSuccess("The post has been removed."))
}

func (rs *PostResource) ListAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	posts, err := rs.Store.ListAll()
	if err != nil {
		api.RespondWithJsonStatus(w, err.(*errors.StatusError))
		return
	}

	api.RespondWithJson(w, posts)
}
