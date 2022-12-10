package internalhttp

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", time.Now().String(), r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
