package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		return nil
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	err = client.Connect(ctx)
	if err != nil {
		return nil
	}
	fmt.Println("Connected to Mongodb")

	return client
}

var Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection = client.Database("cluster0").Collection(collectionName)
	return collection
}
