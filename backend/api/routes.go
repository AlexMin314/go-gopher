package api

const (
	root    = "/"
	apiPath = "/api"
	version = "/v1"
	todo    = "/todo"
)

const (
	todoApiRoute  = apiPath + version + todo
	todoIdPattern = "/{id:[0-9]+}"
)
