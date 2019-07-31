package schema

type Todo struct {
	Title   string `json:"title,omitempty"`
	Checked bool   `json:"checked,omitempty"`
}
