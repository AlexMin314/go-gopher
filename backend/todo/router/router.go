package router

import (
	"net/http"

	"github.com/AlexMin314/go-gopher/backend/db/mongodb"

	"github.com/AlexMin314/go-gopher/backend/todo/constant"
	"github.com/AlexMin314/go-gopher/backend/todo/handler"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, m *mongodb.Mongo) *mux.Router {
	todoSubRouter := r.PathPrefix(constant.TodoApiRoute).Subrouter()

	// with mem ds (leave this as learning history)
	todoSubRouter.HandleFunc(constant.Mem, handler.GetAllTodoMem).Methods(http.MethodGet)
	todoSubRouter.HandleFunc(constant.TodoMemIdPattern, handler.GetTodoMem).Methods(http.MethodGet)
	todoSubRouter.HandleFunc(constant.Mem, handler.PostTodoMem).Methods(http.MethodPost)
	todoSubRouter.HandleFunc(constant.TodoMemIdPattern, handler.PutTodoMem).Methods(http.MethodPut)
	todoSubRouter.HandleFunc(constant.TodoMemIdPattern, handler.DeleteTodoMem).Methods(http.MethodDelete)

	// with Mongo
	todoSubRouter.HandleFunc("", handler.GetAllTodo(m)).Methods(http.MethodGet)
	todoSubRouter.HandleFunc(constant.TodoIdPattern, handler.GetTodo(m)).Methods(http.MethodGet)
	// todoSubRouter.HandleFunc("/", handler.PostTodoController).Methods(http.MethodPost)
	// todoSubRouter.HandleFunc(constant.TodoIdPattern, handler.PutTodoController).Methods(http.MethodPut)
	// todoSubRouter.HandleFunc(constant.TodoIdPattern, handler.DeleteTodoController).Methods(http.MethodDelete)
	todoSubRouter.HandleFunc("", handler.DeleteAllTodo(m)).Methods(http.MethodDelete)
	return r
}
