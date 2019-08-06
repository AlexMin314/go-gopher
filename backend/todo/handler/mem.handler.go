package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlexMin314/go-gopher/backend/todo/constant"
	"github.com/AlexMin314/go-gopher/backend/todo/repository"
	"github.com/AlexMin314/go-gopher/backend/todo/schema"
	"github.com/AlexMin314/go-gopher/backend/todo/service"
)

var memDB = repository.NewMemoryDataAccess()

func GetTodoMem(w http.ResponseWriter, r *http.Request) {
	id := service.GetTodoIdParam(r)
	todo, err := memDB.GetTodo(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(schema.Response{
		Status: constant.Success,
		Data: schema.Data{
			ID:      id,
			Title:   todo.Title,
			Checked: todo.Checked,
		},
	})

	if err != nil {
		http.Error(w, constant.InternalServerError, http.StatusInternalServerError)
	}
}

func GetAllTodoMem(w http.ResponseWriter, r *http.Request) {
	todos, err := memDB.GetAllTodo()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(schema.Response{
		Status: constant.Success,
		Data: schema.Data{
			Todo: todos,
		},
	})

	if err != nil {
		http.Error(w, constant.InternalServerError, http.StatusInternalServerError)
	}
}

func PostTodoMem(w http.ResponseWriter, r *http.Request) {
	todos, err := service.ParseTodoPayload(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, todo := range todos {
		id, _ := memDB.PostTodo(todo)
		err = json.NewEncoder(w).Encode(schema.Response{
			Status: constant.Success,
			Data: schema.Data{
				ID: id,
			},
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func PutTodoMem(w http.ResponseWriter, r *http.Request) {
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
		Status: constant.Success,
		Data: schema.Data{
			ID: id,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteTodoMem(w http.ResponseWriter, r *http.Request) {
	id := service.GetTodoIdParam(r)
	err := memDB.DeleteTodo(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(schema.Response{
		Status: constant.Success,
		Data: schema.Data{
			ID: id,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
