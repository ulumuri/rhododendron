package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/ulumuri/rhododendron/api"
	"github.com/ulumuri/rhododendron/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initiateMongoClient(ctx context.Context, uri string) *mongo.Client {
	var err error
	var client *mongo.Client
	opts := options.Client()
	opts.ApplyURI(uri)
	if client, err = mongo.Connect(ctx, opts); err != nil {
		fmt.Println(err.Error())
	}
	return client
}

func getDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := initiateMongoClient(ctx, getDotEnvVariable("DB_URI"))

	conn := models.Connect(client, "api_test")
	env := api.APIServer{DB: conn}
	router := httprouter.New()
	router.POST("/posts/insert", env.InsertPost)
	router.GET("/posts/get/:id", env.FindPostByID)
	router.DELETE("/posts/delete/:id", env.DeletePostByID)

	log.Fatal(http.ListenAndServe(":8080", router))
}
