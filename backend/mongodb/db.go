package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/AlexMin314/go-gopher/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongo() {
	c := config.InitConfig()

	client, err := mongo.NewClient(options.Client().ApplyURI(c.DB.Mongo.ConnectionUri))
	ctx, cancel := context.WithTimeout(context.Background(), c.DB.Mongo.ConnectionTimeout*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), c.DB.Mongo.PingTimeout*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
		return
	}

	defer cancel()
	log.Println("Connected to MongoDB!")
}
