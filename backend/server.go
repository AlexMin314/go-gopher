package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AlexMin314/go-gopher/backend/api"
	"github.com/AlexMin314/go-gopher/backend/config"
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

	c := config.InitServerConfig()
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(c.Port), RequestLogger(r)))
}
