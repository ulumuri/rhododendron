package DB

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	devMode string = "dev"
	devUri  string = "mongodb://root:dev@mongo:27017/?connect=direct"
)

func ConnectToDB(mode string) (*mongo.Database, error) {
	const databaseName string = "api_test"
	var uri string

	switch mode {
	case devMode:
		uri = devUri
	default:
		var err error
		uri, err = getDotEnvVariable("DB_URI")
		if err != nil {
			return nil, err
		}
	}

	client, err := getMongoClient(context.TODO(), uri)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	return client.Database(databaseName), nil

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
