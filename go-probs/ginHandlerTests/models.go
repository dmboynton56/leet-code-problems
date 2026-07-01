// Problem: test POST /devices without manually building a gin.Context.
//
// This folder is a self-contained mini-app: models + store + handler + router + tests.
// Each piece is duplicated here on purpose — this problem is about httptest, not imports.
package main

import "context"

type CreateDeviceRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Device struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Store interface {
	CreateDevice(ctx context.Context, req CreateDeviceRequest) (Device, error)
}
