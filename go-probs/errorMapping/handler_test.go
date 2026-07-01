package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func performRequest(router http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	return rec
}

func TestDuplicateReturnsConflict(t *testing.T) {
	gin.SetMode(gin.TestMode)

	store := NewMemoryStore()
	router := SetupRouter(store)

	body := `{"id":"device-1","name":"Kitchen Sensor"}`

	first := performRequest(router, http.MethodPost, "/devices", body)
	if first.Code != http.StatusCreated {
		t.Fatalf("first: got %d, want %d", first.Code, http.StatusCreated)
	}

	second := performRequest(router, http.MethodPost, "/devices", body)
	if second.Code != http.StatusConflict {
		t.Fatalf("duplicate: got %d, want %d", second.Code, http.StatusConflict)
	}
}
