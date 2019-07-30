package api

import (
	"github.com/AlexMin314/go-gopher/backend/api/handler"
	"github.com/gorilla/mux"
)

const (
	todo = "todo"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(todo, handler.TodoApiHandler)
	return router
}
