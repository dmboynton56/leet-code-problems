package main

import (
	"testing"
	"time"
)

func TestLatestByDevice(t *testing.T) {
	now := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	events := []Event{
		{DeviceID: "a", Status: "old", CreatedAt: now},
		{DeviceID: "a", Status: "new", CreatedAt: now.Add(time.Minute)},
		{DeviceID: "b", Status: "ok", CreatedAt: now},
	}

	got := LatestByDevice(events)

	if got["a"].Status != "new" {
		t.Fatalf("device a: got %q, want %q", got["a"].Status, "new")
	}
	if got["b"].Status != "ok" {
		t.Fatalf("device b: got %q, want %q", got["b"].Status, "ok")
	}
}
