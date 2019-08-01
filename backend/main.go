package main

import (
	"github.com/AlexMin314/go-gopher/backend/logger"
	"github.com/AlexMin314/go-gopher/backend/mongodb"
	"github.com/AlexMin314/go-gopher/backend/todo"
)

func composeService() {
	todo.InitTodoService(logger.RequestLogger)
}

func connectDB() {
	mongodb.ConnectMongo()
}

func main() {
	composeService()
	connectDB()
}
