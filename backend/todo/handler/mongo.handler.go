package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlexMin314/go-gopher/backend/db/mongodb"

	"github.com/AlexMin314/go-gopher/backend/todo/constant"
	"github.com/AlexMin314/go-gopher/backend/todo/repository"
	"github.com/AlexMin314/go-gopher/backend/todo/schema"
	"github.com/AlexMin314/go-gopher/backend/todo/service"
)

func GetTodo(m *mongodb.Mongo) func(http.ResponseWriter, *http.Request) {
	dao := repository.NewTodoMongoAccess(m)
	return func(w http.ResponseWriter, r *http.Request) {
		id := service.GetTodoIdParam(r)
		todo, err := dao.GetTodo(id)

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
}

func GetAllTodo(m *mongodb.Mongo) func(http.ResponseWriter, *http.Request) {
	dao := repository.NewTodoMongoAccess(m)
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := dao.GetAllTodo()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(schema.Response{
			Status: constant.Success,
			Data: schema.Data{
				Todos: todos,
			},
		})

		if err != nil {
			http.Error(w, constant.InternalServerError, http.StatusInternalServerError)
		}
	}
}

func PostTodo(m *mongodb.Mongo) func(http.ResponseWriter, *http.Request) {
	dao := repository.NewTodoMongoAccess(m)
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := service.ParseTodoPayload(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, postErr := dao.PostTodo(todos)
		log.Println(result, postErr)

		if postErr != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(schema.Response{
			Status: constant.Success,
			Data: schema.Data{
				IDs: result,
			},
		})

		if err != nil {
			http.Error(w, constant.InternalServerError, http.StatusInternalServerError)
		}
	}
}

// func PutTodoMemController(w http.ResponseWriter, r *http.Request) {
// 	id := service.GetTodoIdParam(r)
// 	todos, err := service.ParseTodoPayload(r)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	err = memDB.PutTodo(id, todos[0])
// 	if err != nil {
// 		http.Error(w, constant.InternalServerError, http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(schema.Response{
// 		ID:   id,
// 		Todo: todos[0],
// 	})

// 	if err != nil {
// 		panic(err)
// 	}
// }

// func DeleteTodoMemController(w http.ResponseWriter, r *http.Request) {
// 	id := service.GetTodoIdParam(r)
// 	err := memDB.DeleteTodo(id)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(schema.Response{
// 		ID: id,
// 	})

// 	if err != nil {
// 		panic(err)
// 	}
// }

func DeleteAllTodo(m *mongodb.Mongo) func(http.ResponseWriter, *http.Request) {
	dao := repository.NewTodoMongoAccess(m)
	return func(w http.ResponseWriter, r *http.Request) {
		err := dao.DeleteAllTodo()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(schema.Response{
			Status: constant.Success,
		})

		if err != nil {
			http.Error(w, constant.InternalServerError, http.StatusInternalServerError)
		}
	}
}
