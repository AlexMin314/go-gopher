package todo

import (
	"github.com/AlexMin314/go-gopher/backend/todo/router"
	"github.com/gorilla/mux"
)

func InitTodoService(r *mux.Router) *mux.Router {
	r = router.RegisterRoutes(r)
	return r
}
