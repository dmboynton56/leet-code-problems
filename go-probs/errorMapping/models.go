// Problem: map store errors to the right HTTP status codes (409 for duplicate, 500 for unknown).
//
// Self-contained mini-app — same device domain as other folders, but this problem
// is only about errors.Is + sentinel errors + handler-side HTTP semantics.
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
