package constant

const (
	ApiPath    = "/api"
	ApiVersion = "/v1"
	Todo       = "/todo"
	Mem        = "/m"
	Root       = ""
)

const (
	TodoApiRoute     = ApiPath + ApiVersion + Todo
	TodoIdPattern    = "/{id:[0-9a-z]+}"
	TodoMemIdPattern = Mem + TodoIdPattern
)
