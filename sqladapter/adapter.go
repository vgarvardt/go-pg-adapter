package sqladapter

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/vgarvardt/go-pg-adapter"
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

func NewConn(conn *sql.Conn) *Conn {
	return &Conn{conn}
}

// Exec runs a query and returns an error if any
func (a *DB) Exec(query string, args ...interface{}) error {
	_, err := a.db.Exec(query, args...)
	return err
}

// SelectOne runs a select query and scans the object into a struct or returns an error
func (a *DB) SelectOne(dst interface{}, query string, args ...interface{}) error {
	if err := a.db.Get(dst, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return pgadapter.ErrNoRows
		}
		return err
	}

	return nil
}

// Exec runs a query and returns an error if any
func (a *Conn) Exec(query string, args ...interface{}) error {
	_, err := a.conn.ExecContext(context.Background(), query, args...)
	return err
}

// SelectOne runs a select query and scans the object into a struct or returns an error
func (a *Conn) SelectOne(dst interface{}, query string, args ...interface{}) error {
	if err := a.conn.QueryRowContext(context.Background(), query, args...).Scan(dst); err != nil {
		if err == sql.ErrNoRows {
			return pgadapter.ErrNoRows
		}
		return err
	}

	return nil
}
