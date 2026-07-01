// Problem: how would device storage work with Postgres instead of memory?
//
// Standalone repository layer — no HTTP, no Gin. Know the shape for interviews.
package main

import "context"

type CreateDeviceRequest struct {
	ID   string
	Name string
}

type Device struct {
	ID   string
	Name string
}

// DeviceRepository is the persistence contract. Naming varies by team (Store, Repo, DAO).
// A single-method interface is idiomatic Go — define it where it's consumed, or here
// if the repo package is shared across services.
type DeviceRepository interface {
	CreateDevice(ctx context.Context, req CreateDeviceRequest) (Device, error)
}
