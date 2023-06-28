package db

import (
	"context"
	"fmt"
	"log"

	"github.com/vier21/go-book-api/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

func NewConnection() *Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.GetConfig().MongoDBURL))

	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		Client: client,
	}
}

func (db *Database) Disconnect() {
	if err := db.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func (db *Database) Ping() {
	var result bson.M
	if err := db.Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		fmt.Println("DB not connected")
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
