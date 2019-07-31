package schema

import "github.com/AlexMin314/go-gopher/backend/api/constant"

type Todo struct {
	Title   string           `json:"title,omitempty"`
	Checked constant.Checker `json:"checked,omitempty"`
}
