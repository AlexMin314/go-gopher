package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/AlexMin314/go-gopher/backend/todo/schema"
)

func GetTodoIdParam(r *http.Request) schema.ID {
	return schema.ID(mux.Vars(r)["id"])
}

func ParseTodoPayload(r *http.Request) ([]schema.Todo, error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var payload schema.Request
	if err = json.Unmarshal([]byte(body), &payload); err != nil {
		return nil, err
	}

	return payload.Todo, nil
}

func CastTodoToInterface(todos []schema.Todo) []interface{} {
	s := make([]interface{}, len(todos))
	for i, todo := range todos {
		todo.CreatedAt = time.Now()
		s[i] = todo
	}
	return s
}
