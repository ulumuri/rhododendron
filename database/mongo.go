package database

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() (*mongo.Database, error) {
	dbName := "api_test"
	uri, err := getDotEnvVariable("DB_URI")
	if err != nil {
		return nil, err
	}
	client, err := getMongoClient(context.Background(), uri)
	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}

func getMongoClient(ctx context.Context, uri string) (*mongo.Client, error) {
	opts := options.Client()
	opts.ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getDotEnvVariable(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}
