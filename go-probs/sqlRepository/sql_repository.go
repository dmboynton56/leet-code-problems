package main

import (
	"context"
	"database/sql"
	"errors"
)

/*
TABLE (say this in an interview):

	CREATE TABLE devices (
	    id TEXT PRIMARY KEY,
	    name TEXT NOT NULL
	);

WIRE-UP (conceptual — not in this folder):

	db, _ := sql.Open("postgres", dsn)
	repo := NewSQLRepository(db)
	// pass repo to handler via interface — handler code unchanged from in-memory version
*/

// SQLRepository implements DeviceRepository with database/sql.
//
// database/sql is stdlib. Drivers (pgx, lib/pq) register via blank import in main.
// QueryRowContext propagates context.Context from the HTTP request — if the client
// disconnects, the in-flight query can be cancelled.
type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) CreateDevice(ctx context.Context, req CreateDeviceRequest) (Device, error) {
	query := `
		INSERT INTO devices (id, name)
		VALUES ($1, $2)
		RETURNING id, name
	`

	var device Device

	err := r.db.QueryRowContext(ctx, query, req.ID, req.Name).Scan(&device.ID, &device.Name)
	if err != nil {
		return Device{}, mapInsertError(err)
	}

	return device, nil
}

func mapInsertError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return err
	}

	// With Postgres: check for unique_violation (SQLSTATE 23505) and return a
	// sentinel like ErrDuplicateDevice so the HTTP layer can respond with 409.
	return err
}

/*
WHY Go + SQL REPOS WORK WELL:

- Handler depends on an interface, not *sql.DB — swap memory for Postgres in main
- RETURNING avoids a second SELECT (one round trip)
- Context flows from net/http → store → database/sql without a special framework
- Repo tests can use a real test DB, sqlmock, or testcontainers — all common in Go
*/
