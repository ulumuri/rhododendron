package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/julienschmidt/httprouter"
	"github.com/ulumuri/rhododendron/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostStore interface {
	Create(*database.Post) (*mongo.InsertOneResult, error)
	Get(primitive.ObjectID) (*database.Post, error)
	Delete(primitive.ObjectID) (*mongo.DeleteResult, error)
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
		render.Respond(w, r, err.Error())
		return
	}
	_, err = rs.Store.Create(post)
	if err != nil {
		render.Respond(w, r, err.Error())
		return
	}

	render.Respond(w, r, post)
}

func (rs *PostResource) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		render.Respond(w, r, err.Error())
		return
	}

	post, err := rs.Store.Get(id)
	if err != nil {
		render.Respond(w, r, err.Error())
		return
	}

	render.Respond(w, r, post)
}

func (rs *PostResource) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		render.Respond(w, r, err.Error())
		return
	}

	_, err = rs.Store.Delete(id)
	if err != nil {
		render.Respond(w, r, err.Error())
		return
	}

	render.Respond(w, r, map[string]string{"result": "success"})
}

func (rs *PostResource) ListAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	posts, err := rs.Store.ListAll()
	if err != nil {
		render.Respond(w, r, err.Error())
		return
	}

	render.Respond(w, r, posts)
}
