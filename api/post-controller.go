package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ulumuri/rhododendron/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type APIServer struct {
	DB models.Requester
}

func (api *APIServer) InsertPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	post := &models.Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = api.DB.Insert(post)
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, post)
}

func (api *APIServer) FindPostByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err.Error())
		return
	}

	post, err := api.DB.FindByID(id)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, post)
}

func (api *APIServer) DeletePostByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = api.DB.DeleteByID(id)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
