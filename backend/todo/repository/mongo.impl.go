package repository

import (
	"context"
	"log"

	"github.com/AlexMin314/go-gopher/backend/db/mongodb"
	"github.com/AlexMin314/go-gopher/backend/todo/schema"
	"github.com/AlexMin314/go-gopher/backend/todo/service"
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
	filter := bson.D{{"_id", objID}}
	var result schema.Todo
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	return result, err
}

func (m *MongoRepository) GetAllTodo() ([]*schema.Todo, error) {
	cur, err := m.DB.Collection("todos").Find(
		context.TODO(),
		bson.D{{}},
		// options.Find().SetLimit(2),
	)

	var results []*schema.Todo

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var todo schema.Todo
		if err := cur.Decode(&todo); err != nil {
			log.Fatal(err)
		}
		results = append(results, &todo)
	}

	cur.Close(context.TODO())
	return results, err
}

func (m *MongoRepository) PostTodo(t []schema.Todo) ([]interface{}, error) {
	result, err := m.DB.Collection("todos").InsertMany(context.TODO(), service.CastTodoToInterface(t))
	if err != nil {
		log.Fatal(err)
	}
	return result.InsertedIDs, err
}

func (m *MongoRepository) PutTodo(id schema.ID, t schema.Todo) error {
	//
	return nil
}
func (m *MongoRepository) DeleteTodo(id schema.ID) error {
	//
	return nil
}

func (m *MongoRepository) DeleteAllTodo() error {
	return m.DB.Collection("todos").Drop(context.TODO())
}
