package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ulumuri/rhododendron/api"
	"github.com/ulumuri/rhododendron/database"
)

func main() {
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	newAPI := api.NewAPI(db)
	router := httprouter.New()
	router.POST("/posts/create", newAPI.Post.Create)
	router.GET("/posts/get/:id", newAPI.Post.Get)
	router.DELETE("/posts/delete/:id", newAPI.Post.Delete)
	router.GET("/posts/list_all", newAPI.Post.ListAll)

	log.Fatal(http.ListenAndServe(":8080", router))
}
