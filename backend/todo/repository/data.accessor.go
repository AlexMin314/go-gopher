package repository

import "github.com/AlexMin314/go-gopher/backend/todo/schema"

type DataAccess interface {
	GetTodo(id schema.ID) (schema.Todo, error)
	GetAllTodo() ([]*schema.Todo, error)
	PutTodo(id schema.ID, t schema.Todo) error
	PostTodo(t []schema.Todo) ([]interface{}, error)
	DeleteTodo(id schema.ID) error
	DeleteAllTodo() error
}

type DataMemAccess interface {
	GetTodo(id schema.ID) (schema.Todo, error)
	GetAllTodo() ([]schema.Todo, error)
	PutTodo(id schema.ID, t schema.Todo) error
	PostTodo(t schema.Todo) (schema.ID, error)
	DeleteTodo(id schema.ID) error
}
