package main

import (
	"github.com/AlexMin314/go-gopher/backend/logger"
	"github.com/AlexMin314/go-gopher/backend/todo"
)

func composeService() {
	todo.InitTodoService(logger.RequestLogger)
}

func main() {
	composeService()
}
