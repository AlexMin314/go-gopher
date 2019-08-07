package schema

import "github.com/AlexMin314/go-gopher/backend/todo/constant"

// type ReqTodo struct {
// 	Title   string           `json:"title,omitempty" bson:"title"`
// 	Checked constant.Checker `json:"checked,omitempty" bson:"checked"`
// }

type Payload struct {
	ID       string
	Title    string
	Checked  constant.Checker
	Todos    []Todo
	IDs      []string
	ToggleTo constant.Checker
}
