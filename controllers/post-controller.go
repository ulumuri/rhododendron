package controllers

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/julienschmidt/httprouter"

	"github.com/ulumuri/rhododendron/models"
)

type Env struct {
	DB models.Requester
}

func (env *Env) InsertPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	post := &models.Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = env.DB.Insert(post)
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJson(w, http.StatusOK, post)
}

func (env *Env) FindPostByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	id, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err.Error())
	}

	post, err := env.DB.FindByID(id)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err.Error())
	}

	respondWithJson(w, http.StatusOK, post)
}

func (env *Env) DeletePostByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	_, err := env.DB.DeleteByID(ps.ByName("id"))
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err.Error())
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "succes"})
}
