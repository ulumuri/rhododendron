package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/ulumuri/rhododendron/models"
)

type Env struct {
	db models.Requester
}

func (env *Env) InsertPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	post := &models.Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = env.db.Insert(post)
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJson(w, http.StatusOK, post)
}

func (env *Env) FindPostByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	post, err := env.db.FindByID(ps.ByName("id"))
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err.Error())
	}

	respondWithJson(w, http.StatusOK, post)
}

func (env *Env) DeletePostByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	_, err := env.db.DeleteByID(ps.ByName("id"))
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err.Error())
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "succes"})
}

func (env *Env) DefineEndPoints() {
	router := httprouter.New()
	router.POST("/posts", env.InsertPost)
	router.GET("/posts", env.FindPostByID)
	router.DELETE("/posts", env.DeletePostByID)

	log.Fatal(http.ListenAndServe(":8080", router))
}
