package api_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Victor-Acrani/backend-test-klever/api"
)

func TestHandlerNoRoute(t *testing.T) {
	server := api.NewServer(":8080")
	if server == nil {
		t.Fatalf("failed to create server")
	}
	defer server.Shutdown(context.Background())

	req, err := http.NewRequest(http.MethodGet, "/api/v2/wrongroute", nil)
	if err != nil {
		t.Fatalf("error to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.HandlerNoRoute)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: expected %v, actual %v",
			http.StatusOK, status)
	}

	expected := `{"message": "not found GET /api/v2/wrongroute"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned wrong response body: expected '%v', actual '%v'",
			expected, rr.Body.String())
	}
}
