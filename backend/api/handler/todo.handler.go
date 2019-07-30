package handler

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/AlexMin314/go-gopher/backend/api/repository"
	"github.com/AlexMin314/go-gopher/backend/api/schema"
)

type ResponseError struct {
	Err error
}

type Response struct {
	ID    repository.ID `json:"id,omitempty"`
	Task  schema.Task   `json:"task"`
	Error ResponseError `json:"error"`
}

var memDB = repository.NewMemoryDataAccess()

func TodoApiHandler(w http.ResponseWriter, r *http.Request) {
	//
	requestDump, _ := httputil.DumpRequest(r, true)
	log.Println("-------------------------------")
	log.Println(string(requestDump))
	log.Println("-------------------------------")
}
