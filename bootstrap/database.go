package bootstrap

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(env *Env) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbURI := env.MongoDBURI

	// if dbUser == "" || dbPass == "" {
	// 	mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	// }

	// client, err := mongo.NewClient(mongodbURI)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Fatal(mongodbURI, "dnjd")
	clientOptions := options.Client().ApplyURI(mongodbURI)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client *mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}