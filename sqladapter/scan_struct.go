package sqladapter

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx/reflectx"
)

var mapper = reflectx.NewMapperFunc("db", strings.ToLower)

func scanStruct(r *sql.Rows, dest interface{}) error {
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to ScanStruct destination")
	}
	if v.IsNil() {
		return errors.New("nil pointer passed to ScanStruct destination")
	}

	columns, err := r.Columns()
	if err != nil {
		return err
	}

	fields := mapper.TraversalsByName(v.Type(), columns)

	// if we are not unsafe and are missing fields, return an error
	if f, err := missingFields(fields); err != nil {
		return fmt.Errorf("missing destination name %s in %T", columns[f], dest)
	}
	values := make([]interface{}, len(columns))

	err = fieldsByTraversal(v, fields, values)
	if err != nil {
		return err
	}

	// scan into the struct field pointers and append to our results
	return r.Scan(values...)
}

func missingFields(transversals [][]int) (field int, err error) {
	for i, t := range transversals {
		if len(t) == 0 {
			return i, errors.New("missing field")
		}
	}
	return 0, nil
}

func fieldsByTraversal(v reflect.Value, traversals [][]int, values []interface{}) error {
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return errors.New("argument is not a struct")
	}

	for i, traversal := range traversals {
		if len(traversal) == 0 {
			values[i] = new(interface{})
			continue
		}

		f := reflectx.FieldByIndexes(v, traversal)
		if f.Kind() == reflect.Ptr {
			values[i] = f.Interface()
		} else {
			values[i] = f.Addr().Interface()
		}
	}

	return nil
}
