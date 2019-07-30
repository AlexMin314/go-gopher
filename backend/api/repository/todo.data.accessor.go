package repository

import "github.com/AlexMin314/go-gopher/backend/api/schema"

type ID string

type DataAccess interface {
	GetTodo(id ID) (schema.Todo, error)
	PutTodo(id ID, t schema.Todo) error
	PostTodo(t schema.Todo) (ID, error)
	DeleteTodo(id ID) error
}
