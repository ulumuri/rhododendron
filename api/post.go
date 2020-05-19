package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ulumuri/rhododendron/DB"
	"github.com/ulumuri/rhododendron/status"
	"github.com/ulumuri/rhododendron/util/api"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostStore interface {
	Create(*DB.Post) (*mongo.InsertOneResult, error)
	Get(string) (*DB.Post, error)
	Delete(string) (*DB.Post, error)
	ListAll() (*[]DB.Post, error)
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
	post := &DB.Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		api.RespondWithJsonStatus(w, status.NewBadRequest("TODO", err))
		return
	}

	if len(post.Message) > status.PostMaxSize {
		api.RespondWithJsonStatus(w, status.NewPostMaxSizeExceeded())
		return
	}

	_, err = rs.Store.Create(post)
	if err != nil {
		api.RespondWithJsonStatus(w, err)
		return
	}

	api.RespondWithJson(w, post)
}

func (rs *PostResource) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	post, err := rs.Store.Get(ps.ByName("id"))
	if err != nil {
		api.RespondWithJsonStatus(w, err)
		return
	}

	api.RespondWithJson(w, post)
}

func (rs *PostResource) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := rs.Store.Delete(ps.ByName("id"))
	if err != nil {
		api.RespondWithJsonStatus(w, err)
		return
	}

	api.RespondWithJsonStatus(w, status.NewSuccess("The post has been removed."))
}

func (rs *PostResource) ListAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	posts, err := rs.Store.ListAll()
	if err != nil {
		api.RespondWithJsonStatus(w, err)
		return
	}

	api.RespondWithJson(w, posts)
}
