package main

import (
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create device"})
		return
	}

	c.JSON(http.StatusCreated, device)
}
