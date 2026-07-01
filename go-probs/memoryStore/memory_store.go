package main

import (
	"context"
	"errors"
	"sync"
)

// ErrDuplicateDevice is a sentinel error — a named, comparable error value.
// Go handlers (in a separate problem) can use errors.Is to branch on this
// without string-matching or error codes. Sentinels are idiomatic in Go stdlib
// (e.g. io.EOF, sql.ErrNoRows).
var ErrDuplicateDevice = errors.New("device already exists")

// MemoryStore is a thread-safe map-backed store.
//
// Go maps are NOT safe for concurrent writes. sync.Mutex from the stdlib is
// the straightforward fix — no external locking library, no channel gymnastics
// unless you specifically want them.
type MemoryStore struct {
	mu      sync.Mutex
	devices map[string]Device
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		devices: make(map[string]Device),
	}
}

// CreateDevice inserts a device or returns ErrDuplicateDevice.
//
// context.Context is accepted even though this in-memory impl ignores it.
// Accepting ctx keeps the function signature compatible with a future SQL
// implementation that honors cancellation via QueryRowContext.
func (s *MemoryStore) CreateDevice(ctx context.Context, req CreateDeviceRequest) (Device, error) {
	_ = ctx

	// Lock before touching the map; defer Unlock so we always release on any return path.
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.devices[req.ID]; exists {
		return Device{}, ErrDuplicateDevice
	}

	device := Device{ID: req.ID, Name: req.Name}
	s.devices[device.ID] = device

	return device, nil
}
