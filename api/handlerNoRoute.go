package api

import (
	"fmt"
	"io"
	"net/http"
)

func (s *Server) HandlerNoRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	msg := fmt.Sprintf(`{"message": "not found %s %s"}`, r.Method, r.URL)
	io.WriteString(w, msg)
}
