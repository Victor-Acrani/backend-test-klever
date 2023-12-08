package api_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Victor-Acrani/backend-test-klever/api"
	"github.com/gorilla/mux"
)

func TestHandlerDetails(t *testing.T) {
	t.Run("valid address", func(t *testing.T) {
		// create server
		server := api.NewServer(":8080")
		if server == nil {
			t.Fatalf("failed to create server")
		}
		defer server.Shutdown(context.Background())

		// creater router
		router := mux.NewRouter()
		router.HandleFunc("/api/v2/details/{address}", server.HandlerDetails).Methods("GET")

		// create request
		req := httptest.NewRequest("GET", "/api/v2/details/bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r", nil)
		rr := httptest.NewRecorder()

		// run http request
		router.ServeHTTP(rr, req)

		// check status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: expected %v, actual %v",
				http.StatusOK, status)
		}
	})

	t.Run("invalid address", func(t *testing.T) {
		// create server
		server := api.NewServer(":8080")
		if server == nil {
			t.Fatalf("failed to create server")
		}
		defer server.Shutdown(context.Background())

		// creater router
		router := mux.NewRouter()
		router.HandleFunc("/api/v2/details/{address}", server.HandlerDetails).Methods("GET")

		// create request
		req := httptest.NewRequest("GET", "/api/v2/details/bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh", nil)
		rr := httptest.NewRecorder()

		// run http request
		router.ServeHTTP(rr, req)

		// check status code
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: expected %v, actual %v",
				http.StatusOK, status)
		}
	})
}
