package schema

import "github.com/AlexMin314/go-gopher/backend/todo/constant"

type Todo struct {
	Title   string           `json:"title,omitempty"`
	Checked constant.Checker `json:"checked,omitempty"`
}
