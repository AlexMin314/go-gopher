package sensor

import (
	"github.com/AlexMin314/go-gopher/backend/db/mongodb"
	"github.com/gorilla/mux"
)

func InitSensoirService(r *mux.Router, m *mongodb.Mongo) *mux.Router {
	// r = router.RegisterRoutes(r, m)
	return r
}
