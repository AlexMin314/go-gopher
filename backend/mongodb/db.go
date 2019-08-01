package mongodb

import (
	"context"
	"log"

	"github.com/AlexMin314/go-gopher/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() {
	c := config.InitConfig()
	clientOptions := options.Client().ApplyURI(c.DB.Mongo.ConnectionUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
}
