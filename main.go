package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ulumuri/rhododendron/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitiateMongoClient(ctx context.Context) *mongo.Client {
	var err error
	var client *mongo.Client
	uri := ""
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(ctx, opts); err != nil {
		fmt.Println(err.Error())
	}
	return client
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := InitiateMongoClient(ctx)

	conn := models.Connect(client, ctx, "api_test")
	conn.DefineEndPoints()
}
