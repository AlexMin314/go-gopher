package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AlexMin314/go-gopher/backend/todo/constant"
	"github.com/AlexMin314/go-gopher/backend/todo/repository"
	"github.com/AlexMin314/go-gopher/backend/todo/schema"
	"github.com/AlexMin314/go-gopher/backend/todo/service"
)

var memDB = repository.NewMemoryDataAccess()

func GetTodoMemController(w http.ResponseWriter, r *http.Request) {
	id := service.GetTodoIdParam(r)
	todo, err := memDB.GetTodo(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(schema.Response{
		ID:   id,
		Todo: todo,
	})

	if err != nil {
		http.Error(w, constant.InternalServerError, http.StatusInternalServerError)
	}
}

func PostTodoMemController(w http.ResponseWriter, r *http.Request) {
	todos, err := service.ParseTodoPayload(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, todo := range todos {
		id, _ := memDB.PostTodo(todo)
		err = json.NewEncoder(w).Encode(schema.Response{
			ID:   id,
			Todo: todo,
		})
		if err != nil {
			panic(err)
		}
	}
}

func PutTodoMemController(w http.ResponseWriter, r *http.Request) {
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

	err = json.NewEncoder(w).Encode(schema.Response{
		ID:   id,
		Todo: todos[0],
	})

	if err != nil {
		panic(err)
	}
}

func DeleteTodoMemController(w http.ResponseWriter, r *http.Request) {
	id := service.GetTodoIdParam(r)
	err := memDB.DeleteTodo(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(schema.Response{
		ID: id,
	})

	if err != nil {
		panic(err)
	}
}
