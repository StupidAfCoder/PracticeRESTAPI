package middlewares

import (
	"net/http"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriterHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
