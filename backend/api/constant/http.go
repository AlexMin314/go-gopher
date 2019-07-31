package constant

const (
	InternalServerError = "There was an error on the server and the request could not be completed."
)

const (
	EmptyTodoPayloadError = "Any todo data required to proceed the request"
)

type Checker string

const (
	Yes Checker = "Y"
	No  Checker = "N"
)
