package repository

import (
	"errors"
	"fmt"

	"github.com/AlexMin314/go-gopher/backend/api/schema"
)

// for simaple in-memory case
type MemoryDataAccess struct {
	tasks  map[ID]schema.Task
	nextID int64
}

func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		tasks: map[ID]schema.Task{
			"1": schema.Task{},
			"2": schema.Task{},
		},
		nextID: int64(3),
	}
}

var ErrTaskNotExist = errors.New("task does not exist")

func (m *MemoryDataAccess) GetTask(id ID) (schema.Task, error) {
	task, prs := m.tasks[id]
	if !prs {
		return schema.Task{}, ErrTaskNotExist
	}
	return task, nil
}

func (m *MemoryDataAccess) PostTask(t schema.Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

func (m *MemoryDataAccess) PutTask(id ID, t schema.Task) error {
	if _, prs := m.tasks[id]; !prs {
		return ErrTaskNotExist
	}
	m.tasks[id] = t
	return nil
}

func (m *MemoryDataAccess) DeleteTask(id ID) error {
	if _, prs := m.tasks[id]; !prs {
		return ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}
