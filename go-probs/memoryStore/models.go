// Problem: implement a simple in-memory device store (map + mutex).
//
// This folder focuses on concurrency-safe fake storage — no HTTP layer here.
// In an interview you'd implement this when asked "where does the data go?"
package main

// CreateDeviceRequest is what the store receives after JSON decode/validation elsewhere.
type CreateDeviceRequest struct {
	ID   string
	Name string
}

// Device is the persisted record.
type Device struct {
	ID   string
	Name string
}
