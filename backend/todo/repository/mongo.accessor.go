package repository

import (
	"context"
	"log"

	"github.com/AlexMin314/go-gopher/backend/db/mongodb"
	"github.com/AlexMin314/go-gopher/backend/todo/constant"
	"github.com/AlexMin314/go-gopher/backend/todo/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Session *mongo.Client
	DB      *mongo.Database
}

func NewTodoMongoAccess(d *mongodb.Mongo) DataAccess {
	return &MongoRepository{
		Session: d.Session,
		DB:      d.DB,
	}
}

func (m *MongoRepository) GetTodo(id schema.ID) (schema.Todo, error) {
	collection := m.DB.Collection("todos")

	objID, _ := primitive.ObjectIDFromHex(string(id))
	// objID, _ := primitive.ObjectIDFromHex("5d47d733c66645dbe9b37db3")
	filter := bson.D{{"_id", objID}}
	var result schema.Todo
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	log.Println("---------------------")
	log.Println(err, result)
	log.Println("---------------------")
	if err == nil {
		return result, nil
	}

	//
	return schema.Todo{
		Title:   "hello",
		Checked: constant.Yes,
	}, nil
}
func (m *MongoRepository) PostTodo(t schema.Todo) (schema.ID, error) {
	//
	return schema.ID("10"), nil
}
func (m *MongoRepository) PutTodo(id schema.ID, t schema.Todo) error {
	//
	return nil
}
func (m *MongoRepository) DeleteTodo(id schema.ID) error {
	//
	return nil
}
