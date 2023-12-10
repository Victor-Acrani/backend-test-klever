package api_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Victor-Acrani/backend-test-klever/api"
	"github.com/gorilla/mux"
)

func TestHandlerTransaction(t *testing.T) {
	t.Run("valid transaction", func(t *testing.T) {
		// create server
		server := api.NewServer(":8080")
		if server == nil {
			t.Fatalf("failed to create server")
		}
		defer server.Shutdown(context.Background())

		// create router
		router := mux.NewRouter()
		router.HandleFunc("/api/v2/tx/{tx}", server.HandlerTransaction).Methods(http.MethodGet)

		// create request
		req := httptest.NewRequest(http.MethodGet, "/api/v2/tx/3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f1830ab", nil)
		rr := httptest.NewRecorder()

		// run http request
		router.ServeHTTP(rr, req)

		// check status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: expected %v, actual %v",
				http.StatusOK, status)
		}
	})
	t.Run("invalid transaction", func(t *testing.T) {
		// create server
		server := api.NewServer(":8080")
		if server == nil {
			t.Fatalf("failed to create server")
		}
		defer server.Shutdown(context.Background())

		// create router
		router := mux.NewRouter()
		router.HandleFunc("/api/v2/tx/{tx}", server.HandlerTransaction).Methods(http.MethodGet)

		// create request
		req := httptest.NewRequest(http.MethodGet, "/api/v2/tx/3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f18", nil)
		rr := httptest.NewRecorder()

		// run http request
		router.ServeHTTP(rr, req)

		// check status code
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: expected %v, actual %v",
				http.StatusBadRequest, status)
		}
	})
}
