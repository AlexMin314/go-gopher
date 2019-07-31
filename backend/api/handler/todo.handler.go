package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AlexMin314/go-gopher/backend/api/constant"
	"github.com/AlexMin314/go-gopher/backend/api/repository"
	"github.com/AlexMin314/go-gopher/backend/api/schema"
	"github.com/AlexMin314/go-gopher/backend/api/service"
)

type Response struct {
	ID   repository.ID `json:"id,omitempty"`
	Todo schema.Todo   `json:"todo,omitempty"`
}

var memDB = repository.NewMemoryDataAccess()

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := service.GetTodoIdParam(r)
	todo, err := memDB.GetTodo(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(Response{
		ID:   id,
		Todo: todo,
	})

	if err != nil {
		http.Error(w, constant.InternalServerError, http.StatusInternalServerError)
	}
}

func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := service.ParseTodoPayload(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, todo := range todos {
		id, _ := memDB.PostTodo(todo)
		err = json.NewEncoder(w).Encode(Response{
			ID:   id,
			Todo: todo,
		})
		if err != nil {
			panic(err)
		}
	}
}

func PutTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := service.GetTodoIdParam(r)
	todos, err := service.ParseTodoPayload(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = memDB.PutTodo(id, todos[0])
	if err != nil {
		http.Error(w, constant.InternalServerError, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(Response{
		ID:   id,
		Todo: todos[0],
	})

	if err != nil {
		panic(err)
	}
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := service.GetTodoIdParam(r)
	err := memDB.DeleteTodo(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(Response{
		ID: id,
	})

	if err != nil {
		panic(err)
	}
}
