package schema

import (
	"github.com/AlexMin314/go-gopher/backend/todo/constant"
)

type Todo struct {
	// ID      string           `json:"id" bson:"_id"`
	Title   string           `json:"title,omitempty" bson:"title"`
	Checked constant.Checker `json:"checked,omitempty bson:"checked"`
}
