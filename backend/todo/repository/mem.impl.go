package repository

import (
	"errors"
	"fmt"

	"github.com/AlexMin314/go-gopher/backend/todo/schema"
)

// for simaple in-memory case
type MemoryDataAccess struct {
	todos  map[schema.ID]schema.Todo
	nextID int64
}

func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		todos: map[schema.ID]schema.Todo{
			"1": schema.Todo{},
			"2": schema.Todo{},
		},
		nextID: int64(3),
	}
}

var ErrTaskNotExist = errors.New("task does not exist")

func (m *MemoryDataAccess) GetTodo(id schema.ID) (schema.Todo, error) {
	task, prs := m.todos[id]
	if !prs {
		return schema.Todo{}, ErrTaskNotExist
	}
	return task, nil
}

func (m *MemoryDataAccess) GetAllTodo() ([]schema.Todo, error) {
	var result []schema.Todo
	for _, todo := range m.todos {
		result = append(result, todo)
	}
	return result, nil
}

func (m *MemoryDataAccess) PostTodo(t schema.Todo) (schema.ID, error) {
	id := schema.ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.todos[id] = t
	return id, nil
}

func (m *MemoryDataAccess) PutTodo(id schema.ID, t schema.Todo) error {
	if _, prs := m.todos[id]; !prs {
		return ErrTaskNotExist
	}
	m.todos[id] = t
	return nil
}

func (m *MemoryDataAccess) DeleteTodo(id schema.ID) error {
	if _, prs := m.todos[id]; !prs {
		return ErrTaskNotExist
	}
	delete(m.todos, id)
	return nil
}
