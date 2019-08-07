package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"

	"github.com/AlexMin314/go-gopher/backend/todo/schema"
)

func GetTodoIdParam(r *http.Request) schema.ID {
	return schema.ID(mux.Vars(r)["id"])
}

func ParseTodoPayload(r *http.Request) (schema.Payload, error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return schema.Payload{}, err
	}

	var payload schema.Payload
	if err = json.Unmarshal([]byte(body), &payload); err != nil {
		return schema.Payload{}, err
	}

	return payload, nil
}

func CastTodoToInterface(todos []schema.Todo) []interface{} {
	s := make([]interface{}, len(todos))
	for i, todo := range todos {
		todo.ID = primitive.NewObjectID()
		todo.CreatedAt = time.Now()
		s[i] = todo
	}
	return s
}

func CastStringToObjId(ids []string) []primitive.ObjectID {
	s := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objID, _ := primitive.ObjectIDFromHex(id)
		s[i] = objID
	}
	return s
}
