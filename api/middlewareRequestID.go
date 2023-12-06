package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// key type for passing value trough context
type key string

const (
	keyRequestID     key = "request_id"
	internalErrorMsg     = "500 internal server error"
)

// Middleware to set a requestID.
func middlewareRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create uuid for request
		uuid, err := uuid.NewUUID()
		if err != nil {
			log.Printf("middlewareRequestID() %s", err.Error())
			http.Error(w, fmt.Sprintf(internalErrorMsg), http.StatusInternalServerError)
		}

		// pass uuid trough context
		ctx := context.WithValue(r.Context(), keyRequestID, uuid.String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
