package api

import (
	"log"
	"net/http"
	"time"
)

// LogWriter is a custom ResponseWriter for handling status code.
type LogWriter struct {
	http.ResponseWriter
	done   bool
	status int
}

// ResponseWriter interface signature.
func (lw *LogWriter) WriteHeader(status int) {
	lw.done = true
	lw.status = status
	lw.ResponseWriter.WriteHeader(status)
}

// ResponseWriter interface signature.
func (lw *LogWriter) Write(b []byte) (int, error) {
	return lw.ResponseWriter.Write(b)
}

// Middleware for logging.
func middlwareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("new request. Method: %s | Path: '%s'", r.Method, r.URL.Path)

		// retrieve requestID
		requestID := r.Context().Value(keyRequestID)

		// create logwriter
		lw := &LogWriter{ResponseWriter: w}
		next.ServeHTTP(lw, r)

		// if status code is not set, return internal server error
		if !lw.done {
			lw.ResponseWriter.WriteHeader(http.StatusInternalServerError)
			lw.Write([]byte(internalErrorMsg))
		}
		log.Printf("Path: '%s' | RequestId: %v | Latency: %v | Status code: %v",
			r.URL.Path,
			requestID,
			time.Since(start),
			lw.status)
	})
}
