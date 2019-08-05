package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AlexMin314/go-gopher/backend/config"
	"github.com/AlexMin314/go-gopher/backend/db/mongodb"
	"github.com/AlexMin314/go-gopher/backend/logger"
	"github.com/AlexMin314/go-gopher/backend/todo"
	"github.com/gorilla/mux"
)

func composeService(m *mongodb.Mongo) *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r = todo.InitTodoService(r, m)
	return r
}

func initDB() *mongodb.Mongo {
	m := mongodb.Connect()
	return m
}

func main() {
	m := initDB()
	r := composeService(m)
	http.Handle("/", r)
	c := config.InitConfig()
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(c.Server.Port), logger.RequestLogger(r)))
}
