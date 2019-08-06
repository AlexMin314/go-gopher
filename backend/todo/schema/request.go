package schema

// type ReqTodo struct {
// 	Title   string           `json:"title,omitempty" bson:"title"`
// 	Checked constant.Checker `json:"checked,omitempty" bson:"checked"`
// }

type Request struct {
	Todo []Todo
}
