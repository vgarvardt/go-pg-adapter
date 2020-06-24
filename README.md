# go-pg-adapter

[![GoDoc](https://godoc.org/github.com/vgarvardt/go-pg-adapter?status.svg)](https://godoc.org/github.com/vgarvardt/go-pg-adapter)
[![Coverage Status](https://codecov.io/gh/vgarvardt/go-pg-adapter/branch/master/graph/badge.svg)](https://codecov.io/gh/vgarvardt/go-pg-adapter)
[![ReportCard](https://goreportcard.com/badge/github.com/vgarvardt/go-pg-adapter)](https://goreportcard.com/report/github.com/vgarvardt/go-pg-adapter)
[![License](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)

Simple adapter interface and implementations for different PostgreSQL drivers for Go.

```go
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
```

## Install

```bash
$ go get -u -v github.com/vgarvardt/go-pg-adapter
```

## PostgreSQL drivers

The package bundles the following adapter implementations:

- `database/sql.DB` (e.g. [`github.com/lib/pq`](https://github.com/lib/pq)) - `github.com/vgarvardt/go-pg-adapter/sqladapter.New()`
- `database/sql.Conn` (e.g. [`github.com/lib/pq`](https://github.com/lib/pq)) - `github.com/vgarvardt/go-pg-adapter/sqladapter.NewConn()`
- [`github.com/jmoiron/sqlx.DB`](https://github.com/jmoiron/sqlx) - `github.com/vgarvardt/go-pg-adapter/sqladapter.NewX()`
- [`github.com/jackc/pgx.Conn`](https://github.com/jackc/pgx) (pgx v3) - `github.com/vgarvardt/go-pg-adapter/pgx3adapter.NewConn()`
- [`github.com/jackc/pgx.ConnPool`](https://github.com/jackc/pgx) (pgx v3) - `github.com/vgarvardt/go-pg-adapter/pgx3adapter.NewConnPool()`
- [`github.com/jackc/pgx/v4.Conn`](https://github.com/jackc/pgx) (pgx v4) - `github.com/vgarvardt/go-pg-adapter/pgx4adapter.NewConn()`
- [`github.com/jackc/pgx/v4/pgxpool.Pool`](https://github.com/jackc/pgx) (pgx v4) - `github.com/vgarvardt/go-pg-adapter/pgx4adapter.NewPool()`

## Testing

Linter and tests are running for every Pul Request, but it is possible to run linter
and tests locally using `docker` and `make`.

Run linter: `make link`. This command runs liner in docker container with the project
source code mounted.

Run tests: `make test`. This command runs project dependencies in docker containers
if they are not started yet and runs go tests with coverage.

## MIT License

```
Copyright (c) 2020 Vladimir Garvardt
```
