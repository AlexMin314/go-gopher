package api

import (
	"github.com/AlexMin314/go-gopher/backend/api/handler"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(TodoRoute, handler.TodoApiHandler)
	return r
}
