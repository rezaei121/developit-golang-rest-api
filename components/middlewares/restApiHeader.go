package middlewares

import (
	"net/http"
)

func RestApiHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("accept", "application/json")
		next.ServeHTTP(rw, r)
	})
}
