package main

import (
	"log"
	"net/http"

	"github.com/AlexMin314/go-gopher/backend/api"
)

const (
	apiPath = "/api"
	version = "/v1"
)

func main() {
	r := api.RegisterRoutes()
	r.PathPrefix(apiPath + version)
	// todo: header setting
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
