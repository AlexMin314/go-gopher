package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AlexMin314/go-gopher/backend/config"
	"github.com/AlexMin314/go-gopher/backend/db/mongodb"
	mw "github.com/AlexMin314/go-gopher/backend/middleware"
	"github.com/AlexMin314/go-gopher/backend/sensor"
	"github.com/AlexMin314/go-gopher/backend/todo"
	"github.com/gorilla/mux"
)

func composeService(m *mongodb.Mongo) *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r = todo.InitTodoService(r, m)
	r = sensor.InitSensoirService(r, m)
	return r
}

func initDB(c config.DatabaseConfig) *mongodb.Mongo {
	m := mongodb.Connect(c.Mongo)
	return m
}

func main() {
	c := config.GetConfig()
	m := initDB(c.DB)
	r := composeService(m)
	r.Use(mw.HttpLogger, mw.Response, mw.Cors)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(c.Server.Port), nil))
}
