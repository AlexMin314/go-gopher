package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AlexMin314/go-gopher/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	Session *mongo.Client
	DB      *mongo.Database
}

func Connect() *Mongo {
	c := config.InitConfig().DB.Mongo

	client, err := mongo.NewClient(options.Client().ApplyURI(c.ConnectionUri))
	ctx, cancel := context.WithTimeout(context.Background(), c.ConnectionTimeout*time.Second)
	err = client.Connect(ctx)
	m := &Mongo{}

	if err != nil {
		log.Fatal(err)
		return m
	}

	ctx, cancel = context.WithTimeout(context.Background(), c.PingTimeout*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
		return m
	}

	defer cancel()

	log.Println("Connected to MongoDB!")
	m.Session = client
	m.DB = client.Database(c.DbName)

	recoverTodo(m)

	return m
}

func Close(m *Mongo) {
	err := m.Session.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
