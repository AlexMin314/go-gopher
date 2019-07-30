package service

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/AlexMin314/go-gopher/backend/api/repository"
)

func GetTodoId(r *http.Request) repository.ID {
	return repository.ID(mux.Vars(r)["id"])
}
