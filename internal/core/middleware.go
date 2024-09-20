package core

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// Chain chains multiple middlewares into a single middleware
func Chain(middlewares ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}
