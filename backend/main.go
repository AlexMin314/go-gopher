package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AlexMin314/go-gopher/backend/config"
	"github.com/AlexMin314/go-gopher/backend/logger"
	"github.com/AlexMin314/go-gopher/backend/mongodb"
	"github.com/AlexMin314/go-gopher/backend/todo"
	"github.com/gorilla/mux"
)

func composeService() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r = todo.InitTodoService(r)
	return r
}

func connectDB() {
	mongodb.ConnectMongo()
}

func main() {
	r := composeService()
	connectDB()

	http.Handle("/", r)
	c := config.InitConfig()
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(c.Server.Port), logger.RequestLogger(r)))
}
