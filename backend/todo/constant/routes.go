package constant

const (
	ApiPath    = "/api"
	ApiVersion = "/v1"
	Todo       = "/todo"
)

const (
	TodoApiRoute  = ApiPath + ApiVersion + Todo
	TodoIdPattern = "/{id:[0-9]+}"
)
