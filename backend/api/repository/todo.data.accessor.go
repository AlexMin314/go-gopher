package repository

import "github.com/AlexMin314/go-gopher/backend/api/schema"

type ID string

type DataAccess interface {
	GetTask(id ID) (schema.Task, error)
	PutTask(id ID, t schema.Task) error
	PostTask(t schema.Task) (ID, error)
	DeleteTask(id ID) error
}
