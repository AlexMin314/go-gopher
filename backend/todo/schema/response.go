package schema

import "github.com/AlexMin314/go-gopher/backend/todo/constant"

type ID string
type Response struct {
	Data   Data            `json:"data,omitempty"`
	Status constant.Status `json:"status"`
}

type Data struct {
	ID  ID            `json:"id,omitempty"`
	IDs []interface{} `json:"ids,omitempty"`
	// IDs     []ID             `json:"ids,omitempty"`
	Title   string           `json:"title,omitempty"`
	Checked constant.Checker `json:"checked,omitempty"`
	Todo    []Todo           `json:"todo,omitempty"`
	Todos   []*Todo          `json:"todos,omitempty"`
}
