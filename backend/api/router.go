package api

import (
	"github.com/AlexMin314/go-gopher/backend/api/constant"
	"github.com/AlexMin314/go-gopher/backend/api/handler"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	todoSubRouter := r.PathPrefix(constant.TodoApiRoute).Subrouter()
	todoSubRouter.HandleFunc(constant.TodoIdPattern, handler.GetTodoHandler).Methods("GET")
	todoSubRouter.HandleFunc("", handler.PostTodoHandler).Methods("POST")
	todoSubRouter.HandleFunc(constant.TodoIdPattern, handler.PutTodoHandler).Methods("PUT")
	todoSubRouter.HandleFunc(constant.TodoIdPattern, handler.DeleteTodoHandler).Methods("PUT")
	return r
}
