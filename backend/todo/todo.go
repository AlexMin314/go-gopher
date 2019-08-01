package todo

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AlexMin314/go-gopher/backend/config"
	"github.com/AlexMin314/go-gopher/backend/logger"
	"github.com/AlexMin314/go-gopher/backend/todo/router"
)

func InitTodoService(logger logger.ReqLogger) {
	r := router.RegisterRoutes()
	http.Handle("/", r)
	c := config.InitServerConfig()
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(c.Port), logger(r)))
}
