// Problem: given []Event, return the latest event for each device.
//
// This folder is a standalone algorithm exercise — no HTTP, no database.
// Interview checkboxes: sane struct, simple loop, handle ties via timestamp, state complexity.
package main

import "time"

// Event is a single device reading. Plain structs are Go's default data carrier —
// no classes, no framework. time.Time gives you comparable timestamps for free.
type Event struct {
	DeviceID  string
	Status    string
	CreatedAt time.Time
}

// LatestByDevice walks the slice once and keeps the newest event per DeviceID.
//
// Why Go fits this well:
//   - map[string]Event is a built-in hash map with O(1) average lookup
//   - for range is the idiomatic way to iterate slices
//   - comma-ok lookup (current, ok := latest[id]) is the standard "key exists?" pattern
//
// Time: O(n). Space: O(d) where d = number of distinct devices.
func LatestByDevice(events []Event) map[string]Event {
	latest := make(map[string]Event)

	for _, event := range events {
		current, seen := latest[event.DeviceID]

		if !seen || event.CreatedAt.After(current.CreatedAt) {
			latest[event.DeviceID] = event
		}
	}

	return latest
}
