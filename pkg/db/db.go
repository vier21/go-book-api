package db

import (
	"context"
	"log"

	"github.com/vier21/go-book-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Client

func NewConnection() *mongo.Client {

	url := config.GetConfig().MongoDBURL

	if url == "" {
		log.Fatal("You have to set MONGODB_URI to environment variable")
	}

	var err error
	DB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	return DB
}

func Disconnect() {
	if err := DB.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
