package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	store Store
}

func NewHandler(store Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) CreateDevice(c *gin.Context) {
	var req CreateDeviceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	device, err := h.store.CreateDevice(c.Request.Context(), req)
	if err != nil {
		writeCreateDeviceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, device)
}

// writeCreateDeviceError is where HTTP semantics live.
//
// Go pattern: store returns domain errors (ErrDuplicateDevice); handler translates
// them to status codes. The store never imports net/http — clean layer separation.
//
// errors.Is (not err == ErrDuplicateDevice) unwraps fmt.Errorf("...: %w", err) chains,
// so this keeps working if a service layer wraps the sentinel later.
func writeCreateDeviceError(c *gin.Context, err error) {
	if errors.Is(err, ErrDuplicateDevice) {
		c.JSON(http.StatusConflict, gin.H{"error": "device already exists"})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create device"})
}
