package mw

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func compose(mw []Middleware, finalFn http.HandlerFunc) http.HandlerFunc {
	composed := finalFn
	for i := len(mw) - 1; i >= 0; i-- {
		composed = mw[i](composed)
	}
	return composed
}

func ComposeMiddlewares(mw ...Middleware) Middleware {
	return func(finalHander http.HandlerFunc) http.HandlerFunc {
		composed := compose(mw, finalHander)
		return func(w http.ResponseWriter, r *http.Request) {
			composed(w, r)
		}
	}
}

// left this as learning history, the repo is using mux
