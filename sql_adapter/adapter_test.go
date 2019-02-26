package sql_adapter

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vgarvardt/go-pg-adapter"
)

var uri string

func TestMain(m *testing.M) {
	uri = os.Getenv("PG_URI")
	if uri == "" {
		fmt.Println("Env variable PG_URI is required to run the tests")
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestNewConn(t *testing.T) {
	conn, err := sql.Open("pgx", uri)
	require.NoError(t, err)

	defer func() {
		assert.NoError(t, conn.Close())
	}()

	runTests(t, New(conn), fmt.Sprintf("test_sql_%d", time.Now().UnixNano()))
}

func TestNewConnPool(t *testing.T) {
	conn, err := sql.Open("pgx", uri)
	require.NoError(t, err)

	defer func() {
		assert.NoError(t, conn.Close())
	}()

	runTests(t, NewX(sqlx.NewDb(conn, "")), fmt.Sprintf("test_sqlx_%d", time.Now().UnixNano()))
}

type TestRow struct {
	ID        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	Data      string    `db:"data"`
}

func runTests(t *testing.T, adapter pg_adapter.Adapter, table string) {
	query := fmt.Sprintf(`
CREATE TABLE %[1]s (
  id         SERIAL      NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  data       TEXT        NOT NULL,
  CONSTRAINT %[1]s_pkey PRIMARY KEY (id)
)`, table)
	err := adapter.Exec(query)
	require.NoError(t, err)

	originalRow := TestRow{
		CreatedAt: time.Now(),
		Data:      time.Now().Format(time.RFC1123),
	}
	err = adapter.Exec(fmt.Sprintf("INSERT INTO %s (created_at, data) VALUES ($1, $2)", table), originalRow.CreatedAt, originalRow.Data)
	require.NoError(t, err)

	var selectedRow TestRow
	err = adapter.SelectOne(&selectedRow, fmt.Sprintf("SELECT * FROM %s WHERE data = $1", table), originalRow.Data)
	require.NoError(t, err)

	assert.True(t, selectedRow.ID > 0)
	// time object string format is "2019-02-26 20:37:09.797161 +0100 CET m=+0.024514000" and the one from DB misses the last bit
	assert.Equal(t, originalRow.CreatedAt.Format(time.RFC1123), selectedRow.CreatedAt.Format(time.RFC1123))
	assert.Equal(t, originalRow.Data, selectedRow.Data)

	var unusedRow TestRow
	err = adapter.SelectOne(&unusedRow, fmt.Sprintf("SELECT * FROM %s WHERE data = $1", table), "foo bar")
	require.Error(t, err)
	assert.Equal(t, err, pg_adapter.ErrNoRows)
}
