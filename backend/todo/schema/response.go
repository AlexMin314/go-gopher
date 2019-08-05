package schema

type ID string

type Response struct {
	ID   ID   `json:"id,omitempty"`
	Todo Todo `json:"todo,omitempty"`
}
