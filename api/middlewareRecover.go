package api

import (
	"log"
	"net/http"
)

// Middleware for recovering.
func middlewareRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(w, internalErrorMsg, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
