package schema

import "github.com/AlexMin314/go-gopher/backend/todo/constant"

type ID string

type TodoResponse struct {
	ID    ID     `json:"id,omitempty"`
	Todo  Todo   `json:"todo,omitempty"`
	Todos []Todo `json:"todos,omitempty"`
}
type TodosResponse struct {
	Todos []*Todo `json:"todos,omitempty"`
}

type ResultResponse struct {
	Success constant.Checker `json:"success"`
}
