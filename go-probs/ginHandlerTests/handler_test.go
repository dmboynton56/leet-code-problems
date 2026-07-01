package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// performRequest drives the router like a real HTTP client.
//
// Go's net/http/httptest is stdlib — no Gin-specific test harness. You get:
//   - httptest.NewRequest  → builds *http.Request
//   - httptest.ResponseRecorder → captures status, headers, body
//   - router.ServeHTTP(rec, req) → runs the full handler chain
//
// That last line is why you don't mock gin.Context: the router creates one for you.
func performRequest(router http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	return rec
}

func TestCreateDeviceSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	store := NewMemoryStore()
	router := SetupRouter(store)

	body := `{"id":"device-1","name":"Kitchen Sensor"}`
	rec := performRequest(router, http.MethodPost, "/devices", body)

	if rec.Code != http.StatusCreated {
		t.Fatalf("got status %d, want %d", rec.Code, http.StatusCreated)
	}
}

func TestCreateDeviceMissingName(t *testing.T) {
	gin.SetMode(gin.TestMode)

	store := NewMemoryStore()
	router := SetupRouter(store)

	body := `{"id":"device-1"}`
	rec := performRequest(router, http.MethodPost, "/devices", body)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("got status %d, want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestCreateDeviceDuplicate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	store := NewMemoryStore()
	router := SetupRouter(store)

	body := `{"id":"device-1","name":"Kitchen Sensor"}`

	first := performRequest(router, http.MethodPost, "/devices", body)
	if first.Code != http.StatusCreated {
		t.Fatalf("first: got status %d, want %d", first.Code, http.StatusCreated)
	}

	// Simple handler maps all store errors to 500. The errorMapping folder shows
	// upgrading duplicate ID to 409 Conflict — a separate, focused problem.
	second := performRequest(router, http.MethodPost, "/devices", body)
	if second.Code != http.StatusInternalServerError {
		t.Fatalf("duplicate: got status %d, want %d", second.Code, http.StatusInternalServerError)
	}
}
