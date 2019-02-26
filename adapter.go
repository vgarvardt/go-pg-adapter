package pg_adapter

import "errors"

// ErrNoRows is the driver-agnostic error returned when no record is found
var ErrNoRows = errors.New("sql: no rows in result set")

// Adapter represents DB access layer interface for different PostgreSQL drivers
type Adapter interface {
	Exec(query string, args ...interface{}) error
	SelectOne(dst interface{}, query string, args ...interface{}) error
}
