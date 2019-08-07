package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/AlexMin314/go-gopher/backend/db/mongodb"

	"github.com/AlexMin314/go-gopher/backend/todo/constant"
	"github.com/AlexMin314/go-gopher/backend/todo/repository"
	"github.com/AlexMin314/go-gopher/backend/todo/schema"
	"github.com/AlexMin314/go-gopher/backend/todo/service"
)

func GetTodo(m *mongodb.Mongo) func(http.ResponseWriter, *http.Request) {
	repo := repository.NewTodoMongoAccess(m)
	return func(w http.ResponseWriter, r *http.Request) {
		id := service.GetTodoIdParam(r)
		todo, err := repo.Find(id)

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
	repo := repository.NewTodoMongoAccess(m)
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := repo.FindAll()

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
	repo := repository.NewTodoMongoAccess(m)
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := service.ParseTodoPayload(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, postErr := repo.Save(payload.Todos)

		if postErr != nil {
			http.Error(w, postErr.Error(), http.StatusBadRequest)
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

func PutTodo(m *mongodb.Mongo) func(http.ResponseWriter, *http.Request) {
	repo := repository.NewTodoMongoAccess(m)
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := service.ParseTodoPayload(r)
		empty := schema.Payload{}
		var result int64

		if reflect.DeepEqual(empty, payload) {
			err = errors.New("Empty payload")
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(payload.IDs) != 0 && empty.ToggleTo != payload.ToggleTo {
			result, err = repo.ToggleMany(payload.IDs, payload.ToggleTo)
		} else {
			objId, _ := primitive.ObjectIDFromHex(payload.ID)
			result, err = repo.UpdateOne(schema.Todo{
				ID:      objId,
				Title:   payload.Title,
				Checked: payload.Checked,
			})
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(schema.Response{
			Status: constant.Success,
			Data: schema.Data{
				ModifiedCount: result,
			},
		})

		if err != nil {
			http.Error(w, constant.InternalServerError, http.StatusInternalServerError)
		}
	}
}

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
		err := dao.DeleteAll()

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
