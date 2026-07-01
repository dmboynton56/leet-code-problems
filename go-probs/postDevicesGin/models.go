// Problem: POST /devices — Gin handler that decodes JSON, validates, calls a store, returns JSON.
//
// This folder is ONLY the HTTP handler pattern. The store is an interface — you wire
// a real implementation in main or in tests (see the other standalone folders).
package main

import "context"

// CreateDeviceRequest maps directly from the JSON body.
// json tags tell encoding/json the wire format; binding tags tell Gin's validator
// what's required — validation lives on the struct, not scattered if-checks.
type CreateDeviceRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Device struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Store is the boundary between HTTP and persistence. Depending on an interface
// (not a concrete map or *sql.DB) is core Go design: swap implementations without
// rewriting handlers. Small interfaces (one method here) are preferred in Go.
type Store interface {
	CreateDevice(ctx context.Context, req CreateDeviceRequest) (Device, error)
}
