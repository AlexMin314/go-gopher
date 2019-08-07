package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/AlexMin314/go-gopher/backend/db/mongodb"
	"github.com/AlexMin314/go-gopher/backend/todo/constant"
	"github.com/AlexMin314/go-gopher/backend/todo/schema"
	"github.com/AlexMin314/go-gopher/backend/todo/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	Session *mongo.Client
	DB      *mongo.Database
	// Validation config.Validation
}

func NewTodoMongoAccess(d *mongodb.Mongo) DataAccess {
	return &MongoRepository{
		Session: d.Session,
		DB:      d.DB,
	}
}

func (m *MongoRepository) Find(id schema.ID) (schema.Todo, error) {
	objID, _ := primitive.ObjectIDFromHex(string(id))
	var result schema.Todo
	err := m.DB.Collection("todos").FindOne(
		context.TODO(),
		bson.D{{"_id", objID}},
		options.FindOne().SetProjection(bson.D{{"created_at", 0}, {"updated_at", 0}}),
	).Decode(&result)
	return result, err
}

func (m *MongoRepository) FindAll() ([]*schema.Todo, error) {
	cur, err := m.DB.Collection("todos").Find(
		context.TODO(),
		bson.D{{}},
		options.Find().SetProjection(bson.D{{"created_at", 0}, {"updated_at", 0}}).SetLimit(50),
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

func (m *MongoRepository) Save(t []schema.Todo) ([]interface{}, error) {
	collection := m.DB.Collection("todos")

	if len(t) > 5 {
		return []interface{}{}, errors.New(constant.ExceedMaxTodoInsertLimit)
	}

	if c, _ := collection.CountDocuments(context.TODO(), bson.D{{}}); int(c)+len(t) > 5 {
		return []interface{}{}, errors.New(constant.ExceedMaxTodoInsertLimit)
	}

	result, err := collection.InsertMany(context.TODO(), service.CastTodoToInterface(t))
	if err != nil {
		log.Fatal(err)
	}
	return result.InsertedIDs, err
}

func (m *MongoRepository) UpdateOne(t schema.Todo) (int64, error) {
	result, err := m.DB.Collection("todos").UpdateOne(
		context.TODO(),
		bson.D{{"_id", t.ID}},
		bson.D{{"$set", bson.D{
			{"title", t.Title},
			{"checked", t.Checked},
			{"updated_at", time.Now()},
		}}},
	)
	return result.ModifiedCount, err
}

func (m *MongoRepository) ToggleMany(ids []string, toggleTo constant.Checker) (int64, error) {
	result, err := m.DB.Collection("todos").UpdateMany(
		context.TODO(),
		bson.D{{"_id", bson.D{{"$in", service.CastStringToObjId(ids)}}}},
		bson.D{{"$set", bson.D{
			{"checked", toggleTo},
			{"updated_at", time.Now()},
		}}},
	)
	return result.ModifiedCount, err
}

func (m *MongoRepository) Delete(id schema.ID) error {
	//
	return nil
}

func (m *MongoRepository) DeleteAll() error {
	return m.DB.Collection("todos").Drop(context.TODO())
}
