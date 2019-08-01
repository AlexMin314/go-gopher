package router

import (
	"net/http"

	"github.com/AlexMin314/go-gopher/backend/todo/constant"
	"github.com/AlexMin314/go-gopher/backend/todo/controller"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	todoSubRouter := r.PathPrefix(constant.TodoApiRoute).Subrouter()
	todoSubRouter.HandleFunc(constant.TodoIdPattern, controller.GetTodoController).Methods(http.MethodGet)
	todoSubRouter.HandleFunc("/", controller.PostTodoController).Methods(http.MethodPost)
	todoSubRouter.HandleFunc(constant.TodoIdPattern, controller.PutTodoController).Methods(http.MethodPut)
	todoSubRouter.HandleFunc(constant.TodoIdPattern, controller.DeleteTodoController).Methods(http.MethodDelete)
	return r
}
