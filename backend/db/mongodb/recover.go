package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/AlexMin314/go-gopher/backend/todo/constant"
	"github.com/AlexMin314/go-gopher/backend/todo/schema"
)

func recover(m *Mongo) {
	collection := m.DB.Collection("todos")

	// init data
	todo1 := schema.Todo{"Learning go concurrency", constant.No}
	todo2 := schema.Todo{"Learning mongodb in depth", constant.No}
	todo3 := schema.Todo{"Learning MQTT and more", constant.No}

	insertManyResult, err := collection.InsertMany(context.TODO(), []interface{}{todo1, todo2, todo3})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}
