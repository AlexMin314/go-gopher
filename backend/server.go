package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AlexMin314/go-gopher/backend/api"
)

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}

func main() {
	r := api.RegisterRoutes()
	http.Handle("/", r)

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Printf("Server is starting...")

	log.Fatal(http.ListenAndServe(":8080", RequestLogger(r)))
}
