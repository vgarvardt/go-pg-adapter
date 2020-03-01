# go-pg-adapter

[![Build][Build-Status-Image]][Build-Status-Url] [![Codecov][codecov-image]][codecov-url] [![ReportCard][reportcard-image]][reportcard-url] [![GoDoc][godoc-image]][godoc-url] [![License][license-image]][license-url]

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

## How to run tests

You will need running PostgreSQL instance. E.g. the one running in docker and exposing a port to a host system

```bash
docker run --rm -p 5432:5432 -it -e POSTGRES_PASSWORD=pgadapter -e POSTGRES_USER=pgadapter -e POSTGRES_DB=pgadapter postgres:10
```

Now you can run tests using the running PostgreSQL instance using `PG_URI` environment variable

```bash
PG_URI=postgres://pgadapter:pgadapter@localhost:5432/pgadapter?sslmode=disable go test -cover ./...
```

## MIT License

```
Copyright (c) 2019 Vladimir Garvardt
```

[Build-Status-Url]: https://travis-ci.org/vgarvardt/go-pg-adapter
[Build-Status-Image]: https://travis-ci.org/vgarvardt/go-pg-adapter.svg?branch=master
[codecov-url]: https://codecov.io/gh/vgarvardt/go-pg-adapter
[codecov-image]: https://codecov.io/gh/vgarvardt/go-pg-adapter/branch/master/graph/badge.svg
[reportcard-url]: https://goreportcard.com/report/github.com/vgarvardt/go-pg-adapter
[reportcard-image]: https://goreportcard.com/badge/github.com/vgarvardt/go-pg-adapter
[godoc-url]: https://godoc.org/github.com/vgarvardt/go-pg-adapter
[godoc-image]: https://godoc.org/github.com/vgarvardt/go-pg-adapter?status.svg
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/npm/l/express.svg
