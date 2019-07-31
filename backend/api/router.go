package api

import (
	"net/http"

	"github.com/AlexMin314/go-gopher/backend/api/constant"
	"github.com/AlexMin314/go-gopher/backend/api/handler"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	todoSubRouter := r.PathPrefix(constant.TodoApiRoute).Subrouter()
	todoSubRouter.HandleFunc(constant.TodoIdPattern, handler.GetTodoHandler).Methods(http.MethodGet)
	todoSubRouter.HandleFunc("", handler.PostTodoHandler).Methods(http.MethodPost)
	todoSubRouter.HandleFunc(constant.TodoIdPattern, handler.PutTodoHandler).Methods(http.MethodPut)
	todoSubRouter.HandleFunc(constant.TodoIdPattern, handler.DeleteTodoHandler).Methods(http.MethodDelete)
	return r
}
