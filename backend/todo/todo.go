package todo

import (
	"github.com/AlexMin314/go-gopher/backend/db/mongodb"
	"github.com/AlexMin314/go-gopher/backend/todo/router"
	"github.com/gorilla/mux"
)

func InitTodoService(r *mux.Router, m *mongodb.Mongo) *mux.Router {
	r = router.RegisterRoutes(r, m)
	return r
}
