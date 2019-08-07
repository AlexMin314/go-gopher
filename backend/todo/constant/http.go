package constant

const (
	InternalServerError = "There was an error on the server and the request could not be completed."
)

const (
	EmptyTodoPayloadError    = "Any todo data required to proceed the request."
	ExceedMaxTodoInsertLimit = "Too many todos requested to be inserted."
)

type Checker string
type Status string

const (
	Yes     Checker = "Y"
	No      Checker = "N"
	Success Status  = "success"
	Failure Status  = "failure"
)
