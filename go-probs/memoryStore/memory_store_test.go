package main

import (
	"context"
	"errors"
	"testing"
)

func TestCreateDeviceDuplicate(t *testing.T) {
	store := NewMemoryStore()
	ctx := context.Background()
	req := CreateDeviceRequest{ID: "device-1", Name: "Kitchen Sensor"}

	if _, err := store.CreateDevice(ctx, req); err != nil {
		t.Fatalf("first create: %v", err)
	}

	_, err := store.CreateDevice(ctx, req)
	if !errors.Is(err, ErrDuplicateDevice) {
		t.Fatalf("got %v, want ErrDuplicateDevice", err)
	}
}
