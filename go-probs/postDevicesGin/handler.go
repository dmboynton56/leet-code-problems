package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler receives its dependencies via constructor — no global state, no service locator.
// Explicit struct fields make dependencies visible and testable.
type Handler struct {
	store Store
}

func NewHandler(store Store) *Handler {
	return &Handler{store: store}
}

// CreateDevice handles POST /devices.
//
// Four steps (say this in an interview):
//  1. Decode JSON into a typed struct
//  2. Validate (ShouldBindJSON + binding tags)
//  3. Call store with request context
//  4. Return JSON + status code
//
// Gin wraps net/http — you still get *http.Request inside c.Request. Gin adds routing,
// middleware, and binding; the underlying model stays standard library HTTP.
func (h *Handler) CreateDevice(c *gin.Context) {
	var req CreateDeviceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	device, err := h.store.CreateDevice(c.Request.Context(), req)
	if err != nil {
		// Simple version: every store failure is a 500. A separate problem covers
		// mapping specific errors (like duplicate ID) to 409 Conflict.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create device"})
		return
	}

	c.JSON(http.StatusCreated, device)
}
