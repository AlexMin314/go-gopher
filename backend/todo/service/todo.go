package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/AlexMin314/go-gopher/backend/todo/schema"
)

func GetTodoIdParam(r *http.Request) schema.ID {
	return schema.ID(mux.Vars(r)["id"])
}

type TodoPostBody struct {
	Todo []schema.Todo
}

func ParseTodoPayload(r *http.Request) ([]schema.Todo, error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var todos TodoPostBody
	if err = json.Unmarshal([]byte(body), &todos); err != nil {
		return nil, err
	}

	return todos.Todo, nil
}
