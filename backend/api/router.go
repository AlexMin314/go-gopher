package api

import (
	"github.com/AlexMin314/go-gopher/backend/api/handler"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	todoSubRouter := r.PathPrefix(todoApiRoute).Subrouter()
	todoSubRouter.HandleFunc(todoIdPattern, handler.TodoApiHandler).Methods("GET")
	todoSubRouter.HandleFunc("", handler.TodoApiHandler).Methods("POST")
	todoSubRouter.HandleFunc(todoIdPattern, handler.TodoApiHandler).Methods("PUT")
	todoSubRouter.HandleFunc(todoIdPattern, handler.TodoApiHandler).Methods("PUT")
	return r
}
