package main

import (
	"context"
	"errors"
	"sync"
)

var ErrDuplicateDevice = errors.New("device already exists")

type MemoryStore struct {
	mu      sync.Mutex
	devices map[string]Device
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{devices: make(map[string]Device)}
}

func (s *MemoryStore) CreateDevice(ctx context.Context, req CreateDeviceRequest) (Device, error) {
	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.devices[req.ID]; exists {
		return Device{}, ErrDuplicateDevice
	}

	device := Device{ID: req.ID, Name: req.Name}
	s.devices[device.ID] = device
	return device, nil
}
