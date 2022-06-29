package reference

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect Mongo DB
func Connect() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(GetMongoURL()))
	if err != nil {
		fmt.Println("Hata var", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Hata var", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Ping Var")
	}
	fmt.Printf("Connected Mongodb")

	return client
}

// export mongo db client
var DB *mongo.Client = Connect()

// Get Collection in mongodb
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("constani").Collection(collectionName)
	return collection
}
