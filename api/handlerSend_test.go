package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Victor-Acrani/backend-test-klever/api"
	"github.com/Victor-Acrani/backend-test-klever/types"
	"github.com/gorilla/mux"
)

func TestHandlerSend(t *testing.T) {
	t.Run("address with founds", func(t *testing.T) {
		// create server
		server := api.NewServer(":8080")
		if server == nil {
			t.Fatalf("failed to create server")
		}
		defer server.Shutdown(context.Background())

		// create router
		router := mux.NewRouter()
		router.HandleFunc("/api/v2/send", server.HandlerSend).Methods(http.MethodPost)

		send := types.SendRequestDTO{
			Address: "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh",
			Amount:  "100000",
		}

		// create request
		b, _ := json.Marshal(&send)

		req := httptest.NewRequest(http.MethodPost, "/api/v2/send", bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		// run http request
		router.ServeHTTP(rr, req)

		// check status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: expected %v, actual %v",
				http.StatusOK, status)
		}
	})

	t.Run("address without enough founds", func(t *testing.T) {
		// create server
		server := api.NewServer(":8080")
		if server == nil {
			t.Fatalf("failed to create server")
		}
		defer server.Shutdown(context.Background())

		// create router
		router := mux.NewRouter()
		router.HandleFunc("/api/v2/send", server.HandlerSend).Methods(http.MethodPost)

		send := types.SendRequestDTO{
			Address: "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r",
			Amount:  "100000",
		}

		// create request
		b, _ := json.Marshal(&send)

		req := httptest.NewRequest(http.MethodPost, "/api/v2/send", bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")
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
