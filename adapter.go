package pgadapter

import (
	"context"
	"errors"
)

var (
	// ErrNoRows is the driver-agnostic error returned when no record is found
	ErrNoRows = errors.New("sql: no rows in result set")
	// ErrManyRows is the driver-agnostic error returned when more than one record is found
	// while only one was expected
	ErrManyRows = errors.New("sql: more than one row in result set")
)

// Adapter represents DB access layer interface for different PostgreSQL drivers
type Adapter interface {
	Exec(ctx context.Context, query string, args ...interface{}) error
	SelectOne(ctx context.Context, dst interface{}, query string, args ...interface{}) error
}
