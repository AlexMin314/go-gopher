package todo

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AlexMin314/go-gopher/backend/config"
	"github.com/AlexMin314/go-gopher/backend/logger"
	"github.com/AlexMin314/go-gopher/backend/todo/router"
)

func InitTodoService(lg logger.ReqLogger) {
	r := router.RegisterRoutes()
	http.Handle("/", r)
	c := config.InitConfig()
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(c.Server.Port), lg(r)))
}
