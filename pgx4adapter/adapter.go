package pgx4adapter

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	pgxHelpers "github.com/vgarvardt/pgx-helpers/v4"

	pgAdapter "github.com/vgarvardt/go-pg-adapter"
)

// Pool is the adapter type for PGx pool connection type
type Pool struct {
	conn *pgxpool.Pool
}

// NewPool instantiates PGx pool adapter
func NewPool(conn *pgxpool.Pool) *Pool {
	return &Pool{conn}
}

// Conn is the adapter type for PGx connection connection type
type Conn struct {
	conn *pgx.Conn
}

// NewConn instantiates PGx connection adapter
func NewConn(conn *pgx.Conn) *Conn {
	return &Conn{conn}
}

// Exec runs a query and returns an error if any
func (a *Pool) Exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := a.conn.Exec(ctx, query, args...)
	return err
}

// SelectOne runs a select query and scans the object into a struct or returns an error
func (a *Pool) SelectOne(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	rows, err := a.conn.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	var rowScanned int
	err = pgxHelpers.ScanStructs(rows, func() interface{} {
		return dst
	}, func(r interface{}) {
		rowScanned++
	})

	if rowScanned > 1 {
		return pgAdapter.ErrManyRows
	}

	if rowScanned == 0 || err == pgx.ErrNoRows {
		return pgAdapter.ErrNoRows
	}

	return err
}

// Exec runs a query and returns an error if any
func (a *Conn) Exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := a.conn.Exec(ctx, query, args...)
	return err
}

// SelectOne runs a select query and scans the object into a struct or returns an error
func (a *Conn) SelectOne(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	rows, err := a.conn.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	var rowScanned int
	err = pgxHelpers.ScanStructs(rows, func() interface{} {
		return dst
	}, func(r interface{}) {
		rowScanned++
	})

	if rowScanned > 1 {
		return pgAdapter.ErrManyRows
	}

	if rowScanned == 0 || err == pgx.ErrNoRows {
		return pgAdapter.ErrNoRows
	}

	return err
}
