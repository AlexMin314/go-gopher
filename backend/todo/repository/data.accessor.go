package repository

import (
	"github.com/AlexMin314/go-gopher/backend/todo/constant"
	"github.com/AlexMin314/go-gopher/backend/todo/schema"
)

type DataAccess interface {
	Find(id schema.ID) (schema.Todo, error)
	FindAll() ([]*schema.Todo, error)
	UpdateOne(t schema.Todo) (int64, error)
	ToggleMany(ids []string, t constant.Checker) (int64, error)
	Save(t []schema.Todo) ([]interface{}, error)
	Delete(id schema.ID) error
	DeleteAll() error
}

// keep this as learning history
type DataMemAccess interface {
	GetTodo(id schema.ID) (schema.Todo, error)
	GetAllTodo() ([]schema.Todo, error)
	PutTodo(id schema.ID, t schema.Todo) error
	PostTodo(t schema.Todo) (schema.ID, error)
	DeleteTodo(id schema.ID) error
}
