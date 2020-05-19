package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/ulumuri/rhododendron/DB"
	"github.com/ulumuri/rhododendron/api"
)

func main() {
	mode := os.Args[1]
	db, err := DB.ConnectToDB(mode)
	if err != nil {
		log.Fatal(err)
	}

	newAPI := api.NewAPI(db)
	router := httprouter.New()
	handler := cors.AllowAll().Handler(router)

	router.POST("/posts/create", newAPI.Post.Create)
	router.GET("/posts/get/:id", newAPI.Post.Get)
	router.DELETE("/posts/delete/:id", newAPI.Post.Delete)
	router.GET("/posts/list_all", newAPI.Post.ListAll)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
