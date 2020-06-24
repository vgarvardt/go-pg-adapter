package sqladapter

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	pgAdapter "github.com/vgarvardt/go-pg-adapter"
)

// DB is the adapter type for sqlx.DB connection type
type DB struct {
	db *sqlx.DB
}

// New instantiates sqlx.DB connection adapter from sql.DB connection
func New(db *sql.DB) *DB {
	// The driverName of the original database is required for named query support - we do not use it here
	return &DB{sqlx.NewDb(db, "")}
}

// NewX instantiates sqlx.DB connection adapter
func NewX(db *sqlx.DB) *DB {
	return &DB{db}
}

// Conn is the adapter type for sql.Conn connection type
type Conn struct {
	conn *sql.Conn
}

// NewConn instantiated and returns adapter type for sql.Conn connection type
func NewConn(conn *sql.Conn) *Conn {
	return &Conn{conn}
}

// Exec runs a query and returns an error if any
func (a *DB) Exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := a.db.ExecContext(ctx, query, args...)
	return err
}

// SelectOne runs a select query and scans the object into a struct or returns an error
func (a *DB) SelectOne(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	if err := a.db.GetContext(ctx, dst, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return pgAdapter.ErrNoRows
		}
		return err
	}

	return nil
}

// Exec runs a query and returns an error if any
func (a *Conn) Exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := a.conn.ExecContext(ctx, query, args...)
	return err
}

// SelectOne runs a select query and scans the object into a struct or returns an error
func (a *Conn) SelectOne(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	// QueryRowContext does not work here as Row has very limited usage, we'll handle single scan logic manually
	rows, err := a.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	if !rows.Next() {
		return pgAdapter.ErrNoRows
	}

	return scanStruct(rows, dst)
}
