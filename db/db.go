package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/movies"
const Db = "movies"
const CollName = "movies"

var MongoClient *mongo.Client

func ConnectDatabase() {
	clientOption := options.Client().ApplyURI(connectionString)
	//to establish a connection
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	MongoClient = client

}
