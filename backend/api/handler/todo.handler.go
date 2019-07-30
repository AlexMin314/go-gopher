package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AlexMin314/go-gopher/backend/api/constant"
	"github.com/AlexMin314/go-gopher/backend/api/repository"
	"github.com/AlexMin314/go-gopher/backend/api/schema"
	"github.com/AlexMin314/go-gopher/backend/api/service"
)

type ResponseError struct {
	Code     int
	ErrorMsg string
}

type Response struct {
	ID    repository.ID `json:"id,omitempty"`
	Todo  schema.Todo   `json:"todo"`
	Error ResponseError `json:"error"`
}

var memDB = repository.NewMemoryDataAccess()

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := service.GetTodoId(r)
	todo, err := memDB.GetTodo(id)
	// if err != nil {
	// 	http.Error(w, err.Error(), 404)
	// 	return
	// }

	encodeErr := json.NewEncoder(w).Encode(Response{
		ID:   id,
		Todo: todo,
		// Error: ResponseError{err},
		Error: ResponseError{404, err.Error()},
	})

	if encodeErr != nil {
		http.Error(w, constant.InternalServerError, constant.InternalServerErrorCode)
	}
}

func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	//
}

func PutTodoHandler(w http.ResponseWriter, r *http.Request) {
	//
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	//
}
